# Spans frontend with fresh context

The frontend service now emits a Span with a new context when being called at http://localhost:8081/toastoftheday. 
The [Jaeger UI](http://localhost:16686/) will show distributed traces for the frontend service containing a single Span.

## Run this step

```sh
docker-compose build --build-arg step=2
docker-compose up
```