createdb:
	docker exec -it  daily-quote createdb --username=root --owner=root quote-app

dropdb:
	docker exec -it daily-quote dropdb quote-app

migrateup:
	migrate -database "postgresql://root:123456@localhost?sslmode=disable" -path db/migrations -verbose up

migratedown:
	migrate -database "postgresql://root:123456@localhost?sslmode=disable" -path db/migrations -verbose down



.PHONY: createdb dropdb migrateup migratedown