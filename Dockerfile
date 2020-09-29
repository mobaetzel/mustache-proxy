# Build Executable
FROM golang:1.13-alpine AS Builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
ENV GIT_CURL_VERBOSE=1
ENV GIT_TRACE=1

ENV VERSION=Builder

RUN apk update && \
    apk add make \
            git \
            gcc \
            g++ \
            libc-dev \
    && rm -rf /var/cache/apk/*

RUN adduser -D -g '' appuser

COPY . /source
WORKDIR /source

RUN mkdir ./dist

RUN go mod vendor -v
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o ./dist/service -v ./cmd/main.go


# Build Production Image
FROM alpine:3

COPY --from=Builder /etc/passwd /etc/passwd
COPY --from=Builder /source/dist/service /usr/bin/mustache-proxy

USER appuser

EXPOSE 5555
ENTRYPOINT ["mustache-proxy"]
