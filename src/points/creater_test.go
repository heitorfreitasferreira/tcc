package points

import (
	"reflect"
	"testing"
)

func TestCreateInstancesIsDeterministicForSameSeed(t *testing.T) {
	frequency := map[int]int{14: 3, 10: 3, 12: 3, 11: 3, 13: 3}

	expected := CreateInstances(42, frequency)
	for i := 0; i < 20; i++ {
		got := CreateInstances(42, frequency)
		if !reflect.DeepEqual(expected, got) {
			t.Fatalf("expected deterministic instances for same seed")
		}
	}
}
