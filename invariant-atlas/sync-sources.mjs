import {readdirSync,readFileSync,writeFileSync,mkdirSync} from 'node:fs';
import {fileURLToPath} from 'node:url';
import path from 'node:path';
const root=path.dirname(fileURLToPath(import.meta.url)), repo=path.dirname(root), files={};
mkdirSync(path.join(root,'dist/sources'),{recursive:true});
for(const dir of readdirSync(repo)) {
 const n=Number(dir.slice(0,3));
 if(!/^\d{3}-/.test(dir)||n<82||n>120) continue;
 const name=`${dir}/main.go`;
 writeFileSync(path.join(root,'dist/sources',`${n}.txt`),readFileSync(path.join(repo,name)));
 files[n]=name;
}
writeFileSync(path.join(root,'dist/sources.js'),`const SOURCES = ${JSON.stringify(files,null,2)};\n`);
console.log(`Saved ${Object.keys(files).length} source snapshots.`);
