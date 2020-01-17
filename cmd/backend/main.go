package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"log"
	"net/http"
	"time"

	"github.com/inovex/godays-demo/pkg"
)

var port int
var propagator = pkg.InitHttpHeaderPropagator()

func init() {
	flag.IntVar(&port, "port", 8080, "port to listen on")
}

var toasts = []pkg.Toast{
	{
		Name:    "Hawaii",
		Weekday: time.Monday,
	},
	{
		Name:    "Peperoni",
		Weekday: time.Tuesday,
	},
	{
		Name:    "Cheese",
		Weekday: time.Wednesday,
	},
	{
		Name:    "Ham",
		Weekday: time.Thursday,
	},
	{
		Name:    "Caprese",
		Weekday: time.Friday,
	},
	{
		Name:    "Avocado",
		Weekday: time.Saturday,
	},
	{
		Name:    "Honey",
		Weekday: time.Sunday,
	},
}

func toastsHandler(w http.ResponseWriter, r *http.Request) {
	var span opentracing.Span
	context, err := propagator.Extract(opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		log.Printf("Could not extract SpanContext from request %s", err)
		span = opentracing.GlobalTracer().StartSpan("/toasts")
	} else {
		span = opentracing.GlobalTracer().StartSpan("/toasts", opentracing.ChildOf(context))
	}
	defer span.Finish()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toasts)
}

func main() {
	flag.Parse()
	http.HandleFunc("/toasts/", toastsHandler)
	closer, err := pkg.InitGlobalTracer()
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
