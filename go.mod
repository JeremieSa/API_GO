module TP1

go 1.17

// EXTERNAL MODULES
require github.com/gorilla/mux v1.8.0 // indirect

// INTERNAL MODULES (with REPLACE)
require internal/entities v1.0.0

replace internal/entities => ./internal/entities

require internal/entities/language v1.0.0

replace internal/entities/language => ./internal/entities/language

require internal/entities/student v1.0.0

replace internal/entities/student => ./internal/entities/student

require internal/persistence v1.0.0

replace internal/persistence => ./internal/persistence

require internal/web/rest v1.0.0

require (
	github.com/boltdb/bolt v1.3.1 // indirect
	golang.org/x/sys v0.0.0-20220422013727-9388b58f7150 // indirect
)

replace internal/web/rest => ./internal/web/rest
