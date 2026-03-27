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
func_open_re = re.compile(r'^\t\t[^\t].*\(func\(.*\{\s*$')
block_open_re = re.compile(r'^\t\t[^\t].*\{\s*$')

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

# ── Check 3: empty if blocks (comment-only body with no actual statement) ──
for f in "$DIR"/*_test.go; do
  [ -f "$f" ] || continue

  if ! python3 - "$f" <<'PY2' 2>/dev/null
import sys, re
path = sys.argv[1]
lines = open(path, encoding="utf-8").read().splitlines()
if_re = re.compile(r'^(\t+)if\b.*\{\s*$')
for i, line in enumerate(lines):
    m = if_re.match(line)
    if not m:
        continue
    indent = m.group(1)
    close = indent + '}'
    # scan body: skip blanks and comments, expect at least one real statement
    has_stmt = False
    for j in range(i + 1, min(i + 20, len(lines))):
        body = lines[j]
        stripped = body.strip()
        if body.rstrip() == close or body.rstrip() == close + ')':
            break
        if stripped == '' or stripped.startswith('//'):
            continue
        has_stmt = True
        break
    if not has_stmt:
        print(f'  {path}:{i+1}: empty if block (no statements)')
        sys.exit(1)
sys.exit(0)
PY2
  then
    ISSUES=1
  fi
done

# ── Check 4: func assignment closure must end with `}` (not `})`) ──
for f in "$DIR"/*_test.go; do
  [ -f "$f" ] || continue

  if ! python3 - "$f" <<'PY3' 2>/dev/null
import re
import sys

path = sys.argv[1]
lines = open(path, encoding="utf-8").read().splitlines()
assign_open_re = re.compile(r'^\s*(?:var\s+\w+\s*=\s*func\s*\(|\w+\s*(?::=|=)\s*func\s*\()')

for i, line in enumerate(lines):
    if not assign_open_re.search(line):
        continue

    depth = 0
    started = False
    for j in range(i, min(i + 300, len(lines))):
        l = lines[j]
        open_count = l.count('{')
        close_count = l.count('}')
        if open_count > 0:
            started = True
        depth += open_count
        depth -= close_count

        if started and depth == 0:
            if lines[j].rstrip().endswith('})'):
                print(f'  {path}:{j+1}: func assignment closes with }) (should close with } only)')
                sys.exit(1)
            break

sys.exit(0)
PY3
  then
    ISSUES=1
  fi
done

# ── Check 5: placeholder `...` lines are forbidden ──
for f in "$DIR"/*_test.go; do
  [ -f "$f" ] || continue

  if ! python3 - "$f" <<'PY4' 2>/dev/null
import sys

path = sys.argv[1]
lines = open(path, encoding="utf-8").read().splitlines()
for i, line in enumerate(lines):
    if line.strip() == '...':
        print(f'  {path}:{i+1}: placeholder line ... is not allowed')
        sys.exit(1)

sys.exit(0)
PY4
  then
    ISSUES=1
  fi
done

if [ "$ISSUES" -ne 0 ]; then
  echo ""
  echo "✗ Malformed safeTest boundaries, empty if blocks, func assignment closures, or placeholder lines detected."
  exit 1
else
  echo "✓ All safeTest boundaries, if blocks, assignment closures, and placeholder lines are clean."
  exit 0
fi
