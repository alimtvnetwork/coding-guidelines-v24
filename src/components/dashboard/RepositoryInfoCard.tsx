import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { GitBranch } from "lucide-react";

interface RepositoryInfoCardProps {
  version: string;
  title: string;
  repoSlug: string;
  lastCommitSha: string;
}

const SHORT_SHA_LENGTH = 7;

const shortenSha = (sha: string): string => sha.slice(0, SHORT_SHA_LENGTH);

export function RepositoryInfoCard({
  version,
  title,
  repoSlug,
  lastCommitSha,
}: RepositoryInfoCardProps) {
  return (
    <Card className="border-border/60">
      <CardHeader className="flex flex-row items-center gap-2 space-y-0 pb-3">
        <GitBranch className="h-4 w-4 text-muted-foreground" aria-hidden />
        <CardTitle className="text-sm font-semibold">Repository Info</CardTitle>
      </CardHeader>
      <CardContent className="grid gap-3 text-sm sm:grid-cols-2 lg:grid-cols-4">
        <InfoField label="Version" value={version} />
        <InfoField label="Title" value={title} />
        <InfoField label="RepoSlug" value={repoSlug} mono />
        <InfoField label="LastCommitSha" value={shortenSha(lastCommitSha)} mono title={lastCommitSha} />
      </CardContent>
    </Card>
  );
}

interface InfoFieldProps {
  label: string;
  value: string;
  mono?: boolean;
  title?: string;
}

function InfoField({ label, value, mono, title }: InfoFieldProps) {
  const valueClass = mono === true ? "font-mono text-foreground" : "text-foreground";
  return (
    <div className="space-y-1">
      <p className="text-xs uppercase tracking-wide text-muted-foreground">{label}</p>
      <p className={`truncate ${valueClass}`} title={title ?? value}>
        {value}
      </p>
    </div>
  );
}
