export function findBestSequenceFrame(frames) {
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

export function modeLabel(mode) {
  return mode === "best-sequence" ? "Best Sequence" : "Evolução";
}

export function formatBestSequence(sequence) {
  if (!Array.isArray(sequence) || sequence.length === 0) {
    return "-";
  }

  return sequence.map((value) => String(value)).join(" -> ");
}
