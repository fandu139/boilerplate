FROM golang:1.13-alpine AS builder

RUN apk update && \
    apk upgrade && \
    apk add curl make \
    && curl -fLo /usr/local/bin/air https://git.io/linux_air  \
    && chmod +x /usr/local/bin/air \
    && mkdir -p /app 
    
WORKDIR /app
COPY    . .

RUN make tool-linux \
    && make install \
    && make clean

EXPOSE 3000