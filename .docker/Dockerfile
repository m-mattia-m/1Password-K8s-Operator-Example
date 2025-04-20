FROM golang:1.21 AS builder

ENV GO111MODULE=on

WORKDIR /build

COPY ../go.mod /build
COPY ../go.sum /build
RUN go mod download

COPY ../ /build

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/dist/1p-k8s-example ./main.go

FROM alpine:3.18.4

ENV GIN_MODE=release

WORKDIR /app

COPY --from=builder /build/dist/1p-k8s-example /app/1p-k8s-example

RUN chmod +x /app/1p-k8s-example

EXPOSE 8080

CMD ["/app/1p-k8s-example"]