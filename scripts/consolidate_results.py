#!/usr/bin/env python3
"""Consolidate 5029 experiment results into LaTeX tables and summary stats."""

import json
import os
import glob
import sys
from collections import defaultdict

RESULTS_DIR = os.path.join(os.path.dirname(__file__), "..", "src", "data", "results", "summary")
RESULTS_DIR = os.path.abspath(RESULTS_DIR)


def load_all_summaries():
    runs = []
    for fpath in sorted(glob.glob(os.path.join(RESULTS_DIR, "*.json"))):
        with open(fpath) as f:
            data = json.load(f)
        parts = os.path.basename(fpath).replace(".json", "").split("__")
        instance = parts[0]
        method = parts[1]
        runs.append({
            "instance": instance,
            "method": method,
            "seed": data.get("seed"),
            "makespan": data["result"]["best_makespan"],
        })
    return runs


def compute_stats(runs):
    grouped = defaultdict(list)
    for r in runs:
        key = (r["instance"], r["method"])
        grouped[key].append(r["makespan"])

    stats = {}
    for (inst, method), values in grouped.items():
        values_sorted = sorted(values)
        n = len(values_sorted)
        mean = sum(values_sorted) / n
        median = values_sorted[n // 2] if n % 2 == 1 else (values_sorted[n//2-1] + values_sorted[n//2]) / 2
        std = (sum((v - mean)**2 for v in values_sorted) / n)**0.5
        stats[(inst, method)] = {
            "n": n,
            "best": min(values_sorted),
            "mean": mean,
            "median": median,
            "std": std,
            "worst": max(values_sorted),
        }
    return stats


def get_brute_optimum(stats):
    optimum = {}
    for (inst, method), s in stats.items():
        if method == "bruteforce":
            optimum[inst] = s["best"]
    return optimum


def format_table(stats, instances, methods, optimum=None):
    lines = []
    lines.append(r"\begin{table}[htbp]")
    lines.append(r"\centering")
    lines.append(r"\caption{Resultados consolidados por instância e método.}")
    lines.append(r"\label{tab:resultados-consolidados}")
    cols = "l" + "r" * len(methods)
    lines.append(r"\begin{tabular}{%s}" % cols)
    lines.append(r"\toprule")

    header = "Instância"
    for m in methods:
        header += " & " + {"ga": "GA", "pso": "PSO", "aco": "ACO", "bruteforce": "BF"}.get(m, m.upper())
        if optimum:
            header += r"~(gap\%)"
    header += r" \\"
    lines.append(header)
    lines.append(r"\midrule")

    for inst in instances:
        row = inst
        for m in methods:
            key = (inst, m)
            if key in stats:
                s = stats[key]
                row += f" & {s['best']:.2f}"
                if optimum and inst in optimum and m != "bruteforce":
                    gap = (s["best"] - optimum[inst]) / optimum[inst] * 100
                    row += f" ({gap:.2f}\%)"
            else:
                row += " & ---"
        row += r" \\"
        lines.append(row)

    lines.append(r"\bottomrule")
    lines.append(r"\end{tabular}")
    lines.append(r"\end{table}")
    return "\n".join(lines)


def format_detailed_table(stats, instances, methods, optimum=None):
    lines = []
    lines.append(r"\begin{table}[htbp]")
    lines.append(r"\centering")
    lines.append(r"\caption{Estatísticas descritivas do melhor makespan por instância e método.}")
    lines.append(r"\label{tab:estatisticas-descritivas}")
    cols = "l" + "l" * len(methods)
    lines.append(r"\begin{tabular}{%s}" % cols)
    lines.append(r"\toprule")

    for inst in instances:
        lines.append(r"\midrule")
        lines.append(r"\multicolumn{%d}{l}{\textbf{%s}} \\" % (len(methods) + 1, inst))
        lines.append(r"\midrule")
        header = "Métrica"
        for m in methods:
            header += " & " + {"ga": "GA", "pso": "PSO", "aco": "ACO"}.get(m, m.upper())
        header += r" \\"
        lines.append(header)
        lines.append(r"\midrule")

        metrics = [
            ("Média", "mean"),
            ("Mediana", "median"),
            ("Desv. Pad.", "std"),
            ("Melhor", "best"),
            ("Pior", "worst"),
            ("N", "n"),
        ]
        for label, key in metrics:
            row = label
            for m in methods:
                s = stats.get((inst, m))
                if s:
                    if key == "n":
                        row += f" & {s[key]:.0f}"
                    else:
                        row += f" & {s[key]:.2f}"
                else:
                    row += " & ---"
            row += r" \\"
            lines.append(row)

    lines.append(r"\bottomrule")
    lines.append(r"\end{tabular}")
    lines.append(r"\end{table}")
    return "\n".join(lines)


def main():
    print("Carregando resumos...", file=sys.stderr)
    runs = load_all_summaries()
    print(f"Total de execuções: {len(runs)}", file=sys.stderr)

    stats = compute_stats(runs)
    optimum = get_brute_optimum(stats)

    all_instances = sorted(set(r["instance"] for r in runs))
    all_methods = ["ga", "pso", "aco", "bruteforce"]

    size_groups = defaultdict(list)
    for inst in all_instances:
        size = int(''.join(c for c in inst if c.isdigit()))
        size_groups[size].append(inst)

    for size in sorted(size_groups):
        print(f"\n## Instâncias de tamanho {size}\n")
        instances = size_groups[size]
        for inst in instances:
            print(f"### {inst}")
            for m in ["bruteforce", "ga", "pso", "aco"]:
                s = stats.get((inst, m))
                if s:
                    gap_str = ""
                    if optimum and inst in optimum and m != "bruteforce":
                        gap = (s["best"] - optimum[inst]) / optimum[inst] * 100
                        gap_str = f", gap={gap:.2f}%"
                    print(f"  {m:>12}: best={s['best']:.4f}, mean={s['mean']:.4f}, "
                          f"median={s['median']:.4f}, std={s['std']:.4f}, n={s['n']:.0f}{gap_str}")

    print(f"\n## Tabela de resultados (best makespan)")
    medium_instances = [i for i in all_instances if i not in optimum or not any(
        (i, "bruteforce") in stats for i in [i])]
    small_instances = sorted(optimum.keys())

    print("\n### Instâncias pequenas (com busca exaustiva)\n")
    print(format_table(
        stats,
        small_instances,
        ["bruteforce", "ga", "pso", "aco"],
        optimum,
    ))

    larger = [i for i in all_instances if i not in optimum]
    print("\n### Instâncias maiores (apenas meta-heurísticas)\n")
    print(format_table(
        stats,
        larger,
        ["ga", "pso", "aco"],
    ))

    print("\n## Tabela detalhada (instâncias pequenas)\n")
    print(format_detailed_table(
        stats,
        small_instances,
        ["ga", "pso", "aco"],
        optimum,
    ))


if __name__ == "__main__":
    main()
