FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache 'git=~2'

# Install dependencies
ENV GO111MODULE=on
WORKDIR $GOPATH/src/packages/cycling-tracker-server/
COPY . .

# Fetch dependencies.
# Using go get.
RUN go get -d -v

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/cycling-tracker-server .

############################
FROM alpine:3

WORKDIR /

# Copy our static executable.
COPY --from=builder /go/cycling-tracker-server /go/cycling-tracker-server
#COPY public /go/public
COPY .env /go/.env

ENV PORT 8080
ENV PORT 8081
ENV GIN_MODE release
EXPOSE 8080
EXPOSE 8081

WORKDIR /go

# Run the Go Gin binary.
ENTRYPOINT ["/go/cycling-tracker-server"]