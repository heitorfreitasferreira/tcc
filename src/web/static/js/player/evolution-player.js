import {
  ALLOWED_SPEEDS,
  BASE_ITERATION_DELAY_MS,
  DEFAULT_SPEED,
  MIN_FRAME_DELAY_MS,
} from "../constants.js";
import { drawMapFrame } from "../render/map-renderer.js";
import { updateFrameMetadata } from "../ui/metadata.js";

export function createEvolutionPlayer({ canvas, points, frames, iterations }) {
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
