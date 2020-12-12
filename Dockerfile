FROM golang:alpine

LABEL maintainer="C J Silverio <ceejceej@gmail.com>"

ENV REDIS_HOST=127.0.0.1
ENV REDIS_PORT=6379
ENV REDIS_PASSWORD=
ENV REDIS_KEY=LB
ENV SLACK_TOKEN=
ENV WELCOME_CHANNEL=

RUN mkdir /loudbot
WORKDIR /loudbot
COPY . .
RUN apk update && apk add --no-cache git
RUN apk update && apk add --no-cache bash
RUN go install -v ./...

CMD ["bash", "seed-and-go.sh"]
