'use strict';
const growthSeries=[
 ['O(1)','#7ee787',n=>1],['O(log n)','#56d4dd',n=>Math.log2(n)],
 ['O(n)','#d2e660',n=>n],['O(n log n)','#e3b341',n=>n*Math.log2(n)],
 ['O(n²)','#ffa657',n=>n*n],['O(2ⁿ)','#ff7b72',n=>2**n],
 ['O(n!)','#d2a8ff',n=>{let a=1;for(let i=2;i<=n;i++)a*=i;return a;}]
];
const svg=document.querySelector('#growth-chart');const x=n=>75+(n-1)*850/49;const y=v=>365-v*315/200;
let markup='<title id="growth-title">Growth of seven common complexity classes</title><desc id="growth-desc">Linear axes. Fast-growing curves are clipped above 200 units of work.</desc><defs><clipPath id="plot-clip"><rect x="75" y="50" width="850" height="315"/></clipPath></defs>';
for(const v of [0,50,100,150,200])markup+=`<path d="M75 ${y(v)}H925" stroke="#30363d"/><text x="62" y="${y(v)+5}" text-anchor="end">${v}</text>`;
for(const n of [1,10,20,30,40,50])markup+=`<text x="${x(n)}" y="390" text-anchor="middle">${n}</text>`;
markup+='<text x="500" y="422" text-anchor="middle">Input size (n)</text><text x="20" y="205" transform="rotate(-90 20 205)" text-anchor="middle">Relative work</text><g clip-path="url(#plot-clip)">';
for(const [label,color,fn] of growthSeries){let points=[];for(let n=1;n<=50;n++){points.push(x(n)+','+y(Math.min(fn(n),250)));if(fn(n)>250)break;}markup+=`<polyline points="${points.join(' ')}" fill="none" stroke="${color}" stroke-width="2.5"><title>${label}</title></polyline>`;}
markup+='</g><path d="M75 50V365H925" fill="none" stroke="#9da7b3"/>';svg.innerHTML=markup;
document.querySelector('#growth-legend').innerHTML=growthSeries.map(([label,color])=>`<span><i style="background:${color}"></i>${label}</span>`).join('');
function showGrowth(){const n=Number(document.querySelector('#growth-n').value);document.querySelector('#growth-value').textContent=n;
 document.querySelector('#growth-values').innerHTML=growthSeries.map(([label,color,fn])=>{const value=fn(n);const formatted=value>=1e7?value.toExponential(2):value.toLocaleString('en-US',{maximumFractionDigits:1});return `<div><span style="color:${color}">${label}</span><strong>${formatted}</strong></div>`;}).join('');}
document.querySelector('#growth-n').addEventListener('input',showGrowth);showGrowth();
