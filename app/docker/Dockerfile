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

CMD ["go", "run", "app/main.go"]

# BUILD: docker build -t sample-api:latest -f app/docker/Dockerfile .
# RUN  : docker run -it -p 8080:8080 sample-api:latest