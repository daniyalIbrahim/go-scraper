package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"scraper/internal/models"
)

func FilteredMiddleware(next http.Handler) http.Handler {
	log.Printf("FilteredMiddleware called")
	var response map[string]interface{}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		EnableCors(&w, req)
		var url = req.URL.Query().Get("url")
		if url == "" || IsValidUrl(url) == false {
			log.Printf("Error: URL is either empty or invalid")
			response = models.CreateResponse("failure", "URL is either empty or invalid", http.StatusBadRequest)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response)
			return
		}
		// Call the next handler
		next.ServeHTTP(w, req)
	})
}
