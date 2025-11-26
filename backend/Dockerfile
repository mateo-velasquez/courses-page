FROM golang:1.24

ADD . /backend

WORKDIR /backend
RUN go mod init project
RUN go mod tidy

# Copy environment files
COPY db/.env .env

RUN go build -o project .
RUN chmod +x /backend/project

ENTRYPOINT ["/backend/project"]
