module github.com/GMcD/api-semaphore

go 1.19

require github.com/GMcD/api-semaphore/api v0.0.0

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
)

replace github.com/GMcD/api-semaphore/api => ./api
