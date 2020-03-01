
FROM golang:alpine AS builder
MAINTAINER LamTNB <baolam0307@gmail.com>
# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .
# Build the application

RUN go build -o main .

FROM alpine:3.10
RUN mkdir /app
COPY --from=builder /build/main /app/main
COPY --from=builder /build/configs /app/configs
COPY --from=builder /build/public /app/public
COPY --from=builder /build/views /app/views

RUN apk add --no-cache -U tzdata bash ca-certificates \
    && update-ca-certificates \
    && cp /usr/share/zoneinfo/Asia/Ho_Chi_Minh /etc/localtime \
    && chmod 711 /app/main \
    && rm -rf /var/cache/apk/*

WORKDIR /app
CMD ["/app/main"]
#ENTRYPOINT /app/main