#!/usr/bin/env python3
"""Generate deterministic TeX metrics and tables for the monograph."""

from __future__ import annotations

import argparse
import hashlib
import json
import math
import statistics
import sys
from collections import defaultdict
from dataclasses import dataclass
from pathlib import Path


ROOT = Path(__file__).resolve().parents[1]
DEFAULT_RESULTS_DIR = ROOT / "src" / "data" / "results"
DEFAULT_OUTPUT_DIR = ROOT / "monografia" / "generated"
SCRIPT_NAME = "gerar-metricas-monografia.py"
SCRIPT_VERSION = "2"
ALL_METHODS = ("bruteforce", "lowerbound", "ga", "pso", "aco")
META_METHODS = ("aco", "ga", "pso")
TABLE_METHODS = ("bruteforce", "lowerbound", "ga", "pso", "aco")


@dataclass(frozen=True)
class SummaryRow:
    run_id: str
    instance: str
    method: str
    seed: int | None
    best_makespan: float


@dataclass(frozen=True)
class TimingRow:
    run_id: str
    instance: str
    method: str
    optimize_ms: float
    total_ms: float


def main() -> int:
    args = parse_args()
    results_dir = args.results_dir.resolve()
    output_dir = args.output_dir.resolve()

    try:
        outputs = build_outputs(results_dir)
    except ValueError as exc:
        print(f"error: {exc}", file=sys.stderr)
        return 2

    if args.check:
        return check_outputs(output_dir, outputs)

    write_outputs(output_dir, outputs)
    print(f"generated {len(outputs)} files in {output_dir}")
    return 0


def parse_args() -> argparse.Namespace:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--results-dir", type=Path, default=DEFAULT_RESULTS_DIR)
    parser.add_argument("--output-dir", type=Path, default=DEFAULT_OUTPUT_DIR)
    parser.add_argument("--check", action="store_true", help="fail if generated files are missing or stale")
    return parser.parse_args()


def build_outputs(results_dir: Path) -> dict[str, str]:
    summaries, summary_files = load_summaries(results_dir)
    timings, timing_files = load_timings(results_dir)
    evolution_files = sorted((results_dir / "evolution").glob("*.jsonl"))
    if not summaries:
        raise ValueError(f"no summary JSON files found under {results_dir / 'summary'}")

    derived = derive_metrics(summaries, timings)
    outputs = {
        "metrics.tex": render_metrics(summaries, timings, evolution_files, derived),
        "tables/tab-cobertura-experimental.tex": render_coverage_table(summaries),
        "tables/tab-otimos-bf.tex": render_bruteforce_optimum_table(summaries),
        "tables/tab-gap-bf.tex": render_bruteforce_gap_table(summaries),
        "tables/tab-grandes.tex": render_large_instance_table(summaries),
        "tables/tab-lb-gap.tex": render_lower_bound_gap_table(summaries),
        "tables/tab-tempo-100.tex": render_timing_100_table(summaries, timings),
    }
    outputs["manifest.json"] = render_manifest(results_dir, summary_files, timing_files, evolution_files, summaries, outputs)
    return outputs


def load_summaries(results_dir: Path) -> tuple[list[SummaryRow], list[Path]]:
    summary_dir = results_dir / "summary"
    files = sorted(summary_dir.glob("*.json"))
    rows: list[SummaryRow] = []
    for path in files:
        data = json.loads(path.read_text(encoding="utf-8"))
        if data.get("status", "ok") != "ok":
            continue
        run_id = str(data.get("run_id") or path.stem)
        method = str(data.get("method") or method_from_run_id(run_id))
        best = data.get("result", {}).get("best_makespan")
        if best is None:
            continue
        rows.append(
            SummaryRow(
                run_id=run_id,
                instance=instance_from(data, run_id),
                method=method,
                seed=data.get("seed"),
                best_makespan=float(best),
            )
        )
    return rows, files


def load_timings(results_dir: Path) -> tuple[list[TimingRow], list[Path]]:
    timing_dir = results_dir / "timing"
    files = sorted(timing_dir.glob("*.json"))
    rows: list[TimingRow] = []
    for path in files:
        data = json.loads(path.read_text(encoding="utf-8"))
        run_id = str(data.get("run_id") or path.stem)
        method = str(data.get("method") or method_from_run_id(run_id))
        durations = data.get("durations_ms", {})
        optimize = durations.get("optimize")
        total = durations.get("total")
        if optimize is None or total is None:
            continue
        rows.append(
            TimingRow(
                run_id=run_id,
                instance=instance_from(data, run_id),
                method=method,
                optimize_ms=float(optimize),
                total_ms=float(total),
            )
        )
    return rows, files


def instance_from(data: dict, run_id: str) -> str:
    parts = run_id.split("__")
    if parts and parts[0]:
        return parts[0]
    instance = data.get("instance")
    if instance:
        return Path(str(instance)).stem
    return "unknown"


def method_from_run_id(run_id: str) -> str:
    parts = run_id.split("__")
    return parts[1] if len(parts) > 1 else "unknown"


def instance_sort_key(instance: str) -> tuple[int, str]:
    digits = "".join(ch for ch in instance if ch.isdigit())
    return (int(digits) if digits else -1, instance)


def pct(value: float) -> str:
    return f"{value:.2f}\\%"


def number(value: float) -> str:
    return fmt_pt(value, 2)


def fmt_pt(value: float, digits: int = 2) -> str:
    return f"{value:.{digits}f}".replace(".", "{,}")


def fmt_en(value: float, digits: int = 2) -> str:
    return f"{value:.{digits}f}"


def pct_pt(value: float, digits: int = 2) -> str:
    return f"{fmt_pt(value, digits)}\\%"


def pct_en(value: float, digits: int = 2) -> str:
    return f"{fmt_en(value, digits)}\\%"


def sci_pt(value: float) -> str:
    mantissa, exponent = f"{value:.6e}".split("e")
    return f"{mantissa.replace('.', '{,}')}\\times10^{{{int(exponent)}}}"


def sci_en(value: float) -> str:
    mantissa, exponent = f"{value:.6e}".split("e")
    return f"{mantissa}\\times10^{{{int(exponent)}}}"


def generated_header() -> list[str]:
    return [
        f"% Generated by scripts/{SCRIPT_NAME}; do not edit manually.",
        "% Source: src/data/results via scripts/gerar-metricas-monografia.py.",
    ]


def derive_metrics(summaries: list[SummaryRow], timings: list[TimingRow]) -> dict:
    gap_stats = compute_gap_stats(summaries)
    lb_stats = compute_lower_bound_stats(summaries)
    timing_stats = compute_timing_stats(timings)
    statistical = compute_statistical_test(summaries)
    return {
        "gap_stats": gap_stats,
        "lb_stats": lb_stats,
        "timing_stats": timing_stats,
        "statistical": statistical,
    }


def compute_gap_stats(summaries: list[SummaryRow]) -> dict[str, dict[str, float | int]]:
    best = best_by_key(summaries)
    optimum = {instance: value for (instance, method), value in best.items() if method == "bruteforce"}
    stats: dict[str, dict[str, float | int]] = {}
    for method in META_METHODS:
        gaps: list[float] = []
        hits = 0
        for row in summaries:
            if row.method != method or row.instance not in optimum:
                continue
            gap = (row.best_makespan - optimum[row.instance]) / optimum[row.instance] * 100
            gaps.append(gap)
            if abs(gap) <= 1e-9:
                hits += 1
        if gaps:
            stats[method] = {
                "runs": len(gaps),
                "mean": statistics.fmean(gaps),
                "min": min(gaps),
                "max": max(gaps),
                "success_rate": hits / len(gaps) * 100,
            }
    return stats


def compute_lower_bound_stats(summaries: list[SummaryRow]) -> dict[str, float | int]:
    best = best_by_key(summaries)
    gaps: list[float] = []
    violations = 0
    for (instance, method), optimum in best.items():
        if method != "bruteforce":
            continue
        lower = best.get((instance, "lowerbound"))
        if lower is None:
            continue
        if lower > optimum + 1e-9:
            violations += 1
        gaps.append((optimum - lower) / optimum * 100)
    return {
        "runs": len(gaps),
        "mean": statistics.fmean(gaps) if gaps else 0.0,
        "min": min(gaps) if gaps else 0.0,
        "max": max(gaps) if gaps else 0.0,
        "violations": violations,
    }


def compute_timing_stats(timings: list[TimingRow]) -> dict[str, float]:
    by_method_100: dict[str, list[float]] = defaultdict(list)
    by_instance_method: dict[tuple[str, str], list[float]] = defaultdict(list)
    for row in timings:
        if instance_sort_key(row.instance)[0] >= 100:
            by_method_100[row.method].append(row.optimize_ms)
            by_instance_method[(row.instance, row.method)].append(row.optimize_ms)

    stats: dict[str, float] = {}
    for method in META_METHODS:
        samples = by_method_100.get(method, [])
        stats[f"{method}_mean_100_ms"] = statistics.fmean(samples) if samples else 0.0
        per_instance = [statistics.fmean(samples) for (inst, m), samples in by_instance_method.items() if m == method]
        stats[f"{method}_max_instance_mean_100_ms"] = max(per_instance) if per_instance else 0.0
    return stats


def compute_statistical_test(summaries: list[SummaryRow]) -> dict[str, float | int | dict[str, float]]:
    data: dict[str, dict[str, list[float]]] = defaultdict(lambda: defaultdict(list))
    for row in summaries:
        if row.method in META_METHODS:
            data[row.instance][row.method].append(row.best_makespan)

    instances = sorted([inst for inst, methods in data.items() if all(methods[m] for m in META_METHODS)], key=instance_sort_key)
    n_instances = len(instances)
    k_methods = len(META_METHODS)
    if not instances:
        return {"N": 0, "k": k_methods, "df1": 0, "df2": 0, "F_F": 0.0, "p": 1.0, "CD": 0.0, "avg_ranks": {m: 0.0 for m in META_METHODS}}

    rank_sums = {method: 0.0 for method in META_METHODS}
    for instance in instances:
        medians = {method: statistics.median(data[instance][method]) for method in META_METHODS}
        ordered = sorted(META_METHODS, key=lambda method: medians[method])
        idx = 0
        while idx < len(ordered):
            tied = [ordered[idx]]
            while idx + len(tied) < len(ordered) and abs(medians[ordered[idx + len(tied)]] - medians[tied[0]]) < 1e-9:
                tied.append(ordered[idx + len(tied)])
            avg_rank = idx + 1 + (len(tied) - 1) / 2.0
            for method in tied:
                rank_sums[method] += avg_rank
            idx += len(tied)

    avg_ranks = {method: rank_sums[method] / n_instances for method in META_METHODS}
    sum_ranks_sq = sum(avg_ranks[method] ** 2 for method in META_METHODS)
    chi2 = (12 * n_instances) / (k_methods * (k_methods + 1)) * (sum_ranks_sq - k_methods * (k_methods + 1) ** 2 / 4)
    denom = n_instances * (k_methods - 1) - chi2
    f_stat = (n_instances - 1) * chi2 / denom if denom > 0 else float("inf")
    df1 = k_methods - 1
    df2 = (k_methods - 1) * (n_instances - 1)
    cd = 2.343 * math.sqrt(k_methods * (k_methods + 1) / (6 * n_instances))
    p_value = f_survival(f_stat, df1, df2) if math.isfinite(f_stat) else 0.0
    return {"N": n_instances, "k": k_methods, "df1": df1, "df2": df2, "F_F": f_stat, "p": p_value, "CD": cd, "avg_ranks": avg_ranks}


def f_survival(f_value: float, df1: int, df2: int) -> float:
    if f_value <= 0 or df1 <= 0 or df2 <= 0:
        return 1.0
    x = df2 / (df2 + df1 * f_value)
    return regularized_beta(x, df2 / 2.0, df1 / 2.0)


def regularized_beta(x: float, a: float, b: float) -> float:
    if x <= 0.0:
        return 0.0
    if x >= 1.0:
        return 1.0
    bt = math.exp(math.lgamma(a + b) - math.lgamma(a) - math.lgamma(b) + a * math.log(x) + b * math.log1p(-x))
    if x < (a + 1.0) / (a + b + 2.0):
        return bt * beta_continued_fraction(a, b, x) / a
    return 1.0 - bt * beta_continued_fraction(b, a, 1.0 - x) / b


def beta_continued_fraction(a: float, b: float, x: float) -> float:
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


def summaries_by_key(summaries: list[SummaryRow]) -> dict[tuple[str, str], list[SummaryRow]]:
    grouped: dict[tuple[str, str], list[SummaryRow]] = defaultdict(list)
    for row in summaries:
        grouped[(row.instance, row.method)].append(row)
    return grouped


def best_by_key(summaries: list[SummaryRow]) -> dict[tuple[str, str], float]:
    grouped = summaries_by_key(summaries)
    return {key: min(row.best_makespan for row in rows) for key, rows in grouped.items()}


def render_metrics(summaries: list[SummaryRow], timings: list[TimingRow], evolution_files: list[Path], derived: dict) -> str:
    instances = sorted({row.instance for row in summaries}, key=instance_sort_key)
    small = [inst for inst in instances if any(row.instance == inst and row.method == "bruteforce" for row in summaries)]
    large = [inst for inst in instances if inst not in set(small)]
    method_counts = {method: sum(1 for row in summaries if row.method == method) for method in ALL_METHODS}
    seed_counts = {
        method: len({row.seed for row in summaries if row.method == method and row.seed is not None})
        for method in META_METHODS
    }
    sizes = sorted({instance_sort_key(instance)[0] for instance in instances})
    variants_per_size = max((sum(1 for instance in instances if instance_sort_key(instance)[0] == size) for size in sizes), default=0)
    gap_stats = derived["gap_stats"]
    lb_stats = derived["lb_stats"]
    timing_stats = derived["timing_stats"]
    statistical = derived["statistical"]
    avg_ranks = statistical["avg_ranks"]

    lines = generated_header()
    lines.extend(
        [
            f"\\newcommand{{\\TotalExecucoes}}{{{len(summaries)}}}",
            f"\\newcommand{{\\TotalArquivosResumo}}{{{len(summaries)}}}",
            f"\\newcommand{{\\TotalArquivosEvolucao}}{{{len(evolution_files)}}}",
            f"\\newcommand{{\\TotalArquivosTempo}}{{{len(timings)}}}",
            f"\\newcommand{{\\TotalInstancias}}{{{len(instances)}}}",
            f"\\newcommand{{\\TotalTamanhosInstancia}}{{{len(sizes)}}}",
            f"\\newcommand{{\\TotalVariantesPorTamanho}}{{{variants_per_size}}}",
            f"\\newcommand{{\\TotalInstanciasComOtimo}}{{{len(small)}}}",
            f"\\newcommand{{\\TotalInstanciasGrandes}}{{{len(large)}}}",
            f"\\newcommand{{\\TotalExecucoesBF}}{{{method_counts['bruteforce']}}}",
            f"\\newcommand{{\\TotalExecucoesLB}}{{{method_counts['lowerbound']}}}",
            f"\\newcommand{{\\TotalExecucoesGA}}{{{method_counts['ga']}}}",
            f"\\newcommand{{\\TotalExecucoesPSO}}{{{method_counts['pso']}}}",
            f"\\newcommand{{\\TotalExecucoesACO}}{{{method_counts['aco']}}}",
            f"\\newcommand{{\\TotalSementesGA}}{{{seed_counts['ga']}}}",
            f"\\newcommand{{\\TotalSementesPSO}}{{{seed_counts['pso']}}}",
            f"\\newcommand{{\\TotalSementesACO}}{{{seed_counts['aco']}}}",
            f"\\newcommand{{\\TotalExecucoesOtimoPorMetodo}}{{{int(gap_stats.get('aco', {}).get('runs', 0))}}}",
            f"\\newcommand{{\\GapMedioOtimoACO}}{{{pct_pt(float(gap_stats.get('aco', {}).get('mean', 0.0)), 4)}}}",
            f"\\newcommand{{\\GapMedioOtimoGA}}{{{pct_pt(float(gap_stats.get('ga', {}).get('mean', 0.0)), 4)}}}",
            f"\\newcommand{{\\GapMedioOtimoPSO}}{{{pct_pt(float(gap_stats.get('pso', {}).get('mean', 0.0)), 4)}}}",
            f"\\newcommand{{\\GapMedioOtimoACOEn}}{{{pct_en(float(gap_stats.get('aco', {}).get('mean', 0.0)), 4)}}}",
            f"\\newcommand{{\\GapMedioOtimoGAEn}}{{{pct_en(float(gap_stats.get('ga', {}).get('mean', 0.0)), 4)}}}",
            f"\\newcommand{{\\GapMedioOtimoPSOEn}}{{{pct_en(float(gap_stats.get('pso', {}).get('mean', 0.0)), 4)}}}",
            f"\\newcommand{{\\TaxaAcertoOtimoACO}}{{{pct_pt(float(gap_stats.get('aco', {}).get('success_rate', 0.0)), 2)}}}",
            f"\\newcommand{{\\TaxaAcertoOtimoGA}}{{{pct_pt(float(gap_stats.get('ga', {}).get('success_rate', 0.0)), 2)}}}",
            f"\\newcommand{{\\TaxaAcertoOtimoPSO}}{{{pct_pt(float(gap_stats.get('pso', {}).get('success_rate', 0.0)), 2)}}}",
            f"\\newcommand{{\\GapLBMedio}}{{{pct_pt(float(lb_stats['mean']), 2)}}}",
            f"\\newcommand{{\\GapLBMinimo}}{{{pct_pt(float(lb_stats['min']), 2)}}}",
            f"\\newcommand{{\\GapLBMaximo}}{{{pct_pt(float(lb_stats['max']), 2)}}}",
            f"\\newcommand{{\\ViolacoesLB}}{{{int(lb_stats['violations'])}}}",
            f"\\newcommand{{\\TempoOtimizacaoACOCemMedioSegundos}}{{{fmt_pt(float(timing_stats['aco_mean_100_ms']) / 1000, 2)} s}}",
            f"\\newcommand{{\\TempoOtimizacaoACOCemMedioSegundosEn}}{{{fmt_en(float(timing_stats['aco_mean_100_ms']) / 1000, 2)} s}}",
            f"\\newcommand{{\\TempoOtimizacaoGAMaxCemMs}}{{{fmt_pt(float(timing_stats['ga_max_instance_mean_100_ms']), 2)} ms}}",
            f"\\newcommand{{\\TempoOtimizacaoPSOMaxCemMs}}{{{fmt_pt(float(timing_stats['pso_max_instance_mean_100_ms']), 2)} ms}}",
            f"\\newcommand{{\\TempoOtimizacaoGAMaxCemMsEn}}{{{fmt_en(float(timing_stats['ga_max_instance_mean_100_ms']), 2)} ms}}",
            f"\\newcommand{{\\FriedmanDFOne}}{{{int(statistical['df1'])}}}",
            f"\\newcommand{{\\FriedmanDFTwo}}{{{int(statistical['df2'])}}}",
            f"\\newcommand{{\\FriedmanStatistic}}{{{fmt_pt(float(statistical['F_F']), 4)}}}",
            f"\\newcommand{{\\PValorImanDavenport}}{{{sci_pt(float(statistical['p']))}}}",
            f"\\newcommand{{\\CDNemenyi}}{{{fmt_pt(float(statistical['CD']), 4)}}}",
            f"\\newcommand{{\\RankMedioACO}}{{{fmt_pt(float(avg_ranks.get('aco', 0.0)), 4)}}}",
            f"\\newcommand{{\\RankMedioGA}}{{{fmt_pt(float(avg_ranks.get('ga', 0.0)), 4)}}}",
            f"\\newcommand{{\\RankMedioPSO}}{{{fmt_pt(float(avg_ranks.get('pso', 0.0)), 4)}}}",
        ]
    )
    return "\n".join(lines) + "\n"


def render_coverage_table(summaries: list[SummaryRow]) -> str:
    grouped = summaries_by_key(summaries)
    instances = sorted({row.instance for row in summaries}, key=instance_sort_key)
    labels = {
        "aco": "ACO",
        "ga": "GA",
        "pso": "PSO",
        "lowerbound": "Lower bound",
        "bruteforce": "Busca exaustiva",
    }
    lines = generated_header()
    lines.extend(
        [
            r"\begin{tabular}{lrrr}",
            r"\toprule",
            r"Metodo & Runs & Instancias & Sementes \\",
            r"\midrule",
        ]
    )
    for method in ("aco", "ga", "pso", "lowerbound", "bruteforce"):
        rows = [row for row in summaries if row.method == method]
        method_instances = {row.instance for row in rows}
        seeds = {row.seed for row in rows if row.seed is not None}
        seed_count = len(seeds) if seeds else 1
        lines.append(f"{labels[method]} & {len(rows)} & {len(method_instances)} & {seed_count} " + r"\\")
    lines.extend([r"\bottomrule", r"\end{tabular}"])
    return "\n".join(lines) + "\n"


def render_bruteforce_optimum_table(summaries: list[SummaryRow]) -> str:
    best = best_by_key(summaries)
    instances = sorted({instance for instance, method in best if method == "bruteforce"}, key=instance_sort_key)
    lines = generated_header()
    lines.extend([r"\begin{tabular}{rrrrrr}", r"\toprule", r"Instancia & Otimo & Instancia & Otimo & Instancia & Otimo \\", r"\midrule"])
    for idx in range(0, len(instances), 3):
        cells = []
        for instance in instances[idx : idx + 3]:
            cells.extend([instance, fmt_pt(best[(instance, "bruteforce")], 3)])
        while len(cells) < 6:
            cells.extend(["", ""])
        lines.append(" & ".join(cells) + r" \\")
    lines.extend([r"\bottomrule", r"\end{tabular}"])
    return "\n".join(lines) + "\n"


def render_bruteforce_gap_table(summaries: list[SummaryRow]) -> str:
    stats = compute_gap_stats(summaries)
    labels = {"aco": "ACO", "ga": "GA", "pso": "PSO"}
    lines = generated_header()
    lines.extend(
        [
            r"\begin{tabular}{lrrrrr}",
            r"\toprule",
            r"Metodo & Runs & Gap medio & Gap min. & Gap max. & Taxa de acerto \\",
            r"\midrule",
        ]
    )
    for method in ("aco", "ga", "pso"):
        row = stats.get(method, {"runs": 0, "mean": 0.0, "min": 0.0, "max": 0.0, "success_rate": 0.0})
        lines.append(
            f"{labels[method]} & {int(row['runs'])} & "
            f"{pct_pt(float(row['mean']), 4)} & {pct_pt(float(row['min']), 4)} & "
            f"{pct_pt(float(row['max']), 4)} & {pct_pt(float(row['success_rate']), 2)} " + r"\\"
        )
    lines.extend([r"\bottomrule", r"\end{tabular}"])
    return "\n".join(lines) + "\n"


def render_large_instance_table(summaries: list[SummaryRow]) -> str:
    grouped = summaries_by_key(summaries)
    instances = sorted({row.instance for row in summaries if instance_sort_key(row.instance)[0] in {50, 100}}, key=instance_sort_key)
    lines = generated_header()
    lines.extend([r"\begin{tabular}{lrrrrrr}", r"\toprule", r"Instancia & ACO best & ACO media & GA best & GA media & PSO best & PSO media \\", r"\midrule"])
    for instance in instances:
        cells = [instance]
        for method in ("aco", "ga", "pso"):
            values = [row.best_makespan for row in grouped.get((instance, method), [])]
            if values:
                cells.extend([fmt_pt(min(values), 4), fmt_pt(statistics.fmean(values), 4)])
            else:
                cells.extend(["---", "---"])
        lines.append(" & ".join(cells) + r" \\")
    lines.extend([r"\bottomrule", r"\end{tabular}"])
    return "\n".join(lines) + "\n"


def render_lower_bound_gap_table(summaries: list[SummaryRow]) -> str:
    stats = compute_lower_bound_stats(summaries)
    lines = generated_header()
    lines.extend(
        [
            r"\begin{tabular}{lr}",
            r"\toprule",
            r"Metrica & Valor \\",
            r"\midrule",
            f"Gap medio LB$\\to$BF & {pct_pt(float(stats['mean']), 2)} " + r"\\",
            f"Gap minimo & {pct_pt(float(stats['min']), 2)} " + r"\\",
            f"Gap maximo & {pct_pt(float(stats['max']), 2)} " + r"\\",
            f"Instancias avaliadas & {int(stats['runs'])} " + r"\\",
            f"Violacoes LB $\\leq$ BF & {int(stats['violations'])} " + r"\\",
            r"\bottomrule",
            r"\end{tabular}",
        ]
    )
    return "\n".join(lines) + "\n"


def render_timing_100_table(summaries: list[SummaryRow], timings: list[TimingRow]) -> str:
    summary_instances = {row.instance for row in summaries}
    target_instances = sorted((inst for inst in summary_instances if instance_sort_key(inst)[0] >= 100), key=instance_sort_key)
    grouped: dict[tuple[str, str], list[float]] = defaultdict(list)
    for row in timings:
        grouped[(row.instance, row.method)].append(row.optimize_ms)

    lines = generated_header()
    lines.extend(
        [
            r"\begin{tabular}{lrrrrr}",
            r"\toprule",
            r"Instancia & ACO ms & GA ms & PSO ms & ACO/GA & ACO/PSO \\",
            r"\midrule",
        ]
    )
    for instance in target_instances:
        aco = mean_or_none(grouped.get((instance, "aco"), []))
        ga = mean_or_none(grouped.get((instance, "ga"), []))
        pso = mean_or_none(grouped.get((instance, "pso"), []))
        values = [number_or_dash(aco), number_or_dash(ga), number_or_dash(pso), ratio_or_dash(aco, ga), ratio_or_dash(aco, pso)]
        lines.append(f"{instance} & " + " & ".join(values) + r" \\")
    lines.extend([r"\bottomrule", r"\end{tabular}"])
    return "\n".join(lines) + "\n"


def mean_or_none(values: list[float]) -> float | None:
    return statistics.fmean(values) if values else None


def number_or_dash(value: float | None) -> str:
    return number(value) if value is not None else "---"


def ratio_or_dash(num: float | None, den: float | None) -> str:
    if num is None or den is None or abs(den) <= 1e-12:
        return "---"
    return f"{fmt_pt(num / den, 2)}x"


def render_manifest(
    results_dir: Path,
    summary_files: list[Path],
    timing_files: list[Path],
    evolution_files: list[Path],
    summaries: list[SummaryRow],
    outputs: dict[str, str],
) -> str:
    output_names = sorted(path for path in outputs if path != "manifest.json")
    output_hashes = {path: sha_text(outputs[path]) for path in output_names}
    instances = sorted({row.instance for row in summaries}, key=instance_sort_key)
    method_counts = {method: sum(1 for row in summaries if row.method == method) for method in ALL_METHODS}
    seed_counts = {
        method: len({row.seed for row in summaries if row.method == method and row.seed is not None})
        for method in META_METHODS
    }
    payload = {
        "script": SCRIPT_NAME,
        "script_version": SCRIPT_VERSION,
        "script_sha256": sha_text((ROOT / "scripts" / SCRIPT_NAME).read_text(encoding="utf-8")),
        "inputs": {
            "results_dir": display_path(results_dir),
            "summary_count": len(summary_files),
            "timing_count": len(timing_files),
            "evolution_count": len(evolution_files),
            "input_digest": sha_files(summary_files + timing_files + evolution_files, results_dir),
        },
        "data": {
            "instances": instances,
            "methods": sorted({row.method for row in summaries}),
            "method_counts": method_counts,
            "seed_counts": seed_counts,
            "valid_summary_rows": len(summaries),
        },
        "outputs": output_names,
        "output_hashes": output_hashes,
    }
    return json.dumps(payload, indent=2, sort_keys=True) + "\n"


def display_path(path: Path) -> str:
    try:
        return path.resolve().relative_to(ROOT).as_posix()
    except ValueError:
        return str(path)


def sha_text(text: str) -> str:
    return hashlib.sha256(text.encode("utf-8")).hexdigest()


def sha_files(files: list[Path], base_dir: Path) -> str:
    digest = hashlib.sha256()
    for path in sorted(files):
        try:
            rel = path.resolve().relative_to(base_dir.resolve()).as_posix()
        except ValueError:
            rel = path.name
        digest.update(rel.encode("utf-8"))
        digest.update(b"\0")
        digest.update(path.read_bytes())
        digest.update(b"\0")
    return digest.hexdigest()


def write_outputs(output_dir: Path, outputs: dict[str, str]) -> None:
    for rel_path, content in outputs.items():
        path = output_dir / rel_path
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(content, encoding="utf-8")


def check_outputs(output_dir: Path, outputs: dict[str, str]) -> int:
    stale: list[str] = []
    for rel_path, expected in outputs.items():
        path = output_dir / rel_path
        if not path.exists():
            stale.append(rel_path)
            continue
        if path.read_text(encoding="utf-8") != expected:
            stale.append(rel_path)
    if stale:
        for rel_path in stale:
            print(f"outdated: {output_dir / rel_path}", file=sys.stderr)
        return 1
    print(f"generated metrics are current in {output_dir}")
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
