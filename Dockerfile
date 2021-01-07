FROM golang:1.15.6-alpine3.12 as builder

COPY go.mod go.sum /go/src/gitlab.com/idoko/shikari/
WORKDIR /go/src/gitlab.com/idoko/shikari
RUN go mod download
COPY . /go/src/gitlab.com/idoko/shikari
RUN CGO_ENABLED=0 GOOS=linux go build -a -o build/shikari gitlab.com/idoko/shikari

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/gitlab.com/idoko/shikari/build/shikari /usr/bin/shikari

ENTRYPOINT ["/usr/bin/shikari"]