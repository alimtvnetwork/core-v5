const fs = require('fs');
const path = require('path');

const OLD = 'gitlab.com/auk-go/core';
const NEW = 'github.com/alimtvnetwork/core';

function walk(dir) {
  let results = [];
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    const full = path.join(dir, entry.name);
    if (entry.name === 'node_modules' || entry.name === '.git') continue;
    if (entry.isDirectory()) results.push(...walk(full));
    else if (entry.name.endsWith('.go')) results.push(full);
  }
  return results;
}

let count = 0;
for (const file of walk('.')) {
  const content = fs.readFileSync(file, 'utf8');
  if (content.includes(OLD)) {
    fs.writeFileSync(file, content.replaceAll(OLD, NEW), 'utf8');
    count++;
  }
}
console.log(`Updated ${count} files.`);
