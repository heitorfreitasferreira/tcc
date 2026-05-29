import { formatBestSequence } from "../domain/evolution.js";

export function updateFrameMetadata(frame, totalIterations) {
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
