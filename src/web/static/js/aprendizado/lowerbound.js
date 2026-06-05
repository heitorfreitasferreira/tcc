import { createPlayer, scaleCanvas } from "./aprendizado-base.js";

const W = 700, H = 500;
const POINTS = [
  { x: 100, y: 160, label: "0" },
  { x: 260, y: 100, label: "1" },
  { x: 440, y: 110, label: "2" },
  { x: 530, y: 250, label: "3" },
  { x: 460, y: 400, label: "4" },
  { x: 180, y: 390, label: "5" },
];
const N = POINTS.length;

function dist(a, b) {
  return Math.hypot(POINTS[a].x - POINTS[b].x, POINTS[b].y - POINTS[a].y);
}

function primSteps() {
  const visited = new Set([0]);
  const edges = [];
  const steps = [];
  let total = 0;

  while (visited.size < N) {
    let best = null, bestW = Infinity;
    for (const u of visited) {
      for (let v = 0; v < N; v++) {
        if (visited.has(v)) continue;
        const w = dist(u, v);
        if (w < bestW) { bestW = w; best = [u, v, w]; }
      }
    }
    if (best) {
      edges.push(best);
      visited.add(best[1]);
      total += bestW;
      steps.push({
        type: "mst_build",
        edges: edges.map(e => [...e]),
        visited: new Set(visited),
        total,
        added: best,
      });
    }
  }

  steps.push({
    type: "mst_done",
    edges: edges.map(e => [...e]),
    total,
  });

  const nnRoute = nearestNeighbor();
  const nnCost = routeCost(nnRoute);

  steps.push({
    type: "comparison",
    edges: edges.map(e => [...e]),
    total,
    nnRoute,
    nnCost,
  });

  return steps;
}

function nearestNeighbor() {
  const visited = new Set([0]);
  const route = [0];
  let cur = 0;
  while (visited.size < N) {
    let best = -1, bd = Infinity;
    for (let v = 0; v < N; v++) {
      if (visited.has(v)) continue;
      const d = dist(cur, v);
      if (d < bd) { bd = d; best = v; }
    }
    route.push(best);
    visited.add(best);
    cur = best;
  }
  return route;
}

function routeCost(order) {
  let t = 0;
  for (let i = 1; i < order.length; i++) t += dist(order[i-1], order[i]);
  t += dist(order[order.length-1], order[0]);
  return t;
}

let ctx;

function setupCanvases() {
  const c = document.getElementById("viz-canvas");
  ctx = scaleCanvas(c, W, H);
}

function drawAll(s) {
  ctx.clearRect(0, 0, W, H);
  ctx.fillStyle = "#fafbfc";
  ctx.fillRect(0, 0, W, H);
  drawGrid();

  if (s.type === "mst_build" || s.type === "mst_done") {
    drawMST(s.edges, s.type === "mst_build" ? s.visited : new Set([...Array(N).keys()]));
  }
  if (s.type === "comparison") {
    drawMST(s.edges, new Set([...Array(N).keys()]));
    drawTSPTour(s.nnRoute);
  }

  drawPoints();
  drawTitle(s);
}

function updateInfo(s) {
  if (s.type === "mst_build") {
    const [u, v, w] = s.added;
    document.getElementById("step-info").textContent =
      `Construção da MST — aresta ${s.edges.length} de ${N - 1}`;
    document.getElementById("step-description").innerHTML =
      `<strong>Algoritmo de Prim:</strong> A cada passo, adiciona a aresta <strong>mais leve</strong> ` +
      `que conecta um ponto já visitado a um novo ponto. ` +
      `Aresta adicionada: <strong>${POINTS[u].label} ↔ ${POINTS[v].label}</strong> ` +
      `(peso ${w.toFixed(1)}). Peso parcial: <strong>${s.total.toFixed(1)}</strong>.`;
  } else if (s.type === "mst_done") {
    document.getElementById("step-info").textContent =
      `MST completa — Peso total = ${s.total.toFixed(1)}`;
    document.getElementById("step-description").innerHTML =
      `<strong>MST concluída!</strong> A árvore conecta todos os pontos com o menor custo possível ` +
      `(<strong>${s.total.toFixed(1)}</strong>), usando exatamente ${N-1} arestas e sem formar ciclos. ` +
      `Este valor é um <em>limite inferior</em>: nenhuma rota TSP pode custar menos.`;
  } else {
    document.getElementById("step-info").textContent =
      `Comparação: MST vs. Tour TSP`;
    document.getElementById("step-description").innerHTML =
      `<strong>Por que MST ≤ TSP?</strong> Se você remover <em>qualquer</em> aresta de um tour TSP válido, ` +
      `obtém uma árvore geradora (conecta todos os pontos sem ciclo). ` +
      `Como a MST é a árvore <em>mínima</em>, seu peso (${s.total.toFixed(1)}) é ≤ qualquer tour. ` +
      `Exemplo: tour do vizinho mais próximo = <strong>${s.nnCost.toFixed(1)}</strong> ≥ <strong>${s.total.toFixed(1)}</strong> ✓`;
  }
}

function drawGrid() {
  ctx.strokeStyle = "#e8e8e8"; ctx.lineWidth = 0.5;
  for (let x = 0; x <= W; x += 50) {
    ctx.beginPath(); ctx.moveTo(x, 0); ctx.lineTo(x, H); ctx.stroke();
  }
  for (let y = 0; y <= H; y += 50) {
    ctx.beginPath(); ctx.moveTo(0, y); ctx.lineTo(W, y); ctx.stroke();
  }
}

function drawMST(edges, visited) {
  for (const u of visited) {
    const p = POINTS[u];
    ctx.beginPath();
    ctx.arc(p.x, p.y, 16, 0, Math.PI * 2);
    ctx.fillStyle = "rgba(39,174,96,0.1)";
    ctx.fill();
  }

  for (const [u, v, w] of edges) {
    const pu = POINTS[u], pv = POINTS[v];
    ctx.strokeStyle = "#27ae60";
    ctx.lineWidth = 2.5;
    ctx.lineCap = "round";
    ctx.beginPath();
    ctx.moveTo(pu.x, pu.y);
    ctx.lineTo(pv.x, pv.y);
    ctx.stroke();

    const mx = (pu.x + pv.x) / 2, my = (pu.y + pv.y) / 2;
    const dx = pv.x - pu.x, dy = pv.y - pu.y;
    const len = Math.hypot(dx, dy);
    const nx = -dy / len, ny = dx / len;
    ctx.fillStyle = "#1e8449";
    ctx.font = "bold 11px monospace";
    ctx.textAlign = "center";
    ctx.textBaseline = "middle";
    ctx.fillText(w.toFixed(0), mx + nx * 16, my + ny * 16);
  }
}

function drawTSPTour(route) {
  ctx.strokeStyle = "rgba(142,68,173,0.5)";
  ctx.lineWidth = 2;
  ctx.lineCap = "round";
  ctx.setLineDash([8, 5]);
  ctx.beginPath();
  ctx.moveTo(POINTS[route[0]].x, POINTS[route[0]].y);
  for (let i = 1; i < route.length; i++) {
    ctx.lineTo(POINTS[route[i]].x, POINTS[route[i]].y);
  }
  ctx.lineTo(POINTS[route[0]].x, POINTS[route[0]].y);
  ctx.stroke();
  ctx.setLineDash([]);
}

function drawPoints() {
  for (let i = 0; i < POINTS.length; i++) {
    const p = POINTS[i];
    ctx.beginPath();
    ctx.arc(p.x, p.y, i === 0 ? 11 : 8, 0, Math.PI * 2);
    ctx.fillStyle = i === 0 ? "#c0392b" : "#2980b9";
    ctx.fill();
    ctx.strokeStyle = "#fff";
    ctx.lineWidth = 2.5;
    ctx.stroke();
    ctx.fillStyle = "#fff";
    ctx.font = "bold 11px sans-serif";
    ctx.textAlign = "center";
    ctx.textBaseline = "middle";
    ctx.fillText(i, p.x, p.y);
  }
}

function drawTitle(s) {
  ctx.fillStyle = "#555"; ctx.font = "12px sans-serif"; ctx.textAlign = "left";
  if (s.type === "mst_build") {
    ctx.fillText(`Construção da MST (Prim) — aresta ${s.edges.length} de ${N-1}`, 12, H - 8);
  } else if (s.type === "mst_done") {
    ctx.fillText(`MST completa — Peso = ${s.total.toFixed(1)}`, 12, H - 8);
  } else {
    ctx.fillText(`MST = ${s.total.toFixed(1)}  ≤  TSP ≈ ${s.nnCost.toFixed(1)} (vizinho mais próximo)`, 12, H - 8);
  }
}

const init = () => {
  setupCanvases();
  const steps = primSteps();
  createPlayer({
    genFrames: () => steps,
    drawAll: (s) => drawAll(s),
    updateInfo: (s) => updateInfo(s),
    speed: 1000,
  }).init();
};

export { init };
