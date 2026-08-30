# Subtask 1: Update AppError Spec

1. Open spec/03-error-manage/02-error-architecture/06-apperror-package/01-apperror-reference/02-apperror-struct.md.
2. Add a new subsection defining the Apperror.New creator namespace and the 5 new methods. 
3. Explicitly document that Error, UsingErrorMsg, ErrorVar, and ErrorVars MUST return 
il if the provided err argument is 
il.
4. Open spec/03-error-manage/02-error-architecture/06-apperror-package/09-ai-action-plan-apperror-methods.md.
5. Add explicit instructions for the AI on how to implement this Apperror.New.* pattern (e.g. creating an empty struct 	ype creator struct{} and a global var New = creator{}).
6. Validate formatting (MD022, MD032).
