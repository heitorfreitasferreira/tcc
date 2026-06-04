#!/usr/bin/env python3
"""Baixa PDFs faltantes ou corrompidos para vault/papers/pdfs/<bibtex-key>.pdf.

Fontes (em ordem):
  1. Unpaywall (todas as oa_locations com url_for_pdf)
  2. Semantic Scholar (openAccessPdf)
  3. OpenAlex (best_oa_location.pdf_url e oa_locations)
  4. Sci-Hub (opcional, --use-scihub; requer pacote scihub no PYTHONPATH)

Validação: magic bytes %PDF, pdfinfo com Pages >= --min-pages.

Uso:
  ./scripts/download-pdfs.sh --dry-run
  ./scripts/download-pdfs.sh --limit 5 --min-rating 4
  ./scripts/download-pdfs.sh --keys blum2005acointro,larranaga1999ga
  ./scripts/download-pdfs.sh --use-scihub --limit 10
"""

from __future__ import annotations

import argparse
import json
import re
import subprocess
import sys
import time
import urllib.error
import urllib.parse
import urllib.request
from dataclasses import dataclass
from pathlib import Path
from typing import Callable, Iterable, Optional

REPO_ROOT = Path(__file__).resolve().parent.parent
VAULT_DIR = REPO_ROOT / "vault"
PAPERS_DIR = VAULT_DIR / "papers"
PDFS_DIR = PAPERS_DIR / "pdfs"
CORRUPTED_DIR = PDFS_DIR / "corrupted"


def resolve_pdf_path(pdf_field: str) -> Optional[Path]:
    """Caminhos no frontmatter são relativos à raiz do vault (ex.: papers/pdfs/x.pdf)."""
    if not pdf_field:
        return None
    p = Path(pdf_field)
    if p.is_absolute():
        return p
    if pdf_field.startswith("papers/"):
        return VAULT_DIR / pdf_field
    return PAPERS_DIR / pdf_field

DEFAULT_EMAIL = "heitor@ufu.br"
USER_AGENT = "tcc-vault-pdf-fetch/1.0 (academic; mailto:heitor@ufu.br)"


@dataclass
class PaperMeta:
    key: str
    path: Path
    doi: str
    pdf_field: str
    rating: int
    title: str


def parse_frontmatter(md_path: Path) -> dict[str, str]:
    text = md_path.read_text(encoding="utf-8", errors="replace")
    if not text.startswith("---"):
        return {}
    end = text.find("\n---", 3)
    if end < 0:
        return {}
    block = text[3:end]
    out: dict[str, str] = {}
    for line in block.splitlines():
        m = re.match(r"^([a-zA-Z0-9_-]+):\s*(.*)$", line.strip())
        if not m:
            continue
        key, val = m.group(1), m.group(2).strip()
        if val.startswith('"') and val.endswith('"'):
            val = val[1:-1]
        elif val.startswith("'") and val.endswith("'"):
            val = val[1:-1]
        out[key] = val
    return out


def pdf_is_valid(path: Path, min_pages: int) -> bool:
    if not path.is_file() or path.stat().st_size < 1024:
        return False
    head = path.read_bytes()[:8]
    if not head.startswith(b"%PDF"):
        return False
    try:
        proc = subprocess.run(
            ["pdfinfo", str(path)],
            capture_output=True,
            text=True,
            check=False,
            timeout=15,
        )
    except (FileNotFoundError, subprocess.TimeoutExpired):
        return False
    if proc.returncode != 0:
        return False
    pages = 0
    for line in proc.stdout.splitlines():
        if line.startswith("Pages:"):
            pages = int(line.split(":", 1)[1].strip())
            break
    return pages >= min_pages


def target_pdf_path(key: str) -> Path:
    return PDFS_DIR / f"{key}.pdf"


def needs_download(meta: PaperMeta, min_pages: int) -> bool:
    target = target_pdf_path(meta.key)
    if pdf_is_valid(target, min_pages):
        return False
    if meta.pdf_field and "corrupted/" in meta.pdf_field:
        return True
    if not meta.pdf_field:
        return True
    # YAML aponta para outro caminho, mas o arquivo canônico não existe
    alt = resolve_pdf_path(meta.pdf_field)
    if alt and alt.is_file() and pdf_is_valid(alt, min_pages):
        return True  # renomear para <bibtex-key>.pdf
    return True


def load_papers(keys_filter: Optional[set[str]], min_rating: int) -> list[PaperMeta]:
    papers: list[PaperMeta] = []
    for md in sorted(PAPERS_DIR.glob("*.md")):
        if md.name == "index.md":
            continue
        fm = parse_frontmatter(md)
        key = fm.get("bibtex-key") or md.stem
        if keys_filter and key not in keys_filter:
            continue
        rating = int(fm.get("rating") or 0)
        if rating < min_rating:
            continue
        papers.append(
            PaperMeta(
                key=key,
                path=md,
                doi=(fm.get("doi") or "").strip(),
                pdf_field=(fm.get("pdf") or "").strip(),
                rating=rating,
                title=(fm.get("title") or "").strip(),
            )
        )
    return papers


def http_get_json(url: str, timeout: float = 30.0) -> Optional[dict]:
    req = urllib.request.Request(url, headers={"User-Agent": USER_AGENT})
    try:
        with urllib.request.urlopen(req, timeout=timeout) as resp:
            return json.loads(resp.read().decode("utf-8", errors="replace"))
    except (urllib.error.URLError, json.JSONDecodeError, TimeoutError):
        return None


def download_url(url: str, dest: Path, timeout: float = 120.0) -> bool:
    dest.parent.mkdir(parents=True, exist_ok=True)
    tmp = dest.with_suffix(".pdf.part")
    req = urllib.request.Request(
        url,
        headers={
            "User-Agent": USER_AGENT,
            "Accept": "application/pdf,*/*",
        },
    )
    try:
        with urllib.request.urlopen(req, timeout=timeout) as resp:
            data = resp.read()
        if len(data) < 1024 or not data.startswith(b"%PDF"):
            return False
        tmp.write_bytes(data)
        tmp.replace(dest)
        return True
    except (urllib.error.URLError, OSError):
        if tmp.exists():
            tmp.unlink(missing_ok=True)
        return False


def unpaywall_urls(doi: str, email: str) -> list[str]:
    enc = urllib.parse.quote(doi, safe="")
    data = http_get_json(f"https://api.unpaywall.org/v2/{enc}?email={urllib.parse.quote(email)}")
    if not data:
        return []
    urls: list[str] = []
    for loc in data.get("oa_locations") or []:
        u = loc.get("url_for_pdf")
        if u and u not in urls:
            urls.append(u)
    best = data.get("best_oa_location") or {}
    u = best.get("url_for_pdf")
    if u and u not in urls:
        urls.insert(0, u)
    return urls


def semantic_scholar_url(doi: str) -> Optional[str]:
    enc = urllib.parse.quote(f"DOI:{doi}", safe="")
    data = http_get_json(
        f"https://api.semanticscholar.org/graph/v1/paper/{enc}?fields=openAccessPdf"
    )
    if not data:
        return None
    oa = data.get("openAccessPdf") or {}
    return oa.get("url")


def openalex_urls(doi: str) -> list[str]:
    enc = urllib.parse.quote(f"https://doi.org/{doi}", safe="")
    data = http_get_json(f"https://api.openalex.org/works/{enc}")
    if not data:
        return []
    urls: list[str] = []
    best = data.get("best_oa_location") or {}
    u = best.get("pdf_url")
    if u:
        urls.append(u)
    for loc in data.get("oa_locations") or []:
        u = loc.get("pdf_url")
        if u and u not in urls:
            urls.append(u)
    return urls


def scihub_fetch(doi: str) -> Optional[str]:
    scihub_dir = Path.home() / ".config/opencode/mcp/Sci-Hub-MCP-Server"
    if not scihub_dir.is_dir():
        return None
    sys.path.insert(0, str(scihub_dir))
    try:
        from sci_hub_search import download_paper, search_paper_by_doi  # type: ignore
    except ImportError:
        return None
    result = search_paper_by_doi(doi)
    if result.get("status") != "success":
        return None
    pdf_url = result.get("pdf_url")
    if not pdf_url:
        return None
    tmp = PDFS_DIR / f".{doi.replace('/', '_')}.part.pdf"
    if download_paper(pdf_url, str(tmp)):
        return str(tmp)
    return None


def try_sources(
    meta: PaperMeta,
    *,
    email: str,
    use_scihub: bool,
    min_pages: int,
    dest: Path,
) -> tuple[bool, str]:
    if meta.pdf_field and "corrupted/" not in meta.pdf_field:
        alt = resolve_pdf_path(meta.pdf_field)
        if alt and alt.is_file() and pdf_is_valid(alt, min_pages):
            if alt.resolve() != dest.resolve():
                dest.parent.mkdir(parents=True, exist_ok=True)
                if dest.exists():
                    dest.unlink()
                alt.replace(dest)
            return True, f"renomeado de {meta.pdf_field}"

    candidates: list[tuple[str, str]] = []

    if meta.doi:
        for u in unpaywall_urls(meta.doi, email):
            candidates.append(("unpaywall", u))
        ss = semantic_scholar_url(meta.doi)
        if ss:
            candidates.append(("semantic-scholar", ss))
        for u in openalex_urls(meta.doi):
            candidates.append(("openalex", u))

    seen: set[str] = set()
    for source, url in candidates:
        if url in seen:
            continue
        seen.add(url)
        if download_url(url, dest) and pdf_is_valid(dest, min_pages):
            return True, source

    if use_scihub and meta.doi:
        part = scihub_fetch(meta.doi)
        if part:
            p = Path(part)
            if p.is_file() and pdf_is_valid(p, min_pages):
                p.replace(dest)
                return True, "scihub"
            p.unlink(missing_ok=True)

    return False, "nenhuma fonte"


def update_note_pdf_field(meta: PaperMeta, rel_path: str) -> None:
    text = meta.path.read_text(encoding="utf-8", errors="replace")
    new_pdf = f'pdf: "{rel_path}"'
    if re.search(r"^pdf:\s*", text, flags=re.M):
        text = re.sub(r'^pdf:\s*.*$', new_pdf, text, count=1, flags=re.M)
    else:
        text = text.replace("---\n", f"---\n{new_pdf}\n", 1)
    embed = f"![[{meta.key}.pdf]]"
    if "![[" in text:
        text = re.sub(r"!\[\[[^\]]+\]\]", embed, text, count=1)
    elif "## PDF" in text:
        text = text.replace("## PDF\n", f"## PDF\n\n{embed}\n", 1)
    meta.path.write_text(text, encoding="utf-8")


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--dry-run", action="store_true")
    parser.add_argument("--limit", type=int, default=0)
    parser.add_argument("--min-rating", type=int, default=0)
    parser.add_argument("--min-pages", type=int, default=2)
    parser.add_argument("--keys", type=str, default="")
    parser.add_argument("--email", type=str, default=DEFAULT_EMAIL)
    parser.add_argument("--use-scihub", action="store_true")
    parser.add_argument(
        "--no-update-notes",
        action="store_true",
        help="não alterar o frontmatter das notas após download",
    )
    parser.add_argument("--delay", type=float, default=1.0, help="segundos entre downloads")
    parser.add_argument(
        "--rename-only",
        action="store_true",
        help="apenas normaliza PDFs válidos já existentes (sem baixar)",
    )
    args = parser.parse_args()

    keys_filter: Optional[set[str]] = None
    if args.keys.strip():
        keys_filter = {k.strip() for k in args.keys.split(",") if k.strip()}

    all_meta = load_papers(keys_filter, args.min_rating)
    pending = [m for m in all_meta if needs_download(m, args.min_pages)]
    pending.sort(key=lambda m: (-m.rating, m.key))

    if args.limit > 0:
        pending = pending[: args.limit]

    if not pending:
        print("Nenhum artigo pendente com os filtros informados.")
        return 0

    print(f"Pendentes: {len(pending)}")
    ok = fail = skip = 0

    for i, meta in enumerate(pending, 1):
        dest = target_pdf_path(meta.key)
        rel = f"papers/pdfs/{meta.key}.pdf"
        label = f"[{i}/{len(pending)}] {meta.key}"
        if not meta.doi and not meta.pdf_field:
            print(f"{label}: sem DOI — pular (busca manual)")
            skip += 1
            continue
        if args.dry_run:
            print(f"{label}: doi={meta.doi or '(vazio)'} rating={meta.rating} -> {dest.name}")
            continue

        if args.rename_only:
            alt = resolve_pdf_path(meta.pdf_field)
            if alt and alt.is_file() and pdf_is_valid(alt, args.min_pages):
                if alt.resolve() != dest.resolve():
                    dest.parent.mkdir(parents=True, exist_ok=True)
                    if dest.exists():
                        dest.unlink()
                    alt.replace(dest)
                success, source = True, "renomeado"
            else:
                success, source = False, "sem arquivo válido"
        else:
            success, source = try_sources(
                meta,
                email=args.email,
                use_scihub=args.use_scihub,
                min_pages=args.min_pages,
                dest=dest,
            )
        if success:
            print(f"{label}: OK ({source})")
            if not args.no_update_notes:
                update_note_pdf_field(meta, rel)
            ok += 1
        else:
            print(f"{label}: falhou (doi={meta.doi or 'n/a'})")
            if dest.exists():
                corrupt_name = dest.name
                CORRUPTED_DIR.mkdir(parents=True, exist_ok=True)
                dest.replace(CORRUPTED_DIR / corrupt_name)
            fail += 1

        if (
            i < len(pending)
            and args.delay > 0
            and not args.rename_only
            and not args.dry_run
        ):
            time.sleep(args.delay)

    if args.dry_run:
        return 0

    print(f"\nResumo: {ok} baixados, {fail} falhas, {skip} ignorados")
    return 0 if fail == 0 else 1


if __name__ == "__main__":
    raise SystemExit(main())
