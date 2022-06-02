postgres:
	docker run --name postgreslts -p 5432:5432 -e POSTGRES_USER=bendito -e POSTGRES_PASSWORD=krypt.makn -d postgres

createdb:
	docker exec -it postgreslts createdb --username=bendito --owner=bendito amabilisbank

dropdb:
	docker exec -it postgreslts dropdb --username=bendito --owner=bendito amabilisbank

.PHONY: createdb dropdb postgres

