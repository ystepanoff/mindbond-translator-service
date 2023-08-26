package main

import (
	"flotta-home/mindbond/translator-service/pkg/config"
	"flotta-home/mindbond/translator-service/pkg/openai"
	pb "flotta-home/mindbond/translator-service/pkg/pb"
	services "flotta-home/mindbond/translator-service/pkg/services"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	client := openai.Init(config.OpenAIToken)
	listener, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}
	fmt.Println("Translator service on", config.Port)
	server := services.Server{
		TranslationClient: &client,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTranslatorServiceServer(grpcServer, &server)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
