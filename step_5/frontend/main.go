package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/inovex/godays-demo/pkg"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	opentracingLog "github.com/opentracing/opentracing-go/log"
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
	span := opentracing.GlobalTracer().StartSpan("/toastoftheday")
	defer span.Finish()

	// Set tags for filtering
	ext.SpanKindRPCClient.Set(span)
	ext.HTTPUrl.Set(span, r.URL.String())
	ext.HTTPMethod.Set(span, r.Method)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/toasts", backendUrl), nil)
	if err != nil {
		ext.HTTPStatusCode.Set(span, uint16(http.StatusInternalServerError))
		errStr := fmt.Sprintf("Failed to create httpRequest %s", err)
		span.LogKV("error", errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	err = span.Tracer().Inject(span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))
	if err != nil {
		log.Printf("Could not inject SpanContext from request %s", err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ext.HTTPStatusCode.Set(span, uint16(http.StatusInternalServerError))
		errStr := fmt.Sprintf("Failed to call /toasts on backend service %s, err: %s", backendUrl, err)
		span.LogKV("error", errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	ctx := opentracing.ContextWithSpan(context.Background(), span)
	var toasts []pkg.Toast
	err = json.NewDecoder(resp.Body).Decode(&toasts)
	if err != nil {
		ext.HTTPStatusCode.Set(span, uint16(http.StatusInternalServerError))
		errStr := fmt.Sprintf("Failed to decode JSON, err: %s", err)
		span.LogKV("error", errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}
	toast, err := getToastOfTheDay(ctx, toasts)
	span.LogKV("Toasts", toasts)
	if err != nil {
		ext.HTTPStatusCode.Set(span, uint16(http.StatusInternalServerError))
		errStr := fmt.Sprintf("Failed to find Toast of the Day, err: %s", err)
		span.LogKV("error", errStr)
		http.Error(w, errStr, http.StatusInternalServerError)
		return
	}

	// Set the HTTP response Status Code
	ext.HTTPStatusCode.Set(span, uint16(http.StatusOK))
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(toast.Name))
}

func getToastOfTheDay(ctx context.Context, toasts []pkg.Toast) (pkg.Toast, error) {
	span, _ := opentracing.StartSpanFromContext(ctx, "getToastOfTheDay")
	defer span.Finish()

	weekday := time.Now().Weekday()
	for _, toast := range toasts {
		if toast.Weekday == weekday {
			return toast, nil
		}
	}

	span.LogFields(
		opentracingLog.String("weekday", weekday.String()),
	)

	return pkg.Toast{}, fmt.Errorf("no Toast for Weekday %s in %v", weekday, toasts)
}

func main() {
	flag.Parse()
	http.HandleFunc("/toastoftheday", toastOfTheDayHandler)
	closer, err := pkg.InitGlobalTracer()
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		return
	}
	defer closer.Close()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
