FROM golang:1.24

ADD . /backend

WORKDIR /backend
RUN go mod init project
RUN go mod tidy

# Copy environment files
COPY .env .env
COPY db/.env db/.env

RUN go build -o project .
RUN go test ./...
RUN chmod +x /backend

ENTRYPOINT ["/backend/project"]
