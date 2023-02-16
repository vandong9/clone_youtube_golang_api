generate_proto:
    protoc --go_out=./domain/model/pb ./domain/model/proto/*.proto
    
runapi:
    go run ./cmd/api/main.go