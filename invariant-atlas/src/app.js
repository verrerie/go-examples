'use strict';
const $=s=>document.querySelector(s);
const escapeHTML=s=>String(s).replace(/[&<>"']/g,c=>({'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;',"'":'&#39;'}[c]));
const groups=['All patterns',...new Set(PATTERNS.map(p=>p.group))];
let group='All patterns',query='',recall=false,unreviewed=false,current=null,revealed=false,traceIndex=0;
let reviewed=new Set();
try{const saved=JSON.parse(localStorage.getItem('invariant-atlas-reviewed')||'[]');if(Array.isArray(saved))reviewed=new Set(saved.filter(id=>PATTERNS.some(p=>p.id===id)));}catch{}
function filtered(){return PATTERNS.filter(p=>(group==='All patterns'||p.group===group)&&(!unreviewed||!reviewed.has(p.id))&&[p.title,p.group,p.invariant,p.trap,...p.sources.map(n=>`${n} ${SOURCES[n]}`)].join(' ').toLowerCase().includes(query.toLowerCase()));}
function render(){
 $('#categories').innerHTML=groups.map(g=>`<button class="category ${g===group?'active':''}" data-group="${g}" aria-pressed="${g===group}">${g}<span>${g==='All patterns'?PATTERNS.length:PATTERNS.filter(p=>p.group===g).length}</span></button>`).join('');
 const list=filtered();$('#section-title').innerHTML=`${group} <span>${list.length}</span>`;
 $('#cards').innerHTML=list.map(p=>`<button class="card" data-id="${p.id}" aria-label="Study ${escapeHTML(p.title)}"><span class="card-head"><span class="tag">${p.group}</span><span class="card-number">${String(p.sources[0]).padStart(3,'0')}</span></span><h3>${escapeHTML(p.title)}</h3><p>${recall?'State the invariant before you open the proof.':escapeHTML(p.invariant)}</p><span class="card-foot"><span>${reviewed.has(p.id)?'✓ Reviewed':recall?'Test your recall':'Read the proof'}</span><span aria-hidden="true">↗</span></span></button>`).join('');
 $('#empty').hidden=list.length>0;$('#pattern-count').textContent=PATTERNS.length;$('#review-count').textContent=reviewed.size;$('#progress').max=PATTERNS.length;$('#progress').value=reviewed.size;
 $('#mode-note').textContent=recall?'Recall mode: say the rule first. Then reveal the proof.':'Select a card to see the proof, example, and Go source.';
 $('#status').textContent=`${list.length} patterns shown.`;
}
const traces={
 prefix:[
 {values:[2,4,1,3],colors:[],text:'Input [2,4,1,3]. Build the inclusive prefix from left to right.'},
 {values:[2,'·','·','·'],colors:['done'],text:'Start: prefix[0] = 2. The first prefix is correct.'},
 {values:[2,6,'·','·'],colors:['done','done'],text:'prefix[1] = 2 + 4 = 6.'},
 {values:[2,6,7,'·'],colors:['done','done','done'],text:'prefix[2] = 6 + 1 = 7.'},
 {values:[2,6,7,10],colors:['done','done','done','done'],text:'prefix[3] = 7 + 3 = 10. Query [1,3]: 10 - 2 = 8.'}],
 dnf:[
 {values:[2,0,1,2,0],colors:[],text:'low = 0, scan = 0, high = 4. All values are unknown.'},
 {values:[0,0,1,2,2],colors:['','','','','high'],text:'Move 2 to the right. high = 3. Keep scan = 0: the new value is unknown.'},
 {values:[0,0,1,2,2],colors:['done','','','','high'],text:'Process 0. low = 1, scan = 1.'},
 {values:[0,0,1,2,2],colors:['done','done','','','high'],text:'Process 0. low = 2, scan = 2.'},
 {values:[0,0,1,2,2],colors:['done','done','equal','','high'],text:'Process 1. scan = 3. The middle region contains only 1s.'},
 {values:[0,0,1,2,2],colors:['done','done','equal','high','high'],text:'Process 2. high = 2. scan > high: no unknown values remain.'}]
};
function traceMarkup(){if(!traces[current.id])return '';return '<div class="trace"><div class="trace-head"><span>STEP THROUGH THE STATE</span></div><div id="trace-body"></div><div class="trace-controls"><button id="trace-prev" class="button" aria-label="Previous step">←</button><button id="trace-next" class="button" aria-label="Next step">→</button><span id="trace-count"></span></div></div>';}
function renderTrace(){const steps=traces[current.id];if(!steps)return;const s=steps[traceIndex];$('#trace-body').innerHTML=`<div class="cells">${s.values.map((v,i)=>`<span class="cell ${s.colors[i]||''}">${v}</span>`).join('')}</div><p>${escapeHTML(s.text)}</p>`;$('#trace-count').textContent=`${traceIndex+1} / ${steps.length}`;$('#trace-prev').disabled=traceIndex===0;$('#trace-next').disabled=traceIndex===steps.length-1;}
function detail(){const p=current;
 $('#detail-group').textContent=`${p.group} / ${p.sources.map(n=>String(n).padStart(3,'0')).join(' · ')}`;
 const hidden=recall&&!revealed;
 $('#detail-content').innerHTML=`<h2 id="detail-title" class="detail-title">${escapeHTML(p.title)}</h2>`+(hidden?'<div class="recall-prompt"><p>What stays true after each step?</p><p>Why does the next move preserve it?</p><button id="reveal" class="primary">Reveal the proof</button></div>':`<div class="rule"><p class="eyebrow">${p.id==='scheduler'?'KEY BOUND':'KEY INVARIANT'}</p><p>${escapeHTML(p.invariant)}</p></div><div class="proof-details"><div><h3>01 · Keep it true</h3><p>${escapeHTML(p.step)}</p></div><div><h3>02 · Stop and conclude</h3><p>${escapeHTML(p.stop)}</p></div></div>${traceMarkup()}<div class="example"><h3>A small example</h3><p>${escapeHTML(p.example)}</p></div><div class="trap"><h3>Watch this</h3><p>${escapeHTML(p.trap)}</p></div><p class="cost">${escapeHTML(p.cost)}</p><details class="source-links"><summary>Your Go source</summary>${p.sources.map(n=>`<a href="sources/${n}.txt" target="_blank" rel="noopener">${escapeHTML(SOURCES[n])} ↗</a>`).join('')}<p>Saved snapshot. Notes describe the pattern; source comments can differ.</p></details>`);
 $('#reviewed').textContent=reviewed.has(p.id)?'✓ Reviewed · undo':'Mark as reviewed';
 $('#reviewed').disabled=hidden;
 if(!hidden)renderTrace();
}
function openPattern(id){current=PATTERNS.find(p=>p.id===id);if(!current)return;revealed=false;traceIndex=0;detail();if(!$('#detail').open){$('#detail').showModal();document.body.classList.add('modal-open');}$('#detail').scrollTop=0;history.replaceState(null,'',`#${id}`);}
$('#categories').addEventListener('click',e=>{const b=e.target.closest('[data-group]');if(b){group=b.dataset.group;render();}});
$('#cards').addEventListener('click',e=>{const b=e.target.closest('[data-id]');if(b)openPattern(b.dataset.id);});
$('#search').addEventListener('input',e=>{query=e.target.value;render();});
$('#recall').addEventListener('click',()=>{recall=!recall;$('#recall').setAttribute('aria-pressed',recall);$('#recall').innerHTML=`Recall mode <span>${recall?'●':'○'}</span>`;render();});
$('#review-filter').addEventListener('click',()=>{unreviewed=!unreviewed;$('#review-filter').setAttribute('aria-pressed',unreviewed);render();});
$('#clear').addEventListener('click',()=>{query='';group='All patterns';unreviewed=false;$('#search').value='';$('#review-filter').setAttribute('aria-pressed','false');render();});
$('#close').addEventListener('click',()=>$('#detail').close());
$('#detail').addEventListener('close',()=>{document.body.classList.remove('modal-open');history.replaceState(null,'','#library');});
$('#detail').addEventListener('click',e=>{if(e.target.id==='reveal'){revealed=true;detail();}if(e.target.id==='trace-prev'){traceIndex--;renderTrace();}if(e.target.id==='trace-next'){traceIndex++;renderTrace();}});
$('#reviewed').addEventListener('click',()=>{if(reviewed.has(current.id))reviewed.delete(current.id);else reviewed.add(current.id);try{localStorage.setItem('invariant-atlas-reviewed',JSON.stringify([...reviewed]));}catch{}render();$('#reviewed').textContent=reviewed.has(current.id)?'✓ Reviewed · undo':'Mark as reviewed';});
$('#next-pattern').addEventListener('click',()=>{const list=filtered();if(!list.length){$('#detail').close();return;}const index=list.findIndex(p=>p.id===current.id);openPattern(list[(index+1)%list.length].id);});
document.addEventListener('keydown',e=>{if(e.key==='/'&&!$('#detail').open&&!['INPUT','TEXTAREA'].includes(document.activeElement.tagName)){e.preventDefault();$('#search').focus();}});
render();const initial=location.hash.slice(1);if(PATTERNS.some(p=>p.id===initial))openPattern(initial);
