package main

import (
	"log"
	"net/http"

	"go-demo/api/internal/env"
	"go-demo/api/internal/http/router"
	"go-demo/api/internal/storage/memory"
)

func main() {
	cfg := env.Load()
	repo := memory.NewUserRepo()

	r := router.New(cfg, repo)

	addr := ":" + cfg.APIPort
	log.Printf("API listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}
