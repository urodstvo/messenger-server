gen:
	protoc --go_out=./__generated__/pb/ --go-grpc_out=./__generated__/pb/ ./src/*.proto
