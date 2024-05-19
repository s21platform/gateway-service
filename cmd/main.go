package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/s21platform/gateway-service/internal/config"
	"github.com/s21platform/gateway-service/internal/useCase/RESTHandlers"
	"net/http"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	r := chi.NewRouter()
	RESTHandlers.AttachHandlers(r, cfg)
	fmt.Println(fmt.Sprintf(":%s", cfg.Service.Port))
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Service.Port), r)
}
