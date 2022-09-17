FROM golang:1.18.0-alpine

ENV ROOT=/go/src/app

WORKDIR ${ROOT}

RUN apk add --no-cache gcc alpine-sdk build-base

COPY go.mod go.sum ./
RUN go mod download

COPY app ./app
EXPOSE 80

RUN adduser -D dev && chown -R dev ${ROOT}
USER dev

RUN mkdir bin
RUN go build -o bin ./...

CMD ["./bin/app"]
# BUILD: docker build -t sample-api:latest -f app/docker/Dockerfile .
# RUN  : docker run -it -p 80:80 sample-api:latest
# Build Only DB : docker-compose up -d db
# Build Only App : docker-compose up --build go