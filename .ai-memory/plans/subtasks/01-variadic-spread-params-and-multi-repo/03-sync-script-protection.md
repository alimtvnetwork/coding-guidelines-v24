# Subtask 03: Multi-Repository Synchronization Script Protection Guards & Memory Safety

> **Task ID:** `01-variadic-spread-params-and-multi-repo-subtask-03`  
> **Parent Task:** `01-variadic-spread-params-and-multi-repo`  
> **Status:** READY FOR EXECUTION  
> **Target Script:** `03-ai-scripts/38-sync-prompts-skills-scripts.py`  
> **Specification Reference:** [02-spec/21-app/01-variadic-spread-params-and-multi-repo/02-sync-and-multi-repo-spec.md](../../../02-spec/21-app/01-variadic-spread-params-and-multi-repo/02-sync-and-multi-repo-spec.md)  
> **Learned Architecture Reference:** [.ai-memory/memory/learned/18-cross-repository-sync-rules.md](../../memory/learned/18-cross-repository-sync-rules.md)  

---

## 1. Objective & Scope

Update and harden the central multi-repository synchronization script `03-ai-scripts/38-sync-prompts-skills-scripts.py` to enforce the four non-negotiable synchronization boundaries. Specifically, implement `is_protected_memory_or_plan` and integrate it into `copy_single_file` and `mirror_directory` so that `.ai-memory/memory/` and `.ai-memory/plans/` in target repositories are permanently shielded from upstream modification, deletion, or clobbering.

---

## 2. The Four Non-Negotiable Sync Boundaries

| Boundary | Target Path Pattern | Enforcement Policy | Guard Identifier |
| :--- | :--- | :--- | :--- |
| **Boundary 1: Spec 21 Exclusion** | `02-spec/21-*`, `21-app*` | TOTAL BAN on sync. Upstream must never mirror or copy into child application specs. | `is_spec_21` |
| **Boundary 2: Additive-Only AI Scripts** | `03-ai-scripts/`, `.agents/scripts/` | Only copy new scripts. Existing child scripts are NEVER overwritten. Target files are NEVER deleted. | `is_additive_only` |
| **Boundary 3: Bump Script Guard** | `*bump*`, `bump-version*` | NEVER overwrite version bump scripts in target repositories. Target owns its release mechanics. | `is_bump_script` |
| **Boundary 4: Memory & Plans Guard** | `.ai-memory/memory/*`, `.ai-memory/plans/*` | If target repository contains memory or plans, NEVER modify, overwrite, or delete them. | `is_protected_memory_or_plan` |

---

## 3. Concrete Code Changes in `03-ai-scripts/38-sync-prompts-skills-scripts.py`

### 3.1 Protection Function: `is_protected_memory_or_plan`

Insert the dedicated discriminator function following the project's Boolean and style principles:
- Positive prefix (`is_...`).
- Implicit boolean evaluation (no explicit `== True`).
- No mixed polarity inside conditionals.
- Mandatory vertical line gaps (blank line before `if`, after `}`, before `return`).

```python
def is_protected_memory_or_plan(path: Path) -> bool:
    """Check if file or directory belongs to .ai-memory/memory/ or .ai-memory/plans/.

    Target repositories own their operational memory logs and execution plans.
    These files must NEVER be overwritten, mirrored, or deleted during sync.
    """
    norm = str(path).replace("\\", "/").lower()

    if ".ai-memory/memory" in norm:
        return True

    if ".ai-memory/plans" in norm:
        return True

    for part in path.parts:
        if part in ("memory", "plans"):
            if ".ai-memory" in norm:
                return True

    return False
```

### 3.2 Update `copy_single_file`

Integrate Boundary 4 evaluation directly into `copy_single_file`:

```python
def copy_single_file(
    src_file: Path,
    dst_file: Path,
    is_dry_run: bool = False,
    is_additive_only: bool = False,
) -> int:
    """Copy a single file respecting the 4 non-negotiable boundaries."""
    if not src_file.exists():
        return 0

    # Boundary 1: Spec 21 Exclusion
    if is_spec_21(src_file) or is_spec_21(dst_file):
        return 0

    # Boundary 4: Memory & Plans Protection
    if is_protected_memory_or_plan(dst_file):
        if dst_file.exists():
            return 0

    # Boundary 3: Bump Script Protection
    if is_bump_script(dst_file):
        if dst_file.exists():
            return 0

    # Boundary 2: Additive-Only AI Scripts
    if is_additive_only:
        if dst_file.exists():
            return 0

    is_copy_needed = False

    if not dst_file.exists():
        is_copy_needed = True
    else:
        try:
            if src_file.stat().st_size != dst_file.stat().st_size:
                is_copy_needed = True
            elif src_file.read_bytes() != dst_file.read_bytes():
                is_copy_needed = True
        except Exception:
            is_copy_needed = True

    if is_copy_needed:
        if not is_dry_run:
            dst_file.parent.mkdir(parents=True, exist_ok=True)
            shutil.copy2(src_file, dst_file)

        return 1

    return 0
```

### 3.3 Update `mirror_directory`

Incorporate `is_protected_memory_or_plan` to guarantee that directory traversals neither delete target memory files during stale cleanup nor write files into protected memory/plans directories:

1. In Stale File Cleanup: Skip files and directories satisfying `is_protected_memory_or_plan(dst_file)`.
2. In Directory Traversal: Exclude child directories matching `is_protected_memory_or_plan(Path(root) / d)`.
3. In File Copy: Intercept target paths matching `is_protected_memory_or_plan(dst_file)` when `dst_file.exists()`.

---

## 4. Step-by-Step Implementation & Verification Plan

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ STEP 1: Code Review & Pre-Check                                             │
│         Inspect 03-ai-scripts/38-sync-prompts-skills-scripts.py             │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ STEP 2: Surgical Code Edit                                                  │
│         Add `is_protected_memory_or_plan` and wire into copy & mirror       │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ STEP 3: Automated Unit Test Suite                                           │
│         Run path discrimination tests across 10 positive & negative cases   │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ STEP 4: Syntax & Lint Verification                                          │
│         Verify zero syntax errors via python -m py_compile                  │
└──────────────────────────────────────┬──────────────────────────────────────┘
                                       │
┌──────────────────────────────────────▼──────────────────────────────────────┐
│ STEP 5: Dry-Run Smoke Validation                                            │
│         Execute against sample repo (e.g., movie-cli) with --dry-run        │
└─────────────────────────────────────────────────────────────────────────────┘
```

### 4.1 Test Cases for Path Guards

Create an isolated test assertion checking:

```python
# Boundary 1: Spec 21
assert is_spec_21(Path("02-spec/21-app/01-spec.md")) is True
assert is_spec_21(Path("02-spec/02-coding-guidelines/04-error.md")) is False

# Boundary 2: Additive Scripts
# (Verified via copy_single_file with is_additive_only=True on existing files)

# Boundary 3: Bump Script
assert is_bump_script(Path("03-ai-scripts/37-bump-version.py")) is True
assert is_bump_script(Path("scripts/bump-version.mjs")) is True
assert is_bump_script(Path("03-ai-scripts/11-fast-file-scanner.py")) is False

# Boundary 4: Memory & Plans
assert is_protected_memory_or_plan(Path(".ai-memory/memory/learned/18-rules.md")) is True
assert is_protected_memory_or_plan(Path(".ai-memory/plans/pending/01-task.md")) is True
assert is_protected_memory_or_plan(Path(".ai-memory/plans/subtasks/01-task.md")) is True
assert is_protected_memory_or_plan(Path(".ai-memory/coding-guidelines.md")) is False
assert is_protected_memory_or_plan(Path(".ai-memory/prompts.md")) is False
```

---

## 5. Definition of Done & Quality Gates

- [ ] Function `is_protected_memory_or_plan` implemented and fully documented in `03-ai-scripts/38-sync-prompts-skills-scripts.py`.
- [ ] `copy_single_file` includes the Boundary 4 guard before any file copy operations.
- [ ] `mirror_directory` excludes protected memory and plan directories from traversal and stale file pruning.
- [ ] Test cases demonstrate 100% expected behavior on both positive and negative boundary inputs.
- [ ] `python -m py_compile 03-ai-scripts/38-sync-prompts-scripts.py` executes with zero errors.
- [ ] Strict relative paths mandate verified; zero absolute paths or URI schemes present anywhere.
