package main

import (
	"fmt"
	"github.com/urfave/negroni"
	"github.com/yehorshapanov/BitmediaGo/config"
	db "github.com/yehorshapanov/BitmediaGo/db"
	"github.com/yehorshapanov/BitmediaGo/logger"
	"github.com/yehorshapanov/BitmediaGo/service"
	"strconv"
)

func main() {
	config.Load()

	logger.Init()
	db.Init()

	deps := service.Dependencies{
		DB: db.GetStorer(db.Get()),
	}

	// mux router
	router := service.InitRouter(deps)

	// init web server
	server := negroni.Classic()
	server.UseHandler(router)

	port := config.AppPort() // This should be changed to the service port number via argument or environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	server.Run(addr)
}