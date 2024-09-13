FROM golang:1.22.5

RUN go version

ENV GOPATH=/
WORKDIR ./app
COPY ./ ./

RUN apt-get update

RUN go get "github.com/lib/pq" && go mod download && go build -o app ./cmd/app/main.go

CMD ["./app"]