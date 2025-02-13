package aco

func finished(tour []bool) bool {
	for _, visited := range tour {
		if !visited {
			return false
		}
	}
	return true
}

func cumSum(probs []float64) []float64 {
	cumSum := make([]float64, len(probs))

	cumSum[0] = probs[0]
	for i := 1; i < len(probs); i++ {
		cumSum[i] = cumSum[i-1] + probs[i]
	}
	return cumSum
}
