# build stage
FROM golang:1.10-alpine AS build-env
ARG APP_NAME=goapp

RUN apk add --no-cache curl bash git openssh
RUN go get -u github.com/golang/dep/cmd/dep
    
COPY . /go/src/github.com/govinda-attal/cabride-api
WORKDIR /go/src/github.com/govinda-attal/cabride-api
RUN dep ensure -v && go build -o $APP_NAME

# final stage
FROM alpine:3.7
RUN apk -U add ca-certificates

WORKDIR /app
COPY --from=build-env /go/src/github.com/govinda-attal/cabride-api/$APP_NAME /app/
COPY --from=build-env /go/src/github.com/govinda-attal/cabride-api/config/app-config.yaml /app/config/app-config.yaml

VOLUME [ "/app/config" ]
EXPOSE 9080 10000

# ENTRYPOINT ./cabride serve

