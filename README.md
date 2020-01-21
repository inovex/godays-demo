# GoDays Demo

Each step of the demo can be built using `docker-compose build --build-arg step=4`. Then run the demo: `docker-compose up -d`

The `docker-compose up -d` command will start 3 containers:

- Jaeger all-in-one (contains all services for [Jaeger](https://www.jaegertracing.io/)), exposed at [localhost:16686](localhost:16686)
- Backend Service, exposed at [localhost:8080](localhost:8080)
- Frontend Service, exposed at [localhost:8081](localhost:8081)

We can produce some traces with:

```bash
for i in {1..10}; do curl 'http://localhost:8081/toastoftheday'; done
```

Now go to the [Jaeger UI](localhost:16686) and inspect the traces.

## Cleanup

```bash
docker-compose stop && docker-compose rm -f
```