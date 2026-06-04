#!/usr/bin/env python3
"""Tests for deterministic monograph metric generation."""

from __future__ import annotations

import json
import subprocess
import sys
import tempfile
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
SCRIPT = ROOT / "scripts" / "gerar-metricas-monografia.py"


def write_summary(results_dir: Path, run_id: str, method: str, best: float, seed: int) -> None:
    instance = run_id.split("__", 1)[0]
    path = results_dir / "summary" / f"{run_id}.json"
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(
        json.dumps(
            {
                "schema": "tcc.summary.v1",
                "run_id": run_id,
                "method": method,
                "instance": f"/tmp/{instance}.graph",
                "seed": seed,
                "status": "ok",
                "result": {
                    "best_makespan": best,
                    "best_sequence": [1, 2, 3],
                    "iterations_completed": 1,
                    "evaluations": 1,
                },
            },
            indent=2,
            sort_keys=True,
        ),
        encoding="utf-8",
    )


def write_timing(results_dir: Path, run_id: str, method: str, optimize_ms: float, total_ms: float) -> None:
    instance = run_id.split("__", 1)[0]
    path = results_dir / "timing" / f"{run_id}.json"
    path.parent.mkdir(parents=True, exist_ok=True)
    path.write_text(
        json.dumps(
            {
                "schema": "tcc.timing.v1",
                "run_id": run_id,
                "method": method,
                "instance": f"/tmp/{instance}.graph",
                "durations_ms": {"optimize": optimize_ms, "total": total_ms},
            },
            indent=2,
            sort_keys=True,
        ),
        encoding="utf-8",
    )


def seed_fixture(results_dir: Path) -> None:
    rows = [
        ("10a__bruteforce__s0__hbf", "bruteforce", 10.0, 0, 3.0, 30.0),
        ("10a__lowerbound__s0__hlb", "lowerbound", 5.2, 0, 1.0, 10.0),
        ("10a__ga__s0__hga0", "ga", 10.5, 0, 7.0, 70.0),
        ("10a__ga__s1__hga1", "ga", 10.7, 1, 8.0, 80.0),
        ("10a__pso__s0__hpso", "pso", 11.0, 0, 9.0, 90.0),
        ("10a__aco__s0__haco", "aco", 10.0, 0, 6.0, 60.0),
        ("100a__lowerbound__s0__hlb", "lowerbound", 50.0, 0, 1.0, 10.0),
        ("100a__ga__s0__hga", "ga", 65.0, 0, 7.0, 70.0),
        ("100a__pso__s0__hpso", "pso", 70.0, 0, 9.0, 90.0),
        ("100a__aco__s0__haco", "aco", 62.0, 0, 11.0, 110.0),
    ]
    for run_id, method, best, seed, optimize_ms, total_ms in rows:
        write_summary(results_dir, run_id, method, best, seed)
        write_timing(results_dir, run_id, method, optimize_ms, total_ms)


class GerarMetricasMonografiaTest(unittest.TestCase):
    def run_script(self, results_dir: Path, output_dir: Path, *extra: str) -> subprocess.CompletedProcess[str]:
        return subprocess.run(
            [
                sys.executable,
                str(SCRIPT),
                "--results-dir",
                str(results_dir),
                "--output-dir",
                str(output_dir),
                *extra,
            ],
            cwd=ROOT,
            text=True,
            capture_output=True,
            check=False,
        )

    def test_generates_metrics_manifest_and_table_fragments(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            tmp_path = Path(tmp)
            results_dir = tmp_path / "results"
            output_dir = tmp_path / "generated"
            seed_fixture(results_dir)

            result = self.run_script(results_dir, output_dir)

            self.assertEqual(result.returncode, 0, result.stderr)
            metrics = (output_dir / "metrics.tex").read_text(encoding="utf-8")
            self.assertIn(r"\newcommand{\TotalExecucoes}{10}", metrics)
            self.assertIn(r"\newcommand{\TotalInstancias}{2}", metrics)
            self.assertIn(r"\newcommand{\TotalInstanciasComOtimo}{1}", metrics)
            self.assertIn(r"\newcommand{\TotalInstanciasGrandes}{1}", metrics)
            self.assertIn(r"\newcommand{\GapMedioOtimoGA}{6{,}0000\%}", metrics)
            self.assertIn(r"\newcommand{\TaxaAcertoOtimoACO}{100{,}00\%}", metrics)

            latex_eol = r"\\"
            coverage = (output_dir / "tables" / "tab-cobertura-experimental.tex").read_text(encoding="utf-8")
            self.assertIn("GA & 3 & 2 & 2 " + latex_eol, coverage)
            self.assertIn("Busca exaustiva & 1 & 1 & 1 " + latex_eol, coverage)

            gap_bf = (output_dir / "tables" / "tab-gap-bf.tex").read_text(encoding="utf-8")
            self.assertIn("GA & 2 & 6{,}0000\\% & 5{,}0000\\% & 7{,}0000\\% & 0{,}00\\% " + latex_eol, gap_bf)

            timing = (output_dir / "tables" / "tab-tempo-100.tex").read_text(encoding="utf-8")
            self.assertIn("100a & 11{,}00 & 7{,}00 & 9{,}00 & 1{,}57x & 1{,}22x " + latex_eol, timing)

            manifest = json.loads((output_dir / "manifest.json").read_text(encoding="utf-8"))
            self.assertEqual(manifest["script"], "gerar-metricas-monografia.py")
            self.assertEqual(manifest["inputs"]["summary_count"], 10)
            self.assertIn("script_sha256", manifest)
            self.assertGreaterEqual(set(manifest["outputs"]), {"metrics.tex", "tables/tab-cobertura-experimental.tex", "tables/tab-gap-bf.tex", "tables/tab-tempo-100.tex"})

    def test_check_mode_detects_missing_and_stale_outputs(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            tmp_path = Path(tmp)
            results_dir = tmp_path / "results"
            output_dir = tmp_path / "generated"
            seed_fixture(results_dir)

            missing = self.run_script(results_dir, output_dir, "--check")
            self.assertNotEqual(missing.returncode, 0)

            generated = self.run_script(results_dir, output_dir)
            self.assertEqual(generated.returncode, 0, generated.stderr)

            current = self.run_script(results_dir, output_dir, "--check")
            self.assertEqual(current.returncode, 0, current.stderr)

            write_summary(results_dir, "100a__aco__s0__haco", "aco", 61.0, 0)
            stale = self.run_script(results_dir, output_dir, "--check")
            self.assertNotEqual(stale.returncode, 0)
            self.assertIn("outdated", stale.stderr + stale.stdout)


if __name__ == "__main__":
    unittest.main()
