# Subtask 1: Update AppError Display Specs

1. Open 02-spec/03-error-manage/02-error-architecture/06-apperror-package/01-apperror-reference/02-apperror-struct.md.
2. Expand section 2.3 Display Methods into sub-categories: "Human-Facing Output", "Structured Logger Output", and "Developer Output".
3. Define HumanString() string - returns DisplayError or falls back to a generic safe string.
4. Define LogMap() map[string]any - returns structured key-values (Code, Message, Values, Diagnostics) for use with structured loggers.
5. Define ConsoleString() string (or alias to FullString()) for full terminal prints.
6. Update 02-spec/03-error-manage/02-error-architecture/06-apperror-package/09-ai-action-plan-apperror-methods.md to instruct the AI to write these specific formatting methods.
