postgres:
	docker run --name postgreslts -p 5432:5432 -e POSTGRES_USER=bendito -e POSTGRES_PASSWORD=krypt.makn -d postgres

createdb:
	docker exec -it postgreslts createdb --username=bendito --owner=bendito amabilisbank

dropdb:
	docker exec -it postgreslts dropdb --username=bendito --owner=bendito amabilisbank

migrateup:
	migrate -path databases/migrations -database "postgresql://bendito:krypt.makn@localhost:5432/amabilisbank?sslmode=disable" -verbose up

migratedown:
	migrate -path databases/migrations -database "postgresql://bendito:krypt.makn@localhost:5432/amabilisbank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: createdb dropdb postgres migrateup migratedown sqlc
