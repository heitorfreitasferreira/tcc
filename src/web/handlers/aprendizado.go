package handlers

import (
	"html/template"
	"net/http"
	"strings"
)

// AprendizadoPageData carries the data needed to render an aprendizado-visual page.
type AprendizadoPageData struct {
	Title       string
	Description string
	Slug        string
	IsIndex     bool
	Methods     []AprendizadoMethod
}

// AprendizadoMethod describes one metaheuristic for the listing.
type AprendizadoMethod struct {
	Name        string
	Slug        string
	Description string
}

var aprendizadoMethods = []AprendizadoMethod{
	{
		Name:        "Busca Exaustiva",
		Slug:        "bruteforce",
		Description: "Enumeração completa de todas as permutações possíveis. Garante o ótimo global, mas a complexidade fatorial O(n!) torna-se inviável para instâncias com mais de 10–12 pontos de interesse.",
	},
	{
		Name:        "Algoritmo Genético (GA)",
		Slug:        "ga",
		Description: "Inspirado na evolução natural: seleção por torneio, crossover OX (Crossover por Ordem) e mutação swap atuam sobre uma população de indivíduos a cada geração, com elitismo preservando o melhor.",
	},
	{
		Name:        "Otimização por Colônia de Formigas (ACO)",
		Slug:        "aco",
		Description: "Formigas artificiais constroem rotas TSP incrementalmente, escolhendo a próxima cidade por roleta proporcional a τ^α · η^β. O feromônio nas arestas evapora e é depositado pelas formigas a cada iteração, guiando a busca.",
	},
	{
		Name:        "Otimização por Enxame de Partículas (PSO)",
		Slug:        "pso",
		Description: "Partículas movem-se no espaço de busca ajustando sua velocidade com inércia, componente cognitiva (pbest) e componente social (gbest). A codificação utiliza random keys: o vetor real ordenado produz a permutação.",
	},
	{
		Name:        "Limitante Inferior",
		Slug:        "lowerbound",
		Description: "Calcula um limitante inferior teórico para o makespan usando a Árvore Geradora Mínima (MST) ou a relaxação do Problema de Designação (AP) via algoritmo Hungarian. Nenhuma rota factível pode ter custo inferior.",
	},
}

// AprendizadoHandler serves the /aprendizado-visual/ educational pages.
type AprendizadoHandler struct {
	tmpl *template.Template
}

// NewAprendizadoHandler creates a handler backed by the given template set.
func NewAprendizadoHandler(tmpl *template.Template) *AprendizadoHandler {
	return &AprendizadoHandler{tmpl: tmpl}
}

// RegisterRoutes wires the aprendizado-visual routes into mux.
func (h *AprendizadoHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/aprendizado-visual/", h.Handle)
	mux.HandleFunc("/aprendizado-visual", h.Handle)
}

// Handle dispatches between the index page and individual method pages.
func (h *AprendizadoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	slug := strings.TrimPrefix(r.URL.Path, "/aprendizado-visual")
	slug = strings.TrimPrefix(slug, "/")
	slug = strings.TrimSuffix(slug, "/")

	if slug == "" {
		h.renderIndex(w)
		return
	}

	for _, m := range aprendizadoMethods {
		if m.Slug == slug {
			h.renderMethod(w, m)
			return
		}
	}

	http.NotFound(w, r)
}

func (h *AprendizadoHandler) renderIndex(w http.ResponseWriter) {
	data := AprendizadoPageData{
		Title:   "Aprendizado Visual — Meta-heurísticas Bioinspiradas para o TSP-SD-ATP",
		IsIndex: true,
		Methods: aprendizadoMethods,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.ExecuteTemplate(w, "aprendizado", data); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h *AprendizadoHandler) renderMethod(w http.ResponseWriter, m AprendizadoMethod) {
	data := AprendizadoPageData{
		Title:       m.Name,
		Description: m.Description,
		Slug:        m.Slug,
		Methods:     aprendizadoMethods,
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := h.tmpl.ExecuteTemplate(w, "aprendizado", data); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
