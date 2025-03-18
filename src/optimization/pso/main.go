package pso

import (
	"fmt"
	"math"
	"math/rand"
	"tcc/graph"
	"tcc/shared"
)

type Params struct {
	shared.HyperParams
	C1, C2, W float64
}

type Swarm struct {
	*graph.Graph
	Params

	particles []particle

	gBestPos     []float64
	gBestSeq     []int
	bestMakespan float64

	rng *rand.Rand
}

func (sw *Swarm) Init(g *graph.Graph, p Params, rng *rand.Rand) {
	sw.Graph = g
	sw.Params = p
	sw.rng = rng
	sw.bestMakespan = math.MaxFloat64
	sw.gBestPos = make([]float64, g.NumberOfNodes())
	sw.gBestSeq = make([]int, g.NumberOfNodes())
	sw.particles = make([]particle, 0, p.PopulationSize)
}

func (sw *Swarm) Optimize() *shared.SwarmStats {
	// initialize
	for range sw.PopulationSize {
		x := make([]float64, sw.NumberOfNodes())
		shared.RandomizeSlice(x, sw.rng)
		newParticle := particle{
			x:            x,
			v:            make([]float64, sw.NumberOfNodes()),
			bestX:        make([]float64, sw.NumberOfNodes()),
			sequence:     make([]int, sw.NumberOfNodes()),
			bestMakespan: math.MaxFloat64,
		}

		newParticle.setSequence()

		sw.particles = append(sw.particles, newParticle)
	}

	// run
	sw.evaluate()

	stats := shared.NewStats()
	stats.AddIterData(sw.bestMakespan, sw.gBestPos, sw.gBestSeq)
	for range sw.Iterations {
		sw.update()
		sw.evaluate()
		stats.AddIterData(sw.bestMakespan, sw.gBestPos, sw.gBestSeq)
	}

	return stats
}

func (sw *Swarm) update() {
	for _, p := range sw.particles {
		r1 := sw.rng.Float64()
		r2 := sw.rng.Float64()
		p.update(sw.W, sw.C1, sw.C2, r1, r2, sw.gBestPos)
	}
}

func (sw *Swarm) evaluate() {
	for _, p := range sw.particles {
		p.makespan = sw.Makespan(p.sequence)
		if p.makespan < p.bestMakespan {
			p.bestMakespan = p.makespan
			copy(p.bestX, p.x)
			if p.bestMakespan < sw.bestMakespan {
				fmt.Println("Changing from", sw.bestMakespan, "to", p.bestMakespan)
				sw.bestMakespan = p.bestMakespan
				copy(sw.gBestPos, p.bestX)
				copy(sw.gBestSeq, p.sequence)
			}
		}
	}
}
