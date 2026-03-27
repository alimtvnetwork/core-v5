#!/usr/bin/env bash
# check-safetest-boundaries.sh
# Detects malformed safeTest/function boundaries in corestrtests.
# Exit 1 if issues found, 0 if clean.

set -euo pipefail

DIR="tests/integratedtests/corestrtests"
ISSUES=0

if [ ! -d "$DIR" ]; then
  echo "⚠ Directory $DIR not found, skipping check."
  exit 0
fi

for f in "$DIR"/*_test.go; do
  [ -f "$f" ] || continue

  # Check 1: 3-tab statement immediately followed by \t}) (missing inner block close)
  awk '
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
  ' "$f" 2>/dev/null && true
  if [ $? -ne 0 ]; then
    ISSUES=1
  fi

  # Check 2: closure arg } missing ) — detects \t\t} where a func( opened above
  python3 -c "
import re, sys
lines = open('$f').readlines()
for i, line in enumerate(lines):
    if line.rstrip() == '\t\t}':
        depth = 0
        for j in range(i-1, max(i-30,-1), -1):
            l = lines[j].rstrip()
            if l == '\t\t})': depth += 1
            elif re.search(r'^\t\t\S.*func\(.*\{$', l) or re.search(r'^\t\t\S.*\(func\(', l):
                if depth > 0: depth -= 1
                else: print(f'  $f:{i+1}: closure } missing )'); sys.exit(1)
            elif l.startswith('func Test_') or l.startswith('\tsafeTest('): break
            elif l == '\t\t}': break
" 2>/dev/null && true
  if [ $? -ne 0 ]; then
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
