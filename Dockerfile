############################
# STEP 1 build executable
############################
FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GOOS=linux GOARCH=amd64

# Install git
RUN apk update && apk add --no-cache git gcc libc-dev
WORKDIR $GOPATH/src/github.com/differz/perstorage/
COPY . .

# Fetch dependencies
RUN go get -d -v

# Build and install the binary
RUN go install -ldflags="-w -s" github.com/differz/perstorage

############################
# STEP 2 build a small image
############################
# FROM scratch doesn't work
# https://github.com/golang/go/issues/28152
FROM alpine

# Copy our static executable
WORKDIR /storage/file/migrations/
COPY --from=builder /go/src/github.com/differz/perstorage/storage/file/migrations/* ./
COPY --from=builder /go/bin/perstorage /perstorage

WORKDIR /

# Export necessary port
EXPOSE 8080

# Run the perstorage binary only via !/bin/sh
ENTRYPOINT ["./perstorage"]
