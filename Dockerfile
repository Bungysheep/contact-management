FROM golang:alpine

WORKDIR /app/contact-management

COPY ./go.mod ./

RUN go mod download

COPY . .

EXPOSE 50051

RUN go build -o ./bin/contact-management ./cmd

CMD ["./bin/contact-management"]