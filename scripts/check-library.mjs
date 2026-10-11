import {readFileSync,readdirSync,existsSync,statSync} from 'node:fs';
import path from 'node:path';
import {fileURLToPath} from 'node:url';
import assert from 'node:assert/strict';
const root=path.dirname(path.dirname(fileURLToPath(import.meta.url)));
const {groups,entries}=JSON.parse(readFileSync(path.join(root,'code-library.json'),'utf8'));
assert.equal(new Set(entries.map(e=>e.id)).size,entries.length,'Duplicate reference');
assert.equal(new Set(entries.map(e=>e.path)).size,entries.length,'Duplicate path');
for(const group of groups){
 const members=entries.filter(e=>e.group===group.prefix);
 assert.ok(members.length,'Empty group '+group.prefix);
 members.forEach((e,i)=>{
  const prefix=group.prefix+'-'+String(i+1).padStart(2,'0')+'-';
  assert.ok(e.id.startsWith(prefix),'Nonsequential reference '+e.id);
  assert.equal(e.path,group.folder+'/'+e.id);
  assert.ok(existsSync(path.join(root,e.path,'main.go')),'Missing implementation '+e.id);
 });
 const actual=readdirSync(path.join(root,group.folder)).filter(name=>statSync(path.join(root,group.folder,name)).isDirectory()).map(name=>group.folder+'/'+name).sort();
 assert.deepEqual(actual,members.map(e=>e.path).sort(),'Unindexed folder in '+group.folder);
}
assert.ok(entries.every(e=>groups.some(g=>g.prefix===e.group)),'Unknown group');
assert.ok(!readdirSync(root).some(n=>/^\d{3}-/.test(n)),'Old flat numbered folders remain');
console.log(`Library audit passed: ${entries.length} implementations, ${groups.length} groups, contiguous numbering and valid paths.`);
