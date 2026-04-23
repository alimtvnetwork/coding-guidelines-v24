type TokenKind = "command" | "flag" | "url" | "pipe" | "text";

const ShellOperatorType = {
  Pipe: "|",
  And: "&&",
  Or: "||",
  Semicolon: ";",
} as const;

const SHELL_OPERATORS: ReadonlySet<string> = new Set(Object.values(ShellOperatorType));
const KNOWN_COMMANDS = new Set(["irm", "iex", "curl", "bash", "sh", "wget", "powershell", "pwsh"]);

const TOKEN_CLASS: Record<TokenKind, string> = {
  command: "text-primary font-medium",
  flag: "text-accent-foreground/80",
  url: "text-muted-foreground/90 underline decoration-dotted decoration-muted-foreground/40 underline-offset-2",
  pipe: "text-destructive/80 font-semibold",
  text: "text-foreground/85",
};

function isShellOperator(token: string): boolean {
  return SHELL_OPERATORS.has(token);
}

function isUrlToken(token: string): boolean {
  return /^https?:\/\//i.test(token);
}

function isFlagToken(token: string): boolean {
  return /^-/.test(token);
}

function isCommandToken(token: string, index: number): boolean {
  return index === 0 || KNOWN_COMMANDS.has(token.toLowerCase());
}

function classifyToken(token: string, index: number): TokenKind {
  if (isShellOperator(token)) {
    return "pipe";
  }

  if (isUrlToken(token)) {
    return "url";
  }

  if (isFlagToken(token)) {
    return "flag";
  }

  if (isCommandToken(token, index)) {
    return "command";
  }

  return "text";
}

function isWhitespaceToken(token: string): boolean {
  return /^\s+$/.test(token);
}

function CommandToken({ token, index }: { token: string; index: number }) {
  const kind = classifyToken(token, index);

  return <span className={TOKEN_CLASS[kind]}>{token}</span>;
}

export function HighlightedCommand({ command }: { command: string }) {
  const tokens = command.split(/(\s+)/);
  const cursor = { nonSpaceIndex: -1 };

  return (
    <>
      {tokens.map((token, i) => {
        if (isWhitespaceToken(token)) {
          return <span key={i}>{token}</span>;
        }

        cursor.nonSpaceIndex += 1;

        return <CommandToken key={i} token={token} index={cursor.nonSpaceIndex} />;
      })}
    </>
  );
}