"""CODE-RED-005 — shared helpers for the prefer-8 function-length check.

Pairs with CODE-RED-004 (hard cap, 15 lines). Emits SARIF `warning`
findings when a function body has >PREFER and <=HARD effective lines.

Sibling CODE-RED-004 modules live in ``checks/function-length/`` whose
hyphenated path is not importable directly; ``load_sibling`` loads them
by file path so we can reuse their regex patterns and counters.
"""

from __future__ import annotations

import importlib.util
from pathlib import Path
from types import ModuleType

from _lib.sarif import Finding, Rule
from _lib.walker import relpath


PREFER_LINES = 8
HARD_LINES = 15

RULE = Rule(
    id="CODE-RED-005",
    name="FunctionLengthPrefer8",
    short_description="Prefer function bodies <= 8 effective lines (hard cap 15).",
    help_uri_relative="01-cross-language/04-code-style/00-overview.md",
)


def load_sibling(language: str) -> ModuleType:
    here = Path(__file__).resolve().parent.parent
    target = here / "function-length" / f"{language}.py"
    spec = importlib.util.spec_from_file_location(f"_fl_{language}", target)
    if spec is None or spec.loader is None:
        raise ImportError(f"Cannot load sibling module: {target}")
    module = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(module)
    return module


def is_in_prefer_band(effective: int) -> bool:
    return effective > PREFER_LINES and effective <= HARD_LINES


def make_finding(name: str, effective: int, path: Path, root: str, start_line: int) -> Finding:
    msg = (
        f"Function '{name}' has {effective} effective lines "
        f"(prefer <= {PREFER_LINES}, hard cap {HARD_LINES})."
    )
    return Finding(
        rule_id=RULE.id,
        level="warning",
        message=msg,
        file_path=relpath(path, root),
        start_line=start_line,
    )
