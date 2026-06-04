#!/usr/bin/env python3
"""Smoke tests for the monograph validation gate."""

from __future__ import annotations

import subprocess
import unittest
from pathlib import Path


ROOT = Path(__file__).resolve().parents[2]
SCRIPT = ROOT / "scripts" / "check-monografia.sh"


class CheckMonografiaTest(unittest.TestCase):
    def test_validation_gate_passes_current_repository(self) -> None:
        result = subprocess.run(
            ["bash", str(SCRIPT)],
            cwd=ROOT,
            text=True,
            capture_output=True,
            check=False,
        )

        self.assertEqual(result.returncode, 0, result.stderr + result.stdout)
        self.assertIn("PASS metrics", result.stdout)
        self.assertIn("PASS monografia", result.stdout)


if __name__ == "__main__":
    unittest.main()
