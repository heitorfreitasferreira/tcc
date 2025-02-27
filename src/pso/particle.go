package pso

import (
	"math"
	"sort"
)

type particle struct {
	x            []float64
	v            []float64
	bestX        []float64
	sequence     []int
	clusterId    int
	makespan     float64
	bestMakespan float64
}

func (p *particle) setSequence() {
	indices := make([]int, len(p.x))
	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return p.x[indices[i]] < p.x[indices[j]]
	})

	// Cria um array para os resultados
	seq := make([]int, len(p.x))
	for i, index := range indices {
		seq[index] = i
	}

	p.sequence = seq
}

func newParticle(dimention int) particle {
	return particle{
		x:            make([]float64, dimention),
		v:            make([]float64, dimention),
		bestX:        make([]float64, dimention),
		bestMakespan: math.MaxFloat64,
	}
}

func (p *particle) update(w, c1, c2, r1, r2 float64, gBest *particle) {
	for i := range p.x {
		// Inércia
		p.v[i] *= w
		// Cognitivo
		p.v[i] += c1 * r1 * (p.bestX[i] - p.x[i])
		// Social
		p.v[i] += c2 * r2 * (gBest.x[i] - p.x[i])
		// Att a posição
		p.x[i] += p.v[i]
	}
	p.setSequence()
}
