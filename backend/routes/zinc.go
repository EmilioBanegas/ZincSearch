package routes

import (
  "strings"
  "context"
  "net/http"
  "io"
  "fmt"

  //Router Chi
	"github.com/go-chi/chi/v5"

  //Module helpers
	"example.com/helpers"
)

type zincResorce struct{}

func (rs zincResorce) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/{index}", func(r chi.Router) {
    r.Use(PostCtx)
		r.Get("/map", rs.GetMap)       // GET /api/zinc/{id}/map - Return the map of the DB.
    r.Post("/search/{type}", rs.Search)   // POST /api/zinc/{id}/map - Return the documents of the DB.
	})

	return r
}

func PostCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "index"))
    w.Header().Set("Content-Type", "application/json")
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

// Request Handler - GET /api/zinc/{id} - Get mapping of the {index}.
func (rs zincResorce) GetMap(w http.ResponseWriter, r *http.Request) {
  id := r.Context().Value("id").(string)

  resp, err := http.Get(helpers.GetUrlZinc()+id+"/_mapping" + id)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  defer resp.Body.Close()

  w.Header().Set("Content-Type", "application/json")

  if _, err := io.Copy(w, resp.Body); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}

// Request Hander - POST /api/zinc/{index}/search documents in the {index}.
func (rs zincResorce) Search(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
  api := chi.URLParam(r, "type")
  fmt.Println("param api = "+api);
  
  urlZinc := strings.Replace(helpers.GetUrlZinc(), "api", api, 1)
  fmt.Println("urlZinc = "+urlZinc);

  client := &http.Client{}

  req, err := http.NewRequest("POST", urlZinc + id + "/_search", r.Body)
  req.Header.Add("Content-Type", "application/json")
  req.Header.Set("Authorization", "Basic " + helpers.Getb64())
  req.Body = r.Body
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  resp, err := client.Do(req)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  defer resp.Body.Close()

  w.Header().Set("Content-Type", "application/json")

  if _, err := io.Copy(w, resp.Body); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
}