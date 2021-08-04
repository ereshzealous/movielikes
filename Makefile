migrateup:
	migrate -path db/migration -database "postgresql://developer:postgres@localhost:5432/movieDB?sslmode=disable" -verbose up

sqlc:
	sqlc generate

test:
	go test -v -cover ./...