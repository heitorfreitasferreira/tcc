package lowerbound

import "math"

func hungarian(cost [][]float64) ([]int, float64) {
	n := len(cost)
	if n == 0 {
		return []int{}, 0
	}

	m := n
	u := make([]float64, n+1)
	v := make([]float64, m+1)
	p := make([]int, m+1)
	way := make([]int, m+1)

	for i := 1; i <= n; i++ {
		p[0] = i
		j0 := 0
		minv := make([]float64, m+1)
		for j := range minv {
			minv[j] = math.Inf(1)
		}
		used := make([]bool, m+1)

		for {
			used[j0] = true
			i0 := p[j0]
			delta := math.Inf(1)
			j1 := 0

			for j := 1; j <= m; j++ {
				if !used[j] {
					cur := cost[i0-1][j-1] - u[i0] - v[j]
					if cur < minv[j] {
						minv[j] = cur
						way[j] = j0
					}
					if minv[j] < delta {
						delta = minv[j]
						j1 = j
					}
				}
			}

			for j := 0; j <= m; j++ {
				if used[j] {
					u[p[j]] += delta
					v[j] -= delta
				} else {
					minv[j] -= delta
				}
			}

			j0 = j1
			if p[j0] == 0 {
				break
			}
		}

		for {
			j1 := way[j0]
			p[j0] = p[j1]
			j0 = j1
			if j0 == 0 {
				break
			}
		}
	}

	assign := make([]int, n)
	for j := 1; j <= m; j++ {
		if p[j] > 0 {
			assign[p[j]-1] = j - 1
		}
	}

	return assign, -v[0]
}
