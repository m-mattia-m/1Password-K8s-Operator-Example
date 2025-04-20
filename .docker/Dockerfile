FROM golang:1.23 AS builder

ENV GO111MODULE=on

WORKDIR /build

COPY ../go.mod /build
COPY ../go.sum /build
RUN go mod download

COPY ../ /build

RUN CGO_ENABLED=0 GOOS=linux go build -o /build/dist/one-password-k8s-example ./main.go

FROM alpine:3.18.4

ENV GIN_MODE=release

WORKDIR /app

COPY --from=builder /build/dist/one-password-k8s-example /app/one-password-k8s-example

RUN chmod +x /app/one-password-k8s-example

EXPOSE 8080

CMD ["/app/one-password-k8s-example"]