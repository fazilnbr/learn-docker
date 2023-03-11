FROM golang:1.19.0

WORKDIR /ProjectDirectory/Golang/learning/docker01


RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go get github.com/subosito/gotenv
RUN go mod tidy