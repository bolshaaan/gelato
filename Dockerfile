FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go test -race -v ./...