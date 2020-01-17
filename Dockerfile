FROM golang as builder
ENV GO111MODULE=on
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY pkg/ pkg/
ARG binary
COPY cmd/${binary}/main.go main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o /main main.go

FROM gcr.io/distroless/static
COPY --from=builder /main /main
ENTRYPOINT ["/main"]