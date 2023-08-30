FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o entry

EXPOSE 9000
EXPOSE 9001

ENTRYPOINT [ "./entry" ]