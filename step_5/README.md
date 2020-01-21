# Function Spans, Logs and Tags

In this part we show the advantage of using tags and how to trace function.
Also the Toast backend service will return sometimes an emtpy slice.
After accessing the frontend service at http://localhost:8081/toastoftheday, the [Jaeger UI](http://localhost:16686/)
will show distributed traces composed of the frontend and backend span and the function call.

## Run this step

```sh
docker-compose build --build-arg step=5
docker-compose up -d
```