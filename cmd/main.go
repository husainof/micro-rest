package main

import (
	"flag"
	"fmt"
	"log"
	"micro/internal/config"
	"micro/internal/user"
	"micro/pkg/logging"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func init() {
	flag.Parse()
}

func main() {
	logger := logging.GetLogger()

	config := config.GetConfig()

	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)

	start(router, config)
}

func start(router *httprouter.Router, config *config.Config) {

	logger := logging.GetLogger()
	logger.Info("start application")

	addr := fmt.Sprintf("%s:%s", config.Listen.BindIP, config.Listen.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}

	server := &http.Server{
		Handler:      router,
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
	}
	logger.Info("server is listening port ", addr)
	log.Fatal(server.Serve(listener))
}
