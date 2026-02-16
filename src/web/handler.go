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

	return mux, nil
}
