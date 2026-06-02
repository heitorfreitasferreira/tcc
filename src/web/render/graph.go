package render

import (
	"image"
	"math"
	"strconv"
	"strings"

	"github.com/fogleman/gg"
	"golang.org/x/image/font/basicfont"
)

const (
	worldMin   = -1.0
	worldMax   = 1.0
	worldRange = worldMax - worldMin
)

type Opts struct {
	Width, Height    int
	Padding          float64
	ScaleStep        float64
	ScaleLabel       string
	AngularPenalty   bool
}

type RGB struct {
	R float64
	G float64
	B float64
}

type RouteSpec struct {
	Name     string
	Sequence []int
	Makespan float64
	Color    RGB
}

type Route struct {
	Name     string
	Indices  []int
	Makespan float64
	Color    RGB
}

var defaultRouteNames = []string{"Rota 1", "Rota 2", "Rota 3", "Rota 4"}
var defaultRouteColors = []RGB{
	{R: 0.878, G: 0.478, B: 0.063},
	{R: 0.039, G: 0.4, B: 0.761},
	{R: 0.125, G: 0.569, B: 0.188},
	{R: 0.2, G: 0.2, B: 0.2},
}

func DefaultOpts() Opts {
	return Opts{Width: 800, Height: 800, Padding: 52, ScaleStep: 0.5, ScaleLabel: "unid. coord."}
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

func BuildRoutes(pointsCount int, sequences [][]int) []Route {
	specs := make([]RouteSpec, len(sequences))
	for i, sequence := range sequences {
		specs[i] = RouteSpec{Sequence: sequence}
	}
	return BuildNamedRoutes(pointsCount, specs)
}

func BuildNamedRoutes(pointsCount int, specs []RouteSpec) []Route {
	routes := make([]Route, 0, len(specs))
	for i, spec := range specs {
		indices := BuildRouteIndices(pointsCount, spec.Sequence)
		if len(indices) < 2 {
			continue
		}

		name := strings.TrimSpace(spec.Name)
		if name == "" {
			name = defaultRouteName(i)
		}

		color := spec.Color
		if color == (RGB{}) {
			color = defaultRouteColor(i)
		}

		routes = append(routes, Route{
			Name:     name,
			Indices:  indices,
			Makespan: spec.Makespan,
			Color:    color,
		})
	}
	return routes
}

func defaultRouteName(i int) string {
	if i >= 0 && i < len(defaultRouteNames) {
		return defaultRouteNames[i]
	}
	return "Rota " + strconv.Itoa(i+1)
}

func defaultRouteColor(i int) RGB {
	if i >= 0 && i < len(defaultRouteColors) {
		return defaultRouteColors[i]
	}
	return RGB{R: 0.878, G: 0.478, B: 0.063}
}

type EvolutionFrame struct {
	Iter         int
	EvalCount    int
	BestMakespan float64
	BestSequence []int
}

func FindBestFrame(frames []EvolutionFrame) ([]int, float64, bool) {
	if len(frames) == 0 {
		return nil, 0, false
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
	return bestSeq, bestVal, bestSeq != nil
}

func FindFrameAtOrBefore(frames []EvolutionFrame, targetIter int) ([]int, float64, bool) {
	if len(frames) == 0 {
		return nil, 0, false
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
		return nil, 0, false
	}
	return frames[bestIdx].BestSequence, frames[bestIdx].BestMakespan, true
}

type point2D = [2]float64

func penaltyColor(t float64) (float64, float64, float64) {
	if t < 0.5 {
		u := t / 0.5
		return 0.898 * u, 0.702 - 0.082*u, 0.102 * u
	}
	u := (t - 0.5) / 0.5
	return 0.898 - 0.2*u, 0.620 - 0.520*u, 0.102 - 0.102*u
}

func Render(points []point2D, sequence []int, opts Opts) (image.Image, error) {
	routes := BuildRoutes(len(points), [][]int{sequence})
	return RenderRoutes(points, routes, opts)
}

func RenderRoutes(points []point2D, routes []Route, opts Opts) (image.Image, error) {
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

	drawRoute := func(route Route, routeIndex int) {
		type edgeUse struct{ count int }
		edgeUsage := make(map[[2]int]*edgeUse)

		for i := 1; i < len(route.Indices); i++ {
			start := route.Indices[i-1]
			end := route.Indices[i]
			key := [2]int{start, end}
			if _, ok := edgeUsage[key]; !ok {
				edgeUsage[key] = &edgeUse{}
			}
			edgeUsage[key].count++
		}

		for i := 1; i < len(route.Indices); i++ {
			start := route.Indices[i-1]
			end := route.Indices[i]
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

			jitter := 3.0 + float64(routeIndex)*2.0
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

			if opts.AngularPenalty && len(route.Indices) > 2 {
				next := route.Indices[0]
				if i+1 < len(route.Indices) {
					next = route.Indices[i+1]
				}
				ax, ay := points[start][0], points[start][1]
				bx, by := points[end][0], points[end][1]
				cx, cy := points[next][0], points[next][1]
				v1x, v1y := bx-ax, by-ay
				v2x, v2y := cx-bx, cy-by
				l1 := math.Hypot(v1x, v1y)
				l2 := math.Hypot(v2x, v2y)
				if l1 > 1e-8 && l2 > 1e-8 {
					cosT := (v1x*v2x + v1y*v2y) / (l1 * l2)
					if cosT > 1 {
						cosT = 1
					}
					if cosT < -1 {
						cosT = -1
					}
					theta := math.Acos(cosT)
					penalty := theta / math.Pi
					r, g, b := penaltyColor(penalty)
					dc.SetRGB(r, g, b)
				} else {
					dc.SetRGB(route.Color.R, route.Color.G, route.Color.B)
				}
			} else {
				dc.SetRGB(route.Color.R, route.Color.G, route.Color.B)
			}
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

	drawLegend := func() {
		if len(routes) < 2 {
			return
		}

		x := plotLeft + 8
		y := plotTop + 14
		for _, route := range routes {
			dc.SetRGB(route.Color.R, route.Color.G, route.Color.B)
			dc.SetLineWidth(3)
			dc.DrawLine(x, y-4, x+18, y-4)
			dc.Stroke()

			dc.SetRGB(0.188, 0.188, 0.188)
			label := route.Name
			if route.Makespan > 0 {
				label += " (" + strconv.FormatFloat(route.Makespan, 'f', 2, 64) + ")"
			}
			dc.DrawString(label, x+24, y)
			y += 16
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

		dc.SetRGB(0.25, 0.25, 0.25)
		dc.SetLineWidth(1)
		for w := worldMin; w <= worldMax+1e-9; w += opts.ScaleStep {
			x := toX(w)
			y := toY(w)
			label := strconv.FormatFloat(w, 'f', 1, 64)
			dc.DrawLine(x, plotTop+plotSize, x, plotTop+plotSize+5)
			dc.Stroke()
			dc.DrawStringAnchored(label, x, plotTop+plotSize+18, 0.5, 0.5)

			dc.DrawLine(plotLeft-5, y, plotLeft, y)
			dc.Stroke()
			dc.DrawStringAnchored(label, plotLeft-18, y, 0.5, 0.5)
		}

		dc.DrawStringAnchored("x", plotLeft+plotSize+16, axisY+5, 0.5, 0.5)
		dc.DrawStringAnchored("y", axisX, plotTop-14, 0.5, 0.5)
	}

	drawScaleBar := func() {
		if opts.ScaleStep <= 0 {
			return
		}
		barWorld := opts.ScaleStep
		barPx := barWorld / worldRange * plotSize
		x := plotLeft + plotSize - barPx - 12
		y := plotTop + plotSize - 18
		dc.SetRGB(0.18, 0.18, 0.18)
		dc.SetLineWidth(2)
		dc.DrawLine(x, y, x+barPx, y)
		dc.Stroke()
		dc.DrawLine(x, y-5, x, y+5)
		dc.Stroke()
		dc.DrawLine(x+barPx, y-5, x+barPx, y+5)
		dc.Stroke()
		label := strconv.FormatFloat(barWorld, 'f', 1, 64) + " " + opts.ScaleLabel
		dc.DrawStringAnchored(label, x+barPx/2, y-10, 0.5, 0.5)
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
	drawScaleBar()

	for i, route := range routes {
		if len(route.Indices) > 1 {
			drawRoute(route, i)
		}
	}

	drawPoints()
	drawLegend()

	return dc.Image(), nil
}
