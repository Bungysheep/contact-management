protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. audit.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. contactsystem.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. communicationmethod.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. contact.proto