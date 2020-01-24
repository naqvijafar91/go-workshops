# build stage
FROM golang:alpine AS build-env

ENV BUILD_DIR=/go/src/
RUN mkdir -p $BUILD_DIR/go-workshops/docker/go
ADD . $BUILD_DIR/go-workshops/docker/go

WORKDIR $BUILD_DIR/go-workshops/docker/go
RUN go build -o /go/app

FROM alpine
COPY --from=build-env /go/app /app
ENTRYPOINT ["/app"]
