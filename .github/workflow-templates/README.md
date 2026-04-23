# Workflow Templates — Reusable CI Guards

Three layers, choose what fits:

| Layer | File | When to use |
|-------|------|-------------|
| Composite action | [`action.yml`](./action.yml) | Drop into any existing job as a single `uses:` step |
| Reusable workflow | [`ci-guards.reusable.yml`](./ci-guards.reusable.yml) | Call from another workflow with `workflow_call` |
| Starter workflows | [`starters/`](./starters/) | Copy/paste a complete workflow per language |

All three ultimately call `bash scripts/ci-runner.sh`, which dispatches the
six guards from `spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/`.

---

## Quick start (composite action)

```yaml
jobs:
  guards:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: ./.github/workflow-templates
        with:
          phase: all
          config: ci-guards.yaml
```

## Quick start (reusable workflow)

```yaml
jobs:
  guards:
    uses: <org>/<repo>/.github/workflow-templates/ci-guards.reusable.yml@v1
    with:
      phase: all
      languages: "go"
      file-glob: "*.go"
      rule-set: "go-default"
```

## Inputs

| Input | Default | Notes |
|-------|---------|-------|
| `phase` | `all` | `check` \| `lint` \| `test` \| `all` |
| `guard` | _(empty)_ | Single guard name; overrides `phase` |
| `config` | `ci-guards.yaml` | Skipped if file does not exist |
| `languages` | _(empty)_ | Comma-separated hints, exposed as `CI_GUARDS_LANGUAGES` |
| `file-glob` | _(empty)_ | Exposed as `CI_GUARDS_FILE_GLOB` |
| `rule-set` | _(empty)_ | Exposed as `CI_GUARDS_RULESET` |
| `source-dir` / `baseline` / `results-dir` | _(empty)_ | Direct overrides for the runner |
| `scripts-dir` | `.github/scripts` | Where guard scripts live |
| `fail-on-violation` | `true` | When `false`, exit code `1` is downgraded to `0` |
| `node-version` | `20` | Used to run the config loader |

## Outputs

| Output | Description |
|--------|-------------|
| `exit-code` | Raw exit code from `ci-runner.sh` (`0`/`1`/`2`/`64`) |
| `summary-path` | Path to the uploaded JSON summary artifact |

---

## Spec

- [09-workflow-templates.md](../../spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/09-workflow-templates.md) — full reference + design rationale
- [07-shared-cli-wrapper.md](../../spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/07-shared-cli-wrapper.md) — runner contract
- [08-config-schema.md](../../spec/12-cicd-pipeline-workflows/03-reusable-ci-guards/08-config-schema.md) — `ci-guards.yaml` schema