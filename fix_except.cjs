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
    let content = fs.readFileSync(file, 'utf8');
    let lines = content.split('\n');
    let modified = false;
    
    // We process backwards so line additions don't shift line numbers for the rest of the file
    for (let i = lines.length - 1; i >= 0; i--) {
        const line = lines[i];
        const exceptMatch = line.match(/^(\s*)except(\b(.*?):|:)(.*)/);
        
        if (exceptMatch) {
            const indent = exceptMatch[1];
            let afterExcept = exceptMatch[3] || "";
            let afterColon = exceptMatch[4] || "";
            
            // Check if it already logs within the block
            let hasLog = false;
            let blockEnd = i + 1;
            while (blockEnd < lines.length) {
                const innerIndentMatch = lines[blockEnd].match(/^(\s*)/);
                if (lines[blockEnd].trim() === '' || lines[blockEnd].trim().startsWith('#')) {
                    blockEnd++;
                    continue;
                }
                if (innerIndentMatch[1].length <= indent.length) {
                    break;
                }
                if (lines[blockEnd].includes('print(') || lines[blockEnd].includes('log.') || lines[blockEnd].includes('logger.') || lines[blockEnd].includes('logging.')) {
                    hasLog = true;
                    break;
                }
                blockEnd++;
            }
            
            if (line.includes(' pass') && !line.trim().endsWith(':')) {
                // One-liner
                hasLog = false;
            }
            
            if (!hasLog) {
                // Need to add log
                let varName = "exc";
                let newExceptLine = line;
                
                if (afterExcept.trim() === "") {
                    newExceptLine = `${indent}except Exception as exc:${afterColon}`;
                } else if (!afterExcept.includes(' as ')) {
                    newExceptLine = `${indent}except ${afterExcept.trim()} as exc:${afterColon}`;
                } else {
                    const asMatch = afterExcept.match(/ as (\w+)/);
                    if (asMatch) varName = asMatch[1];
                }
                
                lines[i] = newExceptLine;
                
                // Ensure import sys is at the top if we use sys.stderr
                // Actually we can just print without sys.stderr or use a simple print
                const logLine = `${indent}    import sys; print(f"Error: {${varName}}", file=sys.stderr)`;
                
                if (afterColon.trim() !== "") {
                    // One liner like: except: pass
                    // we split it
                    lines[i] = `${indent}except Exception as exc:`;
                    lines.splice(i + 1, 0, logLine);
                    lines.splice(i + 2, 0, `${indent}    ${afterColon.trim()}`);
                } else {
                    // insert log statement at i + 1
                    lines.splice(i + 1, 0, logLine);
                }
                modified = true;
            }
        }
    }
    
    if (modified) {
        fs.writeFileSync(file, lines.join('\n'), 'utf8');
        console.log(`Modified: ${file}`);
    }
}
