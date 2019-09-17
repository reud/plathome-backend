# Build Stage
FROM golang:1.12.5 AS build
ENV REPOSITORY github.com/reud/plathome-backend

# aiming layer cache(module download)
WORKDIR $GOPATH/src/$REPOSITORY
ENV GO111MODULE on
COPY go.mod go.sum ./
ARG GOARCH="amd64"
RUN go mod download
# If code was updated then execute below without using the cache
ADD . $GOPATH/src/$REPOSITORY
RUN GOOS=linux GOARCH=$GOARCH CGO_ENABLED=0 go build -ldflags '-s -w' -a -installsuffix cgo -o /plathome plathome.go
# memo: -ldflags = https://shinpei.github.io/blog/2014/12/24/make-go-binary-size-small
# memo: -a (if go.mod are get old, a flag is force update packages)


# Runtime Stage
FROM alpine:3.10.1
RUN apk add --no-cache ca-certificates
COPY --from=build /plathome .
ARG HOST="0.0.0.0"
ENV DBHOST="postgres"
CMD ./plathome --port=8080 --host=${HOST}
