proto:
	protoc pkg/pb/*.proto --go_out=plugins=grpc:.

translator-service:
	go run cmd/main.go
