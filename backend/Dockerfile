FROM golang:1.23.4

ADD . /backend

WORKDIR /backend
RUN go mod init project
RUN go mod tidy
COPY db/.env .env
RUN go build -o project .
RUN go test ./...
RUN chmod +x /backend

ENTRYPOINT ["/backend/project"]
