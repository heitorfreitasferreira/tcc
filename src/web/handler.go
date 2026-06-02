// Package web assembles the HTTP server for the experiment visualization dashboard.
//
// Architecture
//
// The package follows a three-layer design:
//
//	Handler (HTTP)  →  Service (business logic)  →  Repository (data access)
//
// All static assets (templates, CSS, JS) and experiment data (.graph, .points,
// results/) are embedded into the binary via Go's embed directive, making the
// server a self-contained executable.
//
// HTMX is used for partial page updates: selecting a map, method, or run in the
// sidebar triggers a GET request that returns HTML fragments. The sidebar tree is
// updated via Out-of-Band (OOB) swap so both panels refresh in one response.
//
// The frontend JavaScript is organized as ES modules (no bundler) under
// static/js/, with separate modules for rendering (Canvas API), animation
// playback (evolution and best-sequence modes), and UI control wiring.
package web

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"net/http"

	dataassets "tcc/data"
	"tcc/web/handlers"
	"tcc/web/repository"
	"tcc/web/service"
)

//go:embed templates/*.html static
var assetsFS embed.FS

// NewHandler builds the full HTTP handler, wiring templates, static file servers,
// repository, service, and handler layers. It embeds all experiment data so the
// binary is self-contained.
func NewHandler() (http.Handler, error) {
	tmpl, err := template.ParseFS(
		assetsFS,
		"templates/base.html",
		"templates/index.html",
		"templates/sidebar.html",
		"templates/main.html",
	)
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}

	staticFS, err := fs.Sub(assetsFS, "static")
	if err != nil {
		return nil, fmt.Errorf("open static root: %w", err)
	}

	cssFS, err := fs.Sub(staticFS, "css")
	if err != nil {
		return nil, fmt.Errorf("open css assets: %w", err)
	}

	jsFS, err := fs.Sub(staticFS, "js")
	if err != nil {
		return nil, fmt.Errorf("open js assets: %w", err)
	}

	imgFS, err := fs.Sub(staticFS, "img")
	if err != nil {
		return nil, fmt.Errorf("open image assets: %w", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.FS(cssFS))))
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.FS(jsFS))))
	mux.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.FS(imgFS))))

	dataRepository := repository.NewEmbeddedDataRepository(dataassets.Files)
	pageService := service.NewPageService(dataRepository, dataRepository, dataRepository)
	pageHandler := handlers.NewPageHandler(tmpl, pageService)
	pageHandler.RegisterRoutes(mux)

	renderHandler := handlers.NewRenderHandler(dataRepository, dataRepository)
	mux.Handle("/api/render", renderHandler)

	return mux, nil
}
