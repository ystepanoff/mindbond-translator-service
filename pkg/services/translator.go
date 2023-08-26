package services

import (
	"context"
	"flotta-home/mindbond/translator-service/pkg/openai"
	"flotta-home/mindbond/translator-service/pkg/pb"
)

type Server struct {
	TranslationClient *openai.Client
}

func (server *Server) Translate(ctx context.Context, request *pb.TranslateRequest) (*pb.TranslateResponse, error) {
	result, err := server.TranslationClient.Translate(
		request.FromLang,
		request.ToLang,
		request.Message,
	)
	if err != nil {
		return nil, err
	}
	return &pb.TranslateResponse{
		Translation: result,
	}, nil
}
