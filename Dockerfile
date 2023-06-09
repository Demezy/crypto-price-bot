FROM alpine as certs
RUN apk --no-cache update &&\
  apk add --no-cache ca-certificates &&\
  rm -rf /var/cache/apk/*

FROM golang:1-alpine as build
WORKDIR /app
# avoid redownloading without dependencies changing
COPY ./go.mod ./go.sum ./
RUN go mod download && go mod verify
# copy sources
COPY ./cmd ./cmd
COPY ./internal ./internal
# build standalone application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM busybox as production
WORKDIR /app
# copy SLL certificates
COPY --from=certs /etc/ssl/certs /etc/ssl/certs
# copy built from previous stage
COPY --from=build /app/main ./
# define entypoint
CMD ["/app/main"]
