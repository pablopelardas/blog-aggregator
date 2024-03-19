module middlewares

go 1.22.0

require github.com/go-chi/chi v1.5.5

require internal/database v1.0.0

replace internal/database v1.0.0 => ../database
require internal/helpers v1.0.0

replace internal/helpers v1.0.0 => ../helpers

require (
	github.com/go-chi/cors v1.2.1
	internal/helpers v1.0.0
)

require github.com/google/uuid v1.6.0

replace internal/helpers v1.0.0 => ../helpers
