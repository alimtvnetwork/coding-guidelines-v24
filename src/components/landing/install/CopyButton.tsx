import { useState } from "react";
import { Check, Copy } from "lucide-react";
import { Button } from "@/components/ui/button";
import { isCopied, isCopyReset } from "@/constants/boolFlags";

const COPY_FEEDBACK_DURATION_MS = 2000;

function useCopyHandler(command: string, setHasCopied: (value: boolean) => void) {
  return async () => {
    await navigator.clipboard.writeText(command);
    setHasCopied(isCopied);
    setTimeout(() => setHasCopied(isCopyReset), COPY_FEEDBACK_DURATION_MS);
  };
}

function CopyButtonView({ hasCopied, onClick }: { hasCopied: boolean; onClick: () => void }) {
  return (
    <Button
      size="sm"
      variant="ghost"
      onClick={onClick}
      className="h-8 shrink-0 px-2 text-muted-foreground hover:text-foreground"
      aria-label="Copy install command"
    >
      {hasCopied ? <Check className="h-4 w-4 text-primary" /> : <Copy className="h-4 w-4" />}
    </Button>
  );
}

export function CopyButton({ command }: { command: string }) {
  const [hasCopied, setHasCopied] = useState(isCopyReset);
  const handleCopy = useCopyHandler(command, setHasCopied);

  return <CopyButtonView hasCopied={hasCopied} onClick={handleCopy} />;
}