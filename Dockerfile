FROM golang:alpine

LABEL maintainer="C J Silverio <ceejceej@gmail.com>"

ARG redis_address=127.0.0.1
ARG redis_port=6379
ARG redis_password
ARG redis_key=LB
ARG slack_token
ARG welcome

ENV REDIS_ADDRESS=$redis_address
ENV REDIS_KEY=$redis_key
ENV SLACK_TOKEN=$slack_token
ENV WELCOME_CHANNEL=$welcome

RUN mkdir /loudbot
WORKDIR /loudbot
COPY . .
RUN apk update && apk add --no-cache git
RUN apk update && apk add --no-cache bash
RUN go install -v ./...

CMD ["bash", "seed-and-go.sh"]
