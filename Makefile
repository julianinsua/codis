postgres:
	docker run --name postgres_db --network codis_net -p 5432:5432 -e POSTGRES_USER=postgres POSTGRES_PASSWORD=password -d postgres:14.1-alpine 

devdb:
	docker-compose -f ~/code/postgreSQL/docker-compose.yml up -d

createdb:
	docker exec -it postgres_db createdb --username=postgres --owner=postgres codis

dropdb:
	docker exec -it postgres_db dropdb --username=postgres codis

migrateup:
	goose -dir ./database/migrations postgres postgresql://postgres:password@localhost:5432/codis up

migratedown:
	goose -dir ./database/migrations postgres postgresql://postgres:password@localhost:5432/codis down

sqlc:
	sqlc generate

serverrun:
	go run main.go

mock:
	mockgen --build_flags=--mod=mod -destination db/mock/store.go -package mock_db github.com/julianinsua/codis/internal/database Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc serverrun mock
