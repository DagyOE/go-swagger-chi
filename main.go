package main

import (
	"flag"
	"fmt"
	"go-swagger-chi/api"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	middleware "github.com/oapi-codegen/nethttp-middleware"
)

func main() {
	port := flag.String("port", "8080", "port to listen on")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	swagger.Servers = nil

	petStore := api.NewPetStore()

	r := chi.NewRouter()

	r.Use(middleware.OapiRequestValidator(swagger))

	api.HandlerFromMux(petStore, r)

	s := &http.Server{
		Handler: r,
		Addr:    net.JoinHostPort("0.0.0.0", *port),
	}

	log.Fatal(s.ListenAndServe())
}
