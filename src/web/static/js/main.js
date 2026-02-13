function parsePoints(rawPoints) {
  if (!rawPoints) {
    return null;
  }

  try {
    const parsed = JSON.parse(rawPoints);
    if (!Array.isArray(parsed)) {
      return null;
    }

    return parsed.filter((point) => {
      return (
        Array.isArray(point) &&
        point.length === 2 &&
        Number.isFinite(point[0]) &&
        Number.isFinite(point[1])
      );
    });
  } catch {
    return null;
  }
}

const WORLD_MIN = -1;
const WORLD_MAX = 1;
const WORLD_RANGE = WORLD_MAX - WORLD_MIN;

function isNearlyZero(value) {
  return Math.abs(value) < 1e-9;
}

function formatTick(value) {
  if (isNearlyZero(value)) {
    return "0";
  }

  return Number(value).toFixed(2).replace(/\.00$/, "").replace(/0$/, "");
}

function drawPoints(canvas, points) {
  if (!canvas || points.length === 0) {
    return;
  }

  const ctx = canvas.getContext("2d");
  if (!ctx) {
    return;
  }

  const width = canvas.width;
  const height = canvas.height;

  const outerPadding = 16;
  const usableWidth = width - outerPadding * 2;
  const usableHeight = height - outerPadding * 2;
  const plotSize = Math.min(usableWidth, usableHeight);
  const plotLeft = (width - plotSize) / 2;
  const plotTop = (height - plotSize) / 2;
  const plotRight = plotLeft + plotSize;
  const plotBottom = plotTop + plotSize;

  const toCanvasX = (x) => {
    return plotLeft + ((x - WORLD_MIN) / WORLD_RANGE) * plotSize;
  };

  const toCanvasY = (y) => {
    return plotTop + ((WORLD_MAX - y) / WORLD_RANGE) * plotSize;
  };

  ctx.clearRect(0, 0, width, height);
  ctx.fillStyle = "#ffffff";
  ctx.fillRect(0, 0, width, height);

  ctx.strokeStyle = "#e6e6e6";
  ctx.lineWidth = 1;

  const gridStep = 0.25;
  for (let world = WORLD_MIN; world <= WORLD_MAX + 1e-9; world += gridStep) {
    if (isNearlyZero(world)) {
      continue;
    }

    const x = toCanvasX(world);
    ctx.beginPath();
    ctx.moveTo(x, plotTop);
    ctx.lineTo(x, plotBottom);
    ctx.stroke();

    const y = toCanvasY(world);
    ctx.beginPath();
    ctx.moveTo(plotLeft, y);
    ctx.lineTo(plotRight, y);
    ctx.stroke();
  }

  ctx.strokeStyle = "#b8b8b8";
  ctx.lineWidth = 1.2;
  ctx.strokeRect(plotLeft, plotTop, plotSize, plotSize);

  const axisX = toCanvasX(0);
  const axisY = toCanvasY(0);

  ctx.strokeStyle = "#666666";
  ctx.lineWidth = 1.6;

  ctx.beginPath();
  ctx.moveTo(plotLeft, axisY);
  ctx.lineTo(plotRight, axisY);
  ctx.stroke();

  ctx.beginPath();
  ctx.moveTo(axisX, plotTop);
  ctx.lineTo(axisX, plotBottom);
  ctx.stroke();


  ctx.fillStyle = "#444";
  ctx.font = "12px sans-serif";



  ctx.fillStyle = "#333";


  ctx.fillStyle = "#0a66c2";
  ctx.font = "12px sans-serif";

  points.forEach((point, index) => {
    const x = toCanvasX(point[0]);
    const y = toCanvasY(point[1]);
    const insideRange =
      point[0] >= WORLD_MIN &&
      point[0] <= WORLD_MAX &&
      point[1] >= WORLD_MIN &&
      point[1] <= WORLD_MAX;

    if (!insideRange) {
      return;
    }

    const pointColor = index === 0 ? "#d62828" : "#0a66c2";

    ctx.fillStyle = pointColor;
    ctx.beginPath();
    ctx.arc(x, y, 4, 0, Math.PI * 2);
    ctx.fill();

    ctx.fillStyle = "#303030";
    ctx.fillText(String(index), x + 7, y - 7);
  });

  ctx.fillStyle = "#333";
  // ctx.fillText("(0,0)", axisX + 8, axisY - 8);
}

function renderPointsFromMain() {
  const canvas = document.getElementById("points-canvas");
  if (!canvas) {
    return;
  }

  const points = parsePoints(canvas.dataset.points);
  if (!points || points.length === 0) {
    return;
  }

  drawPoints(canvas, points);
}

document.addEventListener("DOMContentLoaded", renderPointsFromMain);

document.body.addEventListener("htmx:afterSettle", (event) => {
  const target = event.detail?.target;
  if (target?.id === "main-content") {
    renderPointsFromMain();
  }
});
