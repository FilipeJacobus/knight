# syntax=docker/dockerfile:1

FROM golang:1.17.2


WORKDIR /knight

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /knight


CMD [ "./knight" ]