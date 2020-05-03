package main

import (
	service "github.com/nickwallen/quick-calc-service"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/graph-gophers/graphql-go/relay"
)

const (
	bindAddr = "localhost:8080"
)

func main() {
	schema, err := service.Schema()
	if err != nil {
		panic(err)
	}
	page, err := readFile("cmd/server/index.html")
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))
	mux.Handle("/query", &relay.Handler{Schema: schema})
	log.WithFields(log.Fields{"time": time.Now()}).Info("starting service on http://", bindAddr, "/query")
	log.WithFields(log.Fields{"time": time.Now()}).Info("GraphiQL available at http://", bindAddr)
	log.Fatal(http.ListenAndServe(bindAddr, logged(mux)))
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
			"IP":      r.RemoteAddr,
			"elapsed": time.Now().UTC().Sub(start),
		}).Info()
	})
}
