package shared

import (
	"math/rand"
)

func RandomizeSlice[T ~int | ~int32 | ~int64 | ~float32 | ~float64](slice []T, rnd *rand.Rand) {
	for i := range slice {
		switch any(slice[i]).(type) {
		case int:
			slice[i] = T(rnd.Int())
		case int32:
			slice[i] = T(rnd.Int31())
		case int64:
			slice[i] = T(rnd.Int63())
		case float32:
			slice[i] = T(rnd.Float32())
		case float64:
			slice[i] = T(rnd.Float64())
		}
	}
}

func Shuffle[T any](slice []T, rnd *rand.Rand) {
	for i := len(slice) - 1; i > 0; i-- {
		j := rnd.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
