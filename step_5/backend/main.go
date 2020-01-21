package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	opentracingLog "github.com/opentracing/opentracing-go/log"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/inovex/godays-demo/pkg"
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

	// Set tags for filtering
	ext.SpanKindRPCServer.Set(span)
	ext.HTTPUrl.Set(span, r.URL.String())
	ext.HTTPMethod.Set(span, r.Method)

	w.Header().Set("Content-Type", "application/json")

	if rand.Int()%4 != 0 {
		byteToasts, _ := json.Marshal(pkg.GetToasts())
		ext.HTTPStatusCode.Set(span, uint16(http.StatusOK))
		_, _ = w.Write(byteToasts)
	} else {
		// Ups an error occurred
		ext.HTTPStatusCode.Set(span, uint16(http.StatusInternalServerError))
		span.LogFields(
			opentracingLog.String("error", "could not fetch toasts"),
		)
		byteToasts, _ := json.Marshal([]pkg.Toast{})
		http.Error(w, string(byteToasts), http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/toasts", toastsHandler)
	closer, err := pkg.InitGlobalTracer()
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()
	rand.Seed(time.Now().Unix())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
