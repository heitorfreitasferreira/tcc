#!/usr/bin/env python3
"""Migrate vault flat tags to structured tags. Phases 1-3 in one pass."""

import os
import re
import sys
import yaml
import shutil
from pathlib import Path
from collections import OrderedDict

VAULT = Path("/home/heitor/tcc/vault")

# ── Phase 1: mechanical structured tags based on path + properties ──────────

PATH_TYPE_MAP = {
    "papers":    "tipo/paper",
    "areas":     "tipo/area",
    "projeto":   "tipo/projeto",
    "siglas":    "tipo/sigla",
    "templates": "tipo/template",
}

PATH_EVIDENCIA_MAP = {
    "papers":          "evidencia/referencia",
    "projeto":         "evidencia/codigo",
    "writing/auditorias": "evidencia/auditoria",
}

WRITING_SUBDIR_TYPE = {
    "auditorias":          "tipo/auditoria",
    "capitulos":           "tipo/writing",
    "decisoes":            "tipo/writing",
    "planejamento":        "tipo/writing",
    "review-solicitacoes": "tipo/revisao",
}

STATUS_MAP = {
    "lido":                  "status/lido",
    "resumo-lido":           "status/resumo-lido",
    "lido-parcial":          "status/lido-parcial",
    "pendente":              "status/pendente",
    "disponivel":            "status/pendente",
    "disponível":            "status/pendente",
    "descartado-escopo-imediato": "status/descartado",
    "removido":              "status/removido",
    "analise-posterior":     "status/analise-posterior",
    "validado":              "status/validado",
    "estatistica-validada":  "status/validado",
    "concluido":             "status/concluido",
    "concluido-com-achados": "status/concluido",
    "draft":                 "status/rascunho",
    "ativo":                 "status/ativo",
    "aberto":                "status/aberto",
    "resolvido":             "status/resolvido",
    "pronto-para-revisao":   "status/pronto-revisao",
    "pronto-com-pendencias-bibliograficas-rastreabilidade-e-validacao": "status/pronto-com-pendencias",
    "regenerado-pos-p4":     "status/atualizado",
    "auditado-com-bloqueios": "status/auditado-bloqueios",
    "verificado-externo-encontrou-precedentes": "status/verificado",
}

# ── Phase 2: flat tag → structured tag mapping ──────────────────────────────

TAG_MAP = {
    # Methods
    "aco":                   "metodo/aco",
    "ga":                    "metodo/ga",
    "pso":                   "metodo/pso",
    "sa":                    "metodo/sa",
    "abc":                   "metodo/abc",
    "ts":                    "metodo/ts",
    "gwo":                   "metodo/gwo",
    "cso":                   "metodo/cso",
    "ssa":                   "metodo/ssa",
    "eho":                   "metodo/eho",
    "gp":                    "metodo/gp",
    "algoritmo-genetico":    "metodo/ga",
    "exaustivo":             "metodo/exact",
    "brute-force":           "metodo/exact",
    "metaheuristic":         "metodo/metaheuristic",
    "metaheuristica":        "metodo/metaheuristic",
    "heuristic":             "metodo/heuristic",
    "exact":                 "metodo/exact",
    "exact-methods":         "metodo/exact",
    "hybrid":                "metodo/hybrid",
    "approximation":         "metodo/approximation",
    "branch-and-bound":      "metodo/branch-and-bound",
    "branch-and-cut":        "metodo/branch-and-cut",
    "milp":                  "metodo/milp",
    "held-karp":             "metodo/held-karp",
    "lagrangean":            "metodo/lagrangean",
    "additive-bounding":     "metodo/lower-bound",
    "deep-learning":         "metodo/machine-learning",
    "machine-learning":      "metodo/machine-learning",
    "neural-combinatorial-optimization": "metodo/machine-learning",
    "constraint-programming":"metodo/constraint-programming",
    "random-keys":           "metodo/encoding",
    "encoding":              "metodo/encoding",
    "crossover":             "metodo/crossover",
    "eax":                   "metodo/crossover",
    "edge-assembly":         "metodo/crossover",
    "patching":              "metodo/heuristic",
    "clustering":            "metodo/clustering",
    "decision-diagram":      "metodo/exact",
    "labeling":              "metodo/lower-bound",
    "ng-path":               "metodo/lower-bound",
    "shortest-path":         "metodo/heuristic",
    "graph-algorithm":       "metodo/heuristic",
    "ant-colony":            "area/ant-colony",
    "genetic-algorithms":    "area/genetic-algorithms",
    "particle-swarm":        "area/particle-swarm",

    # Areas / domains
    "tsp":                   "area/tsp",
    "atsp":                  "area/atsp",
    "tdtsp":                 "area/tdtsp",
    "fstsp":                 "area/fstsp",
    "gtsp":                  "area/gtsp",
    "tsp-variants":          "area/tsp-variants",
    "variante-tsp":          "area/tsp-variants",
    "problema-classico":     "area/tsp",
    "drone":                 "area/drone-routing",
    "drone-routing":         "area/drone-routing",
    "routing":               "area/routing",
    "vrp":                   "area/vrp",
    "uav":                   "area/drone-routing",
    "vant":                  "area/drone-routing",
    "lower-bound":           "metodo/lower-bound",
    "lower-bounds":          "area/lower-bounds",
    "bio-inspirado":         "area/bio-inspired-optimization",
    "bio-inspired-optimization": "area/bio-inspired-optimization",
    "comparison":            "area/comparative-studies",
    "comparative-studies":   "area/comparative-studies",
    "angular-cost":          "area/tsp-variants",
    "turn-cost":             "area/tsp-variants",
    "turn-penalty":          "area/tsp-variants",
    "task-allocation":       "area/drone-routing",
    "path-planning":         "area/routing",
    "iot":                   "area/drone-routing",
    "network-coding":        "area/drone-routing",
    "scheduling":            "area/routing",
    "warehousing":           "area/routing",
    "uav":                   "area/drone-routing",
    "combinatorial-optimization": "area/tsp",

    # Roles / paper types
    "survey":                "papel/revisao",
    "review":                "papel/revisao",
    "foundational":          "papel/fundacional",
    "benchmark":             "papel/benchmark",
    "experimental":          "papel/benchmark",
    "comparison":            "papel/comparativo",
    "tradeoff":              "papel/comparativo",
    "feromonio-3d":          "metodo/aco",
    "multi-objetivo":        "topico/multi-objetivo",
    "originalidade":         "topico/originalidade",
    "verificacao":           "evidencia/validacao",

    # Tools / infrastructure
    "statistics":            "evidencia/estatistica",
    "friedman":              "evidencia/estatistica",
    "nemenyi":               "evidencia/estatistica",
    "methodology":           "evidencia/metodologia",
    "metodologia":           "evidencia/metodologia",
    "formulacao":            "topico/formulacao",
    "tensor":                "topico/formulacao",
    "analise":               "evidencia/estatistica",
    "estatistica":           "evidencia/estatistica",
    "go":                    "topico/ferramenta",
    "pacotes":               "topico/arquitetura",
    "cli":                   "topico/ferramenta",
    "pipeline":              "topico/ferramenta",
    "experimentos":          "topico/experimentos",
    "resultados":            "topico/experimentos",
    "implementacao":         "topico/implementacao",
    "arquitetura":           "topico/arquitetura",
    "visao-geral":           "topico/visao-geral",
    "latex":                 "topico/formatacao",
    "abnt":                  "topico/formatacao",
    "facom":                 "topico/formatacao",
    "modelo":                "topico/formatacao",
    "modelo-facom":          "topico/formatacao",
    "codigo":                "evidencia/codigo",
    "dados":                 "evidencia/dados",
    "bibliografia":          "evidencia/referencia",
    "pdf":                   "evidencia/referencia",
    "validacao":             "evidencia/validacao",

    # Writing / project organization
    "monografia":            "topico/monografia",
    "capitulo":              "topico/monografia",
    "capitulos":             "topico/monografia",
    "introducao":            "capitulo/introducao",
    "fundamentacao":         "capitulo/fundamentacao",
    "proposta":              "capitulo/proposta",
    "experimentos":          "capitulo/experimentos",
    "conclusao":             "capitulo/conclusao",
    "glossario":             "topico/glossario",
    "terminologia":          "topico/glossario",
    "figuras":               "topico/figuras",
    "tabelas":               "topico/tabelas",
    "citacoes":              "topico/citacoes",
    "roadmap":               "topico/roadmap",
    "revisao":               "topico/revisao",
    "orientador":            "topico/revisao",
    "banca":                 "topico/revisao",
    "grill":                 "topico/revisao",
    "entrevista":            "topico/revisao",
    "agentes":               "topico/ferramenta",
    "index":                 "tipo/index",
    "catalogo":              "tipo/index",
    "tracker":               "tipo/index",
    "siglas":                "tipo/sigla",
    "projeto":               "tipo/projeto",
    "area":                  "tipo/area",
    "writing":               "tipo/writing",
    "auditoria":             "tipo/auditoria",
    "classificacao":         "topico/classificacao",
    "aplicacao":             "area/drone-routing",
    "otimizacao":            "area/bio-inspired-optimization",
    "complexity":            "area/tsp",
    "exemplo":               "_descarte",
    "neurips":               "_descarte",
    "papers":                "_descarte",
    "claims":                "_descarte",
    "evidencias":            "_descarte",
    "book":                  "papel/livro",
    "theory":                "papel/teorico",
    "algorithm":             "_descarte",
    "swarm":                 "metodo/metaheuristic",
    "assignment":            "metodo/exact",
    "complexity":            "area/tsp",
    "combinatorial-optimization": "area/tsp",

    # Acronym-related (skip, these become tipo/sigla)
    "abnt":                  "topico/formatacao",
    "aco":                   "metodo/aco",
    "ap":                    "_descarte",
    "as":                    "_descarte",
    "cd":                    "_descarte",
    "gp":                    "metodo/gp",
    "gnn":                   "metodo/machine-learning",
    "lb":                    "metodo/lower-bound",
    "lkh":                   "metodo/heuristic",
    "mmas":                  "metodo/aco",
    "nbr":                   "_descarte",
    "np":                    "_descarte",
    "ox":                    "metodo/crossover",
    "poi":                   "_descarte",
    "ppo":                   "_descarte",
    "rl":                    "metodo/machine-learning",
    "tsplib":                "_descarte",
    "vns":                   "metodo/vns",
}

# Tags that should be dropped (internal noise, not useful as structured tags)
DROP_TAGS = {"_descarte"}

# ── Phase 3: property enrichment ────────────────────────────────────────────

def extract_property_values(structured_tags):
    """Given a list of structured tags, extract areas, methods, chapters, etc."""
    areas = set()
    methods = set()
    chapters = set()

    for tag in structured_tags:
        if tag.startswith("area/") and tag != "area/tsp":
            name = tag.split("/", 1)[1]
            areas.add(name)
        elif tag.startswith("metodo/"):
            name = tag.split("/", 1)[1]
            methods.add(name)
        elif tag.startswith("capitulo/"):
            name = tag.split("/", 1)[1]
            chapters.add(name)

    # Always add tsp if area/tsp tag exists (core domain)
    return {
        "areas": sorted(areas),
        "methods": sorted(methods),
        "chapters": sorted(chapters),
    }


def get_role_from_tags(structured_tags):
    """Derive role from structured tags."""
    papel_tags = {t for t in structured_tags if t.startswith("papel/")}
    if "papel/fundacional" in papel_tags:
        return "fundacional"
    if "papel/benchmark" in papel_tags:
        return "benchmark"
    if "papel/comparativo" in papel_tags:
        return "comparativo"
    if "papel/revisao" in papel_tags:
        return "revisao"
    return "revisao"  # safe default for survey/review papers


# ── Frontmatter parsing ─────────────────────────────────────────────────────

def parse_frontmatter(text):
    """Parse YAML frontmatter from markdown text. Returns (fm_dict, body_text, raw_fm)."""
    if not text.startswith("---\n"):
        return None, text, None

    # Find closing ---
    rest = text[4:]
    end = rest.find("\n---\n")
    if end == -1:
        end = rest.find("\n---")
        if end == -1:
            # Check if frontmatter goes to EOF (no closing ---)
            end2 = rest.rfind("\n---")
            if end2 == -1 and rest.strip():
                # Entire file after opening --- is frontmatter (no body)
                try:
                    fm = yaml.safe_load(rest) or {}
                except yaml.YAMLError:
                    return None, text, None
                return fm, "", rest
            return None, text, None

    raw_fm = rest[:end]
    # Determine body offset: \n---\n (5 chars) or \n--- (4 chars)
    body_offset = end + 5 if rest[end:end+5] == "\n---\n" else end + 4
    body = rest[body_offset:]

    try:
        fm = yaml.safe_load(raw_fm) or {}
    except yaml.YAMLError:
        return None, text, None

    return fm, body, raw_fm


def write_frontmatter(fm, body, filepath):
    """Write markdown file with YAML frontmatter."""
    # Use yaml.dump with flow_style=None for clean output
    dumped = yaml.dump(fm, default_flow_style=False, allow_unicode=True,
                       sort_keys=False, width=120)
    content = f"---\n{dumped}---\n{body}"
    with open(filepath, "w", encoding="utf-8") as f:
        f.write(content)

    # Also create backup
    bak = str(filepath) + ".bak"
    if not os.path.exists(bak):
        shutil.copy2(filepath, bak)


def normalize_tags(fm):
    """Ensure tags field is a list. Handle both inline [a,b] and multiline formats."""
    tags = fm.get("tags", [])
    if tags is None:
        return []
    if isinstance(tags, str):
        # Space-separated string
        return [t.strip() for t in tags.split() if t.strip()]
    if isinstance(tags, list):
        return [t.strip() if isinstance(t, str) else str(t) for t in tags]
    return []


def split_multiword_tags(tags):
    """Split multi-word flat tags into individual words.
    E.g., 'aco metaheuristic' -> ['aco', 'metaheuristic']
    But preserve hyphenated terms like 'lower-bound', 'machine-learning'."""
    result = []
    for tag in tags:
        # Skip tags that look like structured (contain /)
        if "/" in tag:
            result.append(tag)
            continue
        # Split on whitespace
        parts = tag.split()
        if len(parts) == 1:
            result.append(tag)
        else:
            result.extend(parts)
    return result


def relative_path(filepath):
    """Get path relative to vault root."""
    return os.path.relpath(filepath, VAULT)


def is_writing_subdir(relpath, subdir):
    """Check if file is in a specific writing subdirectory."""
    return relpath.startswith(f"writing/{subdir}/")


def migrate_file(filepath, phase="all"):
    """Apply all 3 phases to a single file. Returns (changes_made, log_messages)."""
    rel = relative_path(filepath)
    log = []

    with open(filepath, "r", encoding="utf-8") as f:
        text = f.read()

    fm, body, raw_fm = parse_frontmatter(text)
    if fm is None or not fm:
        return False, [f"{rel}: SKIP (no frontmatter)"]

    old_tags = normalize_tags(fm)
    new_tags = list(old_tags)  # start with existing
    changes = False

    # ── Phase 1: directory-based structured tags ──────────────────────────

    # Determine base directory
    top_dir = rel.split("/")[0] if "/" in rel else "root"

    # Don't apply tipo/paper to papers/index.md (it's a catalog, not a paper)
    is_index = rel.endswith("/index.md") or rel == "index.md"

    if top_dir in PATH_TYPE_MAP and not (top_dir == "papers" and is_index):
        tag = PATH_TYPE_MAP[top_dir]
        if tag not in new_tags:
            new_tags.append(tag)
            log.append(f"  +{tag} (path)")
            changes = True

    # Writing subdirectory overrides
    for subdir, tag in WRITING_SUBDIR_TYPE.items():
        if is_writing_subdir(rel, subdir):
            if tag not in new_tags:
                new_tags.append(tag)
                log.append(f"  +{tag} (writing subdir)")
                changes = True

    # Evidence tags
    for path_prefix, eviden_tag in PATH_EVIDENCIA_MAP.items():
        if rel.startswith(path_prefix) and not (path_prefix == "papers" and is_index):
            if eviden_tag not in new_tags:
                new_tags.append(eviden_tag)
                log.append(f"  +{eviden_tag} (path)")
                changes = True

    # Status mapping (use prefix match for atualizado* patterns)
    raw_status = fm.get("status") or fm.get("reading_status")
    if raw_status:
        matched = raw_status in STATUS_MAP
        if not matched:
            # Prefix matching for statuses like atualizado-pos-p4, atualizado-pos-p7, atualizado-2026-06-02
            for prefix, stag in [("atualizado", "status/atualizado"),
                                 ("pronto-com-pendencias", "status/pronto-com-pendencias")]:
                if raw_status.startswith(prefix):
                    if stag not in new_tags:
                        new_tags.append(stag)
                        log.append(f"  +{stag} (status={raw_status})")
                        changes = True
                    matched = True
                    break
        if matched and raw_status in STATUS_MAP:
            stag = STATUS_MAP[raw_status]
            if stag not in new_tags:
                new_tags.append(stag)
                log.append(f"  +{stag} (status={raw_status})")
                changes = True

    # Incluir mapping for siglas/projeto
    incluir = fm.get("incluir")
    if incluir in ("sim", "pendente", "nao"):
        tag = f"incluir/{incluir}"
        if tag not in new_tags:
            new_tags.append(tag)
            log.append(f"  +{tag} (incluir)")
            changes = True

    # ── Phase 2: flat tag → structured tag conversion ────────────────────

    current_flat = [t for t in new_tags if "/" not in t]
    current_structured = {t for t in new_tags if "/" in t}

    # Split multi-word flat tags first
    expanded_flat = split_multiword_tags(current_flat)

    # Track which flat tags were successfully mapped (for Phase 2.5 cleanup)
    mapped_flats = set()

    # Don't map the flat tag 'siglas' to tipo/sigla in non-siglas files
    # (areas/tsp.md had flat tag 'siglas' which shouldn't become tipo/sigla)
    is_siglas_dir = top_dir == "siglas"

    for flat in expanded_flat:
        mapped = TAG_MAP.get(flat.lower())
        if mapped is not None:
            if mapped in DROP_TAGS:
                # Tag explicitly marked for removal - drop the flat tag
                mapped_flats.add(flat)
                continue
            # Skip tipo/sigla for non-siglas files that happen to have 'siglas' flat tag
            if mapped == "tipo/sigla" and not is_siglas_dir:
                # Still remove the flat tag (it's noise in non-siglas files)
                mapped_flats.add(flat)
                continue
            # Always mark as mapped (even if structured tag already exists from prior run)
            mapped_flats.add(flat)
            if mapped not in current_structured:
                current_structured.add(mapped)
                log.append(f"  +{mapped} (from flat '{flat}')")
                changes = True

    # ── Phase 2.5: remove flat tags that were successfully mapped ────────
    # Keep flat tags that have no structured equivalent (unmapped)
    unmapped_flats = [f for f in expanded_flat if f not in mapped_flats]

    # Keep structured tags only (drop flat tags that are now represented)
    new_tags = sorted(current_structured) + sorted(unmapped_flats)

    # ── Phase 3: property enrichment for papers ──────────────────────────

    if top_dir == "papers" and not is_index:
        # Add type if missing
        if not fm.get("type"):
            fm["type"] = "paper"
            log.append(f"  +type: paper")
            changes = True

        # Derive areas, methods, chapters from structured tags
        props = extract_property_values(new_tags)
        for key in ("areas", "methods", "chapters"):
            current_val = fm.get(key, [])
            if isinstance(current_val, str):
                current_val = [current_val]
            if not isinstance(current_val, list):
                current_val = []
            derived = props[key]
            if derived and set(derived) - set(current_val):
                merged = sorted(set(current_val) | set(derived))
                fm[key] = merged
                log.append(f"  +{key}: {merged}")
                changes = True

        # Role
        if not fm.get("role"):
            fm["role"] = get_role_from_tags(new_tags)
            log.append(f"  +role: {fm['role']}")
            changes = True

        # reading_status from status
        if not fm.get("reading_status") and fm.get("status"):
            s = fm["status"]
            if s in ("lido", "resumo-lido", "lido-parcial", "pendente"):
                fm["reading_status"] = s
                changes = True
            elif s in ("disponivel", "disponível"):
                fm["reading_status"] = "pendente"
                changes = True

        # validation_status default
        if not fm.get("validation_status"):
            fm["validation_status"] = "nao-validado"
            changes = True

        # pdf_status from classificacao
        if not fm.get("pdf_status"):
            classificacao = fm.get("classificacao", "")
            if classificacao == "recuperado":
                fm["pdf_status"] = "disponivel"
                changes = True

    # ── Enrichment for other types ────────────────────────────────────────

    if top_dir == "areas":
        if not fm.get("type"):
            fm["type"] = "area"
            changes = True

    if top_dir == "projeto":
        if not fm.get("type"):
            fm["type"] = "projeto"
            changes = True

    if top_dir == "writing":
        if not fm.get("type"):
            fm["type"] = "writing"
            changes = True

    if top_dir == "templates":
        if not fm.get("type"):
            fm["type"] = "template"
            changes = True

    # ── Write back ────────────────────────────────────────────────────────

    # Deduplicate while preserving order
    seen = set()
    deduped = []
    for t in new_tags:
        if t not in seen and t not in DROP_TAGS:
            deduped.append(t)
            seen.add(t)

    # Check if tags actually changed (including flat tag removal)
    old_deduped = []
    seen2 = set()
    for t in old_tags:
        if t not in seen2 and t not in DROP_TAGS:
            old_deduped.append(t)
            seen2.add(t)

    if set(deduped) != set(old_deduped):
        changes = True
        if not log:
            log.append(f"  (cleanup: removed flat tags, kept structured)")
        fm["tags"] = deduped
        write_frontmatter(fm, body, filepath)
        return True, log

    if changes:
        fm["tags"] = deduped
        write_frontmatter(fm, body, filepath)
        return True, log

    return False, [f"{rel}: no changes"]


def main():
    vault_dir = VAULT
    total = 0
    changed = 0
    skipped = 0
    all_logs = []

    for root, dirs, files in os.walk(vault_dir):
        # Skip .obsidian, bases, canvas directories
        dirs[:] = [d for d in dirs if not d.startswith(".") and d not in ("bases", "canvas")]
        for fname in sorted(files):
            if not fname.endswith(".md"):
                continue
            filepath = os.path.join(root, fname)
            total += 1
            did_change, logs = migrate_file(filepath)
            all_logs.extend(logs)
            if did_change:
                changed += 1
            else:
                skipped += 1

    print(f"Total files: {total}")
    print(f"Changed:    {changed}")
    print(f"Unchanged:  {skipped}")
    print()
    for line in all_logs:
        print(line)

    # ── Summary ──────────────────────────────────────────────────────────
    print(f"\n── Backups created as .bak files in vault/")
    print(f"── Run again to verify idempotency (should show 0 changes)")


if __name__ == "__main__":
    main()
