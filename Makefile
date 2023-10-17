proto:
	git clone git@flotta-home:mindbond/proto.git
	protoc proto/translator.proto --go_out=plugins=grpc:./pkg/
	rm -rf proto/

translator-service:
	go run cmd/main.go
