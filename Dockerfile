FROM golang:1.21.1 as builder

WORKDIR /build/http-long-response

COPY . .
RUN CGO_ENABLED=1 go build -o /app/http-long-response /build/http-long-response

FROM debian:12.1

COPY --from=builder /app/http-long-response /app/

ENTRYPOINT ["/app/http-long-response"]
