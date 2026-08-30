export enum TokenKindType {
  Command = "command",
  Flag = "flag",
  Url = "url",
  Pipe = "pipe",
  Text = "text"
}

const ShellOperatorType = {
  Pipe: "|",
  And: "&&",
  Or: "||",
  Semicolon: ";",
} as const;

const SHELL_OPERATORS: ReadonlySet<string> = new Set(Object.values(ShellOperatorType));
const KNOWN_COMMANDS = new Set(["irm", "iex", "curl", "bash", "sh", "wget", "powershell", "pwsh"]);

const TOKEN_CLASS: Record<TokenKindType, string> = {
  [TokenKindType.Command]: "text-primary font-medium",
  [TokenKindType.Flag]: "text-accent-foreground/80",
  [TokenKindType.Url]: "text-muted-foreground/90 underline decoration-dotted decoration-muted-foreground/40 underline-offset-2",
  [TokenKindType.Pipe]: "text-destructive/80 font-semibold",
  [TokenKindType.Text]: "text-foreground/85",
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

function classifyToken(token: string, index: number): TokenKindType {
  if (isShellOperator(token)) {
    return TokenKindType.Pipe;
  }

  if (isUrlToken(token)) {
    return TokenKindType.Url;
  }

  if (isFlagToken(token)) {
    return TokenKindType.Flag;
  }

  if (isCommandToken(token, index)) {
    return TokenKindType.Command;
  }

  return TokenKindType.Text;
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
