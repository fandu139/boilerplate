FROM golang:1.14-alpine AS builder
# FROM golang:1.15.2 AS builder
# FROM golang:1.15.2-alpine AS builder

RUN apk update && \
    apk upgrade && \
    apk add curl make \
    && curl -fLo /usr/local/bin/air https://github.com/cosmtrek/air/releases  \
    && chmod +x /usr/local/bin/air \
    && mkdir -p /app 
    
WORKDIR /app
COPY    . .

EXPOSE 3000