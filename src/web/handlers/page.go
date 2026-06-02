// Package handlers implements the HTTP endpoints for the experiment visualization UI.
// It uses HTMX for partial page updates: selecting a map, method, or run triggers a
// GET request that returns HTML fragments swapped into the page without a full reload.
// The sidebar tree is updated via Out-of-Band (OOB) swap so both main content and
// sidebar are refreshed in a single response.
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
//   - GET /ui/select-map     — select or toggle a map
//   - GET /ui/select-method  — select or toggle a method
//   - GET /ui/select-run     — select a specific run
//   - GET /                  — full page render (also accepts map/method/run query params)
func (h *PageHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ui/select-map", h.SelectMap)
	mux.HandleFunc("/ui/select-method", h.SelectMethod)
	mux.HandleFunc("/ui/select-run", h.SelectRun)
	mux.HandleFunc("/", h.Index)
}

// SelectMap handles /ui/select-map?map=X&current_map=Y&expanded_map=Z.
// Returns HTML for both main content and sidebar (via OOB swap).
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

	h.executeTemplate(w, "main_with_sidebar_oob", data)
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

	h.executeTemplate(w, "main_with_sidebar_oob", data)
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

	h.executeTemplate(w, "main_with_sidebar_oob", data)
}

// Index handles GET / and renders the full page with sidebar and main content.
// Accepts optional query params: map, method, run for deep-linking.
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

// executeTemplate renders the named template with the given data.
// On error it writes a 500 response.
func (h *PageHandler) executeTemplate(w http.ResponseWriter, name string, data service.PageData) {
	if err := h.templates.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

// requireGET rejects non-GET requests with 405 Method Not Allowed.
func requireGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method == http.MethodGet {
		return true
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return false
}

// queryValue returns the trimmed query parameter value for the given key.
func queryValue(r *http.Request, key string) string {
	return strings.TrimSpace(r.URL.Query().Get(key))
}
