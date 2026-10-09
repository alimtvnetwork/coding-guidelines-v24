"""Golden-file tests for ``linter-scripts/check-letterly-format.py``.

Each fixture is a formatter output; the test asserts the validator's verdict.
The FAIL fixtures are the exact misreadings observed in production on
2026-10-09, preserved here as regression tests:

  * ``FENCED`` — output wrapped in a ```markdown fence.
  * ``LITERAL_TITLE`` — title line used the literal words "High Priority
    Instruction" instead of the derived task subject.
  * ``OLD_SLUG_LINE`` — legacy single-line ``slug: <value>`` form.
  * ``SLUG_BEFORE_INPUT`` — slug block placed before the verbatim input.
  * ``MISSING_ITEM`` — fixed actionable item 4 deleted.

Tests are black-box: fixtures are written to temp files and the script is
invoked as a subprocess (exit code + output asserted).
"""
from __future__ import annotations

import subprocess
import sys
import tempfile
import unittest
from pathlib import Path

REPO_ROOT = Path(__file__).resolve().parents[2]
SCRIPT = REPO_ROOT / "linter-scripts" / "check-letterly-format.py"

VALID = """# SEO Writing and Folder Structure Instructions: high priority instruction, non-negotiable task

Hi there. When you started with the SEO writing, before that there was only one or two article write. Follow that article writing pattern.

## slug
SEO Writing and Folder Structure Instructions

# Actionable Items Must Follow Non-Negotiable

1. Write spec under 02-spec/21-app/<slug>/ and enqueue plan task in .ai-memory/plans/<slug>.md (subtasks in .ai-memory/plans/subtasks/<slug>/) first
2. Search codebase exclusively via GitMap (gitmap aum search, gitmap find, gitmap cat, gitmap ps, gitmap py, gitmap llm train); TOTAL BAN on rg, ripgrep, grep, git grep, Select-String
3. Strictly use relative Git paths (02-spec/..., .ai-memory/..., cmd/...); only add the relative paths, never add the absolute path during your work, and ensure this is respected on the release page and in release notes as well
4. Use `gitmap` AI agents to enter data
5. Task completion includes committing and pushing to Git
6. Give every Task-NN in the confirmed task breakdown a `Slug:` sub-item derived from the root slug (`<root-slug> - Task NN`, e.g. `Slug: SEO Writing and Folder Structure Instructions - Task 01`) so GitMap can create and verify subtasks under the root task
7. Follow the existing article writing pattern from the first one or two articles

## Must follow and spawn agent using

[execute-parent-task-with-n-steps-v6](file;.agents/skills/execute-parent-task-with-n-steps-v6)

## Additional Instructions

- [/plan](slashCommand;plan) first before doing the work to reduce the credits.
- [/learn](slashCommand;learn) from [gitmap](file;.agents/skills/gitmap) skill to leverage GitMap high-speed search, toolchain discovery, and caching.
- Only add the relative paths, never add the absolute path during your work; this should be respected on the release page and in release notes as well.
"""

FENCED = "```markdown\n" + VALID + "```\n"

LITERAL_TITLE = VALID.replace(
    "# SEO Writing and Folder Structure Instructions: high priority instruction, non-negotiable task",
    "# High Priority Instruction: high priority instruction, non-negotiable task",
    1,
).replace(
    "## slug\nSEO Writing and Folder Structure Instructions",
    "slug: high-priority-instruction",
    1,
)

OLD_SLUG_LINE = VALID.replace(
    "## slug\nSEO Writing and Folder Structure Instructions",
    "slug: seo-writing-and-folder-structure-instructions",
    1,
)

SLUG_BEFORE_INPUT = VALID.replace(
    """Hi there. When you started with the SEO writing, before that there was only one or two article write. Follow that article writing pattern.

## slug
SEO Writing and Folder Structure Instructions
""",
    """## slug
SEO Writing and Folder Structure Instructions

Hi there. When you started with the SEO writing, before that there was only one or two article write. Follow that article writing pattern.
""",
    1,
)

MISSING_ITEM = "\n".join(
    line for line in VALID.splitlines()
    if not line.startswith("4. Use `gitmap` AI agents to enter data")
) + "\n"


def run_validator(body: str) -> subprocess.CompletedProcess:
    with tempfile.NamedTemporaryFile("w", suffix=".md", delete=False, encoding="utf-8") as fh:
        fh.write(body)
        path = fh.name
    return subprocess.run(
        [sys.executable, str(SCRIPT), path],
        capture_output=True,
        text=True,
    )


class TestCheckLetterlyFormat(unittest.TestCase):
    def test_valid_output_passes(self):
        proc = run_validator(VALID)
        self.assertEqual(proc.returncode, 0, proc.stdout)

    def test_fenced_output_fails(self):
        proc = run_validator(FENCED)
        self.assertEqual(proc.returncode, 1)
        self.assertIn("fenced code block is forbidden", proc.stdout)

    def test_literal_title_fails(self):
        proc = run_validator(LITERAL_TITLE)
        self.assertEqual(proc.returncode, 1)
        self.assertIn('never the literal words "High Priority Instruction"', proc.stdout)

    def test_old_slug_line_fails(self):
        proc = run_validator(OLD_SLUG_LINE)
        self.assertEqual(proc.returncode, 1)
        self.assertIn("single-line form is forbidden", proc.stdout)

    def test_slug_before_input_fails(self):
        proc = run_validator(SLUG_BEFORE_INPUT)
        self.assertEqual(proc.returncode, 1)
        self.assertIn("must come after the verbatim input", proc.stdout)

    def test_missing_fixed_item_fails(self):
        proc = run_validator(MISSING_ITEM)
        self.assertEqual(proc.returncode, 1)
        self.assertIn("fixed actionable item 4 is missing", proc.stdout)


if __name__ == "__main__":
    unittest.main()
