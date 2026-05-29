import {
  ARROW_BASE_JITTER_PX,
  ARROW_REPEAT_JITTER_STEP_PX,
  TRAVERSAL_MARKER_RADIUS,
  WORLD_MAX,
  WORLD_MIN,
} from "../constants.js";
import { buildRouteIndices } from "../domain/route.js";
import { createPlotGeometry } from "./geometry.js";

export function drawMapFrame(canvas, points, sequence) {
  if (!canvas || !Array.isArray(points) || points.length === 0) {
    return;
  }

  const ctx = canvas.getContext("2d");
  if (!ctx) {
    return;
  }

  const geometry = createPlotGeometry(canvas);
  drawMapState(ctx, geometry, points, sequence);
}

export function drawMapState(ctx, geometry, points, sequence) {
  drawBackground(ctx, geometry);
  drawGridAndAxis(ctx, geometry);
  drawRoute(ctx, geometry, points, sequence);
  drawPoints(ctx, geometry, points);
}

function drawBackground(ctx, geometry) {
  ctx.clearRect(0, 0, geometry.width, geometry.height);
  ctx.fillStyle = "#ffffff";
  ctx.fillRect(0, 0, geometry.width, geometry.height);
}

function drawGridAndAxis(ctx, geometry) {
  const gridStep = 0.25;

  ctx.strokeStyle = "#e6e6e6";
  ctx.lineWidth = 1;

  for (let world = WORLD_MIN; world <= WORLD_MAX + 1e-9; world += gridStep) {
    if (Math.abs(world) < 1e-9) {
      continue;
    }

    const x = geometry.toCanvasX(world);
    ctx.beginPath();
    ctx.moveTo(x, geometry.plotTop);
    ctx.lineTo(x, geometry.plotBottom);
    ctx.stroke();

    const y = geometry.toCanvasY(world);
    ctx.beginPath();
    ctx.moveTo(geometry.plotLeft, y);
    ctx.lineTo(geometry.plotRight, y);
    ctx.stroke();
  }

  ctx.strokeStyle = "#b8b8b8";
  ctx.lineWidth = 1.2;
  ctx.strokeRect(geometry.plotLeft, geometry.plotTop, geometry.plotSize, geometry.plotSize);

  const axisX = geometry.toCanvasX(0);
  const axisY = geometry.toCanvasY(0);

  ctx.strokeStyle = "#666666";
  ctx.lineWidth = 1.6;

  ctx.beginPath();
  ctx.moveTo(geometry.plotLeft, axisY);
  ctx.lineTo(geometry.plotRight, axisY);
  ctx.stroke();

  ctx.beginPath();
  ctx.moveTo(axisX, geometry.plotTop);
  ctx.lineTo(axisX, geometry.plotBottom);
  ctx.stroke();
}

function drawRoute(ctx, geometry, points, sequence) {
  if (!Array.isArray(sequence) || sequence.length === 0) {
    return;
  }

  const route = buildRouteIndices(points.length, sequence);
  if (route.length < 2) {
    return;
  }

  ctx.strokeStyle = "#e07a10";
  ctx.fillStyle = "#e07a10";
  ctx.lineWidth = 2.2;
  ctx.lineCap = "round";
  ctx.lineJoin = "round";

  const directedEdgeUsage = new Map();

  for (let index = 1; index < route.length; index += 1) {
    const startIndex = route[index - 1];
    const endIndex = route[index];
    const startPoint = points[startIndex];
    const endPoint = points[endIndex];
    const edgeKey = `${startIndex}:${endIndex}`;
    const occurrence = directedEdgeUsage.get(edgeKey) ?? 0;

    directedEdgeUsage.set(edgeKey, occurrence + 1);

    drawDirectedSegment(
      ctx,
      geometry,
      startPoint,
      endPoint,
      computeDirectedEdgeJitter(occurrence),
    );
  }
}

function computeDirectedEdgeJitter(occurrence) {
  if (occurrence <= 0) {
    return ARROW_BASE_JITTER_PX;
  }

  const lane = Math.ceil(occurrence / 2);
  const signal = occurrence % 2 === 1 ? 1 : -1;

  return ARROW_BASE_JITTER_PX + signal * lane * ARROW_REPEAT_JITTER_STEP_PX;
}

function drawDirectedSegment(ctx, geometry, startPoint, endPoint, jitterOffset) {
  if (!startPoint || !endPoint) {
    return;
  }

  const startX = geometry.toCanvasX(startPoint[0]);
  const startY = geometry.toCanvasY(startPoint[1]);
  const endX = geometry.toCanvasX(endPoint[0]);
  const endY = geometry.toCanvasY(endPoint[1]);

  const deltaX = endX - startX;
  const deltaY = endY - startY;
  const length = Math.hypot(deltaX, deltaY);
  if (length < 1e-6) {
    return;
  }

  const unitX = deltaX / length;
  const unitY = deltaY / length;
  const normalX = -unitY;
  const normalY = unitX;

  const shiftedStartX = startX + normalX * jitterOffset;
  const shiftedStartY = startY + normalY * jitterOffset;
  const shiftedEndX = endX + normalX * jitterOffset;
  const shiftedEndY = endY + normalY * jitterOffset;

  const pointRadius = 4;
  const pointGap = 2;
  const endpointOffset = pointRadius + pointGap;
  const arrowHeadLength = Math.min(12, Math.max(8, length * 0.28));
  const arrowHeadHalfWidth = arrowHeadLength * 0.5;

  const shaftStartX = shiftedStartX + unitX * endpointOffset;
  const shaftStartY = shiftedStartY + unitY * endpointOffset;
  const tipX = shiftedEndX - unitX * endpointOffset;
  const tipY = shiftedEndY - unitY * endpointOffset;
  const shaftEndX = tipX - unitX * arrowHeadLength;
  const shaftEndY = tipY - unitY * arrowHeadLength;

  const shaftLength = Math.hypot(shaftEndX - shaftStartX, shaftEndY - shaftStartY);
  if (shaftLength < 1) {
    return;
  }

  ctx.beginPath();
  ctx.moveTo(shaftStartX, shaftStartY);
  ctx.lineTo(shaftEndX, shaftEndY);
  ctx.stroke();

  ctx.beginPath();
  ctx.moveTo(tipX, tipY);
  ctx.lineTo(shaftEndX + normalX * arrowHeadHalfWidth, shaftEndY + normalY * arrowHeadHalfWidth);
  ctx.lineTo(shaftEndX - normalX * arrowHeadHalfWidth, shaftEndY - normalY * arrowHeadHalfWidth);
  ctx.closePath();
  ctx.fill();
}

function drawPoints(ctx, geometry, points) {
  points.forEach((point, index) => {
    const insideRange =
      point[0] >= WORLD_MIN &&
      point[0] <= WORLD_MAX &&
      point[1] >= WORLD_MIN &&
      point[1] <= WORLD_MAX;

    if (!insideRange) {
      return;
    }

    const x = geometry.toCanvasX(point[0]);
    const y = geometry.toCanvasY(point[1]);
    const pointColor = index === 0 ? "#d62828" : "#0a66c2";

    ctx.fillStyle = pointColor;
    ctx.beginPath();
    ctx.arc(x, y, 4, 0, Math.PI * 2);
    ctx.fill();

    ctx.fillStyle = "#303030";
    ctx.font = "12px sans-serif";
    ctx.fillText(String(index), x + 7, y - 7);
  });
}

export function drawTraversalMarker(ctx, geometry, point) {
  if (!Array.isArray(point) || point.length !== 2) {
    return;
  }

  const x = geometry.toCanvasX(point[0]);
  const y = geometry.toCanvasY(point[1]);

  ctx.fillStyle = "#111111";
  ctx.strokeStyle = "#ffffff";
  ctx.lineWidth = 2;
  ctx.beginPath();
  ctx.arc(x, y, TRAVERSAL_MARKER_RADIUS, 0, Math.PI * 2);
  ctx.fill();
  ctx.stroke();
}
