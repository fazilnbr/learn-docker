FROM golang:1.19.0

WORKDIR /ProjectDirectory/Golang/learning/docker01

COPY . .
RUN go mod tidy