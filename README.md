# GoDays Demo

Build and run the demo: `docker-compose up -d`

## Endpoints

The `docker-compose up` command will start 3 containers:

- Jaeger all-in-one (contains all services for [Jaeger](https://www.jaegertracing.io/)), exposed at [localhost:8081](localhost:16686)
- Backend Service, exposed at [localhost:8080](localhost:8080)
- Frontend Service, exposed at [localhost:8081](localhost:8081)

## Demo

We can produce some traces with:

```bash
hey -z 10s http://localhost:8081/toastoftheday
```