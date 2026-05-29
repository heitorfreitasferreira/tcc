import { ALLOWED_SPEEDS, DEFAULT_SPEED } from "../constants.js";
import { modeLabel } from "../domain/evolution.js";

function updateSpeedButtons(buttons, activeSpeed, enabled) {
  buttons.forEach((button) => {
    const buttonSpeed = Number(button.dataset.speed);
    const supported = ALLOWED_SPEEDS.has(buttonSpeed);

    button.disabled = !enabled || !supported;
    button.classList.toggle("is-active", enabled && supported && buttonSpeed === activeSpeed);
  });
}

export function connectVisualizationControls(controller) {
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
