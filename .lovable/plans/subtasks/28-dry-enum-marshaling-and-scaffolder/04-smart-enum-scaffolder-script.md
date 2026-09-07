# Subtask 28.4: Rebuild Smart Python Enum Scaffolder Script

## Context
Rebuild `03-ai-scripts/30-enum-generator.py` to scaffold complete, idiomatic, 4-file Go enum packages matching all repository standards and using `baseenumer` DRY helpers.

## Requirements
1. Multi-file generation:
   - `variant.go`: Types, constants, predicates, DRY JSON marshaling.
   - `vars.go`: Labels, `baseenumer.CompileMap`, `All()`, `Values()`, `Parse() Result`.
   - `variant_test.go`: Complete unit tests (interfaces, properties, predicates, names, parse, JSON roundtrip).
   - `readme.md`: Clean Markdown documentation catalog.
2. 3 backing types:
   - `byte` (canonical default, `Variant byte`)
   - `int` / `uint16` (`Variant int` or `Variant uint16`)
   - `string` (`Variant string`)
3. 2 input modes:
   - CLI flags: `--name`, `--type`, `--items` (comma-separated), `--package`, `--target-dir`
   - JSON config: `--config` (path to json file) or `--json` (inline json string)
4. Flags:
   - `--dry-run`: Preview files and paths without writing
   - `--overwrite`: Overwrite existing files
5. Code standards:
   - Strictly lowercase filenames.
   - Python functions <= 15 lines.
   - Generated Go code functions <= 15 lines.
   - Implicit booleans only.
   - Strictly relative Git paths in docs and output.

## Verification
- Test CLI dry-run and scaffolding on a temporary test enum.
- Format with `python 03-ai-scripts/26-go-code-formatter.py`.
- Run tests on generated package.
