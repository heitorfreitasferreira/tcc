import { parseEvolution, parsePoints } from "../data/parsers.js";
import { createRunVisualizationController } from "../player/run-controller.js";
import { drawMapFrame } from "../render/map-renderer.js";
import { connectVisualizationControls } from "../ui/controls.js";
import { updateFrameMetadata } from "../ui/metadata.js";

let activeVisualizationController = null;

export function stopActiveVisualizationController() {
  if (!activeVisualizationController) {
    return;
  }

  activeVisualizationController.stop();
  activeVisualizationController = null;
}

export function renderMainContent() {
  stopActiveVisualizationController();

  const canvas = document.getElementById("points-canvas");
  if (!canvas) {
    return;
  }

  const points = parsePoints(canvas.dataset.points);
  if (!points || points.length === 0) {
    return;
  }

  const evolution = parseEvolution(canvas.dataset.evolution);
  if (!evolution || evolution.frames.length === 0) {
    drawMapFrame(canvas, points, []);
    updateFrameMetadata(null, evolution ? evolution.iterations : 0);
    connectVisualizationControls(null);
    return;
  }

  const controller = createRunVisualizationController({
    canvas,
    points,
    evolution,
  });

  activeVisualizationController = controller;
  controller.start();
  connectVisualizationControls(controller);
}

function bindHTMXListeners() {
  if (!document.body) {
    return;
  }

  document.body.addEventListener("htmx:beforeSwap", (event) => {
    const target = event.detail?.target;
    if (target?.id === "main-content") {
      stopActiveVisualizationController();
    }
  });

  document.body.addEventListener("htmx:afterSettle", (event) => {
    const target = event.detail?.target;
    if (target?.id === "main-content") {
      renderMainContent();
    }
  });
}

export function initMainContentApp() {
  if (document.readyState === "loading") {
    document.addEventListener(
      "DOMContentLoaded",
      () => {
        renderMainContent();
        bindHTMXListeners();
      },
      { once: true },
    );
    return;
  }

  renderMainContent();
  bindHTMXListeners();
}
