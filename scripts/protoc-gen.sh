#!/bin/bash

protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. audit.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. communicationmethod.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. communicationmethodfield.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. communicationmethodlabel.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. contact.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. contactcommunicationmethod.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. contactcommunicationmethodfield.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. contactsystem.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:. message.proto

protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:. audit.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:. communicationmethod.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:. contactsystem.proto
protoc --proto_path=api/proto/v1 --proto_path=third_party --grpc-gateway_out=logtostderr=true:. message.proto