generate_proto:
    protoc --go_out=./domain/model/pb ./domain/model/proto/*.proto