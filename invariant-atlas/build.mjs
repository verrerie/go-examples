import {cpSync, mkdirSync} from 'node:fs';
const source = new URL('./src/', import.meta.url);
const output = new URL('./dist/', import.meta.url);
mkdirSync(output, {recursive: true});
cpSync(source, output, {recursive: true});
await import('./sync-sources.mjs');
console.log('Built Invariant Atlas in dist/.');
