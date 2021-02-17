FROM golang:1.15-buster

COPY go.mod go.sum /go/src/gitlab.com/idoko/shikari/
WORKDIR /go/src/gitlab.com/idoko/shikari
RUN go mod download
COPY . /go/src/gitlab.com/idoko/shikari/
RUN go build -o /usr/bin/api gitlab.com/idoko/shikari/cmd

EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/api"]
