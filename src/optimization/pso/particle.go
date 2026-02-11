package pso

import (
	"fmt"
	"sort"
)

type particle struct {
	x            []float64
	v            []float64
	bestX        []float64
	sequence     []int
	makespan     float64
	bestMakespan float64
}

func (p particle) String() string {
	return fmt.Sprintf("Particle{x: %v, v: %v, sequence: %v, makespan: %.2f, bestMakespan: %.2f}",
		p.x, p.v, p.sequence, p.makespan, p.bestMakespan)
}

func (p *particle) setSequence() {
	indices := make([]int, len(p.x))
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return p.x[indices[i]] < p.x[indices[j]]
	})

	// Converte random-keys para rota (nós 1..n-1).
	seq := make([]int, len(p.x))
	for i := range indices {
		seq[i] = indices[i] + 1
	}

	copy(p.sequence, seq)
}

func (p *particle) update(w, c1, c2, r1, r2 float64, gBest []float64) {
	for i := range p.x {
		// Inércia
		p.v[i] *= w
		// Cognitivo
		p.v[i] += c1 * r1 * (p.bestX[i] - p.x[i])
		// Social
		p.v[i] += c2 * r2 * (gBest[i] - p.x[i])
		// Att a posição
		p.x[i] += p.v[i]
	}
	p.setSequence()
}
