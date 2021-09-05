############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
ARG DOCKER_TAG=0.0.0
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git curl unzip
WORKDIR /builder
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /resolver -ldflags="-s -w -extldflags \"-static\" -X main.Version=$DOCKER_TAG"
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /resolver /
# Copy the temlates folder
# COPY templates /templates
# Run the hello binary.
ENTRYPOINT [ "/resolver" ]
CMD [ "" ]
