#!/usr/bin/env python3
"""Testes estatísticos: Friedman + Nemenyi + Wilcoxon/Holm + diagrama CD.

Referência: Demšar (2006) — Statistical Comparisons of Classifiers over
Multiple Data Sets, JMLR 7, 1–30.

Dependências: apenas bibliotecas padrão (sem scipy/matplotlib).
"""

from __future__ import annotations

import itertools
import json
import math
import shutil
import statistics
import subprocess
import sys
from collections import defaultdict
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
SUMMARY = ROOT / "src" / "data" / "results" / "summary"
FIGS = ROOT / "monografia" / "figs"

METHODS = ["aco", "ga", "pso"]
METHOD_LABELS = {"aco": "ACO", "ga": "GA", "pso": "PSO"}
COLORS = {"aco": "#e07810", "ga": "#0a66c2", "pso": "#20913c"}

# Critical values q_alpha for Nemenyi test (Demšar 2006, Table 5a)
# These are the Studentized range critical values DIVIDED by sqrt(2),
# as required by CD = q_alpha * sqrt(k*(k+1)/(6*N)).
NEMENYI_Q = {
    (3, 0.05): 2.343,
    (3, 0.10): 2.052,
    (4, 0.05): 2.569,
    (4, 0.10): 2.291,
    (5, 0.05): 2.728,
    (5, 0.10): 2.465,
    (6, 0.05): 2.850,
    (6, 0.10): 2.601,
}

# F-distribution critical values F(alpha, df1, df2) for Iman-Davenport.
# Here df1 = k-1 = 2, df2 = (k-1)*(N-1) = 2*(N-1).
# We interpolate for N=30 → df2=58.
F_CRIT = {
    (0.05, 2, 30): 3.316,
    (0.05, 2, 40): 3.232,
    (0.05, 2, 60): 3.150,
    (0.05, 2, 120): 3.072,
    (0.01, 2, 30): 5.390,
    (0.01, 2, 40): 5.179,
    (0.01, 2, 60): 4.977,
    (0.01, 2, 120): 4.787,
}

# Wilcoxon signed-rank critical values for two-tailed test at alpha=0.05
# W_crit(n, 0.05) for n pairs (excludes zeros).
WILCOXON_CRIT = {
    6: 0, 7: 2, 8: 4, 9: 6, 10: 8,
    11: 11, 12: 14, 13: 17, 14: 21, 15: 25,
    16: 30, 17: 35, 18: 40, 19: 46, 20: 52,
    21: 59, 22: 66, 23: 73, 24: 81, 25: 89,
    26: 98, 27: 107, 28: 117, 29: 127, 30: 137,
}


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
            '<style>text{font-family:Arial,Helvetica,sans-serif;fill:#303030}.title{font-size:20px;font-weight:700}.label{font-size:13px}.small{font-size:11px}.axis{stroke:#444;stroke-width:1.5}</style>',
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


def load_summary() -> list[dict]:
    rows: list[dict] = []
    for path in sorted(SUMMARY.glob("*.json")):
        with path.open(encoding="utf-8") as f:
            data = json.load(f)
        run_id = data.get("run_id") or path.stem
        parts = run_id.split("__")
        instance = parts[0]
        method = data.get("method") or (parts[1] if len(parts) > 1 else "")
        rows.append(
            {
                "instance": instance,
                "method": method,
                "seed": data.get("seed", 0),
                "makespan": float(data.get("result", {}).get("best_makespan", 0)),
            }
        )
    return rows


def count_seeds(rows: list[dict], instances: list[str]) -> int:
    seeds: set[int] = set()
    for r in rows:
        if r["instance"] in instances and r["method"] in METHODS:
            seeds.add(r["seed"])
    return len(seeds)


def rank_values(values: list[float]) -> dict[float, float]:
    """Assign ranks with tie handling (average rank for ties)."""
    sorted_v = sorted(values)
    ranks: dict[float, float] = {}
    i = 0
    n = len(sorted_v)
    while i < n:
        j = i
        while j < n and abs(sorted_v[j] - sorted_v[i]) < 1e-9:
            j += 1
        avg_rank = i + 1 + (j - i - 1) / 2.0
        for idx in range(i, j):
            ranks[sorted_v[idx]] = avg_rank
        i = j
    return ranks


def wilcoxon_signed_rank(
    x: list[float], y: list[float]
) -> tuple[float, float, int]:
    """Wilcoxon signed-rank test (two-tailed).

    Returns (W_statistic, p_value_approx, n_nonzero).

    Uses exact distribution for n <= 30, normal approximation otherwise.
    """
    diffs = [a - b for a, b in zip(x, y)]
    nonzero = [(d, abs(d)) for d in diffs if abs(d) > 1e-12]
    n = len(nonzero)
    if n == 0:
        return 0.0, 1.0, 0

    abs_vals = [d[1] for d in nonzero]
    rank_map = rank_values(abs_vals)

    w_plus = sum(rank_map[abs_d] for d, abs_d in nonzero if d > 0)
    w_minus = sum(rank_map[abs_d] for d, abs_d in nonzero if d < 0)
    w = min(w_plus, w_minus)

    if n <= 30:
        p = _wilcoxon_exact_p(w, n)
    else:
        mu = n * (n + 1) / 4
        sigma = math.sqrt(n * (n + 1) * (2 * n + 1) / 24)
        z = (w - mu) / sigma
        p = 2.0 * _norm_cdf(z)

    return w, p, n


def _wilcoxon_exact_p(w: float, n: int) -> float:
    """Exact two-tailed p-value for Wilcoxon signed-rank via DP."""
    total = n * (n + 1) // 2
    dp = [0] * (total + 1)
    dp[0] = 1
    for i in range(1, n + 1):
        for s in range(total, i - 1, -1):
            dp[s] += dp[s - i]

    w_int = int(w)
    count_le = sum(dp[: w_int + 1])
    total_combos = 1 << n
    p = 2.0 * count_le / total_combos
    return min(p, 1.0)


def _norm_cdf(z: float) -> float:
    """Standard normal CDF (Abramowitz and Stegun approximation)."""
    if z > 0:
        return 1.0 - _norm_cdf(-z)
    b0, b1, b2, b3, b4, b5 = 0.2316419, 0.319381530, -0.356563782, 1.781477937, -1.821255978, 1.330274429
    t = 1.0 / (1.0 + b0 * (-z))
    poly = b1 + t * (b2 + t * (b3 + t * (b4 + t * b5)))
    phi = math.exp(-z * z / 2) / math.sqrt(2 * math.pi)
    return phi * poly


def holm_correction(p_values: list[tuple[str, str, float]], alpha: float) -> list[tuple[str, str, float, bool]]:
    """Bonferroni-Holm correction for multiple comparisons.

    Returns list of (method1, method2, p_value, significant).
    """
    m = len(p_values)
    sorted_idx = sorted(range(m), key=lambda i: p_values[i][2])
    results: list[tuple[str, str, float, bool, int]] = []
    rejected_so_far = True
    for rank, idx in enumerate(sorted_idx):
        m1, m2, p = p_values[idx]
        adj_alpha = alpha / (m - rank)
        sig = rejected_so_far and p <= adj_alpha
        rejected_so_far = sig
        results.append((m1, m2, p, sig, rank + 1))
    return results


def _betacf(a: float, b: float, x: float) -> float:
    """Continued fraction for incomplete beta, Numerical Recipes style."""
    max_iter = 200
    eps = 3e-14
    fpmin = 1e-300

    qab = a + b
    qap = a + 1.0
    qam = a - 1.0
    c = 1.0
    d = 1.0 - qab * x / qap
    if abs(d) < fpmin:
        d = fpmin
    d = 1.0 / d
    h = d

    for m in range(1, max_iter + 1):
        m2 = 2 * m
        aa = m * (b - m) * x / ((qam + m2) * (a + m2))
        d = 1.0 + aa * d
        if abs(d) < fpmin:
            d = fpmin
        c = 1.0 + aa / c
        if abs(c) < fpmin:
            c = fpmin
        d = 1.0 / d
        h *= d * c

        aa = -(a + m) * (qab + m) * x / ((a + m2) * (qap + m2))
        d = 1.0 + aa * d
        if abs(d) < fpmin:
            d = fpmin
        c = 1.0 + aa / c
        if abs(c) < fpmin:
            c = fpmin
        d = 1.0 / d
        delta = d * c
        h *= delta
        if abs(delta - 1.0) < eps:
            break

    return h


def _regularized_beta(x: float, a: float, b: float) -> float:
    if x <= 0.0:
        return 0.0
    if x >= 1.0:
        return 1.0
    bt = math.exp(math.lgamma(a + b) - math.lgamma(a) - math.lgamma(b) + a * math.log(x) + b * math.log1p(-x))
    if x < (a + 1.0) / (a + b + 2.0):
        return bt * _betacf(a, b, x) / a
    return 1.0 - bt * _betacf(b, a, 1.0 - x) / b


def f_survival(f_value: float, df1: int, df2: int) -> float:
    """Survival function P(F(df1, df2) >= f_value)."""
    if f_value <= 0 or df1 <= 0 or df2 <= 0:
        return 1.0
    x = df2 / (df2 + df1 * f_value)
    return _regularized_beta(x, df2 / 2.0, df1 / 2.0)


def compute_test(rows: list[dict], alpha: float = 0.05) -> dict | None:
    data: dict[str, dict[str, list[float]]] = defaultdict(lambda: defaultdict(list))
    for r in rows:
        if r["method"] in METHODS:
            data[r["instance"]][r["method"]].append(r["makespan"])

    instances = [inst for inst, ms in data.items() if all(len(ms[m]) > 0 for m in METHODS)]
    if not instances:
        print("Nenhuma instância completa para todos os métodos.", file=sys.stderr)
        return None

    N = len(instances)
    k = len(METHODS)
    n_seeds = count_seeds(rows, instances)

    # Per-instance summary: median (Demšar protocol)
    values_per_method: dict[str, list[float]] = {m: [] for m in METHODS}
    rank_sums = {m: 0.0 for m in METHODS}
    inst_ranks: list[tuple[str, str, float]] = []

    for inst in sorted(instances):
        medians = {m: statistics.median(data[inst][m]) for m in METHODS}
        for m in METHODS:
            values_per_method[m].append(medians[m])

        sorted_ms = sorted(METHODS, key=lambda mm: medians[mm])
        i = 0
        while i < len(sorted_ms):
            group = [sorted_ms[i]]
            while i + len(group) < len(sorted_ms) and abs(medians[sorted_ms[i + len(group)]] - medians[group[0]]) < 1e-9:
                group.append(sorted_ms[i + len(group)])
            avg_rank = i + 1 + (len(group) - 1) / 2.0
            for m in group:
                rank_sums[m] += avg_rank
                inst_ranks.append((inst, m, avg_rank))
            i += len(group)

    avg_ranks = {m: rank_sums[m] / N for m in METHODS}

    # Friedman test statistic using MEAN ranks R_j = rank_sums[m] / N
    sum_Rj_sq = sum(avg_ranks[m] ** 2 for m in METHODS)
    chi2 = (12 * N) / (k * (k + 1)) * (sum_Rj_sq - k * (k + 1) ** 2 / 4)

    # Iman-Davenport correction
    denom = N * (k - 1) - chi2
    if denom > 0:
        F_F = (N - 1) * chi2 / denom
        df1 = k - 1
        df2 = (k - 1) * (N - 1)
    else:
        F_F = float("inf")
        df1 = df2 = 0

    # Nemenyi post-hoc
    q = NEMENYI_Q.get((k, alpha))
    CD = q * math.sqrt(k * (k + 1) / (6 * N)) if q else 0.0

    # Wilcoxon signed-rank (pairwise, using median per instance)
    pairs = [(m1, m2) for i, m1 in enumerate(METHODS) for m2 in METHODS[i + 1 :]]
    wilcoxon_results = []
    for m1, m2 in pairs:
        w, p, nz = wilcoxon_signed_rank(values_per_method[m1], values_per_method[m2])
        wilcoxon_results.append({"method1": m1, "method2": m2, "W": w, "p": p, "n": nz})

    holm_results = holm_correction(
        [(r["method1"], r["method2"], r["p"]) for r in wilcoxon_results], alpha
    )

    return {
        "N": N,
        "k": k,
        "alpha": alpha,
        "n_seeds": n_seeds,
        "rank_sums": rank_sums,
        "avg_ranks": avg_ranks,
        "chi2": chi2,
        "F_F": F_F,
        "df1": df1,
        "df2": df2,
        "p_iman_davenport": f_survival(F_F, df1, df2) if math.isfinite(F_F) else 0.0,
        "CD": CD,
        "q": q,
        "inst_ranks": inst_ranks,
        "values_per_method": values_per_method,
        "wilcoxon": wilcoxon_results,
        "holm": holm_results,
    }


def cd_diagram(result: dict, filename: str = "cd-diagram") -> None:
    k = result["k"]
    avg = result["avg_ranks"]
    CD = result["CD"]
    sorted_methods = sorted(METHODS, key=lambda m: avg[m])

    width, height = 700, 110 + k * 45
    left, right = 110, 630
    axis_len = right - left
    rank_lo, rank_hi = 0.4, k + 0.6

    def rx(r: float) -> float:
        return left + (r - rank_lo) / (rank_hi - rank_lo) * axis_len

    body: list[str] = [
        f'<text class="title" x="{width/2}" y="30" text-anchor="middle">'
        f'Diagrama de diferenças críticas (CD = {CD:.3f}, α = {result["alpha"]})</text>'
    ]

    y_axis = 48
    body.append(f'<line x1="{left}" y1="{y_axis}" x2="{right}" y2="{y_axis}" class="axis"/>')

    for r in range(1, k + 1):
        x = rx(r)
        body.append(
            f'<line x1="{x:.1f}" y1="{y_axis - 5}" x2="{x:.1f}" y2="{y_axis + 5}" stroke="#444" stroke-width="1.2"/>'
        )
        body.append(f'<text class="label" x="{x:.1f}" y="{y_axis + 20}" text-anchor="middle">{r}</text>')

    body.append(
        f'<text class="small" x="{left}" y="{y_axis + 35}" text-anchor="middle">melhor</text>'
    )
    body.append(
        f'<text class="small" x="{right}" y="{y_axis + 35}" text-anchor="middle">pior</text>'
    )

    cd_y = y_axis + 42
    cd_x0 = rx(rank_hi - CD)
    cd_x0 = max(left, min(right - 10, cd_x0))
    body.append(
        f'<line x1="{cd_x0:.1f}" y1="{cd_y}" x2="{right}" y2="{cd_y}" stroke="#555" stroke-width="2"/>'
    )
    body.append(
        f'<line x1="{cd_x0:.1f}" y1="{cd_y - 4}" x2="{cd_x0:.1f}" y2="{cd_y + 4}" stroke="#555" stroke-width="1.5"/>'
    )
    body.append(
        f'<line x1="{right}" y1="{cd_y - 4}" x2="{right}" y2="{cd_y + 4}" stroke="#555" stroke-width="1.5"/>'
    )
    body.append(
        f'<text class="small" x="{(cd_x0 + right) / 2:.1f}" y="{cd_y - 6}" text-anchor="middle">CD</text>'
    )

    groups = []
    current = [sorted_methods[0]]
    for m in sorted_methods[1:]:
        if avg[m] - avg[current[0]] < CD:
            current.append(m)
        else:
            groups.append(current)
            current = [m]
    groups.append(current)

    bar_y0 = y_axis + 52
    for g in groups:
        if len(g) > 1:
            y_bar = bar_y0 + sorted_methods.index(g[0]) * 45
            x1 = rx(avg[g[0]])
            x2 = rx(avg[g[-1]])
            body.append(
                f'<line x1="{x1:.1f}" y1="{y_bar}" x2="{x2:.1f}" y2="{y_bar}" stroke="#333" stroke-width="4" stroke-linecap="round"/>'
            )

    for i, m in enumerate(sorted_methods):
        y = y_axis + 68 + i * 45
        x = rx(avg[m])
        body.append(
            f'<line x1="{x:.1f}" y1="{y_axis}" x2="{x:.1f}" y2="{y - 12}" '
            f'stroke="#bbb" stroke-width="1" stroke-dasharray="3,3"/>'
        )
        body.append(
            f'<circle cx="{x:.1f}" cy="{y - 12}" r="7" fill="{COLORS[m]}" '
            f'stroke="white" stroke-width="2"/>'
        )
        body.append(
            f'<text class="label" x="{x + 16}" y="{y - 7}" '
            f'text-anchor="start">{METHOD_LABELS[m]} (rank médio {avg[m]:.3f})</text>'
        )

    write_svg_png(filename, svg(width, height, body))


def main() -> None:
    rows = load_summary()
    result = compute_test(rows, alpha=0.05)
    if not result:
        print("ERRO: não foi possível computar os testes estatísticos.", file=sys.stderr)
        sys.exit(1)

    print("=" * 60)
    print("TESTE DE FRIEDMAN E PÓS-HOC DE NEMENYI")
    print("=" * 60)
    print(f"Referência: Demšar (2006) — JMLR 7, 1–30")
    print(f"Métodos comparados: {', '.join(METHOD_LABELS[m] for m in METHODS)}")
    print(f"Instâncias (blocos): {result['N']}")
    print(f"Sementes por método e instância: {result['n_seeds']}")
    print(f"Valor-resumo por método×instância: mediana")
    print()
    print("Somas de postos (ranks) e postos médios:")
    for m in METHODS:
        print(f"  {METHOD_LABELS[m]:>6}: soma={result['rank_sums'][m]:.1f}  "
              f"médio={result['avg_ranks'][m]:.4f}")
    print()
    print(f"Estatística χ²_F de Friedman: {result['chi2']:.4f}")
    print(f"F_F corrigido (Iman-Davenport): F({result['df1']},{result['df2']}) = {result['F_F']:.4f}")

    p_global = result["p_iman_davenport"]
    print(f"p-valor aproximado (Iman-Davenport): {p_global:.6e}")

    # Critical-value interpretation, retained for transparency.
    df2_approx = result["df2"]
    if df2_approx <= 30:
        f_crit = F_CRIT.get((result["alpha"], result["df1"], 30))
    elif df2_approx <= 40:
        f_crit = F_CRIT.get((result["alpha"], result["df1"], 40))
    elif df2_approx <= 60:
        f_crit = F_CRIT.get((result["alpha"], result["df1"], 60))
    else:
        f_crit = F_CRIT.get((result["alpha"], result["df1"], 120))

    if f_crit and result["F_F"] > f_crit:
        print(f"  → p < {result['alpha']} (rejeita H₀: métodos diferem significativamente)")
    elif f_crit:
        print(f"  → p ≥ {result['alpha']} (não rejeita H₀)")
    else:
        print(f"  → (valor crítico não disponível para comparar)")

    print()
    print(f"Diferença crítica CD (Nemenyi, q={result['q']}, α={result['alpha']}): {result['CD']:.4f}")
    print()
    print("Comparações par-a-par (Nemenyi):")
    for i, m1 in enumerate(METHODS):
        for m2 in METHODS[i + 1 :]:
            diff = abs(result["avg_ranks"][m1] - result["avg_ranks"][m2])
            sig = "SIGNIFICATIVO" if diff >= result["CD"] else "não significativo"
            print(f"  {METHOD_LABELS[m1]} vs {METHOD_LABELS[m2]}: "
                  f"|Δrank| = {diff:.4f}  ({sig})")

    print()
    print("=" * 60)
    print("WILCOXON SIGNED-RANK + CORREÇÃO DE BONFERRONI-HOLM")
    print("=" * 60)
    print(f"Valor-resumo: mediana por instância; α = {result['alpha']}")
    print()

    for w in result["wilcoxon"]:
        m1, m2 = w["method1"], w["method2"]
        print(f"  {METHOD_LABELS[m1]} vs {METHOD_LABELS[m2]}: "
              f"W = {w['W']:.1f}, n = {w['n']}, p = {w['p']:.6e}")

    print()
    print("Correção de Bonferroni-Holm:")
    for r in result["holm"]:
        m1, m2, p, sig, rank = r
        label_sig = "SIGNIFICATIVO" if sig else "não significativo"
        print(f"  Passo {rank}/3 — {METHOD_LABELS[m1]} vs {METHOD_LABELS[m2]}: "
              f"p = {p:.6e}, α/{3-rank+1} = {result['alpha']/(3-rank+1):.6f}  ({label_sig})")

    cd_diagram(result)
    print(f"\nDiagrama CD salvo em {FIGS / 'cd-diagram.svg'}")


if __name__ == "__main__":
    main()
