FROM golang:1.15-alpine

# installing git
RUN apk update && apk upgrade && \
    apk add --no-cache git

