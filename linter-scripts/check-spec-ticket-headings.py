#!/usr/bin/env python3
"""Require the ten executable-spec-ticket headings, in order."""

import sys
from pathlib import Path

TICKET_TITLE = "# Executable spec ticket"
HEADING_CONTEXT = "## Context"
HEADING_DO_NOT_TOUCH = "## Do not touch"
PASS_TEXT = "PASS"
MISSING_LABEL = "missing"
OUT_OF_ORDER_LABEL = "out of order"
DEFAULT_TICKET = "02-spec/01-spec-authoring-guide/15-executable-spec-ticket.md"
REQUIRED_HEADINGS = (
    "## Context",
    "## Current state",
    "## Proposed change",
    "## Acceptance criteria",
    "## Testing plan",
    "## Rollback",
    "## Files",
    "## Out of scope",
    "## Do not touch",
    "## Checklist",
)


def is_exact_line(lines: list[str], text: str) -> bool:
    return text in lines


def is_ticket_text(lines: list[str]) -> bool:
    has_title = is_exact_line(lines, TICKET_TITLE)
    has_context = is_exact_line(lines, HEADING_CONTEXT)
    has_guard = is_exact_line(lines, HEADING_DO_NOT_TOUCH)
    has_pair = has_context and has_guard

    return has_title or has_pair


def find_heading_index(lines: list[str], heading: str) -> int:
    index = 0

    while index < len(lines):
        if lines[index] == heading:

            return index

        index += 1

    return -1


def heading_fault(lines: list[str], heading: str, last_index: int) -> str:
    index = find_heading_index(lines, heading)

    if index < 0:

        return MISSING_LABEL

    if index <= last_index:

        return OUT_OF_ORDER_LABEL

    return ""


def next_index(lines: list[str], heading: str, last_index: int) -> int:
    index = find_heading_index(lines, heading)

    if index <= last_index:

        return last_index

    return index


def collect_faults(raw_path: str, lines: list[str]) -> list[str]:
    faults: list[str] = []
    last_index = -1

    for heading in REQUIRED_HEADINGS:
        fault = heading_fault(lines, heading, last_index)
        if fault:
            faults.append(f"{raw_path}: {fault} {heading}")
        last_index = next_index(lines, heading, last_index)

    return faults


def faults_for_path(raw_path: str) -> list[str]:
    lines = Path(raw_path).read_text(encoding="utf-8").splitlines()

    if is_ticket_text(lines):

        return collect_faults(raw_path, lines)

    return []


def resolve_paths(argv: list[str]) -> list[str]:
    has_paths = len(argv) > 0

    if has_paths:

        return argv

    return [DEFAULT_TICKET]


def main() -> int:
    faults: list[str] = []

    for raw_path in resolve_paths(sys.argv[1:]):
        faults.extend(faults_for_path(raw_path))

    if faults:
        print("\n".join(faults))

        return 1

    print(PASS_TEXT)

    return 0


if __name__ == "__main__":
    sys.exit(main())
