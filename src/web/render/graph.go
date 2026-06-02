package render

import (
	"image"
	"math"
	"strconv"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/basicfont"
)

const (
	worldMin   = -1.0
	worldMax   = 1.0
	worldRange = worldMax - worldMin
)

type Opts struct {
	Width, Height int
	Padding       float64
}

func DefaultOpts() Opts {
	return Opts{Width: 800, Height: 800, Padding: 40}
}

func BuildRouteIndices(pointsCount int, sequence []int) []int {
	if pointsCount <= 0 || len(sequence) == 0 {
		return nil
	}

	normalized := make([]int, 0, len(sequence))
	for _, v := range sequence {
		if v >= 0 && v < pointsCount {
			normalized = append(normalized, v)
		}
	}
	if len(normalized) == 0 {
		return nil
	}

	hasDepot := false
	for _, v := range normalized {
		if v == 0 {
			hasDepot = true
			break
		}
	}

	route := make([]int, 0, len(normalized)+2)
	if !hasDepot {
		route = append(route, 0)
	}
	route = append(route, normalized...)

	if route[len(route)-1] != route[0] {
		route = append(route, route[0])
	}

	return route
}

type EvolutionFrame struct {
	Iter         int
	EvalCount    int
	BestMakespan float64
	BestSequence []int
}

func FindBestFrame(frames []EvolutionFrame) ([]int, bool) {
	if len(frames) == 0 {
		return nil, false
	}

	var bestSeq []int
	bestVal := math.Inf(1)
	for _, f := range frames {
		if len(f.BestSequence) == 0 {
			continue
		}
		if f.BestMakespan < bestVal {
			bestVal = f.BestMakespan
			bestSeq = f.BestSequence
		}
	}
	return bestSeq, bestSeq != nil
}

func FindFrameAtOrBefore(frames []EvolutionFrame, targetIter int) ([]int, bool) {
	if len(frames) == 0 {
		return nil, false
	}

	bestIdx := -1
	for i, f := range frames {
		if f.Iter <= targetIter && len(f.BestSequence) > 0 {
			bestIdx = i
		} else if f.Iter > targetIter {
			break
		}
	}
	if bestIdx < 0 {
		return nil, false
	}
	return frames[bestIdx].BestSequence, true
}

type point2D = [2]float64

func Render(points []point2D, sequence []int, opts Opts) (image.Image, error) {
	dc := gg.NewContext(opts.Width, opts.Height)
	dc.SetFontFace(basicfont.Face7x13)

	plotSize := math.Min(
		float64(opts.Width)-2*opts.Padding,
		float64(opts.Height)-2*opts.Padding,
	)
	plotLeft := (float64(opts.Width) - plotSize) / 2
	plotTop := (float64(opts.Height) - plotSize) / 2

	toX := func(world float64) float64 {
		return plotLeft + ((world-worldMin)/worldRange)*plotSize
	}
	toY := func(world float64) float64 {
		return plotTop + ((worldMax-world)/worldRange)*plotSize
	}

	drawRoute := func() {
		route := BuildRouteIndices(len(points), sequence)
		if len(route) < 2 {
			return
		}

		type edgeUse struct{ count int }
		edgeUsage := make(map[[2]int]*edgeUse)

		for i := 1; i < len(route); i++ {
			start := route[i-1]
			end := route[i]
			key := [2]int{start, end}
			if _, ok := edgeUsage[key]; !ok {
				edgeUsage[key] = &edgeUse{}
			}
			edgeUsage[key].count++
		}

		for i := 1; i < len(route); i++ {
			start := route[i-1]
			end := route[i]
			sp := points[start]
			ep := points[end]

			x1, y1 := toX(sp[0]), toY(sp[1])
			x2, y2 := toX(ep[0]), toY(ep[1])

			dx := x2 - x1
			dy := y2 - y1
			segLen := math.Hypot(dx, dy)
			if segLen < 1e-6 {
				continue
			}

			ux := dx / segLen
			uy := dy / segLen
			nx := -uy
			ny := ux

			key := [2]int{start, end}
			occ := edgeUsage[key].count
			remaining := occ
			edgeUsage[key].count--
			if remaining <= 1 {
				remaining = 1
			}

			jitter := 3.0
			if remaining > 0 {
				lane := math.Ceil(float64(remaining) / 2)
				signal := 1.0
				if remaining%2 == 0 {
					signal = -1.0
				}
				jitter = 3.0 + signal*lane*0.9
			}

			shx := x1 + nx*jitter
			shy := y1 + ny*jitter
			ehx := x2 + nx*jitter
			ehy := y2 + ny*jitter

			pointRadius := 4.0
			pointGap := 2.0
			endOffset := pointRadius + pointGap
			arrowLen := math.Max(8, math.Min(12, segLen*0.28))
			arrowHalf := arrowLen * 0.5

			saX := shx + ux*endOffset
			saY := shy + uy*endOffset
			tipX := ehx - ux*endOffset
			tipY := ehy - uy*endOffset
			shEndX := tipX - ux*arrowLen
			shEndY := tipY - uy*arrowLen

			dc.SetRGB(0.878, 0.478, 0.063)
			dc.SetLineWidth(2.2)

			shaftLen := math.Hypot(shEndX-saX, shEndY-saY)
			if shaftLen >= 1 {
				dc.DrawLine(saX, saY, shEndX, shEndY)
				dc.Stroke()
			}

			lx := shEndX + nx*arrowHalf
			ly := shEndY + ny*arrowHalf
			rx := shEndX - nx*arrowHalf
			ry := shEndY - ny*arrowHalf

			dc.MoveTo(tipX, tipY)
			dc.LineTo(lx, ly)
			dc.LineTo(rx, ry)
			dc.ClosePath()
			dc.Fill()
		}
	}

	drawGrid := func() {
		step := 0.25
		for w := worldMin; w <= worldMax+1e-9; w += step {
			if math.Abs(w) < 1e-9 {
				continue
			}
			x := toX(w)
			dc.SetRGB(0.902, 0.902, 0.902)
			dc.SetLineWidth(1)
			dc.DrawLine(x, plotTop, x, plotTop+plotSize)
			dc.Stroke()

			y := toY(w)
			dc.DrawLine(plotLeft, y, plotLeft+plotSize, y)
			dc.Stroke()
		}

		dc.SetRGB(0.722, 0.722, 0.722)
		dc.SetLineWidth(1.2)
		dc.DrawRectangle(plotLeft, plotTop, plotSize, plotSize)
		dc.Stroke()

		axisX := toX(0)
		axisY := toY(0)
		dc.SetRGB(0.4, 0.4, 0.4)
		dc.SetLineWidth(1.6)
		dc.DrawLine(plotLeft, axisY, plotLeft+plotSize, axisY)
		dc.Stroke()
		dc.DrawLine(axisX, plotTop, axisX, plotTop+plotSize)
		dc.Stroke()
	}

	drawPoints := func() {
		for i, p := range points {
			x, y := toX(p[0]), toY(p[1])
			if i == 0 {
				dc.SetRGB(0.839, 0.157, 0.157)
			} else {
				dc.SetRGB(0.039, 0.4, 0.761)
			}
			dc.DrawCircle(x, y, 4)
			dc.Fill()

			dc.SetRGB(0.188, 0.188, 0.188)
			label := strconv.Itoa(i)
			dc.DrawString(label, x+7, y-7)
		}
	}

	dc.SetRGB(1, 1, 1)
	dc.Clear()

	drawGrid()

	if len(sequence) > 0 {
		drawRoute()
	}

	drawPoints()

	return dc.Image(), nil
}
