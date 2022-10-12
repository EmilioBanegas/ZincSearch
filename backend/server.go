// go mod tidy --> busca y descarga los módulos importados, por defecto descarga la última versión. 
package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	//Router Chi
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	
	//Envivorement Vars
	"github.com/joho/godotenv"

	//Module routes
	"example.com/routes"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	godotenv.Load()

	app := chi.NewRouter()
	
	//Middlewares
	app.Use(middleware.Logger)
	app.Use(middleware.AllowContentType("application/json","text/xml"))

	app.Use(cors.Handler(cors.Options{
    // AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
    AllowedOrigins:   []string{"https://*", "http://*"},
    // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
    AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
    //ExposedHeaders:   []string{"Link"},
    AllowCredentials: false,
    //MaxAge:           300, // Maximum value not ignored by any of major browsers
  }))

	// Adding mores routes to app
	app.Mount("/api", routes.IndexRoutes())
	
	workDir, _ := os.Getwd()
	
	filesDir := http.Dir(filepath.Join(workDir, "dist"))
	FileServer(app, "/", filesDir)


	http.ListenAndServe(":3000", app)
}

// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	r.Get("/email*", http.RedirectHandler(path+"/", 301).ServeHTTP )

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
