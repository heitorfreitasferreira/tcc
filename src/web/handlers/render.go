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
	runID := strings.TrimSpace(r.URL.Query().Get("run"))
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

	var sequence []int

	if runID != "" {
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

		if iterationStr != "" {
			targetIter, err := strconv.Atoi(iterationStr)
			if err == nil && targetIter >= 0 {
				seq, ok := render.FindFrameAtOrBefore(frames, targetIter)
				if ok {
					sequence = seq
				}
			}
		}

		if len(sequence) == 0 {
			seq, ok := render.FindBestFrame(frames)
			if ok {
				sequence = seq
			}
		}
	}

	pts := make([][2]float64, len(points))
	for i, p := range points {
		pts[i] = [2]float64(p)
	}

	opts := render.DefaultOpts()
	img, err := render.Render(pts, sequence, opts)
	if err != nil {
		http.Error(w, "render: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Cache-Control", "no-cache")
	png.Encode(w, img)
}
