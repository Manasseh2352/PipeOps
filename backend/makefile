createdb:
	docker exec -it postgres12 createdb --username=root --owner=root skudoosh

dropdb:
	docker exec -it postgres12 dropdb skudoosh

makeSchema:
	migrate create -ext sql -dir ./db/migrations -seq init

migrateup:
	migrate -path ./internal/db/migrations -database "postgresql://root:secret@localhost:5432/skudoosh?sslmode=disable" -verbose up

migratedown:
	migrate -path ./internal/db/migrations -database "postgresql://root:secret@localhost:5432/skudoosh?sslmode=disable" -verbose down

sqlc:
	sqlc generate

force:
	migrate -path ./internal/db/migrations -database "postgresql://root:secret@localhost:5432/skudoosh?sslmode=disable" force 1

.PHONY: postgresContainer createdb dropdb migrateup migratedown sqlc force  makeSchema
 