module github.com/pablopelardas/blog-aggregator

go 1.22.0

require github.com/joho/godotenv v1.5.1 
require internal/helpers v1.0.0
replace internal/helpers => ./internal/helpers
require internal/middlewares v1.0.0
replace internal/middlewares => ./internal/middlewares
