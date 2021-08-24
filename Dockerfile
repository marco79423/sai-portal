FROM golang:1.14 AS go-builder
ENV CGO_ENABLED=0 GOFLAGS=-mod=vendor
WORKDIR /app
COPY . .
RUN go build -o sai-portal /app/cmd/sai-portal/

FROM alpine:3.12
RUN apk update && \
    apk upgrade && \
    apk add --no-cache ca-certificates && \
    apk add git

WORKDIR /app
COPY --from=go-builder /app/sai-portal /app/sai-portal
COPY --from=go-builder /app/conf.d /app/conf.d

ENTRYPOINT ["/app/sai-portal"]
