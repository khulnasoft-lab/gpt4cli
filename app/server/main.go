package main

import (
	"log"
	"os"
	"gpt4cli-server/routes"
	"gpt4cli-server/setup"

	"github.com/gorilla/mux"
)

func main() {
	// Configure the default logger to include milliseconds in timestamps
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	routes.RegisterHandleGpt4cli(func(router *mux.Router, path string, isStreaming bool, handler routes.Gpt4cliHandler) *mux.Route {
		return router.HandleFunc(path, handler)
	})

	r := mux.NewRouter()
	routes.AddHealthRoutes(r)
	routes.AddApiRoutes(r)
	routes.AddProxyableApiRoutes(r)
	setup.MustLoadIp()
	setup.MustInitDb()
	setup.StartServer(r, nil, nil)
	os.Exit(0)
}
