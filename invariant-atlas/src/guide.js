'use strict';
const escapeHTML=s=>String(s).replace(/[&<>"']/g,c=>({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;',"'":'&#39;'}[c]));
function openPattern(id){location.href='implementations.html#'+encodeURIComponent(id);}
// Keep old links to invariant notes working after moving the catalog.
const oldId=PATTERN_ALIASES[location.hash.slice(1)]||location.hash.slice(1);
if(PATTERNS.some(p=>p.id===oldId))location.replace('implementations.html#'+encodeURIComponent(oldId));
