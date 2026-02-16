const WORLD_MIN = -1;
const WORLD_MAX = 1;
const WORLD_RANGE = WORLD_MAX - WORLD_MIN;
const BASE_ITERATION_DELAY_MS = 100;
const MIN_FRAME_DELAY_MS = 16;
const DEFAULT_SPEED = 1;
const ALLOWED_SPEEDS = new Set([0.5, 1, 2, 4]);
const ARROW_BASE_JITTER_PX = 3;
const ARROW_REPEAT_JITTER_STEP_PX = 0.9;
const BEST_SEQUENCE_EDGE_DURATION_MS = 700;
const TRAVERSAL_MARKER_RADIUS = 6;

let activeVisualizationController = null;

function parsePoints(rawPoints) {
  if (!rawPoints) {
    return null;
  }

  try {
    const parsed = JSON.parse(rawPoints);
    if (!Array.isArray(parsed)) {
      return null;
    }

    return parsed
      .map((point) => {
        if (!Array.isArray(point) || point.length !== 2) {
          return null;
        }

        const x = Number(point[0]);
        const y = Number(point[1]);
        if (!Number.isFinite(x) || !Number.isFinite(y)) {
          return null;
        }

        return [x, y];
      })
      .filter(Boolean);
  } catch {
    return null;
  }
}

function parseEvolution(rawEvolution) {
  if (!rawEvolution) {
    return null;
  }

  try {
    const parsed = JSON.parse(rawEvolution);
    if (!parsed || typeof parsed !== "object") {
      return null;
    }

    const iterations = toNonNegativeInt(parsed.iterations, 0);
    if (!Array.isArray(parsed.frames)) {
      return {
        iterations,
        frames: [],
      };
    }

    const frames = parsed.frames
      .map((frame) => {
        if (!frame || typeof frame !== "object") {
          return null;
        }

        const sequence = Array.isArray(frame.best_sequence)
          ? frame.best_sequence
              .map((value) => Number(value))
              .filter((value) => Number.isFinite(value))
              .map((value) => Math.trunc(value))
          : [];

        return {
          iter: toNonNegativeInt(frame.iter, 0),
          evalCount: toNonNegativeInt(frame.eval_count, 0),
          bestMakespan: Number(frame.best_makespan),
          bestSequence: sequence,
        };
      })
      .filter(Boolean);

    return {
      iterations,
      frames,
    };
  } catch {
    return null;
  }
}

function toNonNegativeInt(value, fallback) {
  const parsed = Number(value);
  if (!Number.isFinite(parsed)) {
    return fallback;
  }

  const intValue = Math.trunc(parsed);
  if (intValue < 0) {
    return fallback;
  }

  return intValue;
}

function createPlotGeometry(canvas) {
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

function drawMapFrame(canvas, points, sequence) {
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

function drawMapState(ctx, geometry, points, sequence) {
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
  ctx.strokeRect(
    geometry.plotLeft,
    geometry.plotTop,
    geometry.plotSize,
    geometry.plotSize,
  );

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
  ctx.lineTo(
    shaftEndX + normalX * arrowHeadHalfWidth,
    shaftEndY + normalY * arrowHeadHalfWidth,
  );
  ctx.lineTo(
    shaftEndX - normalX * arrowHeadHalfWidth,
    shaftEndY - normalY * arrowHeadHalfWidth,
  );
  ctx.closePath();
  ctx.fill();
}

function buildRouteIndices(pointsCount, sequence) {
  if (pointsCount <= 0 || !Array.isArray(sequence) || sequence.length === 0) {
    return [];
  }

  const normalized = [];
  for (const value of sequence) {
    const index = Math.trunc(Number(value));
    if (!Number.isFinite(index) || index < 0 || index >= pointsCount) {
      continue;
    }

    normalized.push(index);
  }

  if (normalized.length === 0) {
    return [];
  }

  const route = [];
  if (!normalized.includes(0)) {
    route.push(0);
  }
  route.push(...normalized);

  if (route.length > 1 && route[route.length - 1] !== route[0]) {
    route.push(route[0]);
  }

  return route;
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

function drawTraversalMarker(ctx, geometry, point) {
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

function interpolateRoutePoint(points, route, segmentIndex, segmentProgress) {
  if (!Array.isArray(route) || route.length < 2) {
    return null;
  }

  const boundedSegmentIndex = Math.max(0, Math.min(route.length - 2, segmentIndex));
  const startPoint = points[route[boundedSegmentIndex]];
  const endPoint = points[route[boundedSegmentIndex + 1]];
  if (!startPoint || !endPoint) {
    return null;
  }

  const progress = Math.max(0, Math.min(1, segmentProgress));
  return [
    startPoint[0] + (endPoint[0] - startPoint[0]) * progress,
    startPoint[1] + (endPoint[1] - startPoint[1]) * progress,
  ];
}

function findBestSequenceFrame(frames) {
  if (!Array.isArray(frames) || frames.length === 0) {
    return null;
  }

  let bestFrame = null;
  let fallbackFrame = null;

  for (const frame of frames) {
    if (!frame || !Array.isArray(frame.bestSequence) || frame.bestSequence.length === 0) {
      continue;
    }

    fallbackFrame = frame;
    if (!Number.isFinite(frame.bestMakespan)) {
      continue;
    }

    if (
      !bestFrame ||
      !Number.isFinite(bestFrame.bestMakespan) ||
      frame.bestMakespan < bestFrame.bestMakespan
    ) {
      bestFrame = frame;
    }
  }

  return bestFrame || fallbackFrame;
}

function modeLabel(mode) {
  return mode === "best-sequence" ? "Best sequence" : "Evolucao";
}

function formatBestSequence(sequence) {
  if (!Array.isArray(sequence) || sequence.length === 0) {
    return "-";
  }

  return sequence.map((value) => String(value)).join(" -> ");
}

function updateFrameMetadata(frame, totalIterations) {
  const iterElement = document.getElementById("frame-iter");
  const evalElement = document.getElementById("frame-eval");
  const makespanElement = document.getElementById("frame-makespan");
  const bestSequenceElement = document.getElementById("frame-best-sequence");

  if (!iterElement || !evalElement || !makespanElement) {
    return;
  }

  if (!frame) {
    iterElement.textContent = "-";
    evalElement.textContent = "-";
    makespanElement.textContent = "-";
    if (bestSequenceElement) {
      bestSequenceElement.textContent = "-";
    }
    return;
  }

  let iterText = String(frame.iter);
  if (totalIterations > 0) {
    iterText = `${iterText} / ${totalIterations}`;
  }

  iterElement.textContent = iterText;
  evalElement.textContent = String(frame.evalCount);
  makespanElement.textContent = Number.isFinite(frame.bestMakespan)
    ? frame.bestMakespan.toFixed(6)
    : "-";

  if (bestSequenceElement) {
    bestSequenceElement.textContent = formatBestSequence(frame.bestSequence);
  }
}

function updateSpeedButtons(buttons, activeSpeed, enabled) {
  buttons.forEach((button) => {
    const buttonSpeed = Number(button.dataset.speed);
    const supported = ALLOWED_SPEEDS.has(buttonSpeed);

    button.disabled = !enabled || !supported;
    button.classList.toggle(
      "is-active",
      enabled && supported && buttonSpeed === activeSpeed,
    );
  });
}

function createEvolutionPlayer({ canvas, points, frames, iterations }) {
  let speedMultiplier = DEFAULT_SPEED;
  let frameIndex = 0;
  let iterationCursor = 0;
  let timeoutID = null;
  let stopped = false;
  let started = false;
  let paused = false;

  const cycleStartIteration = frames.length > 0 ? frames[0].iter : 0;
  const cycleEndIteration = Math.max(
    cycleStartIteration,
    iterations > 0 ? iterations : frames.length > 0 ? frames[frames.length - 1].iter : 0,
  );

  const clearTimer = () => {
    if (timeoutID !== null) {
      window.clearTimeout(timeoutID);
      timeoutID = null;
    }
  };

  const findFrameIndexForIteration = (iter) => {
    if (frames.length === 0) {
      return 0;
    }

    let index = 0;
    while (index < frames.length - 1 && frames[index + 1].iter <= iter) {
      index += 1;
    }

    return index;
  };

  const renderCurrentIteration = () => {
    const frame = frames[frameIndex];
    drawMapFrame(canvas, points, frame.bestSequence);
    updateFrameMetadata(
      {
        ...frame,
        iter: iterationCursor,
      },
      iterations,
    );
  };

  const computeDelay = () =>
    Math.max(MIN_FRAME_DELAY_MS, BASE_ITERATION_DELAY_MS / speedMultiplier);

  const advanceIteration = () => {
    if (iterationCursor >= cycleEndIteration) {
      iterationCursor = cycleStartIteration;
      frameIndex = findFrameIndexForIteration(iterationCursor);
      return;
    }

    iterationCursor += 1;

    while (frameIndex < frames.length - 1 && frames[frameIndex + 1].iter <= iterationCursor) {
      frameIndex += 1;
    }
  };

  const scheduleNextFrame = () => {
    if (stopped || paused || frames.length === 0) {
      return;
    }

    clearTimer();
    timeoutID = window.setTimeout(() => {
      if (stopped || paused) {
        return;
      }

      advanceIteration();
      renderCurrentIteration();
      scheduleNextFrame();
    }, computeDelay());
  };

  return {
    start() {
      if (stopped || started) {
        return;
      }

      started = true;
      paused = false;
      iterationCursor = cycleStartIteration;

      if (frames.length === 0) {
        drawMapFrame(canvas, points, []);
        updateFrameMetadata(null, iterations);
        return;
      }

      frameIndex = findFrameIndexForIteration(iterationCursor);
      renderCurrentIteration();
      scheduleNextFrame();
    },
    pause() {
      if (stopped || !started || paused) {
        return;
      }

      paused = true;
      clearTimer();
    },
    resume() {
      if (stopped || !started || !paused) {
        return;
      }

      paused = false;
      scheduleNextFrame();
    },
    restart() {
      if (stopped || !started) {
        return;
      }

      iterationCursor = cycleStartIteration;
      clearTimer();

      if (frames.length === 0) {
        drawMapFrame(canvas, points, []);
        updateFrameMetadata(null, iterations);
        return;
      }

      frameIndex = findFrameIndexForIteration(iterationCursor);
      renderCurrentIteration();
      if (!paused) {
        scheduleNextFrame();
      }
    },
    stop() {
      stopped = true;
      clearTimer();
    },
    setSpeed(nextSpeed) {
      if (ALLOWED_SPEEDS.has(nextSpeed)) {
        speedMultiplier = nextSpeed;
      }
    },
    isPaused() {
      return paused;
    },
  };
}

function createBestSequenceTraversalPlayer({ canvas, points, frame, iterations }) {
  const sequence = frame && Array.isArray(frame.bestSequence) ? frame.bestSequence : [];
  const route = buildRouteIndices(points.length, sequence);
  const segmentCount = Math.max(0, route.length - 1);

  let speedMultiplier = DEFAULT_SPEED;
  let animationFrameID = null;
  let previousTimestamp = null;
  let segmentIndex = 0;
  let segmentProgress = 0;
  let stopped = false;
  let started = false;
  let paused = false;

  const clearAnimationFrame = () => {
    if (animationFrameID !== null) {
      window.cancelAnimationFrame(animationFrameID);
      animationFrameID = null;
    }
  };

  const resetProgress = () => {
    segmentIndex = 0;
    segmentProgress = 0;
    previousTimestamp = null;
  };

  const renderCurrentState = () => {
    const ctx = canvas.getContext("2d");
    if (!ctx) {
      return;
    }

    const geometry = createPlotGeometry(canvas);
    drawMapState(ctx, geometry, points, sequence);

    const markerPoint = interpolateRoutePoint(points, route, segmentIndex, segmentProgress);
    if (!markerPoint) {
      return;
    }

    drawTraversalMarker(ctx, geometry, markerPoint);
  };

  const step = (timestamp) => {
    if (stopped || paused) {
      return;
    }

    if (previousTimestamp === null) {
      previousTimestamp = timestamp;
    }

    const elapsed = Math.max(0, timestamp - previousTimestamp);
    previousTimestamp = timestamp;

    if (segmentCount > 0) {
      segmentProgress += (elapsed * speedMultiplier) / BEST_SEQUENCE_EDGE_DURATION_MS;

      while (segmentProgress >= 1) {
        segmentProgress -= 1;
        segmentIndex = (segmentIndex + 1) % segmentCount;
      }
    }

    renderCurrentState();
    animationFrameID = window.requestAnimationFrame(step);
  };

  return {
    start() {
      if (stopped || started) {
        return;
      }

      started = true;
      paused = false;
      resetProgress();
      updateFrameMetadata(frame || null, iterations);
      renderCurrentState();

      if (segmentCount > 0) {
        animationFrameID = window.requestAnimationFrame(step);
      }
    },
    pause() {
      if (stopped || !started || paused) {
        return;
      }

      paused = true;
      clearAnimationFrame();
    },
    resume() {
      if (stopped || !started || !paused) {
        return;
      }

      paused = false;
      previousTimestamp = null;
      if (segmentCount > 0) {
        animationFrameID = window.requestAnimationFrame(step);
      }
    },
    restart() {
      if (stopped || !started) {
        return;
      }

      clearAnimationFrame();
      resetProgress();
      updateFrameMetadata(frame || null, iterations);
      renderCurrentState();

      if (!paused && segmentCount > 0) {
        animationFrameID = window.requestAnimationFrame(step);
      }
    },
    stop() {
      stopped = true;
      clearAnimationFrame();
    },
    setSpeed(nextSpeed) {
      if (ALLOWED_SPEEDS.has(nextSpeed)) {
        speedMultiplier = nextSpeed;
      }
    },
    isPaused() {
      return paused;
    },
  };
}

function createRunVisualizationController({ canvas, points, evolution }) {
  const bestFrame = findBestSequenceFrame(evolution.frames);
  const hasBestSequence = Boolean(bestFrame);

  let mode = "evolution";
  let speedMultiplier = DEFAULT_SPEED;
  let paused = false;
  let stopped = false;
  let activePlayer = null;

  const createPlayerForMode = () => {
    if (mode === "best-sequence" && hasBestSequence) {
      return createBestSequenceTraversalPlayer({
        canvas,
        points,
        frame: bestFrame,
        iterations: evolution.iterations,
      });
    }

    return createEvolutionPlayer({
      canvas,
      points,
      frames: evolution.frames,
      iterations: evolution.iterations,
    });
  };

  const replacePlayer = () => {
    if (activePlayer) {
      activePlayer.stop();
      activePlayer = null;
    }

    activePlayer = createPlayerForMode();
    activePlayer.setSpeed(speedMultiplier);
    activePlayer.start();
    if (paused) {
      activePlayer.pause();
    }
  };

  return {
    start() {
      if (stopped) {
        return;
      }

      replacePlayer();
    },
    stop() {
      if (stopped) {
        return;
      }

      stopped = true;
      if (activePlayer) {
        activePlayer.stop();
        activePlayer = null;
      }
    },
    setSpeed(nextSpeed) {
      if (!ALLOWED_SPEEDS.has(nextSpeed)) {
        return;
      }

      speedMultiplier = nextSpeed;
      if (activePlayer) {
        activePlayer.setSpeed(nextSpeed);
      }
    },
    togglePause() {
      if (!activePlayer) {
        return paused;
      }

      paused = !paused;
      if (paused) {
        activePlayer.pause();
      } else {
        activePlayer.resume();
      }

      return paused;
    },
    restart() {
      if (!activePlayer) {
        return;
      }

      activePlayer.restart();
    },
    toggleMode() {
      if (!hasBestSequence) {
        return mode;
      }

      mode = mode === "evolution" ? "best-sequence" : "evolution";
      replacePlayer();
      return mode;
    },
    getMode() {
      return mode;
    },
    getSpeed() {
      return speedMultiplier;
    },
    isPaused() {
      return paused;
    },
    canToggleBestMode() {
      return hasBestSequence;
    },
  };
}

function connectVisualizationControls(controller) {
  const speedButtons = Array.from(document.querySelectorAll(".speed-button[data-speed]"));
  const pauseButton = document.getElementById("playback-pause-button");
  const restartButton = document.getElementById("playback-restart-button");
  const modeElement = document.getElementById("visualization-mode-label");
  const bestSequenceToggle = document.getElementById("best-sequence-toggle");
  const bestSequenceHint = document.getElementById("best-sequence-toggle-hint");

  const refreshControls = () => {
    const enabled = Boolean(controller);
    const activeSpeed = enabled ? controller.getSpeed() : DEFAULT_SPEED;
    const paused = enabled ? controller.isPaused() : false;
    const mode = enabled ? controller.getMode() : "none";
    const canToggleBestMode = enabled && controller.canToggleBestMode();

    updateSpeedButtons(speedButtons, activeSpeed, enabled);

    if (pauseButton) {
      pauseButton.disabled = !enabled;
      pauseButton.textContent = paused ? "Retomar" : "Pausar";
      pauseButton.classList.toggle("is-active", enabled && paused);
    }

    if (restartButton) {
      restartButton.disabled = !enabled;
    }

    if (modeElement) {
      modeElement.textContent = enabled ? modeLabel(mode) : "Indisponivel";
    }

    if (bestSequenceToggle) {
      const bestModeActive = canToggleBestMode && mode === "best-sequence";
      bestSequenceToggle.disabled = !canToggleBestMode;
      bestSequenceToggle.classList.toggle("is-active", bestModeActive);
      bestSequenceToggle.setAttribute("aria-pressed", String(bestModeActive));
    }

    if (bestSequenceHint) {
      if (!canToggleBestMode) {
        bestSequenceHint.textContent = "Sem best sequence disponivel para animacao.";
      } else if (mode === "best-sequence") {
        bestSequenceHint.textContent = "Clique para voltar para a visualizacao da evolucao.";
      } else {
        bestSequenceHint.textContent = "Clique para animar o percurso da best sequence.";
      }
    }
  };

  speedButtons.forEach((button) => {
    button.addEventListener("click", () => {
      if (!controller) {
        return;
      }

      const nextSpeed = Number(button.dataset.speed);
      if (!ALLOWED_SPEEDS.has(nextSpeed)) {
        return;
      }

      controller.setSpeed(nextSpeed);
      refreshControls();
    });
  });

  if (pauseButton) {
    pauseButton.addEventListener("click", () => {
      if (!controller) {
        return;
      }

      controller.togglePause();
      refreshControls();
    });
  }

  if (restartButton) {
    restartButton.addEventListener("click", () => {
      if (!controller) {
        return;
      }

      controller.restart();
      refreshControls();
    });
  }

  if (bestSequenceToggle) {
    bestSequenceToggle.addEventListener("click", () => {
      if (!controller) {
        return;
      }

      controller.toggleMode();
      refreshControls();
    });
  }

  refreshControls();
}

function stopActiveVisualizationController() {
  if (!activeVisualizationController) {
    return;
  }

  activeVisualizationController.stop();
  activeVisualizationController = null;
}

function renderMainContent() {
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

document.addEventListener("DOMContentLoaded", renderMainContent);

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
