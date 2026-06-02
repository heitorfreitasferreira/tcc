package handlers

import (
	"image/png"
	"net/http"
	"strconv"
	"strings"

	"tcc/web/render"
	"tcc/web/repository"
)

type RenderHandler struct {
	mapsRepo      repository.MapsReader
	evolutionRepo repository.EvolutionReader
}

func NewRenderHandler(mapsRepo repository.MapsReader, evolutionRepo repository.EvolutionReader) *RenderHandler {
	return &RenderHandler{mapsRepo: mapsRepo, evolutionRepo: evolutionRepo}
}

func (h *RenderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	mapID := strings.TrimSpace(r.URL.Query().Get("map"))
	runIDs := renderRunIDs(r)
	iterationStr := strings.TrimSpace(r.URL.Query().Get("iteration"))

	if mapID == "" {
		http.Error(w, "missing required param: map", http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	points, err := h.mapsRepo.LoadPoints(ctx, mapID)
	if err != nil {
		http.Error(w, "load points: "+err.Error(), http.StatusInternalServerError)
		return
	}

	routeSpecs := make([]render.RouteSpec, 0, len(runIDs))

	for _, runID := range runIDs {
		rawFrames, err := h.evolutionRepo.LoadEvolution(ctx, runID)
		if err != nil {
			http.Error(w, "load evolution: "+err.Error(), http.StatusInternalServerError)
			return
		}

		frames := make([]render.EvolutionFrame, len(rawFrames))
		for i, f := range rawFrames {
			frames[i] = render.EvolutionFrame{
				Iter:         f.Iter,
				EvalCount:    f.EvalCount,
				BestMakespan: f.BestMakespan,
				BestSequence: f.BestSequence,
			}
		}

		var sequence []int
		var makespan float64
		if iterationStr != "" {
			targetIter, err := strconv.Atoi(iterationStr)
			if err == nil && targetIter >= 0 {
				seq, frameMakespan, ok := render.FindFrameAtOrBefore(frames, targetIter)
				if ok {
					sequence = seq
					makespan = frameMakespan
				}
			}
		}

		if len(sequence) == 0 {
			seq, frameMakespan, ok := render.FindBestFrame(frames)
			if ok {
				sequence = seq
				makespan = frameMakespan
			}
		}

		if len(sequence) > 0 {
			routeSpecs = append(routeSpecs, render.RouteSpec{
				Name:     renderRunName(runID),
				Sequence: sequence,
				Makespan: makespan,
				Color:    renderRunColor(runID),
			})
		}
	}

	pts := make([][2]float64, len(points))
	for i, p := range points {
		pts[i] = [2]float64(p)
	}

	opts := render.DefaultOpts()
	if r.URL.Query().Get("angular") == "1" || r.URL.Query().Get("angular") == "true" {
		opts.AngularPenalty = true
	}
	routes := render.BuildNamedRoutes(len(pts), routeSpecs)
	img, err := render.RenderRoutes(pts, routes, opts)
	if err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-cache")
	png.Encode(w, img)
}

func renderRunIDs(r *http.Request) []string {
	keys := []string{"run", "run2", "run3", "run4"}
	runs := make([]string, 0, len(keys))
	seen := make(map[string]struct{})
	for _, key := range keys {
		runID := strings.TrimSpace(r.URL.Query().Get(key))
		if runID == "" {
			continue
		}
		if _, ok := seen[runID]; ok {
			continue
		}
		seen[runID] = struct{}{}
		runs = append(runs, runID)
	}
	return runs
}

func renderRunName(runID string) string {
	parts := strings.Split(runID, "__")
	if len(parts) >= 2 {
		if parts[1] == "bruteforce" {
			return "BF"
		}
		if parts[1] == "lowerbound" {
			return "LB"
		}
		return strings.ToUpper(parts[1])
	}
	return runID
}

func renderRunColor(runID string) render.RGB {
	parts := strings.Split(runID, "__")
	if len(parts) < 2 {
		return render.RGB{}
	}
	switch parts[1] {
	case "aco":
		return render.RGB{R: 0.878, G: 0.478, B: 0.063}
	case "ga":
		return render.RGB{R: 0.039, G: 0.4, B: 0.761}
	case "pso":
		return render.RGB{R: 0.125, G: 0.569, B: 0.188}
	case "lowerbound":
		return render.RGB{R: 0.482, G: 0.176, B: 0.557}
	case "bruteforce":
		return render.RGB{R: 0.2, G: 0.2, B: 0.2}
	default:
		return render.RGB{}
	}
}
