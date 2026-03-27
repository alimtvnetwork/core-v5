#!/usr/bin/env bash
set -euo pipefail

DIR="tests/integratedtests/corestrtests"
ISSUES=0

if [ ! -d "$DIR" ]; then
  echo "⚠ Directory $DIR not found, skipping check."
  exit 0
fi

for f in "$DIR"/*_test.go; do
  [ -f "$f" ] || continue

  # Check 1: malformed safeTest boundary (missing inner `}` before `\t})`)
  if ! awk '
    NR > 1 {
      if ($0 ~ /^\t\)\}$/ || $0 == "\t})") {
        if (prev ~ /^\t\t\t/ && prev !~ /^\t\t\t[[:space:]]*\}[[:space:]]*$/ && prev !~ /^\t\t\t[[:space:]]*\/\//) {
          printf "  %s:%d: missing closing } before %})\n", FILENAME, NR
          found++
        }
      }
      prev = $0
    }
    NR == 1 { prev = $0 }
    END { exit (found > 0 ? 1 : 0) }
  ' "$f" 2>/dev/null; then
    ISSUES=1
  fi

  # Check 2: closure arg `}` missing `)` (avoid false-positives on if/for/switch/select blocks)
  if ! python3 - "$f" <<'PY' 2>/dev/null
import re
import sys

path = sys.argv[1]
lines = open(path, encoding="utf-8").read().splitlines()

control_re = re.compile(r'^\t\t(?:if|for|switch|select|else)\b')
func_open_re = re.compile(r'^\t\t(?!if\b|for\b|switch\b|select\b|else\b).*(?:\bfunc\(|\(func\().*\{\s*$')
block_open_re = re.compile(r'^\t\t.*\{\s*$')

for i, line in enumerate(lines):
    if line.rstrip() != '\t\t}':
        continue

    depth = 0
    for j in range(i - 1, max(i - 40, -1), -1):
        l = lines[j].rstrip()

        if not l:
            continue

        if l == '\t\t})':
            depth += 1
            continue

        if l.startswith('func Test_') or l.startswith('\tsafeTest('):
            break

        if l == '\t\t}':
            if depth == 0:
                break
            continue

        if block_open_re.search(l):
            if control_re.search(l):
                break

            if func_open_re.search(l):
                if depth > 0:
                    depth -= 1
                else:
                    print(f'  {path}:{i+1}: closure }} missing )')
                    sys.exit(1)
            break

sys.exit(0)
PY
  then
    ISSUES=1
  fi
done

if [ "$ISSUES" -ne 0 ]; then
  echo ""
  echo "✗ Malformed safeTest boundaries detected. Fix missing closing braces."
  exit 1
else
  echo "✓ All safeTest boundaries are clean."
  exit 0
fi
