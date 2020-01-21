# Initial Demo State

Neither the backend nor the frontend send any spans to jaeger. You can access the frontend service at http://localhost:8081/toastoftheday, but the [Jaeger UI](http://localhost:16686/) will show no traces for the demo services.

## Run this step

```sh
docker-compose build --build-arg step=1
docker-compose up
```