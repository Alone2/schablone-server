package main

// This code is heavily inspired by the following example of OpenAPI Code Generation:
// https://github.com/deepmap/oapi-codegen/tree/master/examples/petstore-expanded

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	SchabloneApi "github.com/Alone2/schablone-server/api"
	"github.com/go-chi/chi/v5"
)

func main() {
	var port = flag.Int("port", 8080, "Port for test HTTP server")
	flag.Parse()

	swagger, err := SchabloneApi.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	schabloneServer := SchabloneApi.NewSchabloneServer()

	// This is how you set up a basic chi router
	r := chi.NewRouter()

	// We now register our petStore above as the handler for the interface
	SchabloneApi.HandlerFromMux(schabloneServer, r)

	s := &http.Server{
		Handler: r,
		Addr:    fmt.Sprintf("0.0.0.0:%d", *port),
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
