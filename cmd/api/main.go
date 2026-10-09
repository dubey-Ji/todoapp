package main

import (
	"log"
	"net/http"
	"time"

	"github.com/dubey-Ji/todoapp/internal/config"
	"github.com/dubey-Ji/todoapp/internal/mux"
)

func main() {

	cgf := config.MustLoad()

	server := http.Server{
		Addr:         ":" + cgf.Port,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 60,
		Handler:      mux.Mux(),
	}

	log.Printf("server is running at %v for %v enviroment", server.Addr, cgf.Env)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed %v", err)
	}
}
