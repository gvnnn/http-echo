FROM golang:1.22 AS build

WORKDIR /app

COPY src/go.mod ./
COPY src/*.go ./

RUN CGO_ENABLED=0 go build -o http-echo

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.9

COPY --from=build /app/http-echo .

ENTRYPOINT ["./http-echo"]
