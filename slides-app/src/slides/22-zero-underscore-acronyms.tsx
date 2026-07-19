import { SlideLayout } from "@/components/SlideLayout";
import { ActionPanel } from "@/components/ActionPanel";
import { CodeDiff } from "@/components/CodeDiff";

/**
 * SS-02 task 24: Zero-underscore policy + full-uppercase acronyms.
 *
 * Extends `01-naming-conventions` (NAM-001) which shows the general rule.
 * This slide focuses on the two most-violated corollaries:
 *   1. No underscores anywhere in identifiers (except language-mandated cases,
 *      e.g. Python `snake_case`, C# `_camelCase` private fields).
 *   2. Acronyms are ALWAYS fully uppercase inside PascalCase (URL, HTTP, ID,
 *      API, DB, IO) so `HTTPResponse`, not `HttpResponse`.
 */

const BEFORE = `type Http_Response = { status_code: number };
class Url_Parser {}
const api_key = "...";
function parse_Json_body(raw_input: string) {}`;

const AFTER = `type HTTPResponse = { StatusCode: number };
class URLParser {}
const APIKey = "...";
function ParseJSONBody(RawInput: string) {}`;

export default function ZeroUnderscoreAcronymsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 24 · Naming"
      title="Zero underscores. Acronyms stay fully UPPERCASE."
      subtitle="Two corollaries of the PascalCase rule that PRs violate the most: mid-name underscores and title-cased acronyms."
    >
      <div style={{ display: "flex", flexDirection: "column", gap: 18, marginTop: 4 }}>
        <CodeDiff
          language="typescript"
          before={BEFORE}
          after={AFTER}
          beforeLabel="❌ underscores + Http/Url/Json title-case"
          afterLabel="✅ no underscores, HTTP/URL/JSON/API full-caps"
        />
        <ActionPanel
          slideId="22-zero-underscore-acronyms"
          symptom="A review lands `Http_Response`, `Url_Parser`, `parse_Json_body`. Grep for `Url` misses `URL`, IDE rename skips half the sites, and readers stall on `Json` vs `JSON`."
          rule="No underscores in any identifier (except when the language mandates it: Python `snake_case`, C# `_camelCase` private fields). Acronyms inside names stay fully UPPERCASE: `HTTPResponse`, `URLParser`, `APIKey`, `JSONBody`."
          doThis="When you touch a file with `Http`, `Url`, `Json`, `Api`, `Db`, `Io` or any `_` in an identifier, rename that one symbol to full caps and drop the underscore. Do not bulk-rewrite the repo."
        />
      </div>
    </SlideLayout>
  );
}
