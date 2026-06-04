#!/usr/bin/env bash
# Deterministic validation gate for generated metrics and LaTeX references.
set -uo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
REPO_DIR="$(cd "$SCRIPT_DIR/.." && pwd)"
STATUS=0

run_step() {
  local label="$1"
  shift
  if "$@"; then
    echo "PASS $label"
  else
    echo "FAIL $label" >&2
    STATUS=1
  fi
}

run_step metrics python3 "$SCRIPT_DIR/gerar-metricas-monografia.py" --check

if python3 - "$REPO_DIR" <<'PY'
from __future__ import annotations

import json
import re
import sys
from pathlib import Path


root = Path(sys.argv[1])
monografia = root / "monografia"
bib_file = monografia / "bib" / "abntex2-references.bib"
ignored_dirs = {"ape_comandos"}
generated_dir = monografia / "generated"
image_suffixes = (".pdf", ".png", ".jpg", ".jpeg", ".svg")


def active_text(path: Path) -> str:
    lines = []
    for line in path.read_text(encoding="utf-8").splitlines():
        stripped = line.lstrip()
        if stripped.startswith("%"):
            continue
        lines.append(line.split("%", 1)[0])
    return "\n".join(lines)


def tracked_tex_files() -> list[Path]:
    files = []
    for path in sorted(monografia.rglob("*.tex")):
        rel_parts = path.relative_to(monografia).parts
        if any(part in ignored_dirs for part in rel_parts):
            continue
        files.append(path)
    return files


def source_tex_files() -> list[Path]:
    return [path for path in tracked_tex_files() if generated_dir not in path.parents]


def bib_keys() -> set[str]:
    if not bib_file.exists():
        return set()
    text = bib_file.read_text(encoding="utf-8")
    return set(re.findall(r"@\w+\s*\{\s*([^,\s]+)", text))


def resolve_tex_arg(arg: str) -> Path:
    path = monografia / arg
    if path.suffix:
        return path
    return path.with_suffix(".tex")


def resolve_graphic_arg(arg: str) -> Path | None:
    path = monografia / arg
    if path.suffix:
        return path if path.exists() else None
    for suffix in image_suffixes:
        candidate = path.with_suffix(suffix)
        if candidate.exists():
            return candidate
    return None


errors: list[str] = []
keys = bib_keys()
all_text = ""
source_text = ""
input_args: set[str] = set()
labels: set[str] = set()
refs: list[tuple[str, Path]] = []

for tex_file in tracked_tex_files():
    text = active_text(tex_file)
    all_text += "\n" + text
    if generated_dir not in tex_file.parents:
        source_text += "\n" + text

    for command, arg in re.findall(r"\\(input|include)\{([^{}]+)\}", text):
        if "#" in arg:
            continue
        input_args.add(arg)
        path = resolve_tex_arg(arg)
        if not path.exists():
            errors.append(f"missing \\{command}: {arg} referenced by {tex_file.relative_to(root)}")

    for arg in re.findall(r"\\includegraphics(?:\[[^\]]*\])?\{([^{}]+)\}", text):
        if "#" in arg:
            continue
        if resolve_graphic_arg(arg) is None:
            errors.append(f"missing graphic: {arg} referenced by {tex_file.relative_to(root)}")

    for raw_keys in re.findall(r"\\cite\w*\*?\{([^{}]+)\}", text):
        for key in [item.strip() for item in raw_keys.split(",") if item.strip()]:
            if key not in keys:
                errors.append(f"missing BibTeX key: {key} referenced by {tex_file.relative_to(root)}")

    labels.update(re.findall(r"\\label\{([^{}]+)\}", text))
    for raw_ref in re.findall(r"\\(?:ref|pageref|autoref|cref|Cref)\{([^{}]+)\}", text):
        for ref in [item.strip() for item in raw_ref.split(",") if item.strip()]:
            refs.append((ref, tex_file))

for ref, tex_file in refs:
    if ref not in labels:
        errors.append(f"missing label: {ref} referenced by {tex_file.relative_to(root)}")

if "\\input{generated/metrics.tex}" not in all_text:
    errors.append("generated/metrics.tex is not loaded by the monograph")

for rel in ["generated/metrics.tex", "generated/manifest.json"]:
    if not (monografia / rel).exists():
        errors.append(f"missing generated artifact: monografia/{rel}")

manifest_path = monografia / "generated" / "manifest.json"
if manifest_path.exists():
    manifest = json.loads(manifest_path.read_text(encoding="utf-8"))
    expected_inputs = {"summary_count": 4638, "timing_count": 4638, "evolution_count": 4638}
    for field, expected in expected_inputs.items():
        actual = manifest.get("inputs", {}).get(field)
        if actual != expected:
            errors.append(f"unexpected generated manifest {field}: {actual} != {expected}")
    expected_counts = {"aco": 1530, "ga": 1530, "pso": 1530, "lowerbound": 30, "bruteforce": 18}
    method_counts = manifest.get("data", {}).get("method_counts", {})
    for method, expected in expected_counts.items():
        if method_counts.get(method) != expected:
            errors.append(f"unexpected generated manifest method_counts.{method}: {method_counts.get(method)} != {expected}")
    seed_counts = manifest.get("data", {}).get("seed_counts", {})
    for method in ["aco", "ga", "pso"]:
        if seed_counts.get(method) != 51:
            errors.append(f"unexpected generated manifest seed_counts.{method}: {seed_counts.get(method)} != 51")
    for rel in manifest.get("outputs", []):
        if rel.startswith("tables/"):
            input_name = f"generated/{rel}"
            if input_name not in input_args:
                errors.append(f"generated table not included by monograph: {input_name}")

for literal in [
    "4638", "1530", "918", "0{,}4296", "5{,}2814", "22{,}2915",
    "71{,}79", "26{,}47", "4{,}25", "51{,}36", "40{,}79", "65{,}35",
    "293{,}2222", "4{,}710129", "1{,}1000", "1{,}9000", "3{,}0000",
    "0{,}6050", "4{,}26", "4262", "4276", "4266",
]:
    if literal in source_text:
        errors.append(f"hardcoded generated metric outside monografia/generated: {literal}")

if errors:
    for error in errors:
        print(f"ERROR {error}", file=sys.stderr)
    raise SystemExit(1)
PY
then
  echo "PASS monografia"
else
  echo "FAIL monografia" >&2
  STATUS=1
fi

run_step review-requests bash "$SCRIPT_DIR/check-reviews.sh" --fail-on-high

if [[ "$STATUS" -eq 0 ]]; then
  echo "PASS check-monografia"
else
  echo "FAIL check-monografia" >&2
fi

exit "$STATUS"
