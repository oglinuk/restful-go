package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/oglinuk/restful-go/internal/api/router"
	"github.com/oglinuk/restful-go/internal/pkg/config"
)

// Run the api
func Run() {
	cfg := config.Get()

	r := router.New(nil)

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	log.Printf("Serving at %s ...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
