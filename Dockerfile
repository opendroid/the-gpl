# Dockerfile References: https://docs.docker.com/engine/reference/builder/
# Sample https://cloud.google.com/run/docs/quickstarts/build-and-deploy?_ga=2.91290522.-1679093051.1593441137

# Step 1: Build the application
FROM golang:1.27-alpine AS builder
LABEL maintainer="Open Web <plutoapps@outlook.com>"
# The alpine builder ships no C compiler; pin cgo off so `go test` and `go build`
# behave identically here and keep producing a static binary for the alpine stage.
ENV CGO_ENABLED=0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . ./
RUN go test ./... && GOOS=linux go build -mod=readonly -v -o the-gpl

# Step 2: Build the container image
# Use the official Alpine image (https://hub.docker.com/_/alpine) for a lean production container.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/the-gpl /the-gpl
COPY --from=builder /app/public /public
EXPOSE 8080
CMD ["/the-gpl", "server", "--port=8080"]
