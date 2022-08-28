package pb

//go:generate protoc --proto_path=$PROTOC_INCLUDE_DIR --proto_path=. --go-grpc_out=paths=source_relative:. --go_out=paths=source_relative:. ./producer.proto
