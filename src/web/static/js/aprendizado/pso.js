import { createPlayer, scaleCanvas } from "./aprendizado-base.js";

const W = 700, W2 = 700;
const H_LAND = 440, H_CHROM = 220;
const N_PART = 12;
const N_ITER = 14;
const W_INERTIA = 0.6, C1 = 2.0, C2 = 2.0;
const DIM = 6;

function fitness(x, y) {
  const xc = x * 3, yc = y * 3;
  const p1 = 1.2 * Math.exp(-((xc-1.0)*(xc-1.0) + (yc-0.8)*(yc-0.8)) / 0.25);
  const p2 = 1.8 * Math.exp(-((xc-2.1)*(xc-2.1) + (yc-1.8)*(yc-1.8)) / 0.40);
  const p3 = 2.8 * Math.exp(-((xc-0.7)*(xc-0.7) + (yc-2.5)*(yc-2.5)) / 0.18);
  return p1 + p2 + p3;
}

function mx(x) { return 55 + x * (W - 110); }
function my(y) { return H_LAND - 40 - y * (H_LAND - 80); }

function decodeRK(vec) {
  const indexed = vec.map((v, i) => ({ v, i }));
  indexed.sort((a, b) => a.v - b.v);
  return indexed.map(x => x.i + 1);
}

function genParticles() {
  const parts = [];
  for (let i = 0; i < N_PART; i++) {
    const pos = [];
    for (let d = 0; d < DIM; d++) pos.push(Math.random());
    const vel = new Array(DIM).fill(0);
    const fitVal = fitness(pos[0], pos[1]);
    parts.push({
      pos: [...pos],
      vel: [...vel],
      pbest: [...pos],
      pbestFit: fitVal,
      fit: fitVal,
    });
  }
  return parts;
}

function genFrames() {
  const parts = genParticles();
  let gbest = [...parts[0].pos], gbestFit = parts[0].fit;
  for (const p of parts) { if (p.fit > gbestFit) { gbest = [...p.pos]; gbestFit = p.fit; } }

  const frames = [makeFrame(0, parts, gbest, gbestFit)];

  for (let iter = 1; iter <= N_ITER; iter++) {
    for (const p of parts) {
      for (let d = 0; d < DIM; d++) {
        const r1 = Math.random(), r2 = Math.random();
        p.vel[d] = W_INERTIA * p.vel[d] + C1 * r1 * (p.pbest[d] - p.pos[d]) + C2 * r2 * (gbest[d] - p.pos[d]);
        p.vel[d] = Math.max(-0.06, Math.min(0.06, p.vel[d]));
        p.pos[d] = Math.max(0, Math.min(1, p.pos[d] + p.vel[d]));
      }
      p.fit = fitness(p.pos[0], p.pos[1]);
      if (p.fit > p.pbestFit) { p.pbest = [...p.pos]; p.pbestFit = p.fit; }
      if (p.fit > gbestFit) { gbest = [...p.pos]; gbestFit = p.fit; }
    }

    let vizP = parts[0];
    for (const p of parts) if (p.fit > vizP.fit) vizP = p;

    frames.push(makeFrame(iter, parts, gbest, gbestFit, vizP));
  }

  return frames;
}

function makeFrame(iter, parts, gbest, gbestFit, vizPart) {
  return {
    iter,
    particles: parts.map(p => ({
      pos: [...p.pos], vel: [...p.vel],
      pbest: [...p.pbest], pbestFit: p.pbestFit, fit: p.fit,
    })),
    gbest: [...gbest], gbestFit,
    viz: vizPart ? {
      pos: [...vizPart.pos], vel: [...vizPart.vel],
      pbest: [...vizPart.pbest], gbest: [...gbest],
      route: decodeRK(vizPart.pos),
      pbestRoute: decodeRK(vizPart.pbest),
      gbestRoute: decodeRK(gbest),
    } : null,
  };
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
  document.getElementById("step-info").textContent = `Iteração ${f.iter} / ${N_ITER} (passo ${index+1}/${total})`;
  document.getElementById("step-description").innerHTML =
    `<strong>PSO com Random Keys:</strong> v ← ω·v + c₁·r₁·(pbest − x) + c₂·r₂·(gbest − x). ` +
    `Abaixo: vetor posição da melhor partícula, decodificado para permutação. ` +
    `Fitness gbest: <span style="color:#f39c12;font-weight:700">${f.gbestFit.toFixed(3)}</span>`;
}

function drawLandscape(f) {
  ctx.clearRect(0,0,W,H_LAND);
  ctx.fillStyle="#f8f9fb"; ctx.fillRect(0,0,W,H_LAND);
  drawHeatmap();
  drawPbestMarkers(f);
  drawVelocities(f);
  drawParticles(f);
  drawGbest(f);
  ctx.fillStyle="#555"; ctx.font="11px sans-serif"; ctx.textAlign="right";
  ctx.fillText(`Iteração ${f.iter}/${N_ITER}  |  ★ gbest  |  + pbest  |  → velocidade`, W-12, H_LAND-8);
}

function drawHeatmap() {
  const res=50, cw=W-110, ch=H_LAND-80;
  for (let i=0;i<res;i++) for(let j=0;j<res;j++) {
    const x=i/(res-1),y=j/(res-1),v=fitness(x,y)/3.0;
    ctx.fillStyle=`rgb(${Math.floor(20+v*215)},${Math.floor(80+v*130)},${Math.floor(180-v*100)})`;
    ctx.fillRect(55+i*(cw/res),H_LAND-40-(j+1)*(ch/res),(cw/res)+1,(ch/res)+1);
  }
}

function drawPbestMarkers(f) {
  for(const p of f.particles) {
    const px=mx(p.pbest[0]),py=my(p.pbest[1]);
    ctx.strokeStyle="rgba(142,68,173,0.5)"; ctx.lineWidth=1;
    ctx.beginPath(); ctx.moveTo(px-4,py); ctx.lineTo(px+4,py); ctx.moveTo(px,py-4); ctx.lineTo(px,py+4); ctx.stroke();
  }
}

function drawVelocities(f) {
  for(const p of f.particles) {
    const px=mx(p.pos[0]),py=my(p.pos[1]),vx=p.vel[0]*400,vy=p.vel[1]*400;
    if(Math.hypot(vx,vy)<0.5) continue;
    ctx.strokeStyle="rgba(80,80,80,0.45)"; ctx.lineWidth=1.2;
    ctx.beginPath(); ctx.moveTo(px,py); ctx.lineTo(px+vx,py+vy); ctx.stroke();
    const ang=Math.atan2(vy,vx),al=5;
    ctx.beginPath(); ctx.moveTo(px+vx,py+vy);
    ctx.lineTo(px+vx-al*Math.cos(ang-0.5),py+vy-al*Math.sin(ang-0.5));
    ctx.lineTo(px+vx-al*Math.cos(ang+0.5),py+vy-al*Math.sin(ang+0.5));
    ctx.closePath(); ctx.fillStyle="rgba(80,80,80,0.45)"; ctx.fill();
  }
}

function drawParticles(f) {
  for(const p of f.particles) {
    ctx.beginPath(); ctx.arc(mx(p.pos[0]),my(p.pos[1]),3.5+p.fit*2,0,Math.PI*2);
    ctx.fillStyle="#2c3e50"; ctx.fill(); ctx.strokeStyle="#fff"; ctx.lineWidth=1.5; ctx.stroke();
  }
}

function drawGbest(f) {
  const gx=mx(f.gbest[0]),gy=my(f.gbest[1]);
  ctx.beginPath(); ctx.arc(gx,gy,12,0,Math.PI*2); ctx.fillStyle="rgba(243,156,18,0.2)"; ctx.fill();
  ctx.fillStyle="#f39c12"; ctx.font="bold 20px sans-serif"; ctx.textAlign="center"; ctx.textBaseline="middle";
  ctx.fillText("★",gx,gy-16);
}

function drawChromosome(f) {
  const ctx2 = chromCtx, w = W2, h = H_CHROM;
  ctx2.clearRect(0,0,w,h);
  ctx2.fillStyle="#fafbfc"; ctx2.fillRect(0,0,w,h);
  ctx2.strokeStyle="#ddd"; ctx2.lineWidth=1;
  ctx2.beginPath(); ctx2.moveTo(0,0); ctx2.lineTo(w,0); ctx2.stroke();

  if (!f.viz) {
    ctx2.fillStyle="#888"; ctx2.font="13px sans-serif"; ctx2.textAlign="center";
    ctx2.fillText("Inicializando vetores random-key...", w/2, h/2);
    return;
  }

  const v = f.viz;
  const cellW = 70, startX = 30, y0 = 35;

  ctx2.fillStyle="#333"; ctx2.font="bold 12px sans-serif"; ctx2.textAlign="left";
  ctx2.fillText(`Random Keys — Vetor de Posição (melhor partícula)`, 12, 22);

  ctx2.fillStyle="#2c3e50"; ctx2.font="bold 11px monospace"; ctx2.textAlign="left";
  ctx2.fillText("x (posição):", 12, y0+4);
  for (let d = 0; d < DIM; d++) {
    const cx = startX + d * cellW;
    ctx2.fillStyle="#fff"; ctx2.fillRect(cx, y0-10, cellW-2, 22);
    ctx2.strokeStyle="#2c3e50"; ctx2.lineWidth=1.5; ctx2.strokeRect(cx, y0-10, cellW-2, 22);
    ctx2.fillStyle="#2c3e50"; ctx2.font="bold 12px monospace"; ctx2.textAlign="center"; ctx2.textBaseline="middle";
    ctx2.fillText(v.pos[d].toFixed(3), cx+(cellW-2)/2, y0+1);
  }

  const y1 = y0 + 38;
  ctx2.fillStyle="#27ae60"; ctx2.font="bold 11px monospace"; ctx2.textAlign="left";
  ctx2.fillText("π = ordenar(x):", 12, y1+4);
  for (let d = 0; d < DIM; d++) {
    const cx = startX + d * cellW;
    ctx2.fillStyle="#e8f8f5"; ctx2.fillRect(cx, y1-10, cellW-2, 22);
    ctx2.strokeStyle="#27ae60"; ctx2.lineWidth=1.5; ctx2.strokeRect(cx, y1-10, cellW-2, 22);
    ctx2.fillStyle="#27ae60"; ctx2.font="bold 14px monospace"; ctx2.textAlign="center"; ctx2.textBaseline="middle";
    ctx2.fillText(v.route[d], cx+(cellW-2)/2, y1+1);
  }

  const y2 = y1 + 38;
  ctx2.fillStyle="#8e44ad"; ctx2.font="bold 11px monospace"; ctx2.textAlign="left";
  ctx2.fillText("v (velocidade):", 12, y2+4);
  for (let d = 0; d < DIM; d++) {
    const cx = startX + d * cellW;
    ctx2.fillStyle="#f5eef8"; ctx2.fillRect(cx, y2-10, cellW-2, 22);
    ctx2.strokeStyle="#8e44ad"; ctx2.lineWidth=1.5; ctx2.strokeRect(cx, y2-10, cellW-2, 22);
    ctx2.fillStyle="#8e44ad"; ctx2.font="bold 12px monospace"; ctx2.textAlign="center"; ctx2.textBaseline="middle";
    ctx2.fillText(v.vel[d].toFixed(3), cx+(cellW-2)/2, y2+1);
  }

  const y3 = y2 + 38;
  ctx2.fillStyle="#7f8c8d"; ctx2.font="bold 11px monospace"; ctx2.textAlign="left";
  ctx2.fillText("pbest:", 12, y3+4);
  for (let d = 0; d < DIM; d++) {
    const cx = startX + d * cellW;
    ctx2.fillStyle="#f0f0f0"; ctx2.fillRect(cx, y3-10, cellW-2, 22);
    ctx2.strokeStyle="#bdc3c7"; ctx2.lineWidth=1; ctx2.strokeRect(cx, y3-10, cellW-2, 22);
    ctx2.fillStyle="#555"; ctx2.font="11px monospace"; ctx2.textAlign="center"; ctx2.textBaseline="middle";
    ctx2.fillText(v.pbest[d].toFixed(3), cx+(cellW-2)/2, y3+1);
  }

  const y4 = y3 + 38;
  ctx2.fillStyle="#f39c12"; ctx2.font="bold 11px monospace"; ctx2.textAlign="left";
  ctx2.fillText("gbest:", 12, y4+4);
  for (let d = 0; d < DIM; d++) {
    const cx = startX + d * cellW;
    ctx2.fillStyle="#fff8e1"; ctx2.fillRect(cx, y4-10, cellW-2, 22);
    ctx2.strokeStyle="#f39c12"; ctx2.lineWidth=1.5; ctx2.strokeRect(cx, y4-10, cellW-2, 22);
    ctx2.fillStyle="#e67e22"; ctx2.font="bold 12px monospace"; ctx2.textAlign="center"; ctx2.textBaseline="middle";
    ctx2.fillText(v.gbest[d].toFixed(3), cx+(cellW-2)/2, y4+1);
  }

  const rx = startX + DIM * cellW + 30;
  ctx2.fillStyle="#555"; ctx2.font="10px sans-serif"; ctx2.textAlign="left";
  ctx2.fillText("Atualização:", rx, y0+4);
  ctx2.fillText("vₙ = ω·v + c₁r₁(p−x) + c₂r₂(g−x)", rx, y0+20);
  ctx2.fillText("xₙ = x + vₙ", rx, y0+34);
  ctx2.fillText("π = ordenar índices por x", rx, y0+58);
  ctx2.fillText("(menor valor → 1ª posição)", rx, y0+72);
}

const init = () => {
  setupCanvases();
  const frames = genFrames();
  createPlayer({
    genFrames: () => frames,
    drawAll: (f) => drawAll(f),
    updateInfo: (f) => updateInfo(f),
    speed: 350,
  }).init();
};

export { init };
