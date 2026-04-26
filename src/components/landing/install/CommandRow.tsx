import { useEffect, useLayoutEffect, useRef, useState } from "react";
import { Maximize2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
  isClosed,
  isOpen,
  isOverflowFit,
  isOverflowing,
} from "@/constants/boolFlags";
import { CopyButton } from "./CopyButton";
import { HighlightedCommand } from "./HighlightedCommand";
import { CommandExpandDialog } from "./CommandExpandDialog";

interface CommandRowProps {
  /** Short label rendered in the left gutter (e.g. "$", "bash", "pwsh"). */
  label: string;
  /** The shell-ready command string. */
  command: string;
  /** Title shown in the expand dialog (e.g. platform name or bundle name). */
  expandTitle: string;
  /** Shell name shown as a chip in the expand dialog (e.g. "Bash", "PowerShell"). */
  expandShell: string;
}

function useOverflowDetection(
  ref: React.RefObject<HTMLElement>,
  command: string,
) {
  const [hasOverflow, setHasOverflow] = useState<boolean>(isOverflowFit);

  useLayoutEffect(() => {
    const measure = () => {
      const node = ref.current;
      if (!node) return;
      const isWider = node.scrollWidth > node.clientWidth + 1;
      setHasOverflow(isWider ? isOverflowing : isOverflowFit);
    };
    measure();

    const node = ref.current;
    if (!node) return;
    const observer = new ResizeObserver(measure);
    observer.observe(node);
    return () => observer.disconnect();
  }, [ref, command]);

  useEffect(() => {
    const onResize = () => {
      const node = ref.current;
      if (!node) return;
      const isWider = node.scrollWidth > node.clientWidth + 1;
      setHasOverflow(isWider ? isOverflowing : isOverflowFit);
    };
    window.addEventListener("resize", onResize);
    return () => window.removeEventListener("resize", onResize);
  }, [ref]);

  return hasOverflow;
}

function ExpandButton({ onClick }: { onClick: () => void }) {
  return (
    <Button
      size="sm"
      variant="ghost"
      onClick={onClick}
      className="h-8 shrink-0 px-2 text-muted-foreground hover:text-foreground"
      aria-label="Expand command"
    >
      <Maximize2 className="h-4 w-4" />
    </Button>
  );
}

function CommandRowLabel({ label }: { label: string }) {
  return (
    <span className="select-none text-[10px] font-semibold uppercase tracking-wide text-muted-foreground/80">
      {label}
    </span>
  );
}

export function CommandRow({ label, command, expandTitle, expandShell }: CommandRowProps) {
  const scrollRef = useRef<HTMLDivElement>(null);
  const [isDialogOpen, setIsDialogOpen] = useState<boolean>(isClosed);
  const hasOverflow = useOverflowDetection(scrollRef, command);

  return (
    <>
      <div className="flex items-center gap-2 rounded-md border border-border bg-secondary/60 px-3 py-2 font-mono text-foreground/90">
        <CommandRowLabel label={label} />
        <div className="relative min-w-0 flex-1">
          <div
            ref={scrollRef}
            className="scrollbar-none overflow-x-auto whitespace-nowrap text-[11px] leading-relaxed sm:text-xs"
          >
            <code className="text-foreground/90">
              <HighlightedCommand command={command} />
            </code>
          </div>
          {hasOverflow && (
            <div
              aria-hidden
              className="pointer-events-none absolute inset-y-0 right-0 w-8 bg-gradient-to-l from-secondary/95 to-transparent"
            />
          )}
        </div>
        {hasOverflow && <ExpandButton onClick={() => setIsDialogOpen(isOpen)} />}
        <CopyButton command={command} />
      </div>

      <CommandExpandDialog
        open={isDialogOpen}
        onOpenChange={setIsDialogOpen}
        title={expandTitle}
        shell={expandShell}
        command={command}
      />
    </>
  );
}