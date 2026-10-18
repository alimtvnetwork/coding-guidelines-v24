# Architecture Specification: MIME & Media Type Enum Architecture

## 1. Overview & Strategic Intent
This specification formalizes the architectural standards and migration strategy for eliminating raw magic string literals representing MIME and Media Types across all frontend, backend, CLI, and tooling codebases.
Raw literals such as `'text/html'`, `'text/plain'`, `'application/json'`, `'image/png'` used in:
- Browser Blob constructors: `new Blob([content], { type: 'text/html' })`
- ClipboardItem APIs: `new ClipboardItem({ 'text/html': blob, 'text/plain': textBlob })`
- HTTP request/response headers: `headers: { 'Content-Type': 'application/json' }`
- Backend content negotiation and file streaming responses
lead to silent runtime typing errors, inconsistent naming, broken IDE discovery, and code duplication.

This architecture enforces strongly-typed enums / typed constant dictionaries (`MimeType` / `MediaType`) across TypeScript, Go, Python, PHP, and Rust, updates the `06-constants-and-enums.md` Coding Guideline (CG) execute prompt and companion skills, and establishes automated linting and fleet synchronization.

---

## 2. Root Cause & Code Smell Analysis

### Code Smell Identification (from Clipboard / Blob Screenshot)
In `src/components/instances/PromptTreeView/clipboard.ts`:
```typescript
// ❌ ANTI-PATTERN: Magic strings in Blob & ClipboardItem
const htmlBlob = new Blob([htmlContent], { type: 'text/html' });
const textBlob = new Blob([mdText], { type: 'text/plain' });

const clipboardItem = new ClipboardItem({
  'text/html': htmlBlob,
  'text/plain': textBlob,
});
await navigator.clipboard.write([clipboardItem]);
```

### Architectural Deficiencies
1. **Zero Compiler Verification:** If a developer makes a typographical error (e.g. `'text/htm'` or `'text/palin'`), TypeScript and JavaScript accept the string without error.
2. **Scattered Duplication:** Identical MIME strings are re-typed across dozens of frontend components and backend routes.
3. **Impeded Refactoring:** Changing or inspecting media type usage requires unbounded regex / text searching rather than compiler-backed symbol references.
4. **Non-Conformance with Repository Enum Standards:** Coding guidelines (`02-spec/02-coding-guidelines/01-cross-language/06-constants-and-enums.md` and `36-mime-and-media-type-enum-standards.md`) forbid inline magic literals for domain constants.

---

## 3. Canonical Architecture Contracts

### TypeScript Contract (`MimeType` Object Enum)
```typescript
export const MimeType = {
  TextPlain: 'text/plain',
  TextHtml: 'text/html',
  TextMarkdown: 'text/markdown',
  TextCss: 'text/css',
  TextCsv: 'text/csv',
  ApplicationJson: 'application/json',
  ApplicationXml: 'application/xml',
  ApplicationPdf: 'application/pdf',
  ApplicationZip: 'application/zip',
  ApplicationOctetStream: 'application/octet-stream',
  ImagePng: 'image/png',
  ImageJpeg: 'image/jpeg',
  ImageGif: 'image/gif',
  ImageSvg: 'image/svg+xml',
  ImageWebp: 'image/webp',
  AudioMpeg: 'audio/mpeg',
  VideoMp4: 'video/mp4',
  MultipartFormData: 'multipart/form-data',
} as const;

export type MimeType = typeof MimeType[keyof typeof MimeType];
```

### Go Contract (`mimetype.MimeType` String Type)
```go
package mimetype

type MimeType string

const (
    TextPlain              MimeType = "text/plain"
    TextHTML               MimeType = "text/html"
    TextMarkdown           MimeType = "text/markdown"
    ApplicationJSON        MimeType = "application/json"
    ApplicationXML         MimeType = "application/xml"
    ApplicationPDF         MimeType = "application/pdf"
    ApplicationOctetStream MimeType = "application/octet-stream"
    ImagePNG               MimeType = "image/png"
    ImageJPEG              MimeType = "image/jpeg"
    ImageSVG               MimeType = "image/svg+xml"
)

func (m MimeType) String() string {
    return string(m)
}
```

### Python Contract (`MimeType` StrEnum)
```python
from enum import Enum

class MimeType(str, Enum):
    TEXT_PLAIN = "text/plain"
    TEXT_HTML = "text/html"
    TEXT_MARKDOWN = "text/markdown"
    APPLICATION_JSON = "application/json"
    APPLICATION_XML = "application/xml"
    APPLICATION_PDF = "application/pdf"
    APPLICATION_OCTET_STREAM = "application/octet-stream"
    IMAGE_PNG = "image/png"
    IMAGE_JPEG = "image/jpeg"
    IMAGE_SVG = "image/svg+xml"
```

---

## 4. Architectural Boundaries & Non-Negotiables
1. **Strictly Relative Git Paths:** All documentation, specs, and prompt links use paths relative to repository root (`02-spec/...`, `.ai-memory/...`).
2. **Zero Storage & Zero Build/Test Rules:** No CI artifact storage; zero invocation of build or test runners (`npm run build`, `pytest`, `go test`).
3. **Spec 21 Fleet Isolation:** When synchronizing across 43 repositories via `38-sync-prompts-skills-scripts.py`, `02-spec/21-*` through `25-*` remain strictly isolated in this repository and are never mirrored.
