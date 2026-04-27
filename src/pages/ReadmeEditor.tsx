import { useMemo, useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { toast } from "@/components/ui/sonner";

const MONTHS = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

const formatMalaysiaDateTime = (now: Date): string => {
  // Malaysia: UTC+8, no DST.
  const parts = new Intl.DateTimeFormat("en-GB", {
    timeZone: "Asia/Kuala_Lumpur",
    year: "numeric",
    month: "numeric",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
    second: "2-digit",
    hour12: true,
  }).formatToParts(now);

  const lookup = Object.fromEntries(parts.map((p) => [p.type, p.value]));
  const month = MONTHS[Number(lookup.month) - 1];
  const date = `${lookup.day}-${month}-${lookup.year}`;
  const time = `${lookup.hour}:${lookup.minute}:${lookup.second} ${lookup.dayPeriod.toUpperCase()}`;
  return `${date} ${time}`;
};

const buildContent = (w1: string, w2: string, w3: string, stamp: string): string =>
  `${w1} ${w2} ${w3} ${stamp}\n`;

const downloadReadme = (content: string): void => {
  const blob = new Blob([content], { type: "text/plain;charset=utf-8" });
  const url = URL.createObjectURL(blob);
  const link = document.createElement("a");
  link.href = url;
  link.download = "readme.txt";
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
  URL.revokeObjectURL(url);
};

const ReadmeEditor = () => {
  const [word1, setWord1] = useState("let's");
  const [word2, setWord2] = useState("start");
  const [word3, setWord3] = useState("now");
  const [preview, setPreview] = useState<string>("");

  const allFilled = useMemo(
    () => word1.trim() !== "" && word2.trim() !== "" && word3.trim() !== "",
    [word1, word2, word3],
  );

  const handleGenerate = () => {
    if (!allFilled) {
      toast.error("All three words are required.");
      return;
    }
    const stamp = formatMalaysiaDateTime(new Date());
    const content = buildContent(word1.trim(), word2.trim(), word3.trim(), stamp);
    setPreview(content);
    downloadReadme(content);
    toast.success("readme.txt generated");
  };

  return (
    <main className="min-h-screen bg-background text-foreground py-12 px-4">
      <div className="mx-auto max-w-2xl">
        <Card>
          <CardHeader>
            <CardTitle>Readme Generator</CardTitle>
            <CardDescription>
              Set three words and click Generate. The output appends the current Malaysia
              date (dd-MMM-YYYY) and 12-hour time, then downloads as readme.txt.
            </CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            <div className="grid gap-4 sm:grid-cols-3">
              <div className="space-y-2">
                <Label htmlFor="w1">Word 1</Label>
                <Input id="w1" value={word1} onChange={(e) => setWord1(e.target.value)} />
              </div>
              <div className="space-y-2">
                <Label htmlFor="w2">Word 2</Label>
                <Input id="w2" value={word2} onChange={(e) => setWord2(e.target.value)} />
              </div>
              <div className="space-y-2">
                <Label htmlFor="w3">Word 3</Label>
                <Input id="w3" value={word3} onChange={(e) => setWord3(e.target.value)} />
              </div>
            </div>

            <Button onClick={handleGenerate} disabled={!allFilled} className="w-full">
              Generate readme.txt
            </Button>

            {preview && (
              <div className="space-y-2">
                <Label>Preview</Label>
                <pre className="rounded-md border bg-muted p-4 text-sm font-mono whitespace-pre-wrap">
                  {preview}
                </pre>
              </div>
            )}
          </CardContent>
        </Card>
      </div>
    </main>
  );
};

export default ReadmeEditor;
