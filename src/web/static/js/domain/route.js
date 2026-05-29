export function buildRouteIndices(pointsCount, sequence) {
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

export function interpolateRoutePoint(points, route, segmentIndex, segmentProgress) {
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
