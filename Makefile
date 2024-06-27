generate:
	mkdir libs\grpc\proto\__generated__ 
	protoc --proto_path libs/grpc/proto/src \
	--go_out=libs/grpc/proto/__generated__ --go_opt=paths=source_relative \
	--go-grpc_out=libs/grpc/proto/__generated__ --go-grpc_opt=paths=source_relative \
	libs/grpc/proto/src/*.proto

create-migration:
	migrate create -ext sql -dir libs/migrations/migrations -seq $(NAME)

migrate:
	migrate -database "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable" -path libs/migrations/migrations up

run:
	make generate