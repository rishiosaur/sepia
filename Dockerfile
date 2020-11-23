FROM golang:1.15-alpine as builder

# Meta data:
LABEL maintainer="itsrishikothari@gmail.com"
LABEL description="A minimal interpreted programming language."

# Copying over all files:
COPY . /usr/src/app/
WORKDIR /usr/src/app

# Installing deps and building binary:
RUN go get -v -t -d ./...
RUN go build -o sepia .

# Copying over the binary to a thinner image
# hadolint ignore=DL3006,DL3007
FROM alpine:latest
COPY --from=builder /usr/src/app/sepia/ /usr/bin/

