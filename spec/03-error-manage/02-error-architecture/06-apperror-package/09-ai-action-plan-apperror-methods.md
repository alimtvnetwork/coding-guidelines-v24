# Action Plan for AI: Implementing Contact, Errors, and Msg in AppError

> **Target:** pperror package (e.g., gitlab.com/auk-go/core/apperror)
> **Goal:** Extend AppError with new fields and fluent setters as requested by the spec.

When an AI agent is instructed to implement the new AppError methods (contact, rrors, msg), the agent MUST follow these exact steps to ensure spec compliance.

---

## 1. Modify the AppError Struct
**File to edit:** pperror.go (or wherever AppError is defined)

Add the new fields to the AppError struct, matching the JSON tags required by the spec:

`go
type AppError struct {
    Code         string            json:"Code"
    Message      string            json:"Message"
    DisplayError string            json:"DisplayError,omitempty"
    Details      string            json:"Details,omitempty"
    Contact      string            json:"Contact,omitempty"      // NEW
    Values       map[string]string json:"Values,omitempty"
    Diagnostic   ErrorDiagnostic   json:"Diagnostic,omitempty"
    Errors       []error           json:"Errors,omitempty"       // NEW
    Stack        StackTrace        json:"Stack"
    Cause        error             json:"-"
}
`
*(Note: Ensure you do not accidentally overwrite existing JSON tags if they are explicitly declared in the codebase.)*

## 2. Implement the Fluent Setters
**File to edit:** pperror_setters.go (or append to pperror.go)

Create the fluent setter methods for AppError:

`go
// WithContact sets support contact information (email, URL, Slack channel)
// and returns the modified AppError for chaining.
func (e *AppError) WithContact(contact string) *AppError {
    if e == nil {
        return nil
    }
    e.Contact = contact
    return e
}

// WithErrors appends a list of sub-errors (e.g., batch validation failures)
// and returns the modified AppError for chaining.
func (e *AppError) WithErrors(errs ...error) *AppError {
    if e == nil {
        return nil
    }
    e.Errors = append(e.Errors, errs...)
    return e
}

// WithMsg overrides the developer-facing Message on the fly
// and returns the modified AppError for chaining.
func (e *AppError) WithMsg(msg string) *AppError {
    if e == nil {
        return nil
    }
    e.Message = msg
    return e
}
`

## 3. Update Display Methods (String, FullString)
**File to edit:** pperror_display.go (or equivalent stringification file)

Ensure that Contact and Errors are rendered when the error is printed to logs via FullString().

**Checklist for FullString() / ToClipboard():**
- [ ] If .Contact != "", append it to the output string (e.g., mt.Sprintf(" Contact: %s", e.Contact)).
- [ ] If len(e.Errors) > 0, iterate over .Errors and print each one, indenting them under the main error.

*Example stringification update:*
`go
if e.Contact != "" {
    sb.WriteString(fmt.Sprintf(" | Contact: %s", e.Contact))
}
if len(e.Errors) > 0 {
    sb.WriteString("\n  Sub-errors:")
    for i, subErr := range e.Errors {
        sb.WriteString(fmt.Sprintf("\n    %d. %s", i+1, subErr.Error()))
    }
}
`

## 4. Run Unit Tests & Linting
- Write unit tests in pperror_test.go to ensure chaining WithContact("foo").WithMsg("bar") works.
- Verify json.Marshal(appErr) correctly includes "Contact" and "Errors" if set, and omits them if empty.
- Run go run linter-scripts/validate-guidelines.go --path . to ensure no CODE-RED violations were introduced.
