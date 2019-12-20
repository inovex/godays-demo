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
var backendUrl string

func init() {
	flag.IntVar(&port, "port", 8080, "port to listen on")
	flag.StringVar(&backendUrl, "backend-url", "http://backend:8080", "URL of the backend providing toasts")
}

func toastOfTheDayHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.DefaultClient.Get(fmt.Sprintf("%s/toasts", backendUrl))
	if err != nil {
		_, _ = fmt.Fprintf(w, "Failed to call /toasts on backend service %s, err: %s", backendUrl, err)
		return
	}
	var toasts []pkg.Toast
	err = json.NewDecoder(resp.Body).Decode(&toasts)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Failed to decode JSON, err: %s", err)
		return
	}
	toast, err := getToastOfTheDay(toasts)
	if err != nil {
		_, _ = fmt.Fprintf(w, "Failed to find Toast of the Day, err: %s", toasts, err)
		return
	}
	fmt.Fprint(w, toast.Name)
}

func getToastOfTheDay(toasts []pkg.Toast) (pkg.Toast, error) {
	weekday := time.Now().Weekday()
	for _, toast := range toasts {
		if toast.Weekday == weekday {
			return toast, nil
		}
	}
	return pkg.Toast{}, fmt.Errorf("No Toast for Weekday %s in %v", weekday, toasts)
}

func main() {
	flag.Parse()
	http.HandleFunc("/toastoftheday", toastOfTheDayHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
