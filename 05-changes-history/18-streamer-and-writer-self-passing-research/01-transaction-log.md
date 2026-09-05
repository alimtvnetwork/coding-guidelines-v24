# Transaction Log: Streamer & Writer Self-Passing Architecture & Interface Naming Guidelines

> **Directory:** `05-changes-history/18-streamer-and-writer-self-passing-research/`  
> **Date:** 2026-09-04  
> **Topic:** Streamer & Writer Refinements, Self-Passing Injected Methods (`self`), Lock/Unlock Controls, AppError Return Standard, and Cross-Language Interface Naming  
> **Status:** Completed  

---

## 1. Context & User Directives

1. **Elimination of `Interfacer`:**
   - Remove redundant `Interfacer` / `AsInterfacer()` abstractions.
   - Retain concrete self-binding methods `AsStreamer()` and `AsWriter()` for type compatibility and extraction.

2. **Strict `*appfault.AppError` Return Standard:**
   - All operations (`Stream`, `Write`, `Sync`, `Close`) must return `*appfault.AppError` instead of generic `error` to adhere to repository error management standards.

3. **Locking Lifecycle Methods:**
   - Add explicit lock management methods directly to the interface: `Lock()`, `Unlock()`, `RLock()`, `RUnlock()`, and `IsLocked() bool`.
   - In lockless/unlocked mode, these methods behave as zero-overhead no-ops.

4. **Self-Passing Instance in Injected Functions (`self`):**
   - When injecting higher-order functions or composer functions into a writer or streamer, the instance must pass itself as the first argument (`self Writer`, `self Streamer[T]`).
   - This empowers injected functions and composers to inspect instance state (`Destination()`, `IsLocked()`, `Name()`) without tight coupling or closure scope pollution.

5. **Cross-Language Interface Naming Conventions:**
   - **Go:** `-er` capability suffix (`Writer`, `Streamer`, `Formatter`, `Reader`, `Closer`).
   - **C#:** `I` prefix in PascalCase (`IWriter`, `IStreamer`, `IObservable`, `IDisposable`).
   - **Java:** `-able` / `-ible` capability suffix (`Streamable`, `Writable`, `Observable`, `AutoCloseable`).
   - **TypeScript:** PascalCase nouns or `-able` adjectives without `I` prefix (`Writer`, `Streamer`, `Streamable`).
   - **Rust:** UpperCamelCase traits as action verbs or adjectives (`Write`, `Read`, `Stream`, `Streamable`).
   - **PHP:** `Interface` suffix in PascalCase (`WriterInterface`, `StreamerInterface`).

6. **End-to-End Writer Implementations:**
   - Generic `BaseWriter` (configurable locked vs unlocked modes).
   - Log Type Writer (structured logs with AppError metadata).
   - Unlocked Log Type Writer (lockless high-performance streaming).
   - File Writer Type (managing file destination, flush, and close).
   - Generic JSON Type Writer (configurable field serialization, envelopes, and formatting).
   - Composer Functions (demonstrating how `self` enables dynamic middleware decoration).

7. **Result Wrapper Standard (`result.Wrap[T]`) & Stutter Elimination (`BaseWriterWrap`):**
   - Package stutter (`result.Result`) is eliminated using `result.Wrap[T]` (or `wrapped.Result[T]`).
   - Defined `type BaseWriterWrap = result.Wrap[*BaseWriter]` to encapsulate the wrapped result type cleanly.
   - All writer constructors return `BaseWriterWrap` directly, eliminating generic clutter from signatures and strictly upholding `02-spec/02-coding-guidelines/03-golang/09-wrapped-boolean-results.md` and `CODE-RED-012`.

8. **Default Printing & Pluggable Formatting on Fault & Wrap (`PrintFault()`, `fault.Print()`):**
   - Eliminates manual and repetitive `fmt.Printf("... %s %d", fault.Message(), fault.Type().Code())` at error check sites.
   - `fault.Print()`, `fault.Format(formatter)`, and `fault.PrintWith(formatter)` implemented directly on `*AppError`.
   - `wrap.PrintFault()`, `wrap.Print()`, and `wrap.Format(formatter)` implemented directly on `Result[T]` / `Wrap[T]`.
   - Pluggable formatters allow default terminal output now and custom JSON/colored formatters later.

9. **Enum-Driven File Utility (`FileOpenModeType`, `FilePermType`) & Elimination of Raw `os.OpenFile`:**
   - Raw `os.OpenFile(opts.FilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)` eliminated across the codebase in favor of `fileutil.OpenFile(path, openMode, permMode)`.
   - Implemented `FileOpenModeType` with modes: `FileOpenReadOnly`, `FileOpenWriteOnly`, `FileOpenReadWrite`, `FileOpenAppend`, `FileOpenCreateAppend`, `FileOpenCreateTruncate`, `FileOpenCreateNew`.
   - Expanded `FilePermType` into an exhaustive, industrial-grade POSIX permission enum covering 21 permissions across Owner, Group, Public/World, and special bits (Sticky, SUID, SGID).
   - Equipped `FilePermType` with formatters (`OctalString()`, `PosixString()`), boolean inspections (`IsPrivate()`, `IsPublic()`, `IsExecutable()`, `IsOwnerWritable()`, etc.), modifiers (`WithPrivate()`, `WithReadOnly()`, `WithExecutable()`), and `ParsePerm` parser.

10. **Error Object & Error ID Constructors on Wrap Containers:**
   - Direct construction of failed wrappers with structured error objects or error IDs (`errtype.Variation`).
   - `result.WrapFailureWithId[T](errType, msg)` and `result.WrapFailureFromWrap[T, U](failedWrap)`.
   - `appwriter.WrapWriterFailure(err)`, `appwriter.WrapWriterFailureWithId(errType, msg)`, and `appwriter.WrapWriterFailureFromWrap(wrap)`.

---

## 2. Files Created & Modified

### New Packages & Files
- `04-code/golang/pkg/fileutil/types.go`: `FileOpenModeType` open flags enum.
- `04-code/golang/pkg/fileutil/perm_types.go`: Comprehensive `FilePermType` POSIX permissions enum, inspection methods, modifiers, and parser.
- `04-code/golang/pkg/fileutil/fileutil.go`: `OpenFile`, `Open`, `EnsureDir`, `ReadAll`, `ReadString`, `WriteFile`.
- `04-code/golang/pkg/fileutil/fileutil_test.go`: Test suite for file utilities.
- `04-code/golang/pkg/appwriter/interfaces.go`: `Writer`, `Streamer[T]`, `WriteMethodFunc`, `BaseWriterWrap`.
- `04-code/golang/pkg/appwriter/wrap_constructors.go`: Dedicated constructors for `BaseWriterWrap`.
- `04-code/golang/pkg/appwriter/base_writer.go`: Core `BaseWriter` implementation with `self` passing.
- `04-code/golang/pkg/appwriter/file_writer.go`: Enum-driven `NewFileWriter`.
- `04-code/golang/pkg/appwriter/appwriter_test.go`: Test suite for `appwriter`.

### Modified Files
- `04-code/golang/pkg/result/result.go`: Added `WrapFailureWithId`, `WrapFailureWithCause`, `WrapFailureFromWrap`.
- `04-code/golang/pkg/appfault/result_constructors.go`: Added `NewFailureWithId`, `NewFailureWithCause`, `FailureFromWrap`.
- `04-code/golang/pkg/applogger/file_sink.go`: Replaced raw `os.OpenFile` with `fileutil.OpenFile`.
- `05-changes-history/05-streamer-and-writer-self-passing-research/01-transaction-log.md`: This file.

### Updates to Index
- `05-changes-history/01-index.md`: Registered Task 05.

---

## 3. Verification & Quality Gates

- Verified against repository coding guidelines:
  - Implicit boolean checks (no `== true`).
  - No mixed-polarity conditions.
  - Strict relative git paths.
  - Strict lowercase file naming.
  - Functions cleanly bounded within 8–15 lines.
