export function createPlayer({ genFrames, drawAll, updateInfo, speed = 600 }) {
  let current = 0, playing = false, timer;
  let frames = [];

  function init() {
    frames = genFrames();
    current = 0;
    render();
    wireButtons();
  }

  function render() {
    const f = frames[current];
    drawAll(f, current, frames.length);
    updateInfo(f, current, frames.length);
  }

  function togglePlay() {
    if (playing) { stop(); return; }
    playing = true;
    document.getElementById("btn-play").textContent = "⏸";
    tick();
  }

  function stop() {
    playing = false; clearTimeout(timer);
    document.getElementById("btn-play").textContent = "▶";
  }

  function tick() {
    if (!playing) return;
    if (current < frames.length - 1) { current++; render(); timer = setTimeout(tick, speed); }
    else stop();
  }

  function wireButtons() {
    document.getElementById("btn-reset").onclick = () => { stop(); current = 0; render(); };
    document.getElementById("btn-prev").onclick  = () => { stop(); if (current > 0) current--; render(); };
    document.getElementById("btn-next").onclick  = () => { stop(); if (current < frames.length - 1) current++; render(); };
    document.getElementById("btn-end").onclick   = () => { stop(); current = frames.length - 1; render(); };
    document.getElementById("btn-play").onclick  = togglePlay;
  }

  return { init };
}

export function scaleCanvas(canvas, w, h) {
  canvas.width = w * 2;
  canvas.height = h * 2;
  canvas.style.width = w + "px";
  canvas.style.height = h + "px";
  const ctx = canvas.getContext("2d");
  ctx.scale(2, 2);
  return ctx;
}
