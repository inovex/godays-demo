package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/inovex/godays-demo/pkg"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "port to listen on")
}

func toastsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	byteToasts, _ := json.Marshal(pkg.GetToasts())
	_, _ = w.Write(byteToasts)
}

func main() {
	flag.Parse()
	http.HandleFunc("/toasts/", toastsHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
