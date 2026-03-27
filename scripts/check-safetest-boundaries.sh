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
  # Detect: 3-tab-indented non-brace statement immediately followed by \t})
  # This means a closing } for an inner block is missing.
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
done

if [ "$ISSUES" -ne 0 ]; then
  echo ""
  echo "✗ Malformed safeTest boundaries detected. Fix missing closing braces."
  exit 1
else
  echo "✓ All safeTest boundaries are clean."
  exit 0
fi
