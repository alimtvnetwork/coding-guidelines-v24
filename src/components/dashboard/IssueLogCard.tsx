import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { ListChecks } from "lucide-react";
import type { CiIssue } from "@/types/dashboard";

interface IssueLogCardProps {
  issues: CiIssue[];
}

const isSolved = (status: CiIssue["status"]): boolean => status === "solved";

export function IssueLogCard({ issues }: IssueLogCardProps) {
  const hasIssues = issues.length > 0;

  return (
    <Card className="border-border/60">
      <CardHeader>
        <CardTitle className="flex items-center gap-2 font-heading">
          <ListChecks className="h-5 w-5 text-primary" aria-hidden="true" />
          Tracked Issues
        </CardTitle>
        <p className="text-sm text-muted-foreground">
          Sourced from `.lovable/cicd-issues/`. Each entry is a real validator finding plus its fix.
        </p>
      </CardHeader>
      <CardContent>
        {hasIssues === false ? (
          <p className="text-sm text-muted-foreground">No issues tracked yet.</p>
        ) : (
          <ul className="space-y-2">
            {issues.map((issue) => (
              <IssueRow key={issue.slug} issue={issue} />
            ))}
          </ul>
        )}
      </CardContent>
    </Card>
  );
}

function IssueRow({ issue }: { issue: CiIssue }) {
  const solved = isSolved(issue.status);
  const variant = solved === true ? "secondary" : "destructive";
  const label = solved === true ? "solved" : "open";

  return (
    <li className="rounded-md border border-border/60 bg-card px-3 py-2">
      <div className="flex flex-wrap items-center justify-between gap-2">
        <p className="text-sm font-medium">
          <span className="text-muted-foreground">#{String(issue.number).padStart(2, "0")}</span>{" "}
          {issue.title}
        </p>
        <Badge variant={variant} className="uppercase">{label}</Badge>
      </div>
      <div className="mt-1 flex flex-wrap items-center gap-2 text-xs text-muted-foreground">
        <span>{issue.date}</span>
        {issue.rules.map((rule) => (
          <span key={rule} className="rounded bg-secondary px-2 py-0.5 font-mono text-[11px]">
            {rule}
          </span>
        ))}
      </div>
    </li>
  );
}
