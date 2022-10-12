module example/server

go 1.19

replace example.com/routes => ./routes
replace example.com/helpers => ./helpers

require (
	example.com/routes v0.0.0-00010101000000-000000000000
	example.com/helpers v0.0.0-00010101000000-000000000000 // indirect
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-chi/cors v1.2.1
	github.com/joho/godotenv v1.4.0
)