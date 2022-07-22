package main

import (
	"fmt"
	talkService "github.com/submaline/talk-service"
	"github.com/submaline/talk-service/gen/talk/v1/talkv1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
)

func main() {
	talkServiceHandler := &talkService.TalkService{}

	mux := http.NewServeMux()
	mux.Handle(talkv1connect.NewTalkServiceHandler(
		talkServiceHandler))

	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	addr := fmt.Sprintf(":%s", port)

	log.Printf("Service listenging on %v", port)
	if err := http.ListenAndServe(
		addr,
		h2c.NewHandler(mux, &http2.Server{}),
	); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
