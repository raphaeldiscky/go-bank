makeFileDir := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

DB_LOCAL_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable
DB_AWS_URL=postgresql://root:Iq0EEnRFtuwA1Qeopgiq@simple-bank.c16qeeq24ykj.ap-southeast-1.rds.amazonaws.com/simple_bank

network:
	docker network create bank-network

postgres:
	docker run --name postgres16 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres16 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_LOCAL_URL)" -verbose up

migrateup-aws:
	migrate -path db/migration -database "$(DB_AWS_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_LOCAL_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_LOCAL_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_LOCAL_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	docker run --rm -v $(makeFileDir):/src -w /src sqlc/sqlc generate

test:
	go test -v -cover -short ./...

db_docs:
	dbdocs build doc/db.dbml

db_schema:
	dbml2sql --postgres -o doc/schema.sql doc/db.dbml

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/raphaeldiscky/simple-bank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/raphaeldiscky/simple-bank/worker TaskDistributor

proto:
	rm -f pb/*.go
	rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	--openapiv2_out=doc/swagger --openapiv2_opt=allow_merge=true,merge_file_name=simple_bank \
	proto/*.proto
	statik -src=./doc/swagger -dest=./doc -f

evans:
	evans --host localhost --port 9090 -r repl

redis:
	docker run --name redis -p 6379:6379 -d redis:7-alpine

.PHONY: network postgres createdb dropdb migrateup migrateup-aws migratedown migrateup1 migratedown1 new_migration db_docs db_schema sqlc test server mock proto evans redis