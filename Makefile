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

mockDB:
	mockgen --build_flags=--mod=mod -destination db/mock/store.go -package mock_db github.com/julianinsua/codis/internal/database Store

mockTokenMaker:
	mockgen --build_flags=--mod=mod -destination token/mock/token.go -package mock_token_mkr github.com/julianinsua/codis/token Maker

.PHONY: postgres createdb dropdb migrateup migratedown sqlc serverrun mock
