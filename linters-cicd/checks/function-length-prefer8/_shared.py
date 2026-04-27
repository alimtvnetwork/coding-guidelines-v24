"""CODE-RED-005 — shared helpers for the prefer-8 function-length check.

Pairs with CODE-RED-004 (hard cap, 15 lines). Emits SARIF `warning`
findings when a function body has >PREFER and <=HARD effective lines.
"""

from __future__ import annotations

from pathlib import Path

from _lib.sarif import Finding, Rule


PREFER_LINES = 8
HARD_LINES = 15

RULE = Rule(
    id="CODE-RED-005",
    name="FunctionLengthPrefer8",
    short_description="Prefer function bodies <= 8 effective lines (hard cap 15).",
    help_uri_relative="01-cross-language/04-code-style/00-overview.md",
)


def make_finding(name: str, effective: int, path: Path, root: str, start_line: int) -> Finding:
    msg = (
        f"Function '{name}' has {effective} effective lines "
        f"(prefer <= {PREFER_LINES}, hard cap {HARD_LINES})."
    )
    return Finding(
        rule_id=RULE.id,
        level="warning",
        message=msg,
        file_path=_relpath(path, root),
        start_line=start_line,
    )


def is_in_prefer_band(effective: int) -> bool:
    return effective > PREFER_LINES and effective <= HARD_LINES


def _relpath(path: Path, root: str) -> str:
    from _lib.walker import relpath
    return relpath(path, root)
