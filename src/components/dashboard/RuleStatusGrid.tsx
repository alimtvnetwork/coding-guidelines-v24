import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { ShieldCheck } from "lucide-react";
import type { CiRule } from "@/types/dashboard";

interface RuleStatusGridProps {
  rules: CiRule[];
}

const isCleared = (status: CiRule["status"]): boolean => status === "cleared";

export function RuleStatusGrid({ rules }: RuleStatusGridProps) {
  const hasRules = rules.length > 0;

  return (
    <Card className="border-border/60">
      <CardHeader>
        <CardTitle className="flex items-center gap-2 font-heading">
          <ShieldCheck className="h-5 w-5 text-success" aria-hidden="true" />
          Pass / Fail by Rule
        </CardTitle>
        <p className="text-sm text-muted-foreground">
          Every CODE-RED / STYLE rule that has been triggered at least once. Counts come from `.lovable/cicd-issues/`.
        </p>
      </CardHeader>
      <CardContent>
        {hasRules === false ? (
          <p className="text-sm text-muted-foreground">No rule history recorded.</p>
        ) : (
          <ul className="grid gap-2 md:grid-cols-2">
            {rules.map((rule) => (
              <RuleRow key={rule.rule} rule={rule} />
            ))}
          </ul>
        )}
      </CardContent>
    </Card>
  );
}

function RuleRow({ rule }: { rule: CiRule }) {
  const cleared = isCleared(rule.status);
  const variant = cleared === true ? "secondary" : "destructive";
  const label = cleared === true ? "cleared" : "active";

  return (
    <li className="flex items-center justify-between rounded-md border border-border/60 bg-card px-3 py-2">
      <div>
        <p className="font-mono text-sm font-medium">{rule.rule}</p>
        <p className="text-xs text-muted-foreground">
          {rule.occurrences} historical occurrence{rule.occurrences === 1 ? "" : "s"}
        </p>
      </div>
      <Badge variant={variant} className="uppercase tracking-wide">
        {label}
      </Badge>
    </li>
  );
}
