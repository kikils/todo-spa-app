module github.com/kikils/golang-todo

go 1.14

require (
	cloud.google.com/go/firestore v1.3.0 // indirect
	cloud.google.com/go/storage v1.12.0 // indirect
	firebase.google.com/go v3.13.0+incompatible
	github.com/lib/pq v1.8.0
	github.com/rs/cors v1.7.0
	google.golang.org/api v0.32.0
)

replace (
	github.com/kikils/golang-todo/controllers => ./interfaces/controllers
	github.com/kikils/golang-todo/database => ./interfaces/database
	github.com/kikils/golang-todo/infrastructure => ./infrastructure
	github.com/kikils/golang-todo/model => ./domain/model
	github.com/kikils/golang-todo/usecase => ./usecase
)
