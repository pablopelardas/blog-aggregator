module helpers

require internal/database v1.0.0

require github.com/google/uuid v1.6.0 // indirect

replace internal/database => ../database

go 1.22.0
