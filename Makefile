postgres:
	docker run --name shopchaby-postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine 
createdb:
	docker exec -it shopchaby-postgres16 createdb --username=root --owner=root shopchaby
dropdb:
	docker exec -it shopchaby-postgres16 dropdb shopchaby
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/shopchaby?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/shopchaby?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/shopchaby?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/shopchaby?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store_tx.go github.com/chomchaby/shopchaby-backend-go/db/sqlc StoreTx 

.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server mock