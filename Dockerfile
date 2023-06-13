FROM golang:1.19

WORKDIR /

RUN go mod init teste

COPY . .

RUN go build -o math

CMD ["./math"]