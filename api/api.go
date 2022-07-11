package api

import (
	"log"
	"net/http"
	"scraper/api/handlers"
	"time"
)

func GetMuxAPI() *http.ServeMux {
	log.Print("Initializing Rest Endpoints " + time.Now().String())
	mux := http.NewServeMux()

	filteredImgHandler := http.HandlerFunc(handlers.ImageHandler)
	filteredPdfHandler := http.HandlerFunc(handlers.PdfHandler)

	mux.Handle("/img", handlers.FilteredMiddleware(filteredImgHandler))
	mux.Handle("/pdf", handlers.FilteredMiddleware(filteredPdfHandler))

	return mux
}
