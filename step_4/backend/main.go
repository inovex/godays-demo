package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/inovex/godays-demo/pkg"
	"github.com/opentracing/opentracing-go"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "port to listen on")
}

func toastsHandler(w http.ResponseWriter, r *http.Request) {
	context, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	if err != nil {
		log.Printf("Could not extract SpanContext from request %s", err)
	}
	span := opentracing.GlobalTracer().StartSpan("/toasts", opentracing.ChildOf(context))
	defer span.Finish()
	w.Header().Set("Content-Type", "application/json")
	byteToasts, _ := json.Marshal(pkg.GetToasts())
	_, _ = w.Write(byteToasts)
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
