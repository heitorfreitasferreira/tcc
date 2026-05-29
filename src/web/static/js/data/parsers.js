export function toNonNegativeInt(value, fallback) {
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

export function parsePoints(rawPoints) {
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

export function parseEvolution(rawEvolution) {
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
