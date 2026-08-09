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

for (const file of files) {
    const content = fs.readFileSync(file, 'utf8');
    const lines = content.split('\n');
    let inExcept = false;
    let exceptStart = 0;
    let hasLog = false;
    let exceptLine = 0;
    
    for (let i = 0; i < lines.length; i++) {
        const line = lines[i];
        
        const exceptMatch = line.match(/^(\s*)except(\b|:)/);
        if (exceptMatch) {
            inExcept = true;
            exceptStart = exceptMatch[1].length;
            hasLog = false;
            exceptLine = i + 1;
            if (line.trim().endsWith(':') === false && line.includes(' pass')) {
                inExcept = false;
                hasLog = false;
                console.log(`${file}:${exceptLine}: Missing explicit log in except block (one-liner)`);
            }
        } else if (inExcept) {
            const indentMatch = line.match(/^(\s*)/);
            if (line.trim() === '' || line.trim().startsWith('#')) {
                // ignore
            } else if (indentMatch && indentMatch[1].length <= exceptStart) {
                if (!hasLog) {
                    console.log(`${file}:${exceptLine}: Missing explicit log in except block`);
                }
                inExcept = false;
                
                const reExceptMatch = line.match(/^(\s*)except(\b|:)/);
                if (reExceptMatch) {
                    inExcept = true;
                    exceptStart = reExceptMatch[1].length;
                    hasLog = false;
                    exceptLine = i + 1;
                }
            } else {
                if (line.includes('print(') || line.includes('log.') || line.includes('logger.') || line.includes('logging.')) {
                    hasLog = true;
                }
            }
        }
    }
    if (inExcept && !hasLog) {
        console.log(`${file}:${exceptLine}: Missing explicit log in except block (EOF)`);
    }
}
