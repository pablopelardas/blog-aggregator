module scraper

require internal/database v1.0.0

require github.com/google/uuid v1.6.0 // indirect

replace internal/database => ../database

require internal/api v1.0.0
replace internal/api => ../api

go 1.22.0
