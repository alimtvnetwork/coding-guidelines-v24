import re

lines = [
    "if a && b {",
    "if (a && b) {",
    "if a && b && c {",
    "if a || b || c {",
    "if a && b || c {",
    "if !a && b {",
    "if a && !b {",
    "if a != 1 && b == 2 {",
    "if not a and b:",
    "if a and not b:"
]

for line in lines:
    if "{" in line:
        cond = line[line.find("if")+2:line.rfind("{")].strip()
        if cond.startswith("("): cond = cond[1:]
        if cond.endswith(")"): cond = cond[:-1]
        
        and_count = cond.count("&&")
        or_count = cond.count("||")
        if and_count + or_count > 1:
            print(f"FAIL (too many): {line}")
        elif and_count + or_count == 1:
            op = "&&" if and_count else "||"
            left, right = cond.split(op, 1)
            left_neg = "!" in left or "!=" in left
            right_neg = "!" in right or "!=" in right
            if left_neg != right_neg:
                print(f"FAIL (mixed polarity): {line}")
    elif ":" in line:
        cond = line[line.find("if")+2:line.rfind(":")].strip()
        and_count = cond.count(" and ")
        or_count = cond.count(" or ")
        if and_count + or_count > 1:
            print(f"FAIL (too many): {line}")
        elif and_count + or_count == 1:
            op = " and " if and_count else " or "
            left, right = cond.split(op, 1)
            left_neg = "not " in left or "!=" in left
            right_neg = "not " in right or "!=" in right
            if left_neg != right_neg:
                print(f"FAIL (mixed polarity): {line}")
