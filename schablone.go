package main

// This code is heavily inspired by the following example of "oapi-codegen" OpenAPI Code Generation:
// https://github.com/deepmap/oapi-codegen/tree/master/examples/petstore-expanded

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	SchabloneApi "github.com/Alone2/schablone-server/api"
	"github.com/go-chi/chi/v5"
)

func main() {
	// Check if all arguments here
	argLength := len(os.Args)
	if argLength < 5 {
		fmt.Fprintf(os.Stderr, "Arguments missing, launch the application like this ./schablone-server ${PORT} ${MARIADB_USER} ${MARIADB_PASSWORD} ${MARIADB_HOST}")
		os.Exit(1)
	}

	argumentPort := os.Args[1]
	mariadbUser := os.Args[2]
	mariadbPassword := os.Args[3]
	mariadbHost := os.Args[4]
	fmt.Println("MariaDB:")
	fmt.Println("  Host:", mariadbHost)
	fmt.Println("  Port:", argumentPort)
	fmt.Println("  User:", mariadbUser)

	// Check if port valid
	intPort, err := strconv.Atoi(argumentPort)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Invalid port, launch the application like this ./schablone-server ${PORT} ${MARIADB_USER} ${MARIADB_PASSWORD} ${MARIADB_HOST}")
		os.Exit(1)
	}

	var port = flag.Int("port", intPort, "Port for test HTTP server")
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
	schabloneServer := SchabloneApi.NewSchabloneServer(mariadbUser, mariadbPassword, mariadbHost)

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
