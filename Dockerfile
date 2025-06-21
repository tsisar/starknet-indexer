# Stage 1: Build the Go application with cross-compilation
FROM --platform=$BUILDPLATFORM golang:1.24.2-alpine AS builder

ARG TARGETOS
ARG TARGETARCH
ARG COMPONENT

RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /app/app ./cmd/${COMPONENT}

# Stage 2: Minimal runtime image
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 8040
EXPOSE 8080

ENV METRICS_ENABLED="true" \
    METRICS_PORT="8040" \
    SAVE_LOGS="false" \
    LOG_LEVEL="debug"

ENTRYPOINT ["./app"]