import { SlideLayout } from "@/components/SlideLayout";
import { CodeDiff } from "@/components/CodeDiff";
import { ActionPanel } from "@/components/ActionPanel";

const BEFORE = `const user_id = 1;
const URL_parser = makeParser();
type http_response = { code: number };
function get_user_data(id) { /* ... */ }`;

const AFTER = `const UserId = 1;
const URLParser = makeParser();
type HTTPResponse = { Code: number };
function GetUserData(Id) { /* ... */ }`;

export default function NamingConventionsSlide() {
  return (
    <SlideLayout
      eyebrow="Rule 01 · Naming"
      title="Use PascalCase. Everywhere. No underscores."
      subtitle="One naming rule across identifiers, DB columns, JSON keys and types. Acronyms stay fully uppercase."
    >
      <CodeDiff
        language="typescript"
        before={BEFORE}
        after={AFTER}
        beforeLabel="❌ snake_case mix"
        afterLabel="✅ PascalCase, full-caps acronyms"
      />
      <ActionPanel
        symptom="snake_case, camelCase and lowercase acronyms all coexist — readers can't predict any name."
        rule="PascalCase for every identifier. Acronyms (URL, HTTP, ID) stay fully uppercase inside the name."
        doThis="Rename one snake_case symbol in the file you're touching. Don't bulk-rename — fix as you go."
      />
    </SlideLayout>
  );
}
