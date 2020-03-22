############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git
# Git is required for fetching the dependencies
RUN apk update && apk add --no-cache bash git gcc libc-dev
WORKDIR $GOPATH/src/github.com/differz/perstorage/
COPY . .
# Fetch dependencies
# Using go get
RUN go get -d -v
# Build the binary
RUN go build -o /go/bin/perstorage

############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable
COPY --from=builder /go/bin/perstorage /go/bin/perstorage
# Run the perstorage binary
ENTRYPOINT ["/go/bin/perstorage"]
