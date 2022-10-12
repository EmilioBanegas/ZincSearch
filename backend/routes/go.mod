module example.com/routes

go 1.19

replace example.com/helpers => ../helpers

require (
	example.com/helpers v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.7
)
