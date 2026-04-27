import { Terminal } from "lucide-react";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import versionData from "@/../version.json";
import bundlesManifest from "@/../bundles.json";
import { CommandRow } from "./install/CommandRow";
import { BundleCard, type BundleCommand } from "./install/BundleCard";

interface InstallCommand {
  platform: string;
  shell: string;
  command: string;
}

const installCommands: InstallCommand[] = [
  {
    platform: "Windows",
    shell: "PowerShell",
    command: "irm https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v18/main/install.ps1 | iex",
  },
  {
    platform: "macOS / Linux",
    shell: "Bash",
    command: "curl -fsSL https://raw.githubusercontent.com/alimtvnetwork/coding-guidelines-v18/main/install.sh | bash",
  },
];

const bundleCommands: BundleCommand[] = bundlesManifest.bundles.map((bundle) => ({
  bundle: bundle.name,
  description: bundle.description,
  folders: bundle.folders.map((folder) => folder.dest),
  bash: `curl -fsSL ${bundlesManifest.rawBase}/${bundle.name}-install.sh | bash`,
  ps: `irm ${bundlesManifest.rawBase}/${bundle.name}-install.ps1 | iex`,
  archive: `${bundle.archive.stableName}.zip`,
}));

function InstallCard({ item }: { item: InstallCommand }) {
  return (
    <Card className="overflow-hidden border-border/60 bg-card/50 transition-colors hover:border-primary/40">
      <CardHeader className="pb-3">
        <CardTitle className="flex items-center gap-2 text-base font-semibold text-foreground">
          <Terminal className="h-4 w-4 text-primary" />
          {item.platform}
          <span className="ml-auto rounded-full border border-border bg-secondary px-2 py-0.5 text-xs font-medium text-muted-foreground">
            {item.shell}
          </span>
        </CardTitle>
      </CardHeader>
      <CardContent>
        <CommandRow
          label="$"
          command={item.command}
          expandTitle={item.platform}
          expandShell={item.shell}
        />
      </CardContent>
    </Card>
  );
}

export function InstallSection() {
  return (
    <section className="border-y border-border bg-secondary/20 py-20">
      <div className="mx-auto max-w-6xl px-6">
        <div className="mb-10 text-center">
          <div className="flex items-center justify-center gap-3">
            <h2 className="text-3xl font-bold text-foreground">Release &amp; Migration</h2>
            <Badge variant="secondary" className="text-xs">
              v{versionData.version}
            </Badge>
          </div>
          <p className="mt-3 text-muted-foreground">
            Install the current v16 repo in one line. Bundle installers below point to the same release branch.
          </p>
        </div>

        <div className="grid gap-4 md:grid-cols-2">
          {installCommands.map((item) => (
            <InstallCard key={item.platform} item={item} />
          ))}
        </div>

        <div className="mt-16">
          <div className="mb-6 text-center">
            <h3 className="text-2xl font-bold text-foreground">Named Bundle Installers</h3>
            <p className="mt-2 text-sm text-muted-foreground">
              Install just the spec folders you need — each bundle points at coding-guidelines-v18.
            </p>
          </div>
          <div className="grid gap-4 md:grid-cols-2">
            {bundleCommands.map((bundle) => (
              <BundleCard key={bundle.bundle} bundle={bundle} />
            ))}
          </div>
        </div>
      </div>
    </section>
  );
}
