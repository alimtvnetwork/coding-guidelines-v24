const fs = require('fs');
const path = require('path');

function walk(dir) {
    let results = [];
    const list = fs.readdirSync(dir);
    list.forEach(file => {
        file = path.join(dir, file);
        const stat = fs.statSync(file);
        if (stat && stat.isDirectory()) {
            results = results.concat(walk(file));
        } else if (file.endsWith('.py')) {
            results.push(file);
        }
    });
    return results;
}

const files = [...walk('d:/work/coding-guidelines/linter-scripts'), ...walk('d:/work/coding-guidelines/linters-cicd')];
let found = false;

for (const file of files) {
    const content = fs.readFileSync(file, 'utf8');
    const lines = content.split('\n');
    
    for (let i = 0; i < lines.length; i++) {
        const line = lines[i];
        
        if (line.match(/\bnot\s+is_[a-zA-Z0-9_]*(valid|success|pass)[a-zA-Z0-9_]*\b/i)) {
            console.log(`[INVERTED] ${file}:${i+1}: ${line.trim()}`);
            found = true;
        }
        if (line.match(/\bis_[a-zA-Z0-9_]*(valid|success|pass)[a-zA-Z0-9_]*\s*==\s*False\b/i)) {
            console.log(`[INVERTED] ${file}:${i+1}: ${line.trim()}`);
            found = true;
        }
    }
}
if (!found) console.log('None found!');
