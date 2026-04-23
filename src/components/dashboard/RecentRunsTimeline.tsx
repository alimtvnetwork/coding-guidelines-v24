import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { CheckCircle2, Clock, GitCommitHorizontal } from "lucide-react";
import type { CiRun } from "@/types/dashboard";

interface RecentRunsTimelineProps {
  runs: CiRun[];
}

const isPassed = (status: CiRun["status"]): boolean => status === "passed";

export function RecentRunsTimeline({ runs }: RecentRunsTimelineProps) {
  const hasRuns = runs.length > 0;

  return (
    <Card className="border-border/60">
      <CardHeader>
        <CardTitle className="flex items-center gap-2 font-heading">
          <Clock className="h-5 w-5 text-primary" aria-hidden="true" />
          Recent Runs
        </CardTitle>
        <p className="text-sm text-muted-foreground">
          Synthesised from solved-issue dates and the latest clean validator pass.
        </p>
      </CardHeader>
      <CardContent>
        {hasRuns === false ? (
          <p className="text-sm text-muted-foreground">No recent runs recorded.</p>
        ) : (
          <ol className="relative space-y-4 border-l border-border/60 pl-6">
            {runs.map((run) => (
              <RunRow key={run.id} run={run} />
            ))}
          </ol>
        )}
      </CardContent>
    </Card>
  );
}

function RunRow({ run }: { run: CiRun }) {
  const passed = isPassed(run.status);
  const variant = passed === true ? "secondary" : "destructive";
  const Icon = passed === true ? CheckCircle2 : Clock;
  const tone = passed === true ? "text-success" : "text-destructive";

  return (
    <li className="relative">
      <span className="absolute -left-[31px] top-1 flex h-5 w-5 items-center justify-center rounded-full border border-border bg-card">
        <Icon className={`h-3 w-3 ${tone}`} aria-hidden="true" />
      </span>
      <div className="flex flex-wrap items-center gap-2">
        <span className="font-mono text-sm">{run.date}</span>
        <Badge variant={variant} className="uppercase">{run.status}</Badge>
        <span className="inline-flex items-center gap-1 text-xs text-muted-foreground">
          <GitCommitHorizontal className="h-3 w-3" aria-hidden="true" />
          {run.commitSha}
        </span>
        <span className="text-xs text-muted-foreground">{run.durationSeconds}s</span>
      </div>
      <p className="mt-1 text-sm text-muted-foreground">{run.note}</p>
      <p className="mt-1 text-xs text-muted-foreground">
        {run.filesScanned.toLocaleString()} files · {run.linesScanned.toLocaleString()} lines
      </p>
    </li>
  );
}
