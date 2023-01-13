# First stage: build the executable.
FROM golang:1.17-alpine3.13 as builder
RUN apk update && apk add bash ca-certificates git gcc g++ libc-dev librdkafka-dev pkgconf
WORKDIR /app
COPY . .
RUN go build -tags musl -o main .

FROM alpine:3.13.6
WORKDIR /app
COPY --from=builder /app .
EXPOSE 8874
CMD ["/app/main"]