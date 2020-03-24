############################
# STEP 1 build executable
############################
FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GOOS=linux GOARCH=amd64

# Add Maintainer Info
LABEL maintainer="differz"

# Build Args
ARG PACKAGE=github.com/differz/perstorage

# Install git
RUN apk update && apk add --no-cache git gcc libc-dev
WORKDIR $GOPATH/src/${PACKAGE}
COPY . .

# Fetch dependencies
RUN go get -d -v

# Build and install the binary
RUN go install -ldflags="-w -s" ${PACKAGE}

############################
# STEP 2 build a small image
############################
# FROM scratch doesn't work
# https://github.com/golang/go/issues/28152
FROM alpine

ENV GOPATH=/go

ARG PACKAGE=github.com/differz/perstorage

LABEL maintainer="differz"

RUN apk --no-cache add ca-certificates

# Copy our static executable
WORKDIR /storage/file/migrations/
COPY --from=builder $GOPATH/src/${PACKAGE}/storage/file/migrations/* ./
COPY --from=builder $GOPATH/bin/perstorage /perstorage

WORKDIR /

# Export necessary port
EXPOSE 8443

# Run the perstorage binary only via !/bin/sh
ENTRYPOINT ["./perstorage"]
