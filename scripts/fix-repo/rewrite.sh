#!/usr/bin/env bash
# Rewrite engine for fix-repo.sh

# Echoes space-separated target versions
get_target_versions() {
  local current="$1" span="$2"
  local start=$((current - span))
  if [ "$start" -lt 1 ]; then start=1; fi
  local end=$((current - 1))
  if [ "$end" -lt "$start" ]; then return 0; fi
  seq "$start" "$end" | tr '\n' ' '
}

# Echoes a Perl regex matching {base}-v{n} not followed by a digit
build_pattern() {
  local base="$1" n="$2"
  local escaped
  escaped="$(printf '%s' "$base-v$n" | perl -pe 's/([\\\.\^\$\*\+\?\(\)\[\]\{\}\|])/\\\1/g')"
  printf '%s(?!\\d)' "$escaped"
}

# Echoes the number of replacements made; writes file unless DRY=1
rewrite_file() {
  local path="$1" base="$2" current="$3" dry="$4"; shift 4
  local targets=("$@")
  local total=0
  local tmp
  tmp="$(mktemp)"
  cp "$path" "$tmp"
  local n pattern replacement count
  replacement="$base-v$current"
  for n in "${targets[@]}"; do
    pattern="$(build_pattern "$base" "$n")"
    count="$(perl -ne "while(/$pattern/g){print qq(x\n)}" "$tmp" | wc -l | tr -d ' ')"
    if [ "$count" -gt 0 ]; then
      perl -i -pe "s/$pattern/$replacement/g" "$tmp"
      total=$((total + count))
    fi
  done
  if [ "$total" -gt 0 ] && [ "$dry" != "1" ]; then
    cp "$tmp" "$path"
  fi
  rm -f "$tmp"
  printf '%s' "$total"
}
