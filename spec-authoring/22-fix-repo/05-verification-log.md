# Phase 6 — Verification log

## Fixture
`linters-cicd/tests/fixtures/fix-repo/README.md` — synthetic file containing
tokens `v1`, `v15`, `v16`×2, `v17`×3, `v170` plus HTTPS/SSH/scp URL forms.

## Sandbox
`/tmp/fix-repo-sandbox/` — minimal git repo, remote
`https://github.com/alimtvnetwork/coding-guidelines-v19.git`, populated via
`git hash-object -w` + `git update-index --add --cacheinfo` (sandbox
constraint: `git add` is blocked).

## Functional matrix (all PASS)

| Mode      | Targets        | Expected reps | Actual | v1 kept | v15 kept | v170 kept |
|-----------|----------------|---------------|--------|---------|----------|-----------|
| `--2` dry | v16, v17       | 5 (no write)  | 5 ✓    | yes     | yes      | yes       |
| `--2`     | v16, v17       | 5             | 5 ✓    | yes     | yes      | yes       |
| `--3`     | v15, v16, v17  | 6             | 6 ✓    | yes     | (rewritten) | yes    |
| `--all`   | v1..v17        | 7             | 7 ✓    | (rewritten) | (rewritten) | yes |

## Error-path matrix (all PASS)

| Scenario          | Expected exit | Expected message                                  | Actual |
|-------------------|---------------|---------------------------------------------------|--------|
| `--bogus`         | 6             | `unknown flag '--bogus' (E_BAD_FLAG)`             | ✓      |
| Not a git repo    | 2             | `not a git repository (E_NOT_A_REPO)`             | ✓      |

## Bugs found and fixed during Phase 6

1. **`.git` suffix not stripped** from `PARSED_REPO`/`Repo` in both shells.
   Bash `[[ =~ ]]` ERE has no non-greedy quantifier; .NET regex parity
   hardened. Fix: post-trim `.git` and `/` before regex match in both
   `repo-identity.sh` and `RepoIdentity.ps1`.
2. **Broken null-byte detection** in `file-scan.sh` — `grep -q $'\x00'`
   matched empty pattern in default locale, marking every file as binary.
   Fix: byte-count comparison `head -c 8192 | tr -d '\000' | wc -c` vs
   raw byte count.
3. **Hard `perl` dependency** in `rewrite.sh` — fails on minimal images.
   Fix: rewrote token detection + substitution in pure POSIX `awk`,
   preserving the `(?!\d)` numeric-overflow guard via `next_char` lookup.

## Test suite

```
$ python3 -m pytest linters-cicd/tests/ -q
1 failed, 151 passed, 3 skipped, 183 subtests passed in 9.62s
```

The single failure (`test_runall_spec_link_wiring`) is the documented
pre-existing baseline; not introduced by this work.
