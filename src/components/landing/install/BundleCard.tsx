import { Package } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { CopyButton } from "./CopyButton";
import { HighlightedCommand } from "./HighlightedCommand";

export type BundleCommand = {
  bundle: string;
  description: string;
  folders: string[];
  bash: string;
  ps: string;
  archive: string;
};

function BundleCardHeader({ bundle }: { bundle: BundleCommand }) {
  const folderSuffix = bundle.folders.length === 1 ? "" : "s";

  return (
    <CardHeader className="pb-3">
      <CardTitle className="flex items-center gap-2 text-base font-semibold text-foreground">
        <Package className="h-4 w-4 text-primary" />
        <span className="font-mono text-sm">{bundle.bundle}</span>
        <span className="ml-auto rounded-full border border-border bg-secondary px-2 py-0.5 text-xs font-medium text-muted-foreground">
          {bundle.folders.length} folder{folderSuffix}
        </span>
      </CardTitle>
      <p className="mt-1 text-xs text-muted-foreground">{bundle.description}</p>
    </CardHeader>
  );
}

function BundleFoldersList({ folders }: { folders: string[] }) {
  return (
    <ul className="mt-2 flex flex-wrap gap-1.5">
      {folders.map((folder) => (
        <li
          key={folder}
          className="rounded-md border border-border/60 bg-secondary/40 px-2 py-0.5 font-mono text-[10px] text-muted-foreground"
        >
          {folder}
        </li>
      ))}
    </ul>
  );
}

function BundleCommandRow({ label, command }: { label: string; command: string }) {
  return (
    <div className="flex items-center gap-2 rounded-md border border-border bg-secondary/60 px-3 py-2 font-mono text-foreground/90">
      <span className="select-none text-[10px] font-semibold uppercase tracking-wide text-muted-foreground/80">
        {label}
      </span>
      <code className="flex-1 break-all text-[11px] leading-relaxed sm:text-xs">
        <HighlightedCommand command={command} />
      </code>
      <CopyButton command={command} />
    </div>
  );
}

export function BundleCard({ bundle }: { bundle: BundleCommand }) {
  return (
    <Card className="overflow-hidden border-border/60 bg-card/50 transition-colors hover:border-primary/40">
      <BundleCardHeader bundle={bundle} />
      <CardContent className="space-y-2">
        <BundleCommandRow label="bash" command={bundle.bash} />
        <BundleCommandRow label="pwsh" command={bundle.ps} />
        <BundleFoldersList folders={bundle.folders} />
      </CardContent>
    </Card>
  );
}