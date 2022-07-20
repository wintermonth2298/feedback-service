package main

import (
	"context"
	"flag"
	"log"
	"net"

	fs "github.com/wintermonth2298/feedback-service/api/feedback-service"
	"github.com/wintermonth2298/feedback-service/internal/handler"
	"github.com/wintermonth2298/feedback-service/internal/repository"
	"github.com/wintermonth2298/feedback-service/internal/service"
	"github.com/wintermonth2298/feedback-service/pkg/config"
	"github.com/wintermonth2298/feedback-service/pkg/sql_client"
	"google.golang.org/grpc"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "config", "./cmd/feedback_service/config.json", "path to config")
	flag.Parse()

	cfg, err := config.New(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql_client.New(context.Background(), cfg.SQL)
	if err != nil {
		log.Fatalf("sql_client.New(): %v", err)
	}

	feedbackRepository := repository.NewFeedbackRepository(db)
	feedbackService := service.NewFeedbackService(feedbackRepository)
	feedbackHandler := handler.NewFeedbackHandler(feedbackService)

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("net.Listen(): %v", err)
	}

	server := grpc.NewServer()
	fs.RegisterFeedbackServiceServer(server, feedbackHandler)

	log.Println("starting server...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("server.Serve(listener): %v", err)
	}
}
