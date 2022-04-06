FROM golang:1.17.5-alpine

WORKDIR /app/src
COPY ./src/go.mod ./
COPY ./src/go.sum ./
RUN go mod download
COPY ./src/* ./

CMD ["go", "run", "."]