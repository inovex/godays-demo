FROM golang as builder
ARG binary
WORKDIR /app
COPY . .
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o /main cmd/${binary}/main.go

FROM gcr.io/distroless/static
COPY --from=builder /main /main
ENTRYPOINT ["/main"]