#!/usr/bin/env python3
"""Consolidate experiment summaries into LaTeX tables and summary stats."""

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
        header += " & " + {"ga": "GA", "pso": "PSO", "aco": "ACO", "lowerbound": "LB", "bruteforce": "BF"}.get(m, m.upper())
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
                    row += f" ({gap:.2f}\\%)"
            else:
                row += " & ---"
        row += r" \\"
        lines.append(row)

    lines.append(r"\bottomrule")
    lines.append(r"\end{tabular}")
    lines.append(r"\end{table}")
    return "\n".join(lines)


def format_ap_bound_table(stats, instances, methods):
    lines = []
    lines.append(r"\begin{table}[htbp]")
    lines.append(r"\centering")
    lines.append(r"\caption{Gap do makespan em relação ao limitante inferior (AP bound).}")
    lines.append(r"\label{tab:gap-ap-bound}")
    lines.append(r"\begin{tabular}{l" + "r" * len(methods) + "}")
    lines.append(r"\toprule")
    header = "Instância"
    for m in methods:
        header += " & " + {"ga": "GA", "pso": "PSO", "aco": "ACO"}.get(m, m.upper()) + " gap\\%"
    header += r" \\"
    lines.append(header)
    lines.append(r"\midrule")
    for inst in instances:
        key_lb = (inst, "lowerbound")
        if key_lb not in stats:
            continue
        lb_val = stats[key_lb]["best"]
        row = inst
        for m in methods:
            key = (inst, m)
            if key in stats:
                gap = (stats[key]["best"] - lb_val) / lb_val * 100
                row += f" & {gap:.2f}\\%"
            else:
                row += " & ---"
        row += r" \\"
        lines.append(row)
    lines.append(r"\bottomrule")
    lines.append(r"\end{tabular}")
    lines.append(r"\end{table}")
    return "\n".join(lines)


def format_unified_table(stats, instances, methods, optimum=None):
    lines = []
    lines.append(r"\begin{table}[htbp]")
    lines.append(r"\centering")
    lines.append(r"\caption{Estatísticas descritivas do makespan: melhor, média, pior e desvio.}")
    lines.append(r"\label{tab:estatisticas-descritivas}")
    lines.append(r"\begin{tabular}{llrrrr}")
    lines.append(r"\toprule")
    lines.append(r"Instância & Método & Best & Mean & Worst & Std \\")
    lines.append(r"\midrule")
    for inst in instances:
        first = True
        for m in methods:
            key = (inst, m)
            if key not in stats:
                continue
            s = stats[key]
            label = inst if first else ""
            first = False
            meth_name = {"ga": "GA", "pso": "PSO", "aco": "ACO", "lowerbound": "LB", "bruteforce": "BF"}.get(m, m.upper())
            row = f"{label} & {meth_name} & {s['best']:.2f} & {s['mean']:.2f} & {s['worst']:.2f} & {s['std']:.2f} \\\\"
            lines.append(row)
        if not first:
            lines.append(r"\midrule")
    lines.append(r"\bottomrule")
    lines.append(r"\end{tabular}")
    lines.append(r"\end{table}")
    return "\n".join(lines)


def format_ranking_table(stats, instances, methods):
    lines = []
    lines.append(r"\begin{table}[htbp]")
    lines.append(r"\centering")
    lines.append(r"\caption{Ranking dos métodos por instância (melhor makespan).}")
    lines.append(r"\label{tab:ranking-metodos}")
    lines.append(r"\begin{tabular}{l" + "r" * len(methods) + "}")
    lines.append(r"\toprule")
    header = "Instância"
    for m in methods:
        header += " & " + {"ga": "GA", "pso": "PSO", "aco": "ACO"}.get(m, m.upper())
    header += r" \\"
    lines.append(header)
    lines.append(r"\midrule")
    rank_sums = {m: 0.0 for m in methods}
    count = {m: 0 for m in methods}
    for inst in instances:
        pairs = [(m, stats.get((inst, m), {}).get("best")) for m in methods]
        pairs = [(m, v) for m, v in pairs if v is not None]
        if len(pairs) < 2:
            continue
        sorted_pairs = sorted(pairs, key=lambda p: p[1])
        ranks = {}
        i = 0
        while i < len(sorted_pairs):
            group = [sorted_pairs[i]]
            while i + len(group) < len(sorted_pairs) and abs(sorted_pairs[i + len(group)][1] - group[0][1]) < 1e-9:
                group.append(sorted_pairs[i + len(group)])
            avg_rank = i + 1 + (len(group) - 1) / 2.0
            for m, _ in group:
                ranks[m] = avg_rank
            i += len(group)
        row = inst
        for m in methods:
            r = ranks.get(m)
            if r is not None:
                rank_sums[m] += r
                count[m] += 1
            row += f" & {r if r else '-'}"
        row += r" \\"
        lines.append(row)
    lines.append(r"\midrule")
    row = "Médio"
    for m in methods:
        avg = rank_sums[m] / count[m] if count[m] > 0 else 0
        row += f" & {avg:.2f}"
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
            header += " & " + {"ga": "GA", "pso": "PSO", "aco": "ACO", "lowerbound": "LB"}.get(m, m.upper())
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
    all_methods = ["ga", "pso", "aco", "lowerbound", "bruteforce"]

    size_groups = defaultdict(list)
    for inst in all_instances:
        size = int(''.join(c for c in inst if c.isdigit()))
        size_groups[size].append(inst)

    for size in sorted(size_groups):
        print(f"\n## Instâncias de tamanho {size}\n")
        instances = size_groups[size]
        for inst in instances:
            print(f"### {inst}")
            for m in ["bruteforce", "ga", "pso", "aco", "lowerbound"]:
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
        ["bruteforce", "ga", "pso", "aco", "lowerbound"],
        optimum,
    ))

    larger = [i for i in all_instances if i not in optimum]
    print("\n### Instâncias maiores (apenas meta-heurísticas)\n")
    print(format_table(
        stats,
        larger,
        ["ga", "pso", "aco", "lowerbound"],
        optimum,
    ))

    print("\n## Tabela detalhada (instâncias pequenas)\n")
    print(format_detailed_table(
        stats,
        small_instances,
        ["ga", "pso", "aco", "lowerbound"],
        optimum,
    ))

    print("\n## Gap vs AP bound (instâncias grandes)\n")
    print(format_ap_bound_table(stats, larger, ["ga", "pso", "aco"]))

    print("\n## Tabela unificada best/mean/worst (instâncias pequenas)\n")
    print(format_unified_table(stats, small_instances, ["ga", "pso", "aco", "lowerbound"]))

    print("\n## Tabela unificada best/mean/worst (instâncias grandes)\n")
    print(format_unified_table(stats, larger, ["ga", "pso", "aco", "lowerbound"]))

    print("\n## Ranking métodos (instâncias pequenas)\n")
    print(format_ranking_table(stats, small_instances, ["ga", "pso", "aco"]))

    print("\n## Ranking métodos (instâncias grandes)\n")
    print(format_ranking_table(stats, larger, ["ga", "pso", "aco"]))


if __name__ == "__main__":
    main()
