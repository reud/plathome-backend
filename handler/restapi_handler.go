package handler

import "net/http"
import "github.com/rs/cors"

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func CORShandler(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
		Debug:          true,
	})

	return c.Handler(handler)
}
