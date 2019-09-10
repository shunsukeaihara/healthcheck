FROM golang:1.12.9-alpine3.10 as builder
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN apk add --no-cache ca-certificates git
WORKDIR /go/src/github.com/shunsukeaihara/healthcheck
COPY main.go main.go
RUN mkdir -p bin
RUN go build -o bin/healthcheck main.go

# runtime image
FROM alpine
RUN apk add --no-cache ca-certificates
WORKDIR /var/app
COPY --from=builder /go/src/github.com/shunsukeaihara/healthcheck/bin/healthcheck /var/app/healthcheck
ENTRYPOINT ["./healthcheck"]
