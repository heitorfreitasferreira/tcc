#!/usr/bin/env python3
"""Generate publication-ready scientific figures for the TCC.

Primary standard: the local `scientific-figures` skill. The script keeps a
Tufte-inspired style: explicit claims, high data-ink ratio, direct comparisons,
visible distributions for stochastic methods, shared method colors, SVG as the
editable primary artifact, and PNG previews.

Route images are fetched from the Go web renderer (`/api/render`) instead of
duplicating route drawing logic in Python. Route composition and explanatory
text are emitted as LaTeX so the monograph controls typography and layout.
"""

from __future__ import annotations

import json
import math
import os
import re
import shutil
import statistics
import subprocess
import urllib.parse
import urllib.request
from collections import defaultdict
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
DATA = ROOT / "src" / "data"
RESULTS = DATA / "results"
SUMMARY = RESULTS / "summary"
EVOLUTION = RESULTS / "evolution"
FIGS = ROOT / "monografia" / "figs"

RENDER_BASE = os.environ.get("FIGURE_RENDER_BASE", "").rstrip("/")

WIDTH = 2250
HEIGHT = 1450
FONT = "Arial, Helvetica, sans-serif"
METHODS = ["aco", "ga", "pso", "lowerbound"]
LABELS = {"aco": "ACO", "ga": "GA", "pso": "PSO", "lowerbound": "LB", "bruteforce": "BF"}
COLORS = {
    "aco": "#54a24b",
    "ga": "#4c78a8",
    "pso": "#f58518",
    "lowerbound": "#7b2d8e",
    "bruteforce": "#4c4c4c",
}
MARKERS = {"aco": "circle", "ga": "square", "pso": "triangle", "lowerbound": "diamond"}
TEXT = "#222222"
MUTED = "#666666"
GRID = "#dddddd"
AXIS = "#333333"


def esc(value: object) -> str:
    return (
        str(value)
        .replace("&", "&amp;")
        .replace("<", "&lt;")
        .replace(">", "&gt;")
        .replace('"', "&quot;")
    )


def write_svg_png(name: str, body: list[str], width: int = WIDTH, height: int = HEIGHT) -> None:
    FIGS.mkdir(parents=True, exist_ok=True)
    svg_path = FIGS / f"{name}.svg"
    png_path = FIGS / f"{name}.png"
    height_in = 7.5 * height / width
    svg = "\n".join(
        [
            f'<svg xmlns="http://www.w3.org/2000/svg" width="7.5in" height="{height_in:.3f}in" viewBox="0 0 {width} {height}">',
            f'<rect width="100%" height="100%" fill="#ffffff"/>',
            f'<style>text{{font-family:{FONT};fill:{TEXT}}}.suptitle{{font-size:44px;font-weight:700}}.panel{{font-size:42px;font-weight:700}}.title{{font-size:27px;font-weight:700}}.label{{font-size:23px}}.tick{{font-size:19px;fill:{MUTED}}}.small{{font-size:17px;fill:{MUTED}}}.axis{{stroke:{AXIS};stroke-width:2}}.grid{{stroke:{GRID};stroke-width:1;stroke-dasharray:2 4}}.thin{{stroke-width:1.6}}</style>',
            *body,
            "</svg>",
        ]
    )
    svg_path.write_text(svg, encoding="utf-8")
    converter = shutil.which("rsvg-convert")
    if converter:
        subprocess.run([converter, "-o", str(png_path), str(svg_path)], check=True)


def text(body: list[str], value: str, x: float, y: float, cls: str = "label", anchor: str = "start") -> None:
    body.append(f'<text class="{cls}" x="{x:.1f}" y="{y:.1f}" text-anchor="{anchor}">{esc(value)}</text>')


def line(body: list[str], x1: float, y1: float, x2: float, y2: float, color: str = AXIS, width: float = 2, extra: str = "") -> None:
    body.append(f'<line x1="{x1:.1f}" y1="{y1:.1f}" x2="{x2:.1f}" y2="{y2:.1f}" stroke="{color}" stroke-width="{width}" {extra}/>')


def rect(body: list[str], x: float, y: float, w: float, h: float, fill: str, stroke: str = "none", width: float = 1) -> None:
    body.append(f'<rect x="{x:.1f}" y="{y:.1f}" width="{w:.1f}" height="{h:.1f}" fill="{fill}" stroke="{stroke}" stroke-width="{width}"/>')


def marker(body: list[str], x: float, y: float, method: str, size: float = 9, opacity: float = 1.0) -> None:
    color = COLORS[method]
    shape = MARKERS.get(method, "circle")
    if shape == "circle":
        body.append(f'<circle cx="{x:.1f}" cy="{y:.1f}" r="{size:.1f}" fill="{color}" opacity="{opacity}"/>')
    elif shape == "square":
        rect(body, x - size, y - size, size * 2, size * 2, color)
    else:
        pts = f"{x:.1f},{y-size:.1f} {x-size:.1f},{y+size:.1f} {x+size:.1f},{y+size:.1f}"
        body.append(f'<polygon points="{pts}" fill="{color}" opacity="{opacity}"/>')


def panel_label(body: list[str], label: str, x: float, y: float) -> None:
    text(body, label, x, y, "panel")


def legend_direct(body: list[str], x: float, y: float) -> None:
    for i, method in enumerate(METHODS):
        yy = y + i * 34
        marker(body, x, yy - 5, method, 8)
        text(body, LABELS[method], x + 24, yy, "label")


def instance_key(instance: str) -> tuple[int, str]:
    match = re.match(r"^(\d+)(.*)$", instance)
    return (int(match.group(1)), match.group(2)) if match else (999999, instance)


def load_summary() -> list[dict]:
    rows: list[dict] = []
    for path in sorted(SUMMARY.glob("*.json")):
        data = json.loads(path.read_text(encoding="utf-8"))
        run_id = data.get("run_id") or path.stem
        parts = run_id.split("__")
        instance = parts[0]
        method = data.get("method") or (parts[1] if len(parts) > 1 else "")
        timing_file = data.get("timing_file") or f"timing/{run_id}.json"
        timing_path = RESULTS / timing_file
        optimize_ms = None
        total_ms = None
        if timing_path.exists():
            timing = json.loads(timing_path.read_text(encoding="utf-8"))
            durations = timing.get("durations_ms", {})
            optimize_ms = durations.get("optimize")
            total_ms = durations.get("total")
        rows.append(
            {
                "run_id": run_id,
                "instance": instance,
                "n": instance_key(instance)[0],
                "method": method,
                "seed": int(data.get("seed", 0)),
                "makespan": float(data.get("result", {}).get("best_makespan", 0.0)),
                "optimize_ms": float(optimize_ms) if optimize_ms is not None else None,
                "total_ms": float(total_ms) if total_ms is not None else None,
            }
        )
    return rows


def grouped(rows: list[dict], key: str) -> dict[tuple[str, str], list[float]]:
    out: dict[tuple[str, str], list[float]] = defaultdict(list)
    for row in rows:
        value = row.get(key)
        if value is not None:
            out[(row["instance"], row["method"])].append(float(value))
    return out


def median(values: list[float]) -> float:
    return statistics.median(values) if values else 0.0


def quantile(values: list[float], q: float) -> float:
    values = sorted(values)
    if not values:
        return 0.0
    pos = (len(values) - 1) * q
    lo, hi = math.floor(pos), math.ceil(pos)
    if lo == hi:
        return values[lo]
    return values[lo] * (hi - pos) + values[hi] * (pos - lo)


def rank_instances(rows: list[dict]) -> list[str]:
    return sorted({r["instance"] for r in rows if r["method"] in METHODS}, key=instance_key)


def run_for(rows: list[dict], instance: str, method: str, seed: int = 0) -> dict | None:
    matches = [r for r in rows if r["instance"] == instance and r["method"] == method and r["seed"] == seed]
    return sorted(matches, key=lambda r: r["run_id"])[0] if matches else None


def load_frames(run_id: str) -> list[dict]:
    path = EVOLUTION / f"{run_id}.jsonl"
    if not path.exists():
        return []
    return [json.loads(line) for line in path.read_text(encoding="utf-8").splitlines() if line.strip()]


def axes(body: list[str], x: float, y: float, w: float, h: float, ymin: float, ymax: float, yticks: int = 5, fmt: str = "{:.0f}") -> None:
    line(body, x, y + h, x + w, y + h)
    line(body, x, y, x, y + h)
    for i in range(yticks + 1):
        yy = y + h - h * i / yticks
        line(body, x, yy, x + w, yy, GRID, 1)
        value = ymin + (ymax - ymin) * i / yticks
        text(body, fmt.format(value), x - 10, yy + 6, "tick", "end")


def figure_quality_distribution(rows: list[dict]) -> None:
    claim = "o ACO obtém menor makespan que GA e PSO no conjunto completo de instâncias, com variação explícita entre sementes."
    body: list[str] = []
    text(body, "Figura 1. Qualidade das soluções por instância e semente", 70, 64, "suptitle")
    text(body, f"Afirmação: {claim}", 70, 104, "small")
    panel_label(body, "A", 70, 165)
    panel_label(body, "B", 1150, 165)

    instances = rank_instances(rows)
    values = grouped(rows, "makespan")
    x, y, w, h = 170, 190, 870, 1050
    all_medians = [median(values[(inst, method)]) for inst in instances for method in METHODS]
    xmin, xmax = min(all_medians) * 0.95, max(all_medians) * 1.04

    # Light reference axis: common linear makespan scale.
    for tick in [10, 25, 50, 75, 100, 125, 150]:
        if xmin <= tick <= xmax:
            xx = x + (tick - xmin) / (xmax - xmin) * w
            line(body, xx, y, xx, y + h, GRID, 1)
            text(body, str(tick), xx, y + h + 28, "tick", "middle")
    line(body, x, y + h, x + w, y + h)
    text(body, "makespan final mediano (menor é melhor)", x + w / 2, y + h + 68, "label", "middle")
    text(body, "Mediana de 51 sementes; mesma escala x em todas as instâncias.", x, y - 24, "title")
    for i, inst in enumerate(instances):
        yy = y + i * (h / (len(instances) - 1))
        text(body, inst, x - 16, yy + 6, "tick", "end")
        for method in METHODS:
            vals = values[(inst, method)]
            q1, q3 = quantile(vals, 0.25), quantile(vals, 0.75)
            med = median(vals)
            x1 = x + (q1 - xmin) / (xmax - xmin) * w
            x2 = x + (q3 - xmin) / (xmax - xmin) * w
            xm = x + (med - xmin) / (xmax - xmin) * w
            line(body, x1, yy, x2, yy, COLORS[method], 4)
            marker(body, xm, yy, method, 7)
    legend_direct(body, 900, 220)

    # Panel B: exact optimum success rate for BF-covered instances.
    opt = {r["instance"]: r["makespan"] for r in rows if r["method"] == "bruteforce"}
    small = sorted(opt, key=instance_key)
    x2, y2, w2, h2 = 1280, 220, 770, 900
    for tick in [0, 25, 50, 75, 100]:
        xx = x2 + tick / 100 * w2
        line(body, xx, y2, xx, y2 + h2, GRID, 1)
        text(body, f"{tick}", xx, y2 + h2 + 28, "tick", "middle")
    line(body, x2, y2 + h2, x2 + w2, y2 + h2)
    text(body, "taxa de acerto do ótimo (% de 51 sementes)", x2 + w2 / 2, y2 + h2 + 68, "label", "middle")
    text(body, "Instâncias pequenas com ótimo por busca exaustiva.", x2, y2 - 24, "title")
    bar_gap = h2 / len(small)
    for i, inst in enumerate(small):
        base_y = y2 + i * bar_gap + bar_gap / 2
        text(body, inst, x2 - 16, base_y + 6, "tick", "end")
        for j, method in enumerate(METHODS):
            vals = [r["makespan"] for r in rows if r["instance"] == inst and r["method"] == method]
            rate = sum(1 for v in vals if abs(v - opt[inst]) < 1e-3) / len(vals) * 100
            yy = base_y + (j - 1) * 8
            line(body, x2, yy, x2 + rate / 100 * w2, yy, COLORS[method], 5)
            marker(body, x2 + rate / 100 * w2, yy, method, 5)
    legend_direct(body, 1960, 220)
    write_svg_png("fig-quality-distribution", body)


def figure_runtime_convergence(rows: list[dict]) -> None:
    claim = "o ACO troca maior qualidade de solução por tempo de otimização substancialmente maior."
    body: list[str] = []
    text(body, "Figura 2. Trade-off entre tempo, qualidade e convergência", 70, 64, "suptitle")
    text(body, f"Afirmação: {claim}", 70, 104, "small")
    panel_label(body, "A", 70, 165)
    panel_label(body, "B", 70, 830)
    panel_label(body, "C", 1180, 165)

    data = [r for r in rows if r["method"] in METHODS and r["optimize_ms"] is not None]
    x, y, w, h = 170, 220, 860, 430
    xs = [max(r["optimize_ms"], 0.001) for r in data]
    ys = [r["makespan"] for r in data]
    xmin, xmax = min(xs), max(xs)
    ymin, ymax = min(ys), max(ys)
    axes(body, x, y, w, h, ymin, ymax, 4, "{:.0f}")
    for tick in [1, 10, 100, 1000]:
        if xmin <= tick <= xmax:
            xx = x + (math.log(tick) - math.log(xmin)) / (math.log(xmax) - math.log(xmin)) * w
            line(body, xx, y + h, xx, y + h + 8)
            text(body, str(tick), xx, y + h + 30, "tick", "middle")
    for r in data:
        t = max(r["optimize_ms"], 0.001)
        xx = x + (math.log(t) - math.log(xmin)) / (math.log(xmax) - math.log(xmin)) * w
        yy = y + h - (r["makespan"] - ymin) / (ymax - ymin) * h
        marker(body, xx, yy, r["method"], 4 if r["n"] < 50 else 6, 0.55)
    text(body, "tempo de otimização (ms, escala log)", x + w / 2, y + h + 68, "label", "middle")
    text(body, "best makespan", x, y - 20, "label")
    text(body, "Todas as 4590 execuções estocásticas; marcadores maiores indicam n >= 50.", x, y - 52, "title")
    legend_direct(body, 880, 238)

    # Runtime scaling by size.
    x2, y2, w2, h2 = 170, 890, 860, 370
    n_values = sorted({r["n"] for r in data})
    series = {}
    for method in METHODS:
        series[method] = [(n, median([max(r["optimize_ms"], 0.001) for r in data if r["method"] == method and r["n"] == n])) for n in n_values]
    vals = [v for s in series.values() for _, v in s]
    ymin2, ymax2 = max(min(vals), 0.001), max(vals)
    axes(body, x2, y2, w2, h2, ymin2, ymax2, 4, "{:.0f}")
    min_n, max_n = min(n_values), max(n_values)
    for tick in [10, 20, 30, 50, 100]:
        if tick in n_values:
            xx = x2 + (math.log(tick) - math.log(min_n)) / (math.log(max_n) - math.log(min_n)) * w2
            line(body, xx, y2 + h2, xx, y2 + h2 + 8)
            text(body, str(tick), xx, y2 + h2 + 30, "tick", "middle")
    for method, pts in series.items():
        parts = []
        for i, (n, value) in enumerate(pts):
            xx = x2 + (math.log(n) - math.log(min_n)) / (math.log(max_n) - math.log(min_n)) * w2
            v = max(value, 0.001)
            yy = y2 + h2 - (math.log(v) - math.log(ymin2)) / (math.log(ymax2) - math.log(ymin2)) * h2
            parts.append(("M" if i == 0 else "L") + f" {xx:.1f} {yy:.1f}")
            marker(body, xx, yy, method, 7)
        body.append(f'<path d="{" ".join(parts)}" fill="none" stroke="{COLORS[method]}" stroke-width="3"/>')
    text(body, "n (escala log)", x2 + w2 / 2, y2 + h2 + 68, "label", "middle")
    text(body, "tempo mediano de otimização (ms, escala log)", x2, y2 - 20, "label")

    # Convergence on 100a seed 0.
    x3, y3, w3, h3 = 1280, 220, 850, 520
    conv = {}
    for method in METHODS:
        row = run_for(rows, "100a", method, 0)
        conv[method] = [(int(f["iter"]), float(f["best_makespan"])) for f in load_frames(row["run_id"])] if row else []
    vals3 = [v for s in conv.values() for _, v in s]
    ymin3, ymax3 = min(vals3), max(vals3)
    axes(body, x3, y3, w3, h3, ymin3, ymax3, 5, "{:.0f}")
    for tick in [1, 25, 50, 75, 100]:
        xx = x3 + (tick - 1) / 99 * w3
        line(body, xx, y3 + h3, xx, y3 + h3 + 8)
        text(body, str(tick), xx, y3 + h3 + 30, "tick", "middle")
    for method, pts in conv.items():
        path = []
        for i, (it, value) in enumerate(pts):
            xx = x3 + (it - 1) / 99 * w3
            yy = y3 + h3 - (value - ymin3) / (ymax3 - ymin3) * h3
            path.append(("M" if i == 0 else "L") + f" {xx:.1f} {yy:.1f}")
        body.append(f'<path d="{" ".join(path)}" fill="none" stroke="{COLORS[method]}" stroke-width="4"/>')
    text(body, "iteração", x3 + w3 / 2, y3 + h3 + 68, "label", "middle")
    text(body, "best makespan", x3, y3 - 20, "label")
    text(body, "Execução ilustrativa: instância 100a, semente 0.", x3, y3 - 52, "title")
    legend_direct(body, 1950, 238)
    write_svg_png("fig-runtime-convergence", body)


def tex_escape(value: str) -> str:
    return (
        value.replace("\\", r"\textbackslash{}")
        .replace("&", r"\&")
        .replace("%", r"\%")
        .replace("$", r"\$")
        .replace("#", r"\#")
        .replace("_", r"\_")
        .replace("{", r"\{")
        .replace("}", r"\}")
        .replace("~", r"\textasciitilde{}")
        .replace("^", r"\textasciicircum{}")
    )


def fetch_route_png(instance: str, run_ids: list[str]) -> bytes:
    if not RENDER_BASE:
        raise RuntimeError("FIGURE_RENDER_BASE must be set for route images")
    params = [("map", instance)]
    for i, run_id in enumerate(run_ids):
        params.append(("run" if i == 0 else f"run{i+1}", run_id))
    url = RENDER_BASE + "/api/render?" + urllib.parse.urlencode(params)
    with urllib.request.urlopen(url, timeout=30) as response:
        return response.read()


def write_route_tex(panels: list[tuple[str, str, str]], claim: str) -> None:
    lines = [
        r"% Fragmento gerado automaticamente por scripts/gerar-figuras-publicacao.py.",
        r"% Incluir a partir de monografia/ com: \input{figs/fig-route-small-multiples.tex}",
        r"\newcommand{\rotapainel}[4]{%",
        r"\begin{figure}[htbp]",
        r"\centering",
        r"\includegraphics[width=0.92\textwidth]{#3}",
        r"\caption{#1: #2. Todos os painéis de rota foram gerados pelo renderizador Go e incluem ticks de coordenadas e barra de escala de 0,5 unidade.}",
        r"\label{#4}",
        r"\end{figure}%",
        r"}",
        "",
        rf"% Afirmação: {tex_escape(claim)}",
    ]
    for label, title, image in panels:
        safe_label = image.removeprefix("route-panel-").removesuffix(".png")
        lines.append(rf"\rotapainel{{Painel {tex_escape(label)}}}{{{tex_escape(title)}}}{{figs/{image}}}{{fig:rota-{safe_label}}}")
    (FIGS / "fig-route-small-multiples.tex").write_text("\n".join(lines) + "\n", encoding="utf-8")


def figure_route_small_multiples(rows: list[dict]) -> None:
    claim = "as rotas devem ser comparadas no mesmo sistema de coordenadas para evidenciar diferenças entre métodos."
    instances = ["10a", "30a", "100a"]
    labels = ["A", "B", "C"]
    titles = ["10a: ACO, GA, PSO e ótimo", "30a: métodos estocásticos", "100a: métodos estocásticos"]
    panels: list[tuple[str, str, str]] = []
    for i, instance in enumerate(instances):
        methods = ["aco", "ga", "pso", "lowerbound"] + (["bruteforce"] if instance == "10a" else [])
        run_ids = []
        for method in methods:
            row = run_for(rows, instance, method, 0)
            if row:
                run_ids.append(row["run_id"])
        image_name = f"route-panel-{instance}.png"
        (FIGS / image_name).write_bytes(fetch_route_png(instance, run_ids))
        panels.append((labels[i], titles[i], image_name))
    write_route_tex(panels, claim)


def main() -> None:
    rows = load_summary()
    figure_quality_distribution(rows)
    figure_runtime_convergence(rows)
    if RENDER_BASE:
        figure_route_small_multiples(rows)
    else:
        print("Figura composta de rotas ignorada: FIGURE_RENDER_BASE não foi definido")
    print(f"Figuras científicas de publicação gravadas em {FIGS}")


if __name__ == "__main__":
    main()
