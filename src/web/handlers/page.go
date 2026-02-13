package handlers

import (
	"html/template"
	"net/http"
	"strings"

	"tcc/web/service"
)

type PageHandler struct {
	templates   *template.Template
	pageService *service.PageService
}

func NewPageHandler(templates *template.Template, pageService *service.PageService) *PageHandler {
	return &PageHandler{
		templates:   templates,
		pageService: pageService,
	}
}

func (h *PageHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/ui/select-map", h.SelectMap)
	mux.HandleFunc("/ui/select-method", h.SelectMethod)
	mux.HandleFunc("/ui/select-run", h.SelectRun)
	mux.HandleFunc("/", h.Index)
}

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

	h.executeTemplate(w, "sidebar_tree", data)
}

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

	h.executeTemplate(w, "sidebar_tree", data)
}

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
