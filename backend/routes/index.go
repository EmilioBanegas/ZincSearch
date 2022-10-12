package routes

import (
	"github.com/go-chi/chi/v5"
)

func IndexRoutes() chi.Router {
	routerIndex := chi.NewRouter()	
	routerIndex.Mount("/zinc", zincResorce{}.Routes())
	return routerIndex
}

