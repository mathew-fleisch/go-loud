#!/bin/bash
if ! [[ -f .env ]]; then 
    touch .env
    if [[ -n "$REDIS_HOST" ]]; then
        echo "export REDIS_HOST=$REDIS_HOST" >> .env
    fi
    if [[ -n "$REDIS_PORT" ]]; then
        echo "export REDIS_PORT=$REDIS_PORT" >> .env
    fi
    if [[ -n "$REDIS_PASSWORD" ]]; then
        echo "export REDIS_PASSWORD=$REDIS_PASSWORD" >> .env
    fi
    if [[ -n "$REDIS_KEY" ]]; then
        echo "export REDIS_KEY=$REDIS_KEY" >> .env
    fi
    if [[ -n "$SLACK_TOKEN" ]]; then
        echo "export SLACK_TOKEN=$SLACK_TOKEN" >> .env
    fi
    if [[ -n "$WELCOME_CHANNEL" ]]; then
        echo "export WELCOME_CHANNEL=$WELCOME_CHANNEL" >> .env
    fi
fi
cd cmd/seedlouds
go build
./seedlouds
cd ../..
go build
./go-loud
