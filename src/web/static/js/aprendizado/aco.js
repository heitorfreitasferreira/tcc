import { createPlayer, scaleCanvas } from "./aprendizado-base.js";

const W = 700, W2 = 700;
const H_LAND = 460, H_CHROM = 250;
const N_CITIES = 15;
const N_ANTS = 3;
const N_ITER = 2;
const EVAP = 0.3, ALPHA = 1.0, BETA = 2.0, Q = 20;

const ANT_COLORS = ["#e74c3c", "#2980b9", "#2ecc71"];

const CITIES = (() => {
  const arr = [];
  const cx = 350, cy = 235, radius = 185;
  for (let i = 0; i < N_CITIES; i++) {
    const angle = (i / N_CITIES) * 2 * Math.PI - Math.PI / 2;
    arr.push({
      x: cx + radius * Math.cos(angle) + (i === 0 ? 0 : (Math.random() - 0.5) * 26),
      y: cy + radius * Math.sin(angle) + (i === 0 ? -5 : (Math.random() - 0.5) * 26),
    });
  }
  return arr;
})();

function dist(i, j) {
  return Math.hypot(CITIES[i].x - CITIES[j].x, CITIES[i].y - CITIES[j].y);
}

function routeCost(tour) {
  let cost = 0;
  for (let k = 0; k < tour.length - 1; k++) cost += dist(tour[k], tour[k + 1]);
  if (tour.length > 1) cost += dist(tour[tour.length - 1], tour[0]);
  return cost;
}

function clonePher(p) {
  return p.map(row => [...row]);
}

function genFrames() {
  const frames = [];
  const N = N_CITIES;

  const pher = Array.from({ length: N }, () => Array(N).fill(0));
  const dists = Array.from({ length: N }, () => Array(N).fill(0));
  for (let i = 0; i < N; i++)
    for (let j = 0; j < N; j++)
      dists[i][j] = dist(i, j);

  let globalBest = null;
  let globalBestCost = Infinity;

  frames.push({
    type: "init",
    iteration: 0, round: -1, ants: null,
    pheromone: clonePher(pher),
    bestTour: null, bestCost: Infinity,
    message: "Nenhuma aresta ainda — formigas constroem rotas do zero em paralelo. Pressione ▶.",
  });

  for (let iter = 0; iter < N_ITER; iter++) {
    const antState = [];
    for (let a = 0; a < N_ANTS; a++) {
      antState.push({ tour: [0], visited: new Set([0]) });
    }

    for (let round = 0; round < N - 1; round++) {
      const roundAnts = [];

      for (let a = 0; a < N_ANTS; a++) {
        const st = antState[a];
        const cur = st.tour[st.tour.length - 1];
        const candidates = [];

        for (let next = 0; next < N; next++) {
          if (st.visited.has(next)) continue;
          const tau = Math.pow(pher[cur][next] + 1e-8, ALPHA);
          const eta = Math.pow(1 / (dists[cur][next] + 1e-8), BETA);
          candidates.push({ city: next, tau, eta, raw: tau * eta });
        }

        const total = candidates.reduce((s, c) => s + c.raw, 0);
        candidates.forEach(c => { c.pct = total > 0 ? c.raw / total : 0; });
        candidates.sort((a, b) => b.pct - a.pct);

        let r = Math.random() * total;
        let chosen = candidates[0].city;
        for (const c of candidates) {
          r -= c.raw;
          if (r <= 0) { chosen = c.city; break; }
        }

        st.tour.push(chosen);
        st.visited.add(chosen);

        roundAnts.push({
          antIndex: a,
          tour: [...st.tour],
          currentCity: cur,
          chosen,
          candidates: candidates.slice(0, 7),
        });
      }

      frames.push({
        type: "parallel",
        iteration: iter + 1, round,
        ants: roundAnts,
        pheromone: clonePher(pher),
        bestTour: globalBest ? [...globalBest] : null,
        bestCost: globalBestCost,
        message: `Iteração ${iter + 1} · Rodada ${round + 1}/${N - 1}`,
      });
    }

    const allTours = [];
    const allCosts = [];
    for (let a = 0; a < N_ANTS; a++) {
      const tour = antState[a].tour;
      const cost = routeCost(tour);
      allTours.push([...tour]);
      allCosts.push(cost);
      if (cost < globalBestCost) {
        globalBest = [...tour];
        globalBestCost = cost;
      }
    }

    for (let i = 0; i < N; i++)
      for (let j = 0; j < N; j++)
        pher[i][j] *= 1 - EVAP;

    for (let a = 0; a < N_ANTS; a++) {
      const tour = allTours[a];
      const deposit = Q / allCosts[a];
      for (let k = 0; k < tour.length; k++) {
        const u = tour[k], v = tour[(k + 1) % tour.length];
        pher[u][v] += deposit;
        pher[v][u] += deposit;
      }
    }

    const bestIdx = allCosts.indexOf(Math.min(...allCosts));
    const isNewBest = allCosts[bestIdx] < globalBestCost;
    if (isNewBest) {
      globalBest = [...allTours[bestIdx]];
      globalBestCost = allCosts[bestIdx];
    }

    frames.push({
      type: "update",
      iteration: iter + 1, round: -1, ants: null,
      pheromone: clonePher(pher),
      bestTour: globalBest ? [...globalBest] : null,
      bestCost: globalBestCost,
      allTours: allTours.map(t => [...t]),
      allCosts: [...allCosts],
      message: `Atualização (Iteração ${iter + 1}) — ${
        isNewBest ? "★ Nova melhor rota! " : ""
      }Evaporação ρ=${EVAP} + depósito Q/L. Arestas engrossam onde mais formigas passaram.`,
    });
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
  drawGraph(f);
  drawBottom(f);
}

function updateInfo(f, index, total) {
  const label = f.type === "init" ? "Inicialização"
    : f.type === "parallel" ? `Iteração ${f.iteration} · Rodada ${f.round + 1}`
    : `Atualização feromônio`;
  document.getElementById("step-info").textContent =
    `${label}  (${index + 1}/${total})`;
  document.getElementById("step-description").innerHTML = f.message;
}

function drawGraph(f) {
  ctx.clearRect(0, 0, W, H_LAND);
  ctx.fillStyle = "#fafbfc";
  ctx.fillRect(0, 0, W, H_LAND);

  drawPheromoneEdges(f.pheromone);

  if (f.type === "parallel" && f.ants) {
    for (const ant of f.ants) {
      if (ant.tour.length >= 2) {
        drawTourPath(ant.tour, ANT_COLORS[ant.antIndex], 2.5, false, 0.75);
      }
      const last = ant.tour[ant.tour.length - 1];
      ctx.beginPath();
      ctx.arc(CITIES[last].x, CITIES[last].y, 14, 0, Math.PI * 2);
      ctx.fillStyle = ANT_COLORS[ant.antIndex] + "30";
      ctx.fill();
    }
  }

  if (f.allTours && f.type === "update") {
    for (let a = 0; a < f.allTours.length; a++) {
      drawTourPath(f.allTours[a], ANT_COLORS[a], 2, false, 0.35);
    }
  }

  if (f.bestTour && f.bestTour.length > 1) {
    drawTourPath(f.bestTour, "rgba(243, 156, 18, 0.65)", 2.5, true, 0.9);
  }

  drawCities();

  if (f.type === "parallel" && f.ants) {
    const ant = f.ants[0];
    if (ant && ant.candidates) {
      for (const c of ant.candidates) {
        const px = CITIES[c.city].x, py = CITIES[c.city].y;
        ctx.fillStyle = `rgba(231, 76, 60, ${0.06 + c.pct * 0.5})`;
        ctx.beginPath();
        ctx.arc(px, py, 12, 0, Math.PI * 2);
        ctx.fill();
      }
    }
  }
}

function drawPheromoneEdges(pher) {
  let maxP = 0;
  for (let i = 0; i < N_CITIES; i++)
    for (let j = i + 1; j < N_CITIES; j++)
      if (pher[i][j] > maxP) maxP = pher[i][j];

  if (maxP <= 0) return;

  for (let i = 0; i < N_CITIES; i++) {
    for (let j = i + 1; j < N_CITIES; j++) {
      if (pher[i][j] <= 0) continue;
      const lvl = pher[i][j] / maxP;
      ctx.strokeStyle = `rgba(230, 126, 34, ${0.3 + lvl * 0.55})`;
      ctx.lineWidth = 0.5 + lvl * 4;
      ctx.beginPath();
      ctx.moveTo(CITIES[i].x, CITIES[i].y);
      ctx.lineTo(CITIES[j].x, CITIES[j].y);
      ctx.stroke();
    }
  }
}

function drawTourPath(tour, color, width, dashed, alpha) {
  if (!tour || tour.length < 2) return;
  ctx.save();
  ctx.globalAlpha = alpha;
  ctx.strokeStyle = color;
  ctx.lineWidth = width;
  ctx.lineCap = "round";
  ctx.lineJoin = "round";
  if (dashed) ctx.setLineDash([5, 4]);
  ctx.beginPath();
  ctx.moveTo(CITIES[tour[0]].x, CITIES[tour[0]].y);
  for (let k = 1; k < tour.length; k++)
    ctx.lineTo(CITIES[tour[k]].x, CITIES[tour[k]].y);
  ctx.stroke();
  ctx.restore();
}

function drawCities() {
  for (let i = 0; i < N_CITIES; i++) {
    const p = CITIES[i];
    const r = i === 0 ? 10 : 7;
    ctx.beginPath();
    ctx.arc(p.x, p.y, r, 0, Math.PI * 2);
    ctx.fillStyle = i === 0 ? "#c0392b" : "#2980b9";
    ctx.fill();
    ctx.strokeStyle = "#fff";
    ctx.lineWidth = 2;
    ctx.stroke();
    ctx.fillStyle = "#fff";
    ctx.font = "bold 10px sans-serif";
    ctx.textAlign = "center";
    ctx.textBaseline = "middle";
    ctx.fillText(i, p.x, p.y);
  }
}

function drawBottom(f) {
  const c = chromCtx, w = W2, h = H_CHROM;
  c.clearRect(0, 0, w, h);
  c.fillStyle = "#fafbfc";
  c.fillRect(0, 0, w, h);
  c.strokeStyle = "#ddd";
  c.lineWidth = 1;
  c.beginPath();
  c.moveTo(0, 0);
  c.lineTo(w, 0);
  c.stroke();

  if (f.type === "init") {
    c.fillStyle = "#888";
    c.font = "13px sans-serif";
    c.textAlign = "center";
    c.fillText("Nenhuma aresta ainda — formigas vão construir rotas do zero. Pressione ▶.", w / 2, h / 2);
    return;
  }

  if (f.type === "update") {
    c.fillStyle = "#888";
    c.font = "13px sans-serif";
    c.textAlign = "center";
    c.fillText("Feromônio atualizado — arestas laranjas no grafo ficam mais grossas com o uso.", w / 2, h / 2);
    return;
  }

  if (f.type !== "parallel" || !f.ants) return;

  // ─── Ant status rows ──────────────────────────────────────

  c.fillStyle = "#333";
  c.font = "bold 12px sans-serif";
  c.textAlign = "left";
  c.fillText(`Rodada ${f.round + 1}/${N_CITIES - 1} — 3 formigas em paralelo`, 14, 16);

  for (let a = 0; a < N_ANTS; a++) {
    const ant = f.ants[a];
    const y = 32 + a * 16;
    c.fillStyle = ANT_COLORS[a];
    c.beginPath();
    c.arc(16, y - 3, 5, 0, Math.PI * 2);
    c.fill();
    c.font = "11px monospace";
    c.textAlign = "left";
    c.fillStyle = "#333";
    const path = ant.tour.map(String).join("→");
    const pct = (ant.candidates.find(c => c.city === ant.chosen)?.pct * 100 || 0).toFixed(1);
    c.fillText(`F${a + 1} [${path}]  → ${ant.chosen}  (p=${pct}%)`, 26, y + 4);
  }

  // ─── Detailed decision table for ant 0 ────────────────────

  const ant0 = f.ants[0];
  if (!ant0 || !ant0.candidates || ant0.candidates.length === 0) return;

  const yDiv = 32 + N_ANTS * 16 + 6;
  c.strokeStyle = "#e0e0e0";
  c.lineWidth = 1;
  c.beginPath();
  c.moveTo(10, yDiv);
  c.lineTo(w - 10, yDiv);
  c.stroke();

  c.fillStyle = "#333";
  c.font = "bold 11px sans-serif";
  c.textAlign = "left";
  c.fillText(`Decisão — Formiga 1 (🔴) na cidade ${ant0.currentCity}`, 14, yDiv + 16);

  const col = [18, 68, 130, 200, 290, 352];
  const colW = [48, 58, 65, 85, 55, 340];
  const headers = ["Cid", "τ^α", "η^β", "τ^α·η^β", "Prob", "Roleta Acumulada"];
  const yHdr = yDiv + 20;

  c.fillStyle = "#e8e8e8";
  c.fillRect(10, yHdr, w - 20, 16);
  c.fillStyle = "#333";
  c.font = "bold 9px sans-serif";
  c.textAlign = "left";
  for (let i = 0; i < headers.length; i++) c.fillText(headers[i], col[i], yHdr + 11);

  const rows = ant0.candidates;
  const rowH = 18;

  for (let r = 0; r < rows.length; r++) {
    const n = rows[r];
    const y = yHdr + 18 + r * rowH;

    c.fillStyle = r % 2 === 0 ? "#fff" : "#f8f8f8";
    c.fillRect(10, y - 3, w - 20, rowH);

    c.fillStyle = ANT_COLORS[0];
    c.font = "bold 10px monospace";
    c.textAlign = "center";
    c.fillText(String(n.city), col[0] + colW[0] / 2, y + 8);

    c.fillStyle = "#333";
    c.font = "9px monospace";
    c.textAlign = "right";
    c.fillText(n.tau.toFixed(4), col[1] + colW[1] - 4, y + 8);
    c.fillText(n.eta.toFixed(1), col[2] + colW[2] - 4, y + 8);
    c.fillText(n.raw.toFixed(4), col[3] + colW[3] - 4, y + 8);
    c.textAlign = "right";
    c.fillText((n.pct * 100).toFixed(1) + "%", col[4] + colW[4] - 4, y + 8);

    const barX = col[5], barW = w - col[5] - 14;
    let cumulative = 0;
    for (let k = 0; k <= r; k++) {
      cumulative += rows[k].pct;
    }
    c.fillStyle = "#eee";
    c.fillRect(barX, y, barW, rowH - 4);
    c.fillStyle = ANT_COLORS[0];
    c.fillRect(barX, y, Math.max(1, cumulative * barW), rowH - 4);
  }
}

const init = () => {
  setupCanvases();
  const frames = genFrames();
  createPlayer({
    genFrames: () => frames,
    drawAll: f => drawAll(f),
    updateInfo: (f, i, t) => updateInfo(f, i, t),
    speed: 500,
  }).init();
};

export { init };
