FROM golang:alpine
ENV GOPROXY=https://goproxy.io,direct \
    GOOS=linux \
    GOARCH=amd64
WORKDIR /build/
COPY . .
RUN apk update && apk add libc-dev && apk add gcc
RUN go build -ldflags "-s -w" -o app
RUN apk del libc-dev && apk del gcc
EXPOSE 8080
ENTRYPOINT [ "/build/app" ]