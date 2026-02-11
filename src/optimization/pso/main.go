package pso

import (
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

func (sw *Swarm) Optimize(onImprovement func(shared.Improvement)) shared.OptimizationResult {
	// initialize
	for i := 0; i < sw.PopulationSize; i++ {
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

	result := shared.OptimizationResult{
		BestMakespan: math.MaxFloat64,
	}
	evaluationCount := 0

	// run
	sw.evaluate(0, &result, &evaluationCount, onImprovement)

	for iteration := 1; iteration <= sw.Iterations; iteration++ {
		sw.update()
		sw.evaluate(iteration, &result, &evaluationCount, onImprovement)
	}

	result.IterationsCompleted = sw.Iterations
	result.Evaluations = evaluationCount
	result.BestMakespan = sw.bestMakespan
	result.BestSequence = append([]int(nil), sw.gBestSeq...)

	return result
}

func (sw *Swarm) update() {
	for i := range sw.particles {
		p := &sw.particles[i]
		r1 := sw.rng.Float64()
		r2 := sw.rng.Float64()
		p.update(sw.W, sw.C1, sw.C2, r1, r2, sw.gBestPos)
	}
}

func (sw *Swarm) evaluate(
	iteration int,
	result *shared.OptimizationResult,
	evaluationCount *int,
	onImprovement func(shared.Improvement),
) {
	for i := range sw.particles {
		p := &sw.particles[i]
		p.makespan = sw.Makespan(p.sequence)
		*evaluationCount = *evaluationCount + 1
		if p.makespan < p.bestMakespan {
			p.bestMakespan = p.makespan
			copy(p.bestX, p.x)
			if p.bestMakespan < sw.bestMakespan {
				delta := 0.0
				if sw.bestMakespan < math.MaxFloat64 {
					delta = p.bestMakespan - sw.bestMakespan
				}
				sw.bestMakespan = p.bestMakespan
				copy(sw.gBestPos, p.bestX)
				copy(sw.gBestSeq, p.sequence)

				improvement := shared.Improvement{
					Iteration:    iteration,
					Evaluation:   *evaluationCount,
					BestMakespan: sw.bestMakespan,
					Delta:        delta,
					BestSequence: append([]int(nil), sw.gBestSeq...),
				}
				result.Improvements = append(result.Improvements, improvement)
				if onImprovement != nil {
					onImprovement(improvement)
				}
			}
		}
	}
}
