import {
  ALLOWED_SPEEDS,
  BEST_SEQUENCE_EDGE_DURATION_MS,
  DEFAULT_SPEED,
} from "../constants.js";
import { interpolateRoutePoint, buildRouteIndices } from "../domain/route.js";
import { createPlotGeometry } from "../render/geometry.js";
import { drawMapState, drawTraversalMarker } from "../render/map-renderer.js";
import { updateFrameMetadata } from "../ui/metadata.js";

export function createBestSequenceTraversalPlayer({ canvas, points, frame, iterations }) {
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
