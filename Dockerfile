FROM golang

WORKDIR /app/contact-management

COPY ./go.mod ./

RUN go mod download

COPY . .

EXPOSE 50051

RUN go build -tags dev -o ./bin/contact-management ./cmd

CMD ["./bin/contact-management"]