#recommeded always pulling image using digest
FROM golang:1.21.4-alpine3.18 AS builder
#tweak for disable cache
ARG CACHEBUST=1

WORKDIR $GOPATH/app

COPY ./ ./
COPY go.mod ./
COPY go.sum ./
RUN go mod download
#For small binnary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a -o /go/bin/main .

FROM alpine:latest
RUN apk update && apk add --no-cache ca-certificates bash util-linux tzdata
COPY --from=builder /go/bin/main /

ENV TZ=Asia/Jakarta
# Pick One method whether to be able to access or not
# Execute the binary.
CMD ["/main", "start","-point"]

# Use to debug
#CMD ["/main"]
EXPOSE 1338
