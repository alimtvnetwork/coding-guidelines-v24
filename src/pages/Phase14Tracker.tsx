import { Link } from "react-router-dom";
import { ArrowLeft, FileText, ExternalLink } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import {
  Phase14Meta,
  Phase14Tasks,
  type Phase14Status,
  type Phase14Task,
} from "@/data/phase14Tasks";

const StatusVariant: Record<Phase14Status, "default" | "secondary" | "destructive" | "outline"> = {
  Todo: "outline",
  InProgress: "default",
  Blocked: "destructive",
  Done: "secondary",
};

function TaskRow({ task }: { task: Phase14Task }) {
  const docsHref = `/docs?path=${encodeURIComponent(task.SpecPath)}`;
  return (
    <li className="rounded-lg border border-border bg-card p-4 transition-colors hover:border-primary/40">
      <div className="flex flex-wrap items-center justify-between gap-3">
        <div className="flex items-center gap-3">
          <span className="font-mono text-xs text-muted-foreground">{task.Id}</span>
          <h3 className="font-heading text-base font-semibold">{task.Title}</h3>
          <Badge variant={StatusVariant[task.Status]}>{task.Status}</Badge>
        </div>
        <div className="flex items-center gap-2">
          <Badge variant="outline" className="font-mono text-xs">{task.Language}</Badge>
          <Badge variant="outline" className="font-mono text-xs">{task.Repo}</Badge>
        </div>
      </div>
      <div className="mt-3 flex flex-wrap gap-4 text-xs text-muted-foreground">
        <Link to={docsHref} className="inline-flex items-center gap-1 hover:text-primary">
          <FileText className="h-3.5 w-3.5" aria-hidden="true" />
          <span className="font-mono">{task.SpecPath}</span>
        </Link>
        <a
          href={`https://github.com/mahin/coding-guidelines/blob/main/${task.IssuePath}`}
          target="_blank"
          rel="noopener noreferrer"
          className="inline-flex items-center gap-1 hover:text-primary"
        >
          <ExternalLink className="h-3.5 w-3.5" aria-hidden="true" />
          <span className="font-mono">{task.IssuePath}</span>
        </a>
      </div>
    </li>
  );
}

function Header() {
  return (
    <header className="mb-8">
      <Button asChild variant="ghost" size="sm" className="mb-6">
        <Link to="/">
          <ArrowLeft className="mr-2 h-4 w-4" aria-hidden="true" />
          Back home
        </Link>
      </Button>
      <h1 className="font-heading text-3xl font-bold tracking-tight sm:text-4xl">
        Phase 14 — Implementation Tracker
      </h1>
      <p className="mt-3 text-sm text-muted-foreground">
        {Phase14Meta.SurfaceLabel}. {Phase14Meta.BackupExcluded} Mutation gate{" "}
        <span className="font-mono">{Phase14Meta.MutationGate}</span>. Constraint:{" "}
        <span className="font-mono">{Phase14Meta.ConstraintRef}</span>.
      </p>
    </header>
  );
}

export default function Phase14Tracker() {
  return (
    <main className="min-h-screen bg-background text-foreground">
      <section className="container mx-auto max-w-4xl px-6 py-12">
        <Header />
        <ul className="space-y-3">
          {Phase14Tasks.map((task) => (
            <TaskRow key={task.Id} task={task} />
          ))}
        </ul>
      </section>
    </main>
  );
}
