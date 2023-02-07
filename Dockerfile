FROM golang:alpine
WORKDIR /build/
COPY . .
RUN apk update && apk add libc-dev && apk add gcc
RUN go build -ldflags "-s -w" -o app
EXPOSE 8080
