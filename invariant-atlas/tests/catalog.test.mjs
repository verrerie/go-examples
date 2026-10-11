import {readFileSync,existsSync,readdirSync} from 'node:fs';
import vm from 'node:vm';
import test from 'node:test';
import assert from 'node:assert/strict';
const root=new URL('../',import.meta.url);
const context={};vm.createContext(context);
vm.runInContext(readFileSync(new URL('src/data.js',root),'utf8')+readFileSync(new URL('src/progress.js',root),'utf8')+';globalThis.catalog={PATTERNS,CHEATSHEET,PATTERN_ALIASES,loadReviewed};',context);
const {PATTERNS:patterns,CHEATSHEET:coverage,PATTERN_ALIASES:aliases,loadReviewed}=context.catalog;
function storage(ids,version){const values=new Map([['invariant-atlas-reviewed',JSON.stringify(ids)]]);if(version)values.set('invariant-atlas-review-schema',version);return {getItem:k=>values.get(k)??null,setItem:(k,v)=>values.set(k,v)};}
test('unique cards and complete cheatsheet targets',()=>{
 assert.equal(patterns.length,69);assert.equal(new Set(patterns.map(p=>p.id)).size,patterns.length);
 assert.equal(coverage.length,27);
 for(const [,id] of coverage)assert.ok(patterns.some(p=>p.id===id),id);
 for(const id of Object.values(aliases))assert.ok(patterns.some(p=>p.id===id),id);
 const dirs=readdirSync(new URL('../../',import.meta.url));
 for(const p of patterns){for(const field of ['invariant','step','stop','trap','cost','example'])assert.ok(p[field],p.id+field);for(const n of p.sources){const dir=dirs.find(d=>d.startsWith(String(n).padStart(3,'0')+'-'));assert.ok(dir&&existsSync(new URL('../../'+dir+'/main.go',import.meta.url)),String(n));}}
});
test('merge deduplicates reviews and splits retain review history',()=>{
 const s=storage(['two-pointers','sorted-two-sum','fast-slow','quicksort','inorder','prefix','obsolete']);
 assert.deepEqual([...loadReviewed(s)].sort(),['sorted-two-sum','fast-slow','cycle-detection','quicksort','two-way-partition','inorder','kth-inorder','prefix'].sort());
 assert.equal(s.getItem('invariant-atlas-review-schema'),'2');
 // Undo one split child; reloading must not mark it reviewed again.
 const ids=[...loadReviewed(s)].filter(id=>id!=='cycle-detection');s.setItem('invariant-atlas-reviewed',JSON.stringify(ids));
 assert.equal(loadReviewed(s).has('cycle-detection'),false);
});
test('new users, malformed data, and blocked writes',()=>{
 assert.equal(loadReviewed(storage([])).size,0);
 assert.equal(loadReviewed(storage({bad:true})).size,0);
 const bad=storage([]);bad.setItem('invariant-atlas-reviewed','not json');assert.equal(loadReviewed(bad).size,0);
 const blocked=storage(['two-pointers']);blocked.setItem=()=>{throw Error('blocked')};assert.equal(loadReviewed(blocked).has('sorted-two-sum'),true);
});
