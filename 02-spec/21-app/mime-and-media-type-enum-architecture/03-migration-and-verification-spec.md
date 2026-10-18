# Migration & Verification Specification: MIME & Media Type Enum Architecture

## 1. Migration Strategy & Phase Execution

### Phase 1: Planning & Architecture Specification (Task-01)
- Create cross-language specification: `02-spec/02-coding-guidelines/01-cross-language/36-mime-and-media-type-enum-standards.md`.
- Register spec in `02-spec/02-coding-guidelines/01-cross-language/readme.md`.
- Author application specification suite: `02-spec/21-app/mime-and-media-type-enum-architecture/` (`01-architecture-spec.md`, `02-component-spec.md`, `03-migration-and-verification-spec.md`, `readme.md`).
- Register application in `02-spec/21-app/readme.md`.
- Create parent plan `.ai-memory/plans/mime-and-media-type-enum-architecture.md` and 4 subtask plans in `.ai-memory/plans/subtasks/mime-and-media-type-enum-architecture/`.

### Phase 2: CG Execute Prompts Enhancement (Task-02)
- Update `01-prompts/15-cg-execute/06-constants-and-enums.md` and `01-prompts/v4/15-cg-execute/06-constants-and-enums.md` with explicit MIME/Media type requirements.
- Update `01-prompts/15-cg-execute/readme.md` and `01-prompts/v4/15-cg-execute/readme.md`.

### Phase 3: Companion Skills Enhancement (Task-03)
- Update `.agents/skills/cg-enum-standards/skill.md` and `.cursor/skills/cg-enum-standards/skill.md`.
- Update `.agents/skills/multi-language-enums-and-traits/skill.md` and `.cursor/skills/multi-language-enums-and-traits/skill.md`.

### Phase 4: Fleet Synchronization & Verification (Task-04)
- Synchronize prompts, skills, and specs across all 43 fleet repositories using `python 03-ai-scripts/38-sync-prompts-skills-scripts.py`.
- Consolidate completed subtasks into `.ai-memory/plans/completed/mime-and-media-type-enum-architecture.md` and update `.ai-memory/plans/readme.md`.
- Execute targeted linters:
  - `python linter-scripts/check-relative-paths.py`
  - `python linter-scripts/check-forbidden-strings.py`
  - `python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language --check-only`
- Commit and push atomically via GitMap:
  `gitmap cpf "cg-execute - add mime and media type enum standard and sync skills"`

---

## 2. Verification Commands & Acceptance Criteria

### Verification Commands
```bash
# 1. Relative paths verification
python linter-scripts/check-relative-paths.py

# 2. Forbidden strings verification
python linter-scripts/check-forbidden-strings.py

# 3. Guideline autofixer verification
python 03-ai-scripts/05-guideline-autofixer.py 02-spec/02-coding-guidelines/01-cross-language --check-only
```

### Acceptance Criteria
- [ ] Spec 36 created and registered in `02-spec/02-coding-guidelines/01-cross-language/readme.md`.
- [ ] App spec suite created and registered in `02-spec/21-app/readme.md`.
- [ ] `06-constants-and-enums.md` prompts updated in both `01-prompts/15-cg-execute/` and `01-prompts/v4/15-cg-execute/`.
- [ ] Companion skills updated in `.agents/skills/` and `.cursor/skills/`.
- [ ] Fleet synchronization executed across all repositories without errors.
- [ ] All relative path and forbidden string linters exit 0.
- [ ] Clean working tree and atomic commit pushed upstream.
