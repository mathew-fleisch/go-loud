FROM golang:alpine

LABEL maintainer="C J Silverio <ceejceej@gmail.com>"

ARG REDIS_HOST=127.0.0.1
ARG REDIS_PORT=6379
ARG REDIS_PASSWORD
ARG REDIS_KEY=LB
ARG SLACK_TOKEN
ARG WELCOME_CHANNEL

RUN mkdir /loudbot
WORKDIR /loudbot
COPY . .
RUN apk update && apk add --no-cache git
RUN apk update && apk add --no-cache bash
RUN go install -v ./...

CMD ["bash", "seed-and-go.sh"]
