import glob
import re
from collections import defaultdict

headings = defaultdict(int)
for cg_file in glob.glob("01-prompts/15-cg-execute/*.md"):
    with open(cg_file, "r", encoding="utf-8") as f:
        for line in f:
            if line.startswith("## "):
                clean_heading = re.sub(r'^## \d+\.\s*', '## ', line.strip())
                headings[clean_heading] += 1

with open("headings.txt", "w", encoding="utf-8") as f:
    for k, v in sorted(headings.items(), key=lambda x: -x[1]):
        f.write(f"{v:3d}  {k}\n")
