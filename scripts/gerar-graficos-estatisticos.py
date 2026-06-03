#!/usr/bin/env python3
"""Gera graficos estatisticos dos experimentos em SVG e PNG.

O script evita dependencias externas (pandas/matplotlib/seaborn). Ele le os JSONs
em src/data/results e cria artefatos em monografia/figs.
"""

from __future__ import annotations

import json
import math
import os
import re
import shutil
import statistics
import subprocess
from collections import defaultdict
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
RESULTS = ROOT / "src" / "data" / "results"
SUMMARY = RESULTS / "summary"
EVOLUTION = RESULTS / "evolution"
TIMING = RESULTS / "timing"
FIGS = ROOT / "monografia" / "figs"

METHODS = ["aco", "ga", "pso", "lowerbound"]
METHOD_LABELS = {"aco": "ACO", "ga": "GA", "pso": "PSO", "lowerbound": "LB", "bruteforce": "BF"}
COLORS = {"aco": "#e07810", "ga": "#0a66c2", "pso": "#20913c", "lowerbound": "#7b2d8e", "bruteforce": "#333333"}


def esc(text: str) -> str:
    return (
        str(text)
        .replace("&", "&amp;")
        .replace("<", "&lt;")
        .replace(">", "&gt;")
        .replace('"', "&quot;")
    )


def svg(width: int, height: int, body: list[str]) -> str:
    return "\n".join(
        [
            f'<svg xmlns="http://www.w3.org/2000/svg" width="{width}" height="{height}" viewBox="0 0 {width} {height}">',
            '<rect width="100%" height="100%" fill="white"/>',
            '<style>text{font-family:Arial,Helvetica,sans-serif;fill:#303030}.title{font-size:22px;font-weight:700}.label{font-size:12px}.small{font-size:10px}.axis{stroke:#444;stroke-width:1}.grid{stroke:#e5e5e5;stroke-width:1}.legend{font-size:13px}</style>',
            *body,
            "</svg>",
        ]
    )


def write_svg_png(name: str, content: str) -> None:
    FIGS.mkdir(parents=True, exist_ok=True)
    svg_path = FIGS / f"{name}.svg"
    png_path = FIGS / f"{name}.png"
    svg_path.write_text(content, encoding="utf-8")
    converter = shutil.which("rsvg-convert")
    if converter:
        subprocess.run([converter, "-o", str(png_path), str(svg_path)], check=True)


def quantile(values: list[float], q: float) -> float:
    if not values:
        return 0.0
    values = sorted(values)
    pos = (len(values) - 1) * q
    lo = math.floor(pos)
    hi = math.ceil(pos)
    if lo == hi:
        return values[lo]
    return values[lo] * (hi - pos) + values[hi] * (pos - lo)


def median(values: list[float]) -> float:
    return statistics.median(values) if values else 0.0


def instance_key(instance: str) -> tuple[int, str]:
    match = re.match(r"^(\d+)(.*)$", instance)
    if not match:
        return (999999, instance)
    return (int(match.group(1)), match.group(2))


def load_summary() -> list[dict]:
    rows: list[dict] = []
    for path in sorted(SUMMARY.glob("*.json")):
        with path.open(encoding="utf-8") as handle:
            data = json.load(handle)
        run_id = data.get("run_id") or path.stem
        parts = run_id.split("__")
        instance = parts[0]
        method = data.get("method") or (parts[1] if len(parts) > 1 else "")
        seed = data.get("seed", 0)
        result = data.get("result", {})
        timing_file = data.get("timing_file") or f"timing/{run_id}.json"
        total_ms = None
        optimize_ms = None
        timing_path = RESULTS / timing_file
        if timing_path.exists():
            with timing_path.open(encoding="utf-8") as handle:
                timing = json.load(handle)
            durations = timing.get("durations_ms", {})
            total_ms = durations.get("total")
            optimize_ms = durations.get("optimize")
        rows.append(
            {
                "run_id": run_id,
                "instance": instance,
                "n": instance_key(instance)[0],
                "method": method,
                "seed": seed,
                "makespan": float(result.get("best_makespan", 0.0)),
                "total_ms": float(total_ms) if total_ms is not None else None,
                "optimize_ms": float(optimize_ms) if optimize_ms is not None else None,
            }
        )
    return rows


def grouped_values(rows: list[dict], value_key: str) -> dict[tuple[str, str], list[float]]:
    groups: dict[tuple[str, str], list[float]] = defaultdict(list)
    for row in rows:
        value = row.get(value_key)
        if value is None:
            continue
        groups[(row["instance"], row["method"])].append(float(value))
    return groups


def bf_optima(rows: list[dict]) -> dict[str, float]:
    opts: dict[str, float] = {}
    for row in rows:
        if row["method"] == "bruteforce":
            opts[row["instance"]] = row["makespan"]
    return opts


def heatmap(name: str, title: str, rows_labels: list[str], cols_labels: list[str], values: dict[tuple[str, str], float], fmt: str, color_mode: str = "normal") -> None:
    cell_w, cell_h = 98, 28
    left, top = 130, 70
    width = left + cell_w * len(cols_labels) + 60
    height = top + cell_h * len(rows_labels) + 60
    vals = [v for v in values.values() if math.isfinite(v)]
    vmin = min(vals) if vals else 0
    vmax = max(vals) if vals else 1
    if vmax == vmin:
        vmax = vmin + 1

    def color(v: float) -> str:
        t = (v - vmin) / (vmax - vmin)
        if color_mode == "success":
            t = 1 - t
        r = int(245 * t + 46 * (1 - t))
        g = int(80 * t + 164 * (1 - t))
        b = int(60 * t + 92 * (1 - t))
        return f"#{r:02x}{g:02x}{b:02x}"

    body = [f'<text class="title" x="{width/2}" y="32" text-anchor="middle">{esc(title)}</text>']
    for j, col in enumerate(cols_labels):
        body.append(f'<text class="label" x="{left + j*cell_w + cell_w/2}" y="58" text-anchor="middle">{esc(col)}</text>')
    for i, row in enumerate(rows_labels):
        y = top + i * cell_h
        body.append(f'<text class="label" x="{left-8}" y="{y+18}" text-anchor="end">{esc(row)}</text>')
        for j, col in enumerate(cols_labels):
            x = left + j * cell_w
            value = values.get((row, col))
            fill = "#f5f5f5" if value is None else color(value)
            text = "-" if value is None else fmt.format(value)
            body.append(f'<rect x="{x}" y="{y}" width="{cell_w}" height="{cell_h}" fill="{fill}" stroke="white"/>')
            body.append(f'<text class="small" x="{x+cell_w/2}" y="{y+18}" text-anchor="middle">{esc(text)}</text>')
    write_svg_png(name, svg(width, height, body))


def chart_success_rate(rows: list[dict]) -> None:
    opts = bf_optima(rows)
    instances = sorted(opts, key=instance_key)
    values: dict[tuple[str, str], float] = {}
    for inst in instances:
        opt = opts[inst]
        for method in METHODS:
            vals = [row["makespan"] for row in rows if row["instance"] == inst and row["method"] == method]
            if vals:
                values[(inst, METHOD_LABELS[method])] = sum(1 for v in vals if abs(v - opt) < 1e-3) / len(vals) * 100
    heatmap("heatmap-success-rate", "Taxa de sucesso: encontrou o otimo (%)", instances, [METHOD_LABELS[m] for m in METHODS], values, "{:.1f}%", "success")


def chart_gap_vs_bf(rows: list[dict]) -> None:
    opts = bf_optima(rows)
    instances = sorted(opts, key=instance_key)
    values: dict[tuple[str, str], float] = {}
    for inst in instances:
        opt = opts[inst]
        for method in METHODS:
            vals = [row["makespan"] for row in rows if row["instance"] == inst and row["method"] == method]
            if vals:
                values[(inst, METHOD_LABELS[method])] = median([(v - opt) / opt * 100 for v in vals])
    heatmap("heatmap-gap-vs-bf", "Gap mediano vs brute-force (%)", instances, [METHOD_LABELS[m] for m in METHODS], values, "{:.2f}%")


def chart_makespan_heatmap(rows: list[dict]) -> None:
    instances = sorted({row["instance"] for row in rows if row["method"] in METHODS}, key=instance_key)
    groups = grouped_values(rows, "makespan")
    values: dict[tuple[str, str], float] = {}
    for inst in instances:
        for method in METHODS:
            vals = groups.get((inst, method), [])
            if vals:
                values[(inst, METHOD_LABELS[method])] = median(vals)
    heatmap("heatmap-makespan-median", "Makespan mediano por instancia e metodo", instances, [METHOD_LABELS[m] for m in METHODS], values, "{:.2f}")


def draw_axes(body: list[str], x: int, y: int, w: int, h: int, y_ticks: int = 5, y_min: float | None = None, y_max: float | None = None, y_fmt: str = "{:.0f}") -> None:
    body.append(f'<line class="axis" x1="{x}" y1="{y+h}" x2="{x+w}" y2="{y+h}"/>')
    body.append(f'<line class="axis" x1="{x}" y1="{y}" x2="{x}" y2="{y+h}"/>')
    for i in range(y_ticks + 1):
        yy = y + h - h * i / y_ticks
        body.append(f'<line class="grid" x1="{x}" y1="{yy:.1f}" x2="{x+w}" y2="{yy:.1f}"/>')
        if y_min is not None and y_max is not None:
            value = y_min + (y_max - y_min) * i / y_ticks
            body.append(f'<text class="small" x="{x-8}" y="{yy+4:.1f}" text-anchor="end">{esc(y_fmt.format(value))}</text>')


def draw_x_ticks(body: list[str], ticks: list[tuple[float, str]], y: float) -> None:
    for x, label in ticks:
        body.append(f'<line x1="{x:.1f}" y1="{y}" x2="{x:.1f}" y2="{y+5}" stroke="#444"/>')
        body.append(f'<text class="small" x="{x:.1f}" y="{y+18}" text-anchor="middle">{esc(label)}</text>')


def chart_scalability(rows: list[dict]) -> None:
    width, height = 860, 520
    left, top, plot_w, plot_h = 80, 70, 720, 360
    points: dict[str, list[tuple[int, float]]] = defaultdict(list)
    for method in METHODS:
        for n in sorted({row["n"] for row in rows if row["method"] == method}):
            vals = [max(row["optimize_ms"], 0.001) for row in rows if row["method"] == method and row["n"] == n and row["optimize_ms"] is not None]
            if vals:
                points[method].append((n, median(vals)))
    all_x = [n for series in points.values() for n, _ in series]
    all_y = [v for series in points.values() for _, v in series]
    min_x, max_x = min(all_x), max(all_x)
    min_y, max_y = max(1e-6, min(all_y)), max(all_y)

    def sx(n: int) -> float:
        return left + (math.log(n) - math.log(min_x)) / (math.log(max_x) - math.log(min_x)) * plot_w

    def sy(v: float) -> float:
        return top + plot_h - (math.log(v) - math.log(min_y)) / (math.log(max_y) - math.log(min_y)) * plot_h

    body = [f'<text class="title" x="{width/2}" y="32" text-anchor="middle">Escalabilidade: tempo de otimizacao vs n</text>']
    draw_axes(body, left, top, plot_w, plot_h, y_min=min_y, y_max=max_y, y_fmt="{:.0f}")
    draw_x_ticks(body, [(sx(n), str(n)) for n in sorted({10, 20, 30, 50, 100} & {n for series in points.values() for n, _ in series})], top + plot_h)
    for method in METHODS:
        series = points[method]
        path = " ".join(("M" if i == 0 else "L") + f" {sx(n):.1f} {sy(v):.1f}" for i, (n, v) in enumerate(series))
        body.append(f'<path d="{path}" fill="none" stroke="{COLORS[method]}" stroke-width="3"/>')
        for n, v in series:
            body.append(f'<circle cx="{sx(n):.1f}" cy="{sy(v):.1f}" r="4" fill="{COLORS[method]}"/>')
    body.extend(legend(620, 88))
    body.append(f'<text class="label" x="{left+plot_w/2}" y="{top+plot_h+45}" text-anchor="middle">n (escala log)</text>')
    body.append(f'<text class="label" x="18" y="{top+plot_h/2}" transform="rotate(-90 18 {top+plot_h/2})" text-anchor="middle">ms (escala log)</text>')
    write_svg_png("scalability-runtime", svg(width, height, body))


def legend(x: int, y: int) -> list[str]:
    body: list[str] = []
    for i, method in enumerate(METHODS):
        yy = y + i * 22
        body.append(f'<rect x="{x}" y="{yy-11}" width="14" height="14" fill="{COLORS[method]}"/>')
        body.append(f'<text class="legend" x="{x+20}" y="{yy}">{METHOD_LABELS[method]}</text>')
    return body


def chart_tradeoff(rows: list[dict]) -> None:
    width, height = 900, 560
    left, top, plot_w, plot_h = 90, 70, 720, 390
    data = [row for row in rows if row["method"] in METHODS and row["optimize_ms"] and row["makespan"]]
    xs = [row["optimize_ms"] for row in data]
    ys = [row["makespan"] for row in data]
    min_x, max_x = max(1e-6, min(xs)), max(xs)
    min_y, max_y = min(ys), max(ys)

    def sx(v: float) -> float:
        return left + (math.log(v) - math.log(min_x)) / (math.log(max_x) - math.log(min_x)) * plot_w

    def sy(v: float) -> float:
        return top + plot_h - (v - min_y) / (max_y - min_y) * plot_h

    body = [f'<text class="title" x="{width/2}" y="32" text-anchor="middle">Trade-off: qualidade vs tempo</text>']
    draw_axes(body, left, top, plot_w, plot_h, y_min=min_y, y_max=max_y, y_fmt="{:.0f}")
    time_ticks = [1, 10, 100, 1000]
    time_ticks = [t for t in time_ticks if min_x <= t <= max_x]
    draw_x_ticks(body, [(sx(t), str(t)) for t in time_ticks], top + plot_h)
    for row in data:
        radius = 2.2 if row["n"] < 50 else 3.4
        body.append(f'<circle cx="{sx(row["optimize_ms"]):.1f}" cy="{sy(row["makespan"]):.1f}" r="{radius}" fill="{COLORS[row["method"]]}" opacity="0.42"/>')
    body.extend(legend(660, 88))
    body.append(f'<text class="label" x="{left+plot_w/2}" y="{top+plot_h+45}" text-anchor="middle">tempo de otimizacao (ms, escala log)</text>')
    body.append(f'<text class="label" x="18" y="{top+plot_h/2}" transform="rotate(-90 18 {top+plot_h/2})" text-anchor="middle">best makespan</text>')
    write_svg_png("scatter-quality-vs-time", svg(width, height, body))


def chart_boxplot_stability(rows: list[dict]) -> None:
    instances = ["10a", "15a", "20a", "30a", "50a", "100a"]
    width, height = 1000, 620
    left, top = 55, 70
    panel_w, panel_h = 150, 220
    body = [f'<text class="title" x="{width/2}" y="32" text-anchor="middle">Distribuicao de makespan por metodo</text>']
    for idx, inst in enumerate(instances):
        col = idx % 3
        row = idx // 3
        x0 = left + col * 315
        y0 = top + row * 270
        vals_all = [r["makespan"] for r in rows if r["instance"] == inst and r["method"] in METHODS]
        mn, mx = min(vals_all), max(vals_all)
        if mx == mn:
            mx = mn + 1
        body.append(f'<text class="label" x="{x0+panel_w/2}" y="{y0-12}" text-anchor="middle">{inst}</text>')
        draw_axes(body, x0, y0, panel_w, panel_h, 4, mn, mx, "{:.1f}")
        for j, method in enumerate(METHODS):
            vals = sorted(r["makespan"] for r in rows if r["instance"] == inst and r["method"] == method)
            q1, q2, q3 = quantile(vals, 0.25), quantile(vals, 0.5), quantile(vals, 0.75)
            lo, hi = vals[0], vals[-1]
            cx = x0 + 30 + j * 45
            def yy(v: float) -> float:
                return y0 + panel_h - (v - mn) / (mx - mn) * panel_h
            body.append(f'<line x1="{cx}" y1="{yy(lo):.1f}" x2="{cx}" y2="{yy(hi):.1f}" stroke="{COLORS[method]}"/>')
            body.append(f'<rect x="{cx-12}" y="{yy(q3):.1f}" width="24" height="{max(1, yy(q1)-yy(q3)):.1f}" fill="{COLORS[method]}" opacity="0.45" stroke="{COLORS[method]}"/>')
            body.append(f'<line x1="{cx-13}" y1="{yy(q2):.1f}" x2="{cx+13}" y2="{yy(q2):.1f}" stroke="{COLORS[method]}" stroke-width="2"/>')
            body.append(f'<text class="small" x="{cx}" y="{y0+panel_h+16}" text-anchor="middle">{METHOD_LABELS[method]}</text>')
    write_svg_png("boxplot-estabilidade", svg(width, height, body))


def read_evolution(run_id: str) -> list[dict]:
    path = EVOLUTION / f"{run_id}.jsonl"
    frames: list[dict] = []
    if not path.exists():
        return frames
    with path.open(encoding="utf-8") as handle:
        for line in handle:
            line = line.strip()
            if line:
                frames.append(json.loads(line))
    return frames


def chart_convergence(rows: list[dict], instance: str) -> None:
    width, height = 860, 520
    left, top, plot_w, plot_h = 80, 70, 720, 360
    bands: dict[str, tuple[list[int], list[float], list[float], list[float]]] = {}
    for method in METHODS:
        runs_for_method = [r for r in rows if r["instance"] == instance and r["method"] == method]
        traces: list[list[dict]] = []
        for run in runs_for_method:
            frames = read_evolution(run["run_id"])
            if frames:
                traces.append(frames)
        if not traces:
            continue
        by_iter: dict[int, list[float]] = {}
        for trace in traces:
            for f in trace:
                it = int(f["iter"])
                if it not in by_iter:
                    by_iter[it] = []
                by_iter[it].append(float(f["best_makespan"]))
        iters = sorted(by_iter.keys())
        if len(iters) < 2:
            continue
        min_vals = [min(by_iter[i]) for i in iters]
        med_vals = [statistics.median(by_iter[i]) for i in iters]
        max_vals = [max(by_iter[i]) for i in iters]
        bands[method] = (iters, min_vals, med_vals, max_vals)
    all_x = [x for iters, _, _, _ in bands.values() for x in iters]
    all_y = [y for _, mins, meds, maxs in bands.values() for y in mins + meds + maxs]
    if not all_x or not all_y:
        return
    min_x, max_x = min(all_x), max(all_x)
    min_y, max_y = min(all_y), max(all_y)
    if max_y == min_y:
        max_y = min_y + 1
    def sx(v: int) -> float:
        return left + (v - min_x) / (max_x - min_x) * plot_w
    def sy(v: float) -> float:
        return top + plot_h - (v - min_y) / (max_y - min_y) * plot_h
    body = [f'<text class="title" x="{width/2}" y="32" text-anchor="middle">Convergencia - {instance}</text>']
    draw_axes(body, left, top, plot_w, plot_h, y_min=min_y, y_max=max_y, y_fmt="{:.1f}")
    draw_x_ticks(body, [(sx(x), str(x)) for x in [1, 25, 50, 75, 100] if min_x <= x <= max_x], top + plot_h)
    for method, (iters, mins, meds, maxs) in bands.items():
        pts_min = [(sx(iters[i]), sy(mins[i])) for i in range(len(iters))]
        pts_max = [(sx(iters[i]), sy(maxs[i])) for i in range(len(iters))]
        band_d = "M {:.1f} {:.1f}".format(*pts_min[0])
        for p in pts_min[1:]:
            band_d += " L {:.1f} {:.1f}".format(*p)
        for p in reversed(pts_max):
            band_d += " L {:.1f} {:.1f}".format(*p)
        band_d += " Z"
        body.append(f'<path d="{band_d}" fill="{COLORS[method]}" opacity="0.12" stroke="none"/>')
    for method, (iters, _, meds, _) in bands.items():
        path = " ".join(("M" if i == 0 else "L") + f" {sx(iters[i]):.1f} {sy(meds[i]):.1f}" for i in range(len(iters)))
        body.append(f'<path d="{path}" fill="none" stroke="{COLORS[method]}" stroke-width="3"/>')
    body.extend(legend(630, 88))
    body.append(f'<text class="label" x="{left+plot_w/2}" y="{top+plot_h+45}" text-anchor="middle">iteracao</text>')
    body.append(f'<text class="label" x="18" y="{top+plot_h/2}" transform="rotate(-90 18 {top+plot_h/2})" text-anchor="middle">best makespan</text>')
    write_svg_png(f"convergence-overlay-{instance}", svg(width, height, body))


def chart_performance_profile(rows: list[dict]) -> None:
    opts = bf_optima(rows)
    best_per_inst: dict[str, float] = {}
    for inst in {r["instance"] for r in rows}:
        inst_vals = [r["makespan"] for r in rows if r["instance"] == inst and r["method"] in METHODS]
        if not inst_vals:
            continue
        if inst in opts:
            best_per_inst[inst] = opts[inst]
        else:
            best_per_inst[inst] = min(inst_vals)

    ratios: dict[str, list[float]] = defaultdict(list)
    for method in METHODS:
        for r in rows:
            if r["method"] != method or r["instance"] not in best_per_inst:
                continue
            ref = best_per_inst[r["instance"]]
            if ref > 0:
                ratios[method].append(r["makespan"] / ref)

    width, height = 860, 520
    left, top, plot_w, plot_h = 80, 70, 720, 360
    max_tau = 3.0
    steps = 200

    def sx(tau: float) -> float:
        return left + (tau - 1.0) / (max_tau - 1.0) * plot_w

    def sy(p: float) -> float:
        return top + plot_h - p * plot_h

    body = [f'<text class="title" x="{width/2}" y="32" text-anchor="middle">Perfil de desempenho (Dolan-Moré)</text>']
    body.append(f'<line x1="{left}" y1="{top+plot_h}" x2="{left+plot_w}" y2="{top+plot_h}" class="axis"/>')
    body.append(f'<line x1="{left}" y1="{top}" x2="{left}" y2="{top+plot_h}" class="axis"/>')
    for i in range(6):
        tau = 1.0 + i * (max_tau - 1.0) / 5
        xx = sx(tau)
        body.append(f'<line class="grid" x1="{xx:.1f}" y1="{top}" x2="{xx:.1f}" y2="{top+plot_h}"/>')
        body.append(f'<text class="small" x="{xx:.1f}" y="{top+plot_h+18}" text-anchor="middle">{tau:.1f}</text>')
    for i in range(6):
        p = i / 5
        yy = sy(p)
        body.append(f'<line class="grid" x1="{left}" y1="{yy:.1f}" x2="{left+plot_w}" y2="{yy:.1f}"/>')
        body.append(f'<text class="small" x="{left-10}" y="{yy+4:.1f}" text-anchor="end">{p*100:.0f}</text>')
    body.append(f'<text class="label" x="{left+plot_w/2}" y="{top+plot_h+45}" text-anchor="middle">τ (fator de razão)</text>')
    body.append(f'<text class="label" x="18" y="{top+plot_h/2}" transform="rotate(-90 18 {top+plot_h/2})" text-anchor="middle">P(τ) (%)</text>')

    for method in METHODS:
        if method not in ratios or len(ratios[method]) < 10:
            continue
        all_r = sorted(ratios[method])
        n = len(all_r)
        prev_tau, prev_p = 1.0, 0.0
        d = ""
        for i_r, r in enumerate(all_r):
            tau = max(1.0, r)
            p = (i_r + 1) / n
            if tau > prev_tau + 1e-9:
                if d:
                    d += " L {:.1f} {:.1f}".format(sx(prev_tau + 1e-9), sy(prev_p))
                d += " L {:.1f} {:.1f}".format(sx(tau), sy(prev_p))
            prev_tau, prev_p = tau, p
        if d:
            d += " L {:.1f} {:.1f}".format(sx(max_tau), sy(prev_p))
            body.append(f'<path d="M {sx(1.0):.1f} {sy(0):.1f} {d}" fill="none" stroke="{COLORS[method]}" stroke-width="2.5"/>')
            body.append(f'<circle cx="{sx(1.0):.1f}" cy="{sy(0):.1f}" r="3" fill="{COLORS[method]}"/>')

    body.extend(legend(660, 88))
    write_svg_png("performance-profile", svg(width, height, body))


def chart_last_improvement(rows: list[dict]) -> None:
    values: dict[str, list[float]] = defaultdict(list)
    for row in rows:
        if row["method"] not in METHODS:
            continue
        frames = read_evolution(row["run_id"])
        if not frames:
            continue
        last = max(int(f["iter"]) for f in frames if f.get("best_sequence"))
        values[row["method"]].append(float(last))
    width, height = 640, 460
    left, top, plot_w, plot_h = 80, 60, 460, 310
    all_vals = [v for vals in values.values() for v in vals]
    mn, mx = min(all_vals), max(all_vals)
    def yy(v: float) -> float:
        return top + plot_h - (v - mn) / (mx - mn) * plot_h
    body = [f'<text class="title" x="{width/2}" y="32" text-anchor="middle">Iteracao da ultima melhoria</text>']
    draw_axes(body, left, top, plot_w, plot_h, y_min=mn, y_max=mx, y_fmt="{:.0f}")
    for i, method in enumerate(METHODS):
        vals = sorted(values[method])
        q1, q2, q3 = quantile(vals, 0.25), quantile(vals, 0.5), quantile(vals, 0.75)
        lo, hi = vals[0], vals[-1]
        cx = left + 90 + i * 140
        body.append(f'<line x1="{cx}" y1="{yy(lo):.1f}" x2="{cx}" y2="{yy(hi):.1f}" stroke="{COLORS[method]}"/>')
        body.append(f'<rect x="{cx-28}" y="{yy(q3):.1f}" width="56" height="{max(1, yy(q1)-yy(q3)):.1f}" fill="{COLORS[method]}" opacity="0.45" stroke="{COLORS[method]}"/>')
        body.append(f'<line x1="{cx-30}" y1="{yy(q2):.1f}" x2="{cx+30}" y2="{yy(q2):.1f}" stroke="{COLORS[method]}" stroke-width="2"/>')
        body.append(f'<text class="label" x="{cx}" y="{top+plot_h+35}" text-anchor="middle">{METHOD_LABELS[method]}</text>')
    write_svg_png("convergence-last-improvement", svg(width, height, body))


def main() -> None:
    rows = load_summary()
    chart_makespan_heatmap(rows)
    chart_gap_vs_bf(rows)
    chart_success_rate(rows)
    chart_scalability(rows)
    chart_tradeoff(rows)
    chart_boxplot_stability(rows)
    chart_last_improvement(rows)
    chart_performance_profile(rows)
    for instance in ["10a", "30a", "50a", "100a"]:
        chart_convergence(rows, instance)
    print(f"Graficos estatisticos gerados em {FIGS}")


if __name__ == "__main__":
    main()
