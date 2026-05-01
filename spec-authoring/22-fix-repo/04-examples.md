# 04 — Worked Examples

All examples assume `git remote get-url origin` returns
`https://github.com/alimtvnetwork/coding-guidelines-v20.git`.

## Default mode (last 2)

```
$ ./fix-repo.sh --dry-run
fix-repo  base=coding-guidelines  current=v18  mode=--2
targets:  v16, v17
host:     github.com  owner=alimtvnetwork

scanned: 614 files
changed: 0 files (0 replacements)
mode:    dry-run
```

## Three-version sweep

```
$ ./fix-repo.sh -3 --verbose
fix-repo  base=coding-guidelines  current=v18  mode=--3
targets:  v15, v16, v17
host:     github.com  owner=alimtvnetwork

modified: spec/02-coding-guidelines/06-cicd-integration/01-sarif-contract.md (4 replacements)
modified: readme.md (12 replacements)
...

scanned: 614 files
changed: 87 files (542 replacements)
mode:    write
```

## Full sweep

```
$ ./fix-repo.sh --all
fix-repo  base=coding-guidelines  current=v18  mode=--all
targets:  v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17
host:     github.com  owner=alimtvnetwork

scanned: 614 files
changed: 87 files (542 replacements)
mode:    write
```

## PowerShell parity

```
PS> .\fix-repo.ps1 -3 -DryRun -Verbose
PS> .\fix-repo.ps1 -all
PS> .\fix-repo.ps1            # default = last 2
```

## Detection failures

```
$ ./fix-repo.sh
fix-repo: ERROR no -vN suffix on repo name 'my-app' (E_NO_VERSION_SUFFIX)
$ echo $?
4
```

```
$ cd /tmp/not-a-repo && /path/to/fix-repo.sh
fix-repo: ERROR not a git repository (E_NOT_A_REPO)
$ echo $?
2
```

## URL preservation

Input file `docs/install.md`:
```
curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v20/main/install.sh | bash
```

After `./fix-repo.sh`:
```
curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v20/main/install.sh | bash
```

Host (`raw.githubusercontent.com`), owner (`alimtvnetwork`), and the
rest of the path (`/main/install.sh`) are preserved automatically —
only the `coding-guidelines-v20` segment changes because that
segment is the literal token.

## Numeric-overflow guard in action

Input:
```
The legacy coding-guidelines-v170 channel is unrelated.
See coding-guidelines-v20/spec.md for details.
```

After `./fix-repo.sh`:
```
The legacy coding-guidelines-v170 channel is unrelated.
See coding-guidelines-v20/spec.md for details.
```

`v170` is left alone; `v17` is rewritten.
