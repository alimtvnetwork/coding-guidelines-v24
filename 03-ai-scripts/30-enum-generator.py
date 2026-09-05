#!/usr/bin/env python3
"""
30-enum-generator.py — Generates type-safe Go enums conforming to BaseEnum.

Generates idiomatic Go enum types that implement:
  - BaseEnum (Name, String, ValueString, IsValid, IsEnum)
  - NumberEnum (for integer-backed enums: Int, Code)
  - IsCompare(target)
  - json.Marshaler (MarshalJSON)
  - json.Unmarshaler (UnmarshalJSON)
  - All*() slice of all values
  - Parse*(string) case-insensitive parser

Usage:
  python 03-ai-scripts/30-enum-generator.py --name ConnectionState --type string --members Disconnected,Connecting,Connected,Reconnecting --package connection
  python 03-ai-scripts/30-enum-generator.py --name Priority --type int --members Low,Medium,High,Critical --package task
  python 03-ai-scripts/30-enum-generator.py --name DeliveryMode --members Fast,Standard,Scheduled --out 04-code/golang/pkg/streamwriter/delivery_mode.go
"""

from __future__ import annotations

import argparse
from importlib import import_module
from pathlib import Path
import sys

sys.path.insert(0, str(Path(__file__).parent))
engine = import_module("02-shared-engine")

ExitCodeType = engine.ExitCodeType


def generate_string_enum(name: str, pkg: str, members: list[str]) -> str:
    enum_type = f"{name}Type" if not name.endswith("Type") else name
    registry_name = f"{enum_type[0].lower() + enum_type[1:]}Registry"
    all_func = f"All{name}s" if not name.endswith("s") else f"All{name}"
    parse_func = f"Parse{name}"

    lines = [
        f"package {pkg}",
        "",
        'import (',
        '\t"encoding/json"',
        '\t"strings"',
        ')',
        "",
        f"// {enum_type} represents string-backed enum values conforming to BaseEnum.",
        f"type {enum_type} string",
        "",
        "const (",
    ]

    for m in members:
        lines.append(f'\t{name}{m} {enum_type} = "{m}"')

    lines.append(f'\t{name}Unknown {enum_type} = "Unknown"')
    lines.append(")")
    lines.append("")

    lines.append(f"var {registry_name} = map[{enum_type}]bool{{")
    for m in members:
        lines.append(f"\t{name}{m}: true,")

    lines.append("}")
    lines.append("")

    lines.extend([
        f"// Name returns the uppercase identifier.",
        f"func (e {enum_type}) Name() string {{",
        "\treturn string(e)",
        "}",
        "",
        f"// String implements fmt.Stringer.",
        f"func (e {enum_type}) String() string {{",
        "\treturn string(e)",
        "}",
        "",
        f"// ValueString returns the string representation of value.",
        f"func (e {enum_type}) ValueString() string {{",
        "\treturn string(e)",
        "}",
        "",
        f"// Value returns the raw string value.",
        f"func (e {enum_type}) Value() string {{",
        "\treturn string(e)",
        "}",
        "",
        f"// IsValid returns true if this enum value is recognized.",
        f"func (e {enum_type}) IsValid() bool {{",
        f"\treturn {registry_name}[e]",
        "}",
        "",
        f"// IsEnum returns true if this enum exists in the registry.",
        f"func (e {enum_type}) IsEnum() bool {{",
        f"\treturn {registry_name}[e]",
        "}",
        "",
        f"// IsCompare checks equality against another {enum_type}.",
        f"func (e {enum_type}) IsCompare(target {enum_type}) bool {{",
        "\treturn e == target",
        "}",
        "",
        f"// MarshalJSON implements json.Marshaler.",
        f"func (e {enum_type}) MarshalJSON() ([]byte, error) {{",
        "\treturn json.Marshal(string(e))",
        "}",
        "",
        f"// UnmarshalJSON implements json.Unmarshaler.",
        f"func (e *{enum_type}) UnmarshalJSON(data []byte) error {{",
        "\tvar raw string",
        "\tif err := json.Unmarshal(data, &raw); err != nil {",
        "\t\treturn err",
        "\t}",
        "",
        f"\t*e = {parse_func}(raw)",
        "",
        "\treturn nil",
        "}",
        "",
        f"// {all_func} returns all valid {enum_type} values.",
        f"func {all_func}() []{enum_type} {{",
        f"\treturn []{enum_type}{{",
    ])

    for m in members:
        lines.append(f"\t\t{name}{m},")

    lines.extend([
        "\t}",
        "}",
        "",
        f"// {parse_func} parses string into {enum_type} case-insensitively.",
        f"func {parse_func}(val string) {enum_type} {{",
        f"\tfor _, candidate := range {all_func}() {{",
        "\t\tif strings.EqualFold(string(candidate), strings.TrimSpace(val)) {",
        "\t\t\treturn candidate",
        "\t\t}",
        "\t}",
        "",
        f"\treturn {name}Unknown",
        "}",
        "",
    ])

    return "\n".join(lines)


def generate_int_enum(name: str, pkg: str, members: list[str]) -> str:
    enum_type = f"{name}Type" if not name.endswith("Type") else name
    names_map = f"{enum_type[0].lower() + enum_type[1:]}Names"
    all_func = f"All{name}s" if not name.endswith("s") else f"All{name}"
    parse_func = f"Parse{name}"

    lines = [
        f"package {pkg}",
        "",
        'import (',
        '\t"encoding/json"',
        '\t"fmt"',
        '\t"strings"',
        ')',
        "",
        f"// {enum_type} represents integer-backed enum values conforming to BaseEnum and NumberEnum.",
        f"type {enum_type} uint16",
        "",
        "const (",
    ]

    for idx, m in enumerate(members, start=1):
        lines.append(f"\t{name}{m} {enum_type} = {idx}")

    lines.append(")")
    lines.append("")

    lines.append(f"var {names_map} = map[{enum_type}]string{{")
    for m in members:
        lines.append(f'\t{name}{m}: "{m}",')

    lines.append("}")
    lines.append("")

    lines.extend([
        f"// Name returns the uppercase identifier.",
        f"func (e {enum_type}) Name() string {{",
        f"\tif name, ok := {names_map}[e]; ok {{",
        "\t\treturn name",
        "\t}",
        "",
        f'\treturn fmt.Sprintf("{name}(%d)", uint16(e))',
        "}",
        "",
        f"// String implements fmt.Stringer.",
        f"func (e {enum_type}) String() string {{",
        "\treturn e.Name()",
        "}",
        "",
        f"// ValueString returns code as string.",
        f"func (e {enum_type}) ValueString() string {{",
        '\treturn fmt.Sprintf("%d", uint16(e))',
        "}",
        "",
        f"// Code returns uint16 code value.",
        f"func (e {enum_type}) Code() uint16 {{",
        "\treturn uint16(e)",
        "}",
        "",
        f"// Int returns int representation.",
        f"func (e {enum_type}) Int() int {{",
        "\treturn int(e)",
        "}",
        "",
        f"// IsValid returns true if this enum value is registered.",
        f"func (e {enum_type}) IsValid() bool {{",
        f"\t_, ok := {names_map}[e]",
        "",
        "\treturn ok",
        "}",
        "",
        f"// IsEnum returns true if this enum exists in the registry.",
        f"func (e {enum_type}) IsEnum() bool {{",
        f"\t_, ok := {names_map}[e]",
        "",
        "\treturn ok",
        "}",
        "",
        f"// IsCompare checks equality against another {enum_type}.",
        f"func (e {enum_type}) IsCompare(target {enum_type}) bool {{",
        "\treturn e == target",
        "}",
        "",
        f"// MarshalJSON implements json.Marshaler.",
        f"func (e {enum_type}) MarshalJSON() ([]byte, error) {{",
        "\treturn json.Marshal(e.Name())",
        "}",
        "",
        f"// UnmarshalJSON implements json.Unmarshaler.",
        f"func (e *{enum_type}) UnmarshalJSON(data []byte) error {{",
        "\tvar raw string",
        "\tif err := json.Unmarshal(data, &raw); err == nil {",
        f"\t\t*e = {parse_func}(raw)",
        "",
        "\t\treturn nil",
        "\t}",
        "",
        "\tvar code uint16",
        "\tif err := json.Unmarshal(data, &code); err != nil {",
        "\t\treturn err",
        "\t}",
        "",
        f"\t*e = {enum_type}(code)",
        "",
        "\treturn nil",
        "}",
        "",
        f"// {all_func} returns all valid {enum_type} values.",
        f"func {all_func}() []{enum_type} {{",
        f"\treturn []{enum_type}{{",
    ])

    for m in members:
        lines.append(f"\t\t{name}{m},")

    lines.extend([
        "\t}",
        "}",
        "",
        f"// {parse_func} parses string into {enum_type} case-insensitively.",
        f"func {parse_func}(val string) {enum_type} {{",
        "\tcleaned := strings.TrimSpace(val)",
        f"\tfor code, name := range {names_map} {{",
        "\t\tif strings.EqualFold(name, cleaned) {",
        "\t\t\treturn code",
        "\t\t}",
        "\t}",
        "",
        "\treturn 0",
        "}",
        "",
    ])

    return "\n".join(lines)


def main() -> int:
    parser = argparse.ArgumentParser(description="Generate type-safe Go enum conforming to BaseEnum")
    parser.add_argument("--name", required=True, help="Base name of the enum (e.g. ProcessState, Priority)")
    parser.add_argument("--type", choices=["string", "int"], default="string", help="Backing type (default: string)")
    parser.add_argument("--members", required=True, help="Comma-separated member names (e.g. Pending,Running,Failed)")
    parser.add_argument("--package", default="enums", help="Target Go package name (default: enums)")
    parser.add_argument("--out", help="Optional output file path")

    args = parser.parse_args()
    members = [m.strip() for m in args.members.split(",") if m.strip()]

    if not members:
        print("Error: at least one member name must be provided", file=sys.stderr)
        return ExitCodeType.ARGUMENT_ERROR.value

    if args.type == "string":
        code = generate_string_enum(args.name, args.package, members)
    else:
        code = generate_int_enum(args.name, args.package, members)

    if args.out:
        out_path = Path(args.out)
        out_path.parent.mkdir(parents=True, exist_ok=True)
        out_path.write_text(code, encoding="utf-8")
        print(f"Generated enum {args.name} in {out_path}")
    else:
        print(code)

    return ExitCodeType.SUCCESS.value


if __name__ == "__main__":
    sys.exit(main())
