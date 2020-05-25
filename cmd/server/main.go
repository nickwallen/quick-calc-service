package main

import (
	"flag"
	"fmt"
	service "github.com/nickwallen/quick-calc-service"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"path/filepath"
	"strings"
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
	logLevel := flag.String(
		"logLevel",
		"Info",
		"the log level; Trace, Debug, Info, Warn, Error, Fatal, Panic")
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

	setupLogging(*logLevel)
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

func setupLogging(logLevel string) {
	switch strings.ToLower(logLevel) {
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	default:
		panic(fmt.Errorf("unknown log level: %s", logLevel))
	}
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
		reqDump, err := httputil.DumpRequest(r, true)
		if err != nil {
			fmt.Println(err)
		}
		next.ServeHTTP(w, r)
		log.WithFields(log.Fields{
			"ip":      r.RemoteAddr,
			"elapsed": time.Since(start),
			"request": fmt.Sprintf("%s", reqDump),
		}).Info()
	})
}
