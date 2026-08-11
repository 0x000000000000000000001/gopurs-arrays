const fs = require('fs');
const data = JSON.parse(fs.readFileSync('output/Data.Ord/corefn.json', 'utf8'));
const ord = data.classDecls.find(c => c.name === 'Ord');
console.log(JSON.stringify(ord, null, 2));
