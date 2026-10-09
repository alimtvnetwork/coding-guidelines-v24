#!/usr/bin/env python3
"""Mechanical validator for the desktop Letterly formatter output shape.

The Letterly/Muse prompts describe the required output in prose, and prose is
lossy: the same misreadings recur (fenced output, literal "High Priority
Instruction" as the title, slug on the wrong line / in the wrong case). This
script enforces the shape mechanically instead of hoping the formatter got it
right. It is the executable counterpart of
``01-prompts/22-letterly/02-desktop-letterly.md`` ("Output Format").

Checks, in order:
  1. No fenced code blocks anywhere (``` fence is forbidden, even ```markdown).
  2. First content line is the single-line title:
     ``# <task title>: high priority instruction, non-negotiable task``
     where <task title> is derived from the input subject and is NEVER the
     literal words "High Priority Instruction".
  3. Slug block: a bare ``## slug`` subheader line (no colon, no trailing
     value), with the slug VALUE on the next non-empty line in Title Case
     (never ``high-priority-instruction``, never "High Priority Instruction").
  4. Slot order: title < ``## slug`` < ``# Actionable Items`` <
     ``## Must follow`` < ``## Additional Instructions``.
  5. Fixed actionable items 1-6 present verbatim (item 1 allows a substituted
     slug in place of the ``<slug>`` placeholder).
  6. Mandatory agent-invocation suffix link present.
  7. ``## Additional Instructions`` carries the /plan and /learn directives.

Usage:
    gitmap py linter-scripts/check-letterly-format.py <formatted-output.md>
    gitmap py linter-scripts/check-letterly-format.py --stdin < formatted-output.md

Exit code 0 when the output conforms, 1 otherwise (violations on stdout).
"""

from __future__ import annotations

import re
import sys

TITLE_RE = re.compile(r"^# (.+): high priority instruction, non-negotiable task$")
FENCE_RE = re.compile(r"^\s*```")
ACTIONABLE_HEADER = "# Actionable Items Must Follow Non-Negotiable"
MUST_FOLLOW_HEADER = "## Must follow and spawn agent using"
ADDITIONAL_HEADER = "## Additional Instructions"
AGENT_LINK = "[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)"

FIXED_ITEM_PREFIXES = [
    "1. Write spec under 02-spec/21-app/",
    "2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String",
    "3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well",
    "4. Use `gitmap` AI agents to enter data",
    "5. Task completion includes committing and pushing to Git",
    "6. Give every Task-NN in the confirmed task breakdown a `Slug:` sub-item",
]


def _nonempty_lines(text):
    return [(i + 1, line) for i, line in enumerate(text.splitlines()) if line.strip()]


def check_no_fences(numbered):
    violations = []
    for lineno, line in numbered:
        if FENCE_RE.match(line):
            violations.append(f"line {lineno}: fenced code block is forbidden (found `{line.strip()}`)")
    return violations


def check_title(numbered):
    violations = []
    if not numbered:
        return ["empty output: missing title line"]
    lineno, first = numbered[0]
    match = TITLE_RE.match(first.strip())
    if not match:
        return [f"line {lineno}: first line must be `# <task title>: high priority instruction, non-negotiable task`"]
    title = match.group(1).strip()
    if not title:
        violations.append(f"line {lineno}: task title is empty")
    if title.lower() == "high priority instruction":
        violations.append(
            f"line {lineno}: title must be derived from the input subject, "
            'never the literal words "High Priority Instruction"'
        )
    return violations


def _find_header(numbered, header):
    for lineno, line in numbered:
        if line.strip() == header:
            return lineno
    return None


def check_slug_block(numbered):
    violations = []
    slug_lineno = None
    for lineno, line in numbered:
        stripped = line.strip()
        if stripped == "## slug":
            slug_lineno = lineno
            break
        if stripped.startswith("## slug"):
            violations.append(
                f"line {lineno}: slug label must be a bare `## slug` subheader "
                f"(no colon, no trailing value; found `{stripped}`)"
            )
            slug_lineno = lineno
            break
        if re.match(r"^slug\s*:", stripped):
            violations.append(
                f"line {lineno}: old `slug: <value>` single-line form is forbidden; "
                "use a `## slug` subheader with the value on the next line"
            )
    if slug_lineno is None and not violations:
        return ["missing `## slug` subheader"], None
    value = None
    if slug_lineno is not None and not any(v.startswith(f"line {slug_lineno}: old") for v in violations):
        after = [(n, l) for n, l in numbered if n > slug_lineno]
        if not after:
            violations.append("`## slug` subheader has no value line after it")
        else:
            value_lineno, value = after[0][0], after[0][1].strip()
            if not value:
                violations.append(f"line {value_lineno}: slug value line is empty")
            elif value.lower() == "high priority instruction":
                violations.append(
                    f"line {value_lineno}: slug must be the Title Case task title, "
                    'never the literal words "High Priority Instruction"'
                )
            elif re.fullmatch(r"[a-z0-9]+(?:-[a-z0-9]+)*", value):
                violations.append(
                    f"line {value_lineno}: slug value must be Title Case "
                    f"(found hyphenated-lowercase `{value}`)"
                )
    return violations, slug_lineno


def check_slot_order(numbered, slug_lineno):
    violations = []
    title_lineno = numbered[0][0] if numbered else None
    actionable = _find_header(numbered, ACTIONABLE_HEADER)
    must_follow = _find_header(numbered, MUST_FOLLOW_HEADER)
    additional = _find_header(numbered, ADDITIONAL_HEADER)
    for name, lineno in [
        ("`# Actionable Items Must Follow Non-Negotiable`", actionable),
        ("`## Must follow and spawn agent using`", must_follow),
        ("`## Additional Instructions`", additional),
    ]:
        if lineno is None:
            violations.append(f"missing {name} section")
    ordered = [
        ("title line", title_lineno),
        ("`## slug`", slug_lineno),
        ("`# Actionable Items`", actionable),
        ("`## Must follow`", must_follow),
        ("`## Additional Instructions`", additional),
    ]
    present = [(name, n) for name, n in ordered if n is not None]
    for (prev_name, prev_n), (name, n) in zip(present, present[1:]):
        if n < prev_n:
            violations.append(f"slot order violated: {name} (line {n}) appears before {prev_name} (line {prev_n})")
    if title_lineno is not None and slug_lineno is not None:
        between = [
            line for n, line in numbered
            if title_lineno < n < slug_lineno
            and line.strip() and not line.strip().startswith("#")
        ]
        if not between:
            violations.append(
                f"`## slug` (line {slug_lineno}) must come after the verbatim input, "
                f"not directly after the title line (line {title_lineno})"
            )
    return violations


def check_fixed_items(numbered, actionable_lineno):
    violations = []
    if actionable_lineno is None:
        return violations
    items = {}
    for lineno, line in numbered:
        if lineno <= actionable_lineno:
            continue
        stripped = line.strip()
        m = re.match(r"^([1-6])\.\s", stripped)
        if m:
            items[int(m.group(1))] = (lineno, stripped)
        if stripped.startswith("## "):
            break
    for idx, prefix in enumerate(FIXED_ITEM_PREFIXES, start=1):
        if idx not in items:
            violations.append(f"fixed actionable item {idx} is missing")
            continue
        lineno, text = items[idx]
        if idx == 1:
            if not text.startswith(prefix) or "enqueue plan task in .ai-memory/plans/" not in text:
                violations.append(f"line {lineno}: fixed item 1 text was modified")
        elif not text.startswith(prefix):
            violations.append(f"line {lineno}: fixed item {idx} text was modified")
    return violations


def check_suffix_and_instructions(text):
    violations = []
    if AGENT_LINK not in text:
        violations.append("missing mandatory agent-invocation suffix link (execute-parent-task-with-n-steps-v6)")
    if "[/plan](slashCommand;plan)" not in text:
        violations.append("`## Additional Instructions` is missing the /plan directive")
    if "[/learn](slashCommand;learn)" not in text:
        violations.append("`## Additional Instructions` is missing the /learn directive")
    return violations


def validate(text):
    numbered = _nonempty_lines(text)
    violations = []
    violations.extend(check_no_fences(numbered))
    violations.extend(check_title(numbered))
    slug_violations, slug_lineno = check_slug_block(numbered)
    violations.extend(slug_violations)
    violations.extend(check_slot_order(numbered, slug_lineno))
    actionable_lineno = _find_header(numbered, ACTIONABLE_HEADER)
    violations.extend(check_fixed_items(numbered, actionable_lineno))
    violations.extend(check_suffix_and_instructions(text))
    return violations


def main(argv):
    if "--stdin" in argv:
        text = sys.stdin.read()
    elif len(argv) == 2:
        with open(argv[1], "r", encoding="utf-8") as fh:
            text = fh.read()
    else:
        print(__doc__.strip().splitlines()[-6], file=sys.stderr)
        print("usage: check-letterly-format.py <file> | --stdin", file=sys.stderr)
        return 2
    violations = validate(text)
    if not violations:
        print("OK: Letterly output conforms to the required format.")
        return 0
    print(f"FAIL: {len(violations)} violation(s):")
    for v in violations:
        print(f"  - {v}")
    return 1


if __name__ == "__main__":
    sys.exit(main(sys.argv))
