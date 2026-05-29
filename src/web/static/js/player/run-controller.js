import { ALLOWED_SPEEDS, DEFAULT_SPEED } from "../constants.js";
import { findBestSequenceFrame } from "../domain/evolution.js";
import { createBestSequenceTraversalPlayer } from "./best-sequence-player.js";
import { createEvolutionPlayer } from "./evolution-player.js";

export function createRunVisualizationController({ canvas, points, evolution }) {
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
