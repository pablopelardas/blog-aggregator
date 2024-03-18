module github.com/pablopelardas/blog-aggregator

go 1.22.0

require github.com/joho/godotenv v1.5.1

require internal/helpers v1.0.0

replace internal/helpers => ./internal/helpers

require internal/middlewares v1.0.0

require (
	github.com/go-chi/chi v1.5.5 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
)

replace internal/middlewares => ./internal/middlewares
