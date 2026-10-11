// Versioned migration keeps old review marks without repeating a split after undo.
function loadReviewed(storage) {
 let ids=[];
 try {
  const saved=JSON.parse(storage.getItem('invariant-atlas-reviewed')||'[]');
  if(!Array.isArray(saved)) return new Set();
  ids=storage.getItem('invariant-atlas-review-schema')==='2'
   ? saved.filter(id=>PATTERNS.some(p=>p.id===id)) : migrateReviewed(saved);
  storage.setItem('invariant-atlas-reviewed',JSON.stringify(ids));
  storage.setItem('invariant-atlas-review-schema','2');
 } catch {} // Reading still works if writes are blocked.
 return new Set(ids);
}
