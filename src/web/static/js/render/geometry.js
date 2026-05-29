import { WORLD_MIN, WORLD_MAX, WORLD_RANGE } from "../constants.js";

export function createPlotGeometry(canvas) {
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

  return {
    width,
    height,
    plotSize,
    plotLeft,
    plotTop,
    plotRight,
    plotBottom,
    toCanvasX: (x) => plotLeft + ((x - WORLD_MIN) / WORLD_RANGE) * plotSize,
    toCanvasY: (y) => plotTop + ((WORLD_MAX - y) / WORLD_RANGE) * plotSize,
  };
}
