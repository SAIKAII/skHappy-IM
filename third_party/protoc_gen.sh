protoc --proto_path=protocols --proto_path=third_party --go_out=plugins=grpc:protocols cli-interface.proto
protoc --proto_path=protocols --proto_path=third_party --go_out=plugins=grpc:protocols conn.proto
protoc --proto_path=protocols --proto_path=third_party --go_out=plugins=grpc:protocols tcp-conn.proto