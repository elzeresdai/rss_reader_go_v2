# base go image
FROM golang:1.19-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o rssReader ./cmd

RUN chmod +x /app/rssReader

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/rssReader /app

CMD [ "/app/rssReader" ]