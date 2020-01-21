# Spans from backend and frontend with separate contexts

Both services now emit Spans with a new context when being called. After accessing http://localhost:8081/toastoftheday, 
the [Jaeger UI](http://localhost:16686/) will show distributed traces for each service, which all contain the single Span of their service.

## Run this step

```sh
docker-compose build --build-arg step=3
docker-compose up
```