import { createPlayer, scaleCanvas } from "./aprendizado-base.js";

const W = 700, W2 = 700;
const H_LAND = 440, H_CHROM = 210;

function fitness(x, y) {
  const xc = x * 3, yc = y * 3;
  const p1 = 1.2 * Math.exp(-((xc-1.0)*(xc-1.0) + (yc-0.8)*(yc-0.8)) / 0.25);
  const p2 = 1.8 * Math.exp(-((xc-2.1)*(xc-2.1) + (yc-1.8)*(yc-1.8)) / 0.40);
  const p3 = 2.8 * Math.exp(-((xc-0.7)*(xc-0.7) + (yc-2.5)*(yc-2.5)) / 0.18);
  return p1 + p2 + p3;
}

function mx(x) { return 55 + x * (W - 110); }
function my(y) { return H_LAND - 40 - y * (H_LAND - 80); }

const N_CITIES = 6;
const POP = 12;
const GENS = 6;

function randomPerm() {
  const arr = [1,2,3,4,5];
  for (let i = arr.length-1; i > 0; i--) {
    const j = Math.floor(Math.random() * (i+1));
    [arr[i], arr[j]] = [arr[j], arr[i]];
  }
  return [0, ...arr];
}

function randInd() {
  return { x: Math.random(), y: Math.random(), fit: 0, chrom: randomPerm() };
}

function evalPop(pop) {
  for (const p of pop) p.fit = fitness(p.x, p.y);
  return [...pop].sort((a,b) => b.fit - a.fit);
}

function tournamentSelect(sorted, k) {
  const cand = [];
  for (let i = 0; i < k; i++) cand.push(sorted[Math.floor(Math.random()*sorted.length)]);
  cand.sort((a,b) => b.fit - a.fit);
  return cand[0];
}

function oxCrossover(p1, p2) {
  const n = p1.length;
  const a = 1 + Math.floor(Math.random()*(n-2));
  const b = 1 + Math.floor(Math.random()*(n-2));
  const lo = Math.min(a,b), hi = Math.max(a,b);
  const seg1 = p1.slice(lo, hi+1);
  const seg2 = p2.slice(lo, hi+1);

  const c1 = new Array(n).fill(-1);
  const c2 = new Array(n).fill(-1);
  c1[0] = 0; c2[0] = 0;
  for (let i = lo; i <= hi; i++) { c1[i] = p1[i]; c2[i] = p2[i]; }

  const fill = (child, parent) => {
    let idx = (hi+1) % n;
    if (idx === 0) idx = 1;
    for (let i = 0; i < n; i++) {
      const gene = parent[(hi+1+i) % n];
      if (gene === 0) continue;
      if (!child.includes(gene)) {
        child[idx] = gene;
        idx = (idx+1) % n;
        if (idx === 0) idx = 1;
      }
    }
  };
  fill(c1, p2); fill(c2, p1);

  const missing = (c, parent) => {
    for (let i = 1; i < n; i++) {
      if (c[i] === -1) {
        for (const g of parent) {
          if (g !== 0 && !c.includes(g)) { c[i] = g; break; }
        }
      }
    }
  };
  missing(c1, p1); missing(c2, p2);

  return { c1, c2, lo, hi, p1: [...p1], p2: [...p2] };
}

function swapMutate(chrom) {
  const c = [...chrom];
  const i = 1 + Math.floor(Math.random()*(c.length-1));
  let j = 1 + Math.floor(Math.random()*(c.length-1));
  if (i === j) j = (j % (c.length-1)) + 1;
  [c[i], c[j]] = [c[j], c[i]];
  return { before: [...chrom], after: c, i, j };
}

function genEvolution() {
  const frames = [];
  let pop = Array.from({length: POP}, randInd);
  pop = evalPop(pop);

  for (let g = 0; g < GENS; g++) {
    const sorted = [...pop].sort((a,b) => b.fit - a.fit);
    const best = sorted[0];

    frames.push({ type: "population", gen: g,
      pop: sorted.map(p=>({x:p.x,y:p.y,fit:p.fit,chrom:[...p.chrom]})),
      best: {x:best.x,y:best.y,fit:best.fit,chrom:[...best.chrom]} });

    const elite = sorted.slice(0,2);
    const selected = [...elite];
    for (let i=2; i<POP; i++) selected.push(tournamentSelect(sorted,3));

    frames.push({ type: "selection", gen: g,
      pop: sorted.map(p=>({x:p.x,y:p.y,fit:p.fit,chrom:[...p.chrom]})),
      selected: selected.map(p=>({x:p.x,y:p.y,fit:p.fit,chrom:[...p.chrom]})),
      best: {x:best.x,y:best.y,fit:best.fit,chrom:[...best.chrom]} });

    const children = [elite[0], elite[1]];
    let vizXover = null;
    for (let i=2; i<POP; i+=2) {
      const p1 = selected[Math.floor(Math.random()*POP)];
      const p2 = selected[Math.floor(Math.random()*POP)];
      const xo = oxCrossover(p1.chrom, p2.chrom);
      children.push({x:0,y:0,fit:0,chrom:xo.c1});
      if (children.length < POP) children.push({x:0,y:0,fit:0,chrom:xo.c2});
      if (!vizXover) vizXover = xo;
    }
    while (children.length > POP) children.pop();

    frames.push({ type: "crossover", gen: g,
      pop: sorted.map(p=>({x:p.x,y:p.y,fit:p.fit,chrom:[...p.chrom]})),
      children: children.map(p=>({x:p.x,y:p.y,fit:p.fit,chrom:[...p.chrom]})),
      best: {x:best.x,y:best.y,fit:best.fit,chrom:[...best.chrom]},
      vizXover });

    const mutated = [];
    let vizMut = null;
    for (const ch of children) {
      const m = swapMutate(ch.chrom);
      mutated.push({x:0,y:0,fit:0,chrom:m.after});
      if (!vizMut) vizMut = m;
    }
    pop = evalPop(mutated.map(m => ({x:Math.random(),y:Math.random(),fit:0,chrom:m.chrom})));

    frames.push({ type: "mutation", gen: g,
      pop: sorted.map(p=>({x:p.x,y:p.y,fit:p.fit,chrom:[...p.chrom]})),
      mutated: pop.map(p=>({x:p.x,y:p.y,fit:p.fit,chrom:[...p.chrom]})),
      best: {x:best.x,y:best.y,fit:best.fit,chrom:[...best.chrom]},
      vizMut });
  }

  return frames;
}

let ctx, chromCtx;

function setupCanvases() {
  const c = document.getElementById("viz-canvas");
  ctx = scaleCanvas(c, W, H_LAND);

  document.getElementById("chrom-wrapper").style.display = "block";
  const cc = document.getElementById("chrom-canvas");
  chromCtx = scaleCanvas(cc, W2, H_CHROM);
}

function drawAll(f) {
  drawLandscape(f);
  drawChromosome(f);
}

function updateInfo(f, index, total) {
  const labels = {
    population: "População — cada círculo é um indivíduo. Abaixo: 4 melhores cromossomos (permutações).",
    selection: "Seleção por torneio — pais destacados em dourado. Abaixo: torneio entre 3, vence o de maior fitness.",
    crossover: "Crossover OX: segmento laranja do Pai 1 é copiado para o Filho 1; o resto é preenchido na ordem do Pai 2.",
    mutation: "Mutação Swap: duas posições aleatórias do cromossomo são trocadas para explorar novas soluções.",
  };
  document.getElementById("step-info").textContent = `Geração ${f.gen} — ${f.type} (passo ${index+1}/${total})`;
  document.getElementById("step-description").innerHTML =
    `<strong>${labels[f.type]}</strong> | Melhor fitness: <span style="color:#0a66c2;font-weight:700">${f.best.fit.toFixed(3)}</span>`;
}

function drawLandscape(f) {
  ctx.clearRect(0, 0, W, H_LAND);
  ctx.fillStyle = "#f8f9fb"; ctx.fillRect(0, 0, W, H_LAND);
  drawHeatmap();
  if (f.type==="population") drawInds(f.pop,f.best,false,false);
  else if (f.type==="selection") { drawInds(f.pop,f.best,false,true); drawSelHighlight(f.selected); }
  else if (f.type==="crossover") { drawInds(f.pop,f.best,true,true); drawInds(f.children,f.best,false,false); }
  else if (f.type==="mutation") { drawInds(f.pop,f.best,true,true); drawInds(f.mutated,f.best,false,false); }
  drawLandTitle(f);
}

function drawHeatmap() {
  const res=50, cw=W-110, ch=H_LAND-80;
  for (let i=0;i<res;i++) for (let j=0;j<res;j++) {
    const x=i/(res-1),y=j/(res-1),v=fitness(x,y)/3.0;
    ctx.fillStyle = `rgb(${Math.floor(20+v*215)},${Math.floor(80+v*130)},${Math.floor(180-v*100)})`;
    ctx.fillRect(55+i*(cw/res),H_LAND-40-(j+1)*(ch/res),(cw/res)+1,(ch/res)+1);
  }
}

function drawInds(inds,best,faded,showBest) {
  for (const p of inds) {
    const px=mx(p.x),py=my(p.y),isBest=showBest&&Math.abs(p.x-best.x)<0.001&&Math.abs(p.y-best.y)<0.001;
    ctx.beginPath(); ctx.arc(px,py,4+(p.fit||0)*3,0,Math.PI*2);
    ctx.fillStyle=faded?"rgba(180,180,180,0.5)":(isBest?"#c0392b":"#2c3e50"); ctx.fill();
    ctx.strokeStyle=isBest?"#f1c40f":"#fff"; ctx.lineWidth=isBest?2.5:1.5; ctx.stroke();
  }
}

function drawSelHighlight(sel) {
  for (const p of sel) {
    ctx.beginPath(); ctx.arc(mx(p.x),my(p.y),7,0,Math.PI*2);
    ctx.strokeStyle="#f39c12"; ctx.lineWidth=3; ctx.setLineDash([4,2]); ctx.stroke(); ctx.setLineDash([]);
  }
}

function drawLandTitle(f) {
  ctx.fillStyle="#555"; ctx.font="12px sans-serif"; ctx.textAlign="left";
  const t={population:`Geração ${f.gen} — População`,selection:`Geração ${f.gen} — Seleção`,crossover:`Geração ${f.gen} — Crossover`,mutation:`Geração ${f.gen} — Mutação`};
  ctx.fillText(t[f.type],12,H_LAND-8);
}

function drawChromosome(f) {
  const ctx2 = chromCtx, w = W2, h = H_CHROM;
  ctx2.clearRect(0,0,w,h);
  ctx2.fillStyle="#fafbfc"; ctx2.fillRect(0,0,w,h);

  const colors = ["#c0392b","#2980b9","#27ae60","#8e44ad","#e67e22","#16a085"];

  if (f.type === "population") {
    drawChromTitle(ctx2, w, `População — melhores cromossomos (permutações)`);
    const top = sortedByFit(f.pop).slice(0, 6);
    for (let r=0; r<top.length; r++) {
      drawChromRow(ctx2, w, 40 + r*26, top[r].chrom, colors, top[r].fit, r===0);
    }
  } else if (f.type === "selection") {
    drawChromTitle(ctx2, w, `Seleção por Torneio (k=3)`);
    const cand = f.pop.slice(0,6).sort(()=>Math.random()-0.5).slice(0,3);
    cand.sort((a,b)=>b.fit-a.fit);
    const winner = cand[0];
    for (let r=0; r<cand.length; r++) {
      const isWin = cand[r] === winner;
      const y = 50 + r*32;
      ctx2.fillStyle=isWin?"#fff3cd":"#fff"; ctx2.fillRect(30,y-14,w-60,28);
      ctx2.strokeStyle=isWin?"#f39c12":"#ddd"; ctx2.lineWidth=isWin?2.5:1;
      ctx2.strokeRect(30,y-14,w-60,28);
      drawChromRowAt(ctx2, 50, y, cand[r].chrom, colors, cand[r].fit, isWin);
      if (isWin) { ctx2.fillStyle="#f39c12"; ctx2.font="bold 11px sans-serif"; ctx2.textAlign="right"; ctx2.fillText("← vencedor", w-20, y+4); }
    }
  } else if (f.type === "crossover" && f.vizXover) {
    const xo = f.vizXover;
    drawChromTitle(ctx2, w, `Crossover OX (Order Crossover)`);
    drawChromRowAt(ctx2, 40, 50, xo.p1, colors, null, false);
    ctx2.fillStyle="#555"; ctx2.font="10px sans-serif"; ctx2.textAlign="left"; ctx2.fillText("Pai 1", 12, 54);
    drawChromRowAt(ctx2, 40, 82, xo.p2, colors, null, false);
    ctx2.fillText("Pai 2", 12, 86);

    const segX1 = 40 + xo.lo * 32, segW = (xo.hi - xo.lo + 1) * 32;
    ctx2.strokeStyle="#f39c12"; ctx2.lineWidth=2.5;
    ctx2.strokeRect(segX1, 50-10, segW, 20);
    ctx2.strokeRect(segX1, 82-10, segW, 20);

    drawChromRowAt(ctx2, 40, 130, xo.c1, colors, null, false);
    ctx2.fillText("Filho 1", 12, 134);
    drawChromRowAt(ctx2, 40, 162, xo.c2, colors, null, false);
    ctx2.fillText("Filho 2", 12, 166);

    ctx2.strokeStyle="#f39c12"; ctx2.lineWidth=1.5;
    drawArrow(ctx2, segX1+segW/2, 70, segX1+segW/2, 118);
    ctx2.fillStyle="#f39c12"; ctx2.font="9px sans-serif"; ctx2.textAlign="center";
    ctx2.fillText("segmento", segX1+segW/2, 100);
    ctx2.fillText("copiado", segX1+segW/2, 112);

    ctx2.fillStyle="#555"; ctx2.font="10px sans-serif"; ctx2.textAlign="left";
    ctx2.fillText("restante preenchido na ordem do outro pai", 300, 148);
    ctx2.fillText("(mantendo a ordem relativa dos genes)", 300, 162);

  } else if (f.type === "mutation" && f.vizMut) {
    const m = f.vizMut;
    drawChromTitle(ctx2, w, `Mutação Swap`);
    drawChromRowAt(ctx2, 40, 55, m.before, colors, null, false);
    ctx2.fillStyle="#555"; ctx2.font="10px sans-serif"; ctx2.textAlign="left"; ctx2.fillText("Antes", 12, 59);

    drawChromRowAt(ctx2, 40, 100, m.after, colors, null, false);
    ctx2.fillText("Depois", 12, 104);

    const x1 = 40 + m.i * 32, x2 = 40 + m.j * 32;
    ctx2.strokeStyle="#e74c3c"; ctx2.lineWidth=2;
    ctx2.setLineDash([3,2]);
    ctx2.strokeRect(x1, 45, 32, 20);
    ctx2.strokeRect(x2, 45, 32, 20);
    ctx2.strokeRect(x1, 90, 32, 20);
    ctx2.strokeRect(x2, 90, 32, 20);
    ctx2.setLineDash([]);

    ctx2.fillStyle="#e74c3c"; ctx2.font="10px sans-serif"; ctx2.textAlign="center";
    ctx2.fillText("pos "+(m.i), x1+16, 78);
    ctx2.fillText("↕ swap", (x1+x2)/2+16, 82);
    ctx2.fillText("pos "+(m.j), x2+16, 78);

    ctx2.fillStyle="#555"; ctx2.font="10px sans-serif"; ctx2.textAlign="left";
    ctx2.fillText("Duas posições são trocadas aleatoriamente — promove diversidade.", 300, 70);
  }

  chromCtx.strokeStyle = "#ddd"; chromCtx.lineWidth = 1;
  chromCtx.beginPath(); chromCtx.moveTo(0,0); chromCtx.lineTo(w,0); chromCtx.stroke();
}

function sortedByFit(pop) { return [...pop].sort((a,b)=>b.fit-a.fit); }

function drawChromRow(ctx2, w, y, chrom, colors, fit, highlight) {
  drawChromRowAt(ctx2, 40, y, chrom, colors, fit, highlight);
}

function drawChromRowAt(ctx2, x, y, chrom, colors, fit, highlight) {
  const cellW = 30;
  for (let i=0; i<chrom.length; i++) {
    const gene = chrom[i];
    const cx = x + i*cellW;
    ctx2.fillStyle = colors[gene % colors.length];
    ctx2.fillRect(cx, y-10, cellW-2, 20);
    ctx2.fillStyle = "#fff";
    ctx2.font = "bold 10px monospace";
    ctx2.textAlign = "center";
    ctx2.textBaseline = "middle";
    ctx2.fillText(gene, cx+(cellW-2)/2, y);
  }
  if (fit !== null && fit !== undefined) {
    ctx2.fillStyle = "#333";
    ctx2.font = "10px monospace";
    ctx2.textAlign = "left";
    ctx2.fillText(`f=${fit.toFixed(3)}`, x + chrom.length*cellW + 10, y+4);
  }
  if (highlight) {
    ctx2.strokeStyle = "#f1c40f";
    ctx2.lineWidth = 2;
    ctx2.strokeRect(x-3, y-13, chrom.length*cellW+4, 26);
  }
}

function drawChromTitle(ctx2, w, text) {
  ctx2.fillStyle = "#333";
  ctx2.font = "bold 12px sans-serif";
  ctx2.textAlign = "left";
  ctx2.fillText(text, 12, 24);
}

function drawArrow(ctx2, x1, y1, x2, y2) {
  ctx2.beginPath(); ctx2.moveTo(x1,y1); ctx2.lineTo(x2,y2); ctx2.stroke();
  const ang = Math.atan2(y2-y1, x2-x1);
  const al = 6;
  ctx2.beginPath();
  ctx2.moveTo(x2,y2);
  ctx2.lineTo(x2-al*Math.cos(ang-0.5), y2-al*Math.sin(ang-0.5));
  ctx2.lineTo(x2-al*Math.cos(ang+0.5), y2-al*Math.sin(ang+0.5));
  ctx2.closePath();
  ctx2.fillStyle = "#f39c12";
  ctx2.fill();
}

const init = () => {
  setupCanvases();
  const frames = genEvolution();
  createPlayer({
    genFrames: () => frames,
    drawAll: (f) => drawAll(f),
    updateInfo: (f) => updateInfo(f),
    speed: 800,
  }).init();
};

export { init };
