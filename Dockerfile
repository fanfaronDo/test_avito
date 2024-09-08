FROM golang:1.22

RUN go version

ENV GOPATH=/
WORKDIR ./app
COPY ./ ./

RUN apt-get update
#RUN apt-get -y install postgresql-client

RUN go mod download && go build -o app ./cmd/app/main.go

CMD ["./app"]