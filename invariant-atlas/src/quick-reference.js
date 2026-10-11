// Concise study cues adapted from LeetCode resource 4725; links use stable atlas IDs.
const QUICK_CHOICES = [
 ['Sorted values; find a boundary','Binary search','Keep the answer inside the remaining bounds.',['lower-bound','upper-bound']],
 ['Sorted values; find a pair','Two pointers','Move the end that cannot produce the target.',['sorted-two-sum']],
 ['Contiguous range with a valid shrink rule','Sliding window','Update counts as each item enters or leaves.',['window','minimum-window']],
 ['Many range-sum queries','Prefix sum','Subtract the total before the range.',['prefix']],
 ['Count ranges with an exact sum','Prefix frequencies','Count earlier totals equal to current total minus target.',['prefix-counts']],
 ['Need membership or counts','Hash map or set','Store the key needed by the next lookup.',['anagrams','random-set']],
 ['Need the next smaller item','Monotonic stack','Keep unresolved candidates in order.',['monotonic-stack']],
 ['Need the newest pending value','Stack','The last stored item is processed first.',['rpn']],
 ['Need repeated min/max or top k','Heap','Keep the next priority at the root.',['top-k-heap','rooms']],
 ['Need a window maximum','Monotonic deque','Remove expired and dominated candidates.',['monotonic-queue']],
 ['Need word-prefix lookup','Trie','Each edge extends the matched prefix.',['trie']],
 ['Need the middle or a list cycle','Fast and slow pointers','For cycles, compare node addresses after moving.',['fast-slow','cycle-detection']],
 ['Need to reverse list links','Previous and next pointers','Save the next node before changing its link.',['reverse-list']],
 ['Need levels or fewest unweighted edges','BFS','Mark on discovery; process the queue in order.',['levels','graph-bfs']],
 ['Need subtree results or reachability','DFS','Track unfinished work; mark graph nodes as visited.',['tree-recursive','graph-dfs-stack']],
 ['Feasibility changes once across answers','Binary search on the answer','First prove that the check is monotonic.',['search-minimum','search-maximum']]
];
function renderQuickChoices(){
 const query=document.querySelector('#quick-search').value.toLowerCase().trim();
 const rows=QUICK_CHOICES.filter(row=>query.split(/ +/).every(term=>row.flat().join(' ').toLowerCase().includes(term)));
 document.querySelector('#quick-results').textContent=rows.length+' matching cues';
 document.querySelector('#quick-choices').innerHTML=rows.map(([cue,name,rule,ids])=>`<article class="quick-choice"><p class="eyebrow">${escapeHTML(cue)}</p><h3>${escapeHTML(name)}</h3><p>${escapeHTML(rule)}</p><div>${ids.map(id=>{const p=PATTERNS.find(p=>p.id===id);return `<button class="text-button" data-id="${id}">${escapeHTML(p.title)} →</button>`;}).join('')}</div></article>`).join('');
}
document.querySelector('#quick-search').addEventListener('input',renderQuickChoices);
document.querySelector('#quick-choices').addEventListener('click',event=>{const button=event.target.closest('[data-id]');if(button)openPattern(button.dataset.id);});
renderQuickChoices();
document.querySelector('.quick-nav').addEventListener('click',event=>{
 const link=event.target.closest('a');if(!link)return;
 const target=document.querySelector(link.getAttribute('href'));
 if(target?.tagName==='DETAILS')target.open=true;
});
