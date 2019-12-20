package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/inovex/godays-demo/pkg"
	"log"
	"net/http"
	"time"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "port to listen on")
}

var toasts = []pkg.Toast{
	{
		Name: "Hawaii",
		Weekday: time.Monday,
	},
	{
		Name: "Peperoni",
		Weekday: time.Tuesday,
	},
	{
		Name: "Cheese",
		Weekday: time.Wednesday,
	},
	{
		Name: "Ham",
		Weekday: time.Thursday,
	},
	{
		Name: "Caprese",
		Weekday: time.Friday,
	},
	{
		Name: "Avocado",
		Weekday: time.Saturday,
	},
	{
		Name: "Honey",
		Weekday: time.Sunday,
	},
}


func toastsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toasts)
}

func main() {
	flag.Parse()
	http.HandleFunc("/toasts/", toastsHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
