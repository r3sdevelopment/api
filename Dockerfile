# Use go 1.x based on alpine image.
FROM golang:1.17.0-alpine AS build

# Install build tools.
RUN apk add --update gcc musl-dev

# Cache dependencies
COPY . /go/src/r3s.dev/api
ENV GO111MODULE on
WORKDIR /go/src/r3s.dev/api
RUN go mod download

RUN go build \
        -o /go/src/r3s.dev/api/bin/api-r3s-dev \
        -v -x \
        cmd/*.go

###
FROM alpine:3.14.2

# Make sure /etc/hosts is resolved before DNS
RUN echo "hosts: files dns" > /etc/nsswitch.conf

COPY --from=build /go/src/r3s.dev/api/bin/api-r3s-dev /usr/local/bin/api-r3s-dev

# Add non-privileged user
RUN adduser -D -u 1001 appuser && \
        chown appuser:appuser /usr/local/bin/api-r3s-dev && \
        chmod +x /usr/local/bin/api-r3s-dev
USER appuser

CMD ["/usr/local/bin/api-r3s-dev"]

