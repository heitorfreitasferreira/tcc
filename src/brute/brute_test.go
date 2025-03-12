package brute

import (
	"reflect"
	"testing"
)

func TestPermutation(t *testing.T) {
	testCases := []struct {
		desc  string
		input []int
		out   [][]int
	}{
		{
			desc:  "Basic",
			input: []int{1, 2},
			out: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			desc:  "Three elements",
			input: []int{1, 2, 3},
			out: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			desc:  "Single element",
			input: []int{1},
			out: [][]int{
				{1},
			},
		},
		{
			desc:  "Empty input",
			input: []int{},
			out:   [][]int{},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			perms := generatePermutations(tC.input)
			if !containsAll(perms, tC.out) {
				t.Errorf("expected %v, got %v", tC.out, perms)
			}
		})
	}
}

func containsAll(perms, expected [][]int) bool {
	for _, exp := range expected {
		found := false
		for _, perm := range perms {
			if reflect.DeepEqual(perm, exp) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
