gen:
    protoc --proto_path=domain/model proto/*.proto --go_out=plugins=grpc:pb

clean:
    rm pb/*.go 

run:
    go run main.go