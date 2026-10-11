import {readFileSync,writeFileSync,mkdirSync,readdirSync,unlinkSync,existsSync} from 'node:fs';
import {fileURLToPath} from 'node:url';
import path from 'node:path';
const root=path.dirname(fileURLToPath(import.meta.url)),repo=path.dirname(root);
const catalog=JSON.parse(readFileSync(path.join(repo,'code-library.json'),'utf8'));
const sources={},metadata={},output=path.join(root,'dist/sources');
mkdirSync(output,{recursive:true});
for(const entry of catalog.entries){
 const name=entry.path+'/main.go';const full=path.resolve(repo,name);
 if(!full.startsWith(repo+path.sep))throw Error('Source outside repository');
 let text=readFileSync(full,'utf8');
 const scaffold=path.join(repo,entry.path,'scaffold.go');
 if(existsSync(scaffold))text+='\n\n// --- scaffold.go (original teaching template) ---\n'+readFileSync(scaffold,'utf8');
 sources[entry.id]=name;metadata[entry.id]=entry;
 writeFileSync(path.join(output,entry.id+'.txt'),text);
}
// Remove obsolete generated snapshots only inside this exact output directory.
for(const file of readdirSync(output)){if(file.endsWith('.txt')&&!Object.hasOwn(sources,file.slice(0,-4)))unlinkSync(path.join(output,file));}
writeFileSync(path.join(root,'dist/sources.js'),`const SOURCES = ${JSON.stringify(sources,null,2)};\nconst SOURCE_META = ${JSON.stringify(metadata,null,2)};\nconst SOURCE_GROUPS = ${JSON.stringify(catalog.groups)};\n`);
console.log(`Saved ${catalog.entries.length} source snapshots across ${catalog.groups.length} groups.`);
