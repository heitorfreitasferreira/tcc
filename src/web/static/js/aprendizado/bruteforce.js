import { createPlayer, scaleCanvas } from "./aprendizado-base.js";

const W = 700, H = 500;
const XL = 0, XR = 5;
const N = 80;
const PAD = 0.15;

function f(x) {
  return Math.sin(2.5 * x) * Math.cos(1.8 * x) + 0.4 * Math.sin(0.7 * x + 1) + 0.6;
}

let ctx, yMin, yMax, yRange, stepsCache;

function mx(x) { return 60 + ((x - XL) / (XR - XL)) * (W - 120); }
function my(y) {
  return 40 + (1 - (y - yMin) / yRange) * (H - 90);
}

function computeYRange(steps) {
  const vals = steps.map(s => s.y);
  const rawMin = Math.min(...vals);
  const rawMax = Math.max(...vals);
  const pad = (rawMax - rawMin) * PAD || 0.5;
  yMin = Math.floor((rawMin - pad) / 0.5) * 0.5;
  yMax = Math.ceil((rawMax + pad) / 0.5) * 0.5;
  yRange = yMax - yMin;
}

function genSteps() {
  const steps = [];
  let bestX = XL, bestY = Infinity;
  const dx = (XR - XL) / N;

  for (let i = 0; i <= N; i++) {
    const x = XL + i * dx;
    const y = f(x);
    if (y < bestY) { bestY = y; bestX = x; }
    steps.push({ x, y, bestX, bestY, step: i, total: N + 1 });
  }
  computeYRange(steps);
  return steps;
}

function setupCanvases() {
  const c = document.getElementById("viz-canvas");
  ctx = scaleCanvas(c, W, H);
}

function drawAll(s) {
  ctx.clearRect(0, 0, W, H);
  ctx.fillStyle = "#fafbfc";
  ctx.fillRect(0, 0, W, H);

  drawAxes();
  drawCurve();
  drawEvaluated(s);
  drawCurrentPoint(s);
  drawBestPoint(s);
}

function updateInfo(s) {
  document.getElementById("step-info").textContent =
    `Avaliação ${s.step} / ${s.total - 1}`;
  document.getElementById("step-description").innerHTML =
    `<strong>x = ${s.x.toFixed(2)}</strong> → f(x) = ${s.y.toFixed(4)} &nbsp;|&nbsp; ` +
    `<strong>Melhor até agora:</strong> f(${s.bestX.toFixed(2)}) = <span style="color:#0a66c2;font-weight:700">${s.bestY.toFixed(4)}</span>` +
    (s.step === s.total - 1 ? `<br><br><strong>✓ Busca exaustiva concluída!</strong> O mínimo da função neste intervalo é f(${s.bestX.toFixed(2)}) = ${s.bestY.toFixed(4)}. A Busca Exaustiva testa <em>todos</em> os pontos possíveis — garante o ótimo global, mas é muito lenta para instâncias grandes (${s.total} avaliações aqui).` : '');
}

function drawAxes() {
  // y=0 reference line
  const y0 = Math.max(30, Math.min(H - 40, my(0)));
  ctx.strokeStyle = "#ccc"; ctx.lineWidth = 1;
  ctx.setLineDash([3, 3]);
  ctx.beginPath(); ctx.moveTo(50, y0); ctx.lineTo(W - 40, y0); ctx.stroke();
  ctx.setLineDash([]);

  // y axis
  ctx.strokeStyle = "#aaa"; ctx.lineWidth = 1;
  ctx.beginPath(); ctx.moveTo(mx(0), 25); ctx.lineTo(mx(0), H - 40); ctx.stroke();

  ctx.fillStyle = "#888"; ctx.font = "12px sans-serif";
  ctx.textAlign = "center";
  ctx.fillText("f(x)", mx(0) - 6, 18);

  // y ticks
  ctx.fillStyle = "#666"; ctx.font = "11px sans-serif";
  ctx.textAlign = "right";
  const yStep = yRange <= 1.5 ? 0.25 : yRange <= 4 ? 0.5 : 1;
  for (let y = Math.ceil(yMin / yStep) * yStep; y <= yMax; y += yStep) {
    const cy = my(y);
    if (cy < 25 || cy > H - 45) continue;
    ctx.fillText(y.toFixed(yStep < 1 ? 2 : 1), mx(0) - 6, cy + 4);
  }

  // x axis
  ctx.strokeStyle = "#aaa"; ctx.lineWidth = 1;
  ctx.beginPath(); ctx.moveTo(mx(XL), H - 40); ctx.lineTo(mx(XR), H - 40); ctx.stroke();

  // x ticks
  ctx.fillStyle = "#666"; ctx.font = "11px sans-serif";
  ctx.textAlign = "center";
  for (let x = 0; x <= XR; x++) {
    ctx.fillText(x, mx(x), H - 40 + 18);
  }
  ctx.fillStyle = "#888"; ctx.font = "12px sans-serif";
  ctx.fillText("x", W - 20, H - 40 + 18);
}

function drawCurve() {
  ctx.strokeStyle = "#bbb"; ctx.lineWidth = 1.5;
  ctx.beginPath();
  for (let i = 0; i <= 500; i++) {
    const x = XL + (i / 500) * (XR - XL);
    const y = f(x);
    const px = mx(x), py = my(y);
    if (py < 20 || py > H - 35) continue;
    if (i === 0) ctx.moveTo(px, py);
    else ctx.lineTo(px, py);
  }
  ctx.stroke();
}

function drawEvaluated(s) {
  for (let i = 0; i < s.step; i++) {
    const si = stepsCache[i];
    const px = mx(si.x), py = my(si.y);
    ctx.beginPath(); ctx.arc(px, py, 3, 0, Math.PI * 2);
    ctx.fillStyle = "rgba(150,150,150,0.5)";
    ctx.fill();
  }

  ctx.strokeStyle = "rgba(10,102,194,0.25)";
  ctx.lineWidth = 2;
  ctx.setLineDash([4, 4]);
  ctx.beginPath();
  let started = false;
  for (let i = 0; i <= s.step; i++) {
    const si = stepsCache[i];
    const px = mx(si.bestX), py = my(si.bestY);
    if (!started) { ctx.moveTo(px, py); started = true; }
    else ctx.lineTo(px, py);
  }
  ctx.stroke();
  ctx.setLineDash([]);
}

function drawCurrentPoint(s) {
  const px = mx(s.x), py = my(s.y);
  ctx.strokeStyle = "rgba(200,50,50,0.4)";
  ctx.lineWidth = 1;
  ctx.setLineDash([3, 3]);
  ctx.beginPath(); ctx.moveTo(px, my(0)); ctx.lineTo(px, py); ctx.stroke();
  ctx.setLineDash([]);

  ctx.beginPath(); ctx.arc(px, py, 5.5, 0, Math.PI * 2);
  ctx.fillStyle = "#e74c3c";
  ctx.fill();
  ctx.strokeStyle = "#fff"; ctx.lineWidth = 2; ctx.stroke();

  ctx.fillStyle = "#333"; ctx.font = "bold 11px monospace";
  ctx.textAlign = "left";
  ctx.fillText(`f(${s.x.toFixed(2)})=${s.y.toFixed(3)}`, px + 10, py - 6);
}

function drawBestPoint(s) {
  const px = mx(s.bestX), py = my(s.bestY);
  ctx.beginPath(); ctx.arc(px, py, 7, 0, Math.PI * 2);
  ctx.fillStyle = "#0a66c2";
  ctx.fill();
  ctx.strokeStyle = "#fff"; ctx.lineWidth = 2.5; ctx.stroke();
  ctx.fillStyle = "#0a66c2"; ctx.font = "bold 11px monospace";
  ctx.textAlign = "right";
  ctx.fillText(`★ melhor: f(${s.bestX.toFixed(2)})=${s.bestY.toFixed(3)}`, px - 10, py - 12);
}

const init = () => {
  setupCanvases();
  stepsCache = genSteps();
  const steps = stepsCache;
  createPlayer({
    genFrames: () => steps,
    drawAll: (s) => drawAll(s),
    updateInfo: (s) => updateInfo(s),
    speed: 60,
  }).init();
};

export { init };
