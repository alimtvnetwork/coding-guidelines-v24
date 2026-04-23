import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { LucideIcon } from "lucide-react";

interface MetricCardProps {
  icon: LucideIcon;
  label: string;
  value: string;
  caption: string;
  tone?: "default" | "success" | "warning";
}

const TONE_CLASS: Record<NonNullable<MetricCardProps["tone"]>, string> = {
  default: "text-primary",
  success: "text-success",
  warning: "text-warning",
};

export function MetricCard({ icon: Icon, label, value, caption, tone = "default" }: MetricCardProps) {
  const accentClass = TONE_CLASS[tone];

  return (
    <Card className="border-border/60">
      <CardHeader className="flex flex-row items-center justify-between space-y-0 pb-2">
        <CardTitle className="text-sm font-medium text-muted-foreground">{label}</CardTitle>
        <Icon className={`h-4 w-4 ${accentClass}`} aria-hidden="true" />
      </CardHeader>
      <CardContent>
        <div className={`text-3xl font-heading font-semibold ${accentClass}`}>{value}</div>
        <p className="mt-1 text-xs text-muted-foreground">{caption}</p>
      </CardContent>
    </Card>
  );
}
