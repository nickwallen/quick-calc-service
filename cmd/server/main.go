package main

import (
	"flag"
	"fmt"
	service "github.com/nickwallen/quick-calc-service"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/graph-gophers/graphql-go/relay"
	log "github.com/sirupsen/logrus"

	"github.com/rs/cors"
)

func main() {
	hostname := flag.String(
		"host",
		"",
		"the hostname to bind to")
	port := flag.Int(
		"port",
		8080,
		"the port number to bind to")
	html := flag.String(
		"html",
		"cmd/server/index.html",
		"path to the iGraphQL html")
	allowedOrigins := flag.String(
		"allowedOrigins",
		"*",
		"origins for CORS requests")
	allowedHeaders := flag.String(
		"allowedHeaders",
		"*",
		"headers allowed for CORS requests")
	flag.Parse()

	schema, err := service.Schema()
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", iGraphQL(html))
	router.Handle("/query", &relay.Handler{Schema: schema}).Methods("GET", "POST", "OPTIONS", "PUT", "HEAD")

	bindAddr := fmt.Sprintf("%s:%d", *hostname, *port)
	log.WithFields(log.Fields{"time": time.Now()}).Info("Endpoint at http://", bindAddr, "/query")
	log.WithFields(log.Fields{"time": time.Now()}).Info("GraphiQL at http://", bindAddr)

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{*allowedOrigins},
		AllowedHeaders: []string{*allowedHeaders},
	})
	log.Fatal(http.ListenAndServe(bindAddr, logged(corsHandler.Handler(router))))
}

func iGraphQL(htmlPath *string) func(w http.ResponseWriter, r *http.Request) {
	page, err := readFile(*htmlPath)
	if err != nil {
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}
}

func readFile(relativePath string) (contents []byte, err error) {
	path, _ := filepath.Abs(relativePath)
	if err != nil {
		return contents, err
	}
	page, err := ioutil.ReadFile(path)
	if err != nil {
		return contents, err
	}
	return page, nil
}

func logged(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()
		next.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"path":    r.RequestURI,
			"ip":      r.RemoteAddr,
			"elapsed": time.Now().UTC().Sub(start),
		}).Info()
	})
}
