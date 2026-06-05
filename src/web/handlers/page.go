package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"tcc/web/service"
)

// PageHandler serves the experiment visualization pages and HTMX partial endpoints.
type PageHandler struct {
	templates   *template.Template
	pageService *service.PageService
}

// NewPageHandler creates a handler with the given template set and service layer.
func NewPageHandler(templates *template.Template, pageService *service.PageService) *PageHandler {
	return &PageHandler{
		templates:   templates,
		pageService: pageService,
	}
}

// RegisterRoutes wires the HTMX endpoints to the given mux:
//   - GET /ui/select-map     — select an instance
//   - GET /ui/select-method  — select a method
//   - GET /ui/select-run     — select an execution
//   - GET /                  — full page render (also accepts map/method/run query params)
func (h *PageHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ui/select-map", h.SelectMap)
	mux.HandleFunc("/ui/select-method", h.SelectMethod)
	mux.HandleFunc("/ui/select-run", h.SelectRun)
	mux.HandleFunc("/", h.Index)
}

// MethodLabel maps internal method IDs to display names.
func MethodLabel(method string) string {
	switch method {
	case "aco":
		return "ACO"
	case "pso":
		return "PSO"
	case "ga":
		return "GA"
	case "bruteforce":
		return "Busca Exaustiva"
	case "lowerbound":
		return "Limitante Inferior"
	default:
		return strings.ToUpper(method)
	}
}

// SelectMap handles /ui/select-map?map=X&current_map=Y&expanded_map=Z.
// Returns HTML for main content and nav bar (via OOB swap).
func (h *PageHandler) SelectMap(w http.ResponseWriter, r *http.Request) {
	if !requireGET(w, r) {
		return
	}

	currentState := service.SelectionState{
		SelectedMap: queryValue(r, "current_map"),
		ExpandedMap: queryValue(r, "expanded_map"),
	}

	selection := h.pageService.ResolveMapSelection(currentState, queryValue(r, "map"))

	data, err := h.pageService.BuildPage(r.Context(), selection)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	h.executeTemplate(w, "main_with_nav_oob", data)
}

// SelectMethod handles /ui/select-method?map=X&method=Y&current_method=Z&expanded_method=W.
func (h *PageHandler) SelectMethod(w http.ResponseWriter, r *http.Request) {
	if !requireGET(w, r) {
		return
	}

	selectedMap := queryValue(r, "map")
	currentState := service.SelectionState{
		SelectedMap:    selectedMap,
		ExpandedMap:    selectedMap,
		SelectedMethod: queryValue(r, "current_method"),
		ExpandedMethod: queryValue(r, "expanded_method"),
	}

	selection := h.pageService.ResolveMethodSelection(currentState, queryValue(r, "method"))

	data, err := h.pageService.BuildPage(r.Context(), selection)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	h.executeTemplate(w, "main_with_nav_oob", data)
}

// SelectRun handles /ui/select-run?map=X&method=Y&run=Z.
func (h *PageHandler) SelectRun(w http.ResponseWriter, r *http.Request) {
	if !requireGET(w, r) {
		return
	}

	selection := h.pageService.ResolveRunSelection(
		queryValue(r, "map"),
		queryValue(r, "method"),
		queryValue(r, "run"),
	)

	data, err := h.pageService.BuildPage(r.Context(), selection)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	h.executeTemplate(w, "main_with_nav_oob", data)
}

// Index handles GET / and renders the full page with nav bar and main content.
func (h *PageHandler) Index(w http.ResponseWriter, r *http.Request) {
	if !requireGET(w, r) {
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	selection := service.SelectionState{
		SelectedMap:    queryValue(r, "map"),
		ExpandedMap:    queryValue(r, "map"),
		SelectedMethod: queryValue(r, "method"),
		ExpandedMethod: queryValue(r, "method"),
		SelectedRun:    queryValue(r, "run"),
	}

	data, err := h.pageService.BuildPage(r.Context(), selection)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	h.executeTemplate(w, "base", data)
}

func (h *PageHandler) executeTemplate(w http.ResponseWriter, name string, data service.PageData) {
	if err := h.templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func requireGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == http.MethodGet {
		return true
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return false
}

func queryValue(r *http.Request, key string) string {
	return strings.TrimSpace(r.URL.Query().Get(key))
}
