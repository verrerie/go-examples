import {readFileSync,existsSync,readdirSync} from 'node:fs';
import vm from 'node:vm';
import test from 'node:test';
import assert from 'node:assert/strict';
const root=new URL('../',import.meta.url);
const context={};vm.createContext(context);
vm.runInContext(readFileSync(new URL('dist/sources.js',root),'utf8')+readFileSync(new URL('src/data.js',root),'utf8')+readFileSync(new URL('src/progress.js',root),'utf8')+';globalThis.catalog={PATTERNS,CHEATSHEET,PATTERN_ALIASES,loadReviewed,SOURCE_META,SOURCES};',context);
const {PATTERNS:patterns,CHEATSHEET:coverage,PATTERN_ALIASES:aliases,loadReviewed,SOURCE_META:metadata,SOURCES:sources}=context.catalog;
function storage(ids,version){const values=new Map([['invariant-atlas-reviewed',JSON.stringify(ids)]]);if(version)values.set('invariant-atlas-review-schema',version);return {getItem:k=>values.get(k)??null,setItem:(k,v)=>values.set(k,v)};}
test('unique cards and complete cheatsheet targets',()=>{
 assert.equal(patterns.length,69);assert.equal(new Set(patterns.map(p=>p.id)).size,patterns.length);
 assert.equal(coverage.length,27);
 for(const [,id] of coverage)assert.ok(patterns.some(p=>p.id===id),id);
 for(const id of Object.values(aliases))assert.ok(patterns.some(p=>p.id===id),id);
 for(const p of patterns){for(const field of ['invariant','step','stop','trap','cost','example'])assert.ok(p[field],p.id+field);for(const id of p.sources){assert.ok(metadata[id],id);assert.ok(existsSync(new URL('../../'+sources[id],import.meta.url)),id);assert.ok(existsSync(new URL('dist/sources/'+id+'.txt',root)),id);}}
 assert.equal(Object.keys(metadata).length,153);
 assert.ok(patterns.every(p=>p.sources.every(id=>typeof id==='string')));

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

test('new references, old numbers, and group filters find code',()=>{
 const app=readFileSync(new URL('src/app.js',root),'utf8');
 const helpers=app.slice(app.indexOf('function matches('),app.indexOf('function renderCodeIndex('));
 vm.runInContext('var query="";var group="All patterns";var unreviewed=false;var reviewed=new Set();'+helpers+';globalThis.findCode=codeMatches;globalThis.findNotes=filtered;',context);
 context.query='129';assert.deepEqual(Array.from(context.findCode(),m=>m.id),['hash-03-subarray-sum-count']);
 context.query='arr-04-prefix-sum';assert.equal(context.findNotes()[0].id,'prefix');
 context.query='119';assert.ok(context.findCode().some(m=>m.id==='arr-03-consecutive-ones'));
 context.query='';context.group='Go language and features';assert.equal(context.findCode().length,85);assert.equal(context.findNotes().length,0);
 context.query='closures';assert.equal(context.findCode()[0].id,'golang-14-closures');
 context.group='All patterns';context.query='cycle';assert.ok(context.findNotes().some(p=>p.id==='cycle-detection'));
});
