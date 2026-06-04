#!/usr/bin/env python3
"""Gerenciar Lista de Siglas: scan .tex, manter notas em vault/ (siglas/ + projeto/ + areas/), gerar/validar Abreviaturas.tex."""

from __future__ import annotations

import argparse
import re
import sys
from collections import Counter
from dataclasses import dataclass, field
from pathlib import Path
from typing import Dict, List, Set, Tuple

ROOT = Path(__file__).resolve().parents[1]
DEFAULT_SIGLAS_DIR = ROOT / "vault" / "siglas"
DEFAULT_ABREV_TEX = ROOT / "monografia" / "abrev" / "Abreviaturas.tex"
SCRIPT_NAME = "gerar-lista-siglas.py"

TEX_DIRS: List[Path] = [
    ROOT / "monografia" / "cap_introducao",
    ROOT / "monografia" / "cap_fundamentacao",
    ROOT / "monografia" / "cap_proposta",
    ROOT / "monografia" / "cap_experimentos",
    ROOT / "monografia" / "cap_conclusao",
    ROOT / "monografia" / "ape_metodos",
    ROOT / "monografia" / "ape_sobre",
]

# Additional vault directories to search for siglas-tagged notes
VAULT_SCAN_DIRS: List[Path] = [
    ROOT / "vault" / "siglas",
    ROOT / "vault" / "projeto",
    ROOT / "vault" / "areas",
    ROOT / "vault" / "writing",
]

AC_CMD_RE = re.compile(r'\\(ac[sfpl]?)\{([A-Za-z0-9_-]+)\}')
ACRO_RE = re.compile(r'\\acro\{([A-Za-z0-9_-]+)\}\s*\{([^}]*)\}')
PAREN_ACRO_RE = re.compile(r'\(([A-Z][A-Z0-9][A-Za-z0-9]{0,10}(?:-[A-Z0-9][A-Za-z0-9]*)*)\)')

CANONICAL: Dict[str, str] = {
    "ABNT": "Associação Brasileira de Normas Técnicas",
    "ACO": "Otimização por Colônia de Formigas (Ant Colony Optimization)",
    "ACS": "Ant Colony System",
    "AP": "Problema de Designação (Assignment Problem)",
    "AS": "Ant System",
    "BF": "Busca Exaustiva (Brute Force)",
    "CD": "Diferença Crítica (Critical Difference)",
    "EAX": "Edge Assembly Crossover",
    "FSTSP": "Flying Sidekick Traveling Salesman Problem",
    "GA": "Algoritmo Genético (Genetic Algorithm)",
    "GNN": "Graph Neural Network",
    "GP": "Programação Genética (Genetic Programming)",
    "LB": "Limitante Inferior (Lower Bound)",
    "LKH": "Heurística Lin--Kernighan (Lin--Kernighan Heuristic)",
    "MILP": "Programação Linear Inteira Mista (Mixed-Integer Linear Programming)",
    "MMAS": "MAX-MIN Ant System",
    "NBR": "Norma Brasileira",
    "NP": "Nondeterministic Polynomial Time",
    "OX": "Crossover por Ordem (Order Crossover)",
    "POI": "Ponto de Interesse (Point of Interest)",
    "PPO": "Proximal Policy Optimization",
    "PSO": "Otimização por Enxame de Partículas (Particle Swarm Optimization)",
    "RL": "Aprendizagem por Reforço (Reinforcement Learning)",
    "TSP": "Problema do Caixeiro Viajante (Traveling Salesman Problem)",
    "TSP-SD-ATP": "Traveling Salesman Problem with Sequence-Dependent Angular Turn Penalties",
    "TSPLIB": "TSP Library",
    "UAV": "Veículo Aéreo Não Tripulado (Unmanned Aerial Vehicle)",
    "VANT": "Veículo Aéreo Não Tripulado",
    "VNS": "Busca em Vizinhança Variável (Variable Neighborhood Search)",
    "VRP": "Problema de Roteamento de Veículos (Vehicle Routing Problem)",
}


def _filename_for(key: str) -> str:
    return key.lower().replace("-", "_").replace(".", "_")


def _note_path(siglas_dir: Path, key: str) -> Path:
    return siglas_dir / f"{_filename_for(key)}.md"


# ---------------------------------------------------------------------------
# frontmatter helpers
# ---------------------------------------------------------------------------

def _parse_frontmatter(text: str) -> Tuple[Dict[str, str], str]:
    """Extract YAML frontmatter as flat dict and remaining body."""
    lines = text.splitlines()
    if not lines or lines[0].strip() != "---":
        return {}, text
    end = -1
    for i in range(1, len(lines)):
        if lines[i].strip() == "---":
            end = i
            break
    if end == -1:
        return {}, text

    fm: Dict[str, str] = {}
    in_block_list = False
    block_key = ""
    for line in lines[1:end]:
        stripped = line.strip()
        if not stripped:
            continue
        if stripped.startswith("- ") and in_block_list:
            # block list item
            val = stripped[2:].strip().strip('"').strip("'")
            if val:
                prev = fm.get(block_key, "")
                if prev:
                    fm[block_key] = prev + ", " + val
                else:
                    fm[block_key] = val
            continue
        if ":" not in stripped:
            in_block_list = False
            continue
        k, _, v = stripped.partition(":")
        k = k.strip()
        v = v.strip().strip('"').strip("'")
        fm[k] = v
        in_block_list = False
        # detect inline list like tags: [a, b, c]
        if v.startswith("[") and v.endswith("]"):
            items = [x.strip().strip('"').strip("'") for x in v[1:-1].split(",")]
            fm[k] = ", ".join(items)
        elif v == "":
            # might be block list next
            in_block_list = True
            block_key = k

    body = "\n".join(lines[end + 1:])
    return fm, body


def _has_tag(fm: Dict[str, str], tag: str) -> bool:
    """Check if frontmatter dict has a given tag in tags/area/methods fields."""
    for key in ("tags", "tag", "areas", "methods"):
        val = fm.get(key, "")
        if tag in [t.strip() for t in val.replace(",", " ").split()]:
            return True
    return False


# ---------------------------------------------------------------------------
# scanning vault for sigla notes
# ---------------------------------------------------------------------------

def _scan_vault_for_siglas() -> Dict[str, Tuple[Path, Dict[str, str]]]:
    """Find all vault notes with 'siglas' tag and sigla frontmatter.

    Returns {key: (path, frontmatter_dict)}.
    """
    result: Dict[str, Tuple[Path, Dict[str, str]]] = {}
    seen: Set[Path] = set()
    for d in VAULT_SCAN_DIRS:
        if not d.exists():
            continue
        for f in d.glob("*.md"):
            if f in seen:
                continue
            seen.add(f)
            text = f.read_text(encoding="utf-8")
            fm, _ = _parse_frontmatter(text)
            if not _has_tag(fm, "siglas"):
                continue
            if "sigla" not in fm:
                continue
            key = fm["sigla"]
            result[key] = (f, fm)
    return result


def _is_managed(path: Path) -> bool:
    """True if this note lives in vault/siglas/ (managed by script)."""
    return path.parent.name == "siglas" and path.parent.parent.name == "vault"


# ---------------------------------------------------------------------------
# scanning .tex
# ---------------------------------------------------------------------------

def scan_ac_commands() -> Dict[str, Counter]:
    usage: Dict[str, Counter] = {}
    for d in TEX_DIRS:
        if not d.exists():
            continue
        for f in sorted(d.glob("*.tex")):
            text = f.read_text(encoding="utf-8")
            for m in AC_CMD_RE.finditer(text):
                _, key = m.groups()
                usage.setdefault(key, Counter())[f.name] += 1
    return usage


def scan_paren_acronyms() -> Dict[str, Counter]:
    usage: Dict[str, Counter] = {}
    for d in TEX_DIRS:
        if not d.exists():
            continue
        for f in sorted(d.glob("*.tex")):
            text = f.read_text(encoding="utf-8")
            for m in PAREN_ACRO_RE.finditer(text):
                key = m.group(1)
                usage.setdefault(key, Counter())[f.name] += 1
    return usage


def parse_abreviaturas_tex(path: Path) -> Dict[str, str]:
    if not path.exists():
        return {}
    text = path.read_text(encoding="utf-8")
    return {m.group(1): m.group(2).strip() for m in ACRO_RE.finditer(text)}


# ---------------------------------------------------------------------------
# read / write sigla notes
# ---------------------------------------------------------------------------

def read_all_sigla_entries() -> Dict[str, dict]:
    """Read sigla entries from all siglas-tagged notes across the vault.

    Returns {key: {"definition": str, "include": str, "path": Path, "managed": bool, ...}}.
    """
    entries: Dict[str, dict] = {}
    for key, (path, fm) in _scan_vault_for_siglas().items():
        entries[key] = {
            "definition": fm.get("definicao", ""),
            "include": fm.get("incluir", "pendente"),
            "ocorrencias_ac": int(fm.get("ocorrencias_ac", 0)),
            "ocorrencias_texto": int(fm.get("ocorrencias_texto", 0)),
            "arquivos_ac": fm.get("arquivos_ac", ""),
            "path": path,
            "managed": _is_managed(path),
        }
    return entries


def write_sigla_note(siglas_dir: Path, key: str, definition: str, include: str,
                     ac_occ: Counter, paren_occ: Counter) -> None:
    """Create or update a managed sigla note in vault/siglas/."""
    siglas_dir.mkdir(parents=True, exist_ok=True)
    path = _note_path(siglas_dir, key)
    ac_total = sum(ac_occ.values())
    paren_total = sum(paren_occ.values())
    ac_files = ", ".join(sorted(ac_occ.keys()))
    body = (
        f"---\n"
        f'sigla: "{key}"\n'
        f'definicao: "{definition}"\n'
        f"incluir: {include}\n"
        f"ocorrencias_ac: {ac_total}\n"
        f"ocorrencias_texto: {paren_total}\n"
        f'arquivos_ac: "{ac_files}"\n'
        f"tags:\n"
        f"  - siglas\n"
        f"---\n"
        f"\n"
        f"# {key}\n"
        f"\n"
        f"**Definição:** {definition}\n"
        f"\n"
        f"## Ocorrências\n"
        f"\n"
        f"| Tipo | Contagem |\n"
        f"|---|---|\n"
        f"| `\\ac{{{key}}}` | {ac_total} |\n"
        f"| Texto corrido | {paren_total} |\n"
    )
    if ac_files:
        body += f"| Arquivos | {ac_files} |\n"
    body += "\n"
    path.write_text(body, encoding="utf-8")


# ---------------------------------------------------------------------------
# commands
# ---------------------------------------------------------------------------

def cmd_scan(args: argparse.Namespace) -> int:
    """Scan .tex files, cross-reference, update managed sigla notes."""
    ac_usage = scan_ac_commands()
    paren_usage = scan_paren_acronyms()
    tex_defs = parse_abreviaturas_tex(args.abrev_tex)
    existing = read_all_sigla_entries()

    all_keys: set[str] = set()
    all_keys.update(ac_usage.keys())
    all_keys.update(paren_usage.keys())
    all_keys.update(tex_defs.keys())
    all_keys.update(CANONICAL.keys())

    n_sim = n_nao = n_pend = 0
    n_managed = 0
    n_external = 0
    for key in sorted(all_keys):
        if key in existing:
            incl = existing[key]["include"]
            defn = existing[key]["definition"]
            if not existing[key]["managed"]:
                n_external += 1
                if incl == "sim":
                    n_sim += 1
                elif incl == "nao":
                    n_nao += 1
                else:
                    n_pend += 1
                continue  # don't write to external notes
        else:
            incl = "pendente"
            defn = tex_defs.get(key, CANONICAL.get(key, "---"))

        if incl == "sim":
            n_sim += 1
        elif incl == "nao":
            n_nao += 1
        else:
            n_pend += 1
        n_managed += 1
        write_sigla_note(args.siglas_dir, key, defn, incl,
                         ac_usage.get(key, Counter()),
                         paren_usage.get(key, Counter()))

    total = len(all_keys)
    print(f"[scan] {total} siglas catalogadas ({n_sim} incluir, {n_nao} excluir, {n_pend} pendente)")
    print(f"[scan] {n_managed} notas gerenciadas em vault/siglas/, {n_external} externas (projeto/areas/)")
    print(f"[scan] {sum(1 for k in ac_usage if k in all_keys)} usadas com \\ac{{}}")

    if n_pend:
        print(f"[scan] AVISO: {n_pend} siglas pendentes — edite as notas e marque incluir como sim/nao")
    return 0 if n_pend == 0 else 1


def cmd_generate(args: argparse.Namespace) -> int:
    """Generate Abreviaturas.tex from sigla notes across the vault."""
    entries = read_all_sigla_entries()
    included = [(k, v) for k, v in entries.items() if v["include"] == "sim"]
    if not included:
        print("[generate] ERRO: nenhuma sigla marcada como 'sim'")
        return 1

    sources = set()
    for k, v in entries.items():
        sources.add(str(v["path"].parent))

    abrev_path = args.abrev_tex
    abrev_path.parent.mkdir(parents=True, exist_ok=True)
    sources_str = ", ".join(sorted(sources))
    lines = [
        f"% Gerado por scripts/{SCRIPT_NAME} a partir de: {sources_str}",
        "% Edite as notas e re-execute; não edite este arquivo manualmente.",
        "\\begin{acronym}",
    ]
    for key, v in sorted(included, key=lambda x: x[0]):
        lines.append(f"\\acro{{{key}}}{{{v['definition']}}}")
    lines.append("\\end{acronym}")
    lines.append("")
    abrev_path.write_text("\n".join(lines), encoding="utf-8")
    print(f"[generate] {len(included)} siglas escritas em {abrev_path}")
    return 0


def cmd_validate(args: argparse.Namespace) -> int:
    """Validate that all \\ac{{}} commands have definitions."""
    ac_usage = scan_ac_commands()
    entries = read_all_sigla_entries()
    tex_defs = parse_abreviaturas_tex(args.abrev_tex)

    defined: set[str] = set()
    for key, v in entries.items():
        if v["include"] == "sim":
            defined.add(key)
    defined.update(tex_defs.keys())

    errors = 0
    for key, files in sorted(ac_usage.items()):
        if key not in defined:
            flist = ", ".join(sorted(files.keys()))
            print(f"[validate] ERRO: \\ac{{{key}}} usado em {flist} mas não definido")
            errors += 1

    unused = defined - set(ac_usage.keys()) - set(scan_paren_acronyms().keys())
    if unused:
        for k in sorted(unused):
            print(f"[validate] INFO: '{k}' definido mas não usado no texto")

    if errors:
        print(f"[validate] FAIL: {errors} erro(s)")
        return 1
    print("[validate] PASS: todos os \\ac{} têm definição")
    return 0


def main() -> int:
    parser = argparse.ArgumentParser(
        description="Gerenciar Lista de Siglas via notas com tag 'siglas' no vault."
    )
    parser.add_argument(
        "--siglas-dir",
        type=Path,
        default=DEFAULT_SIGLAS_DIR,
        help=f"Diretório das notas gerenciadas (default: {DEFAULT_SIGLAS_DIR})",
    )
    parser.add_argument(
        "--abrev-tex",
        type=Path,
        default=DEFAULT_ABREV_TEX,
        help=f"Caminho do Abreviaturas.tex (default: {DEFAULT_ABREV_TEX})",
    )

    sub = parser.add_subparsers(dest="subcommand", required=True)

    sp = sub.add_parser("scan", help="Examinar .tex, atualizar notas gerenciadas em vault/siglas/")
    sp.set_defaults(func=cmd_scan)

    gp = sub.add_parser("generate", help="Gerar Abreviaturas.tex de todas as notas com tag siglas")
    gp.set_defaults(func=cmd_generate)

    vp = sub.add_parser("validate", help="Validar consistência \\ac{} vs definições")
    vp.set_defaults(func=cmd_validate)

    args = parser.parse_args()
    return args.func(args)


if __name__ == "__main__":
    sys.exit(main())
