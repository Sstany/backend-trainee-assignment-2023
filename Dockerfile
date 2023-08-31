FROM golang:1.20.3-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN apk update && \
    apk add build-base

RUN go mod download

COPY ./ ./

RUN go build -o avito /app/cmd/main.go

FROM alpine

WORKDIR /app

COPY --from=builder /app/segmenty /app/segmenty

CMD [ "./segmenty" ]

EXPOSE 8090