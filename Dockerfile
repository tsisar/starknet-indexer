# Minimal runtime image
FROM alpine:latest

ARG BIN_PATH=bin/indexer
ARG CONFIG_PATH=config

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY ${BIN_PATH} /app/app
COPY ${CONFIG_PATH} /app/config

ENTRYPOINT ["./app"]

EXPOSE 8040
EXPOSE 8080

ENV METRICS_ENABLED="true" \
    METRICS_PORT="8040" \
    SAVE_LOGS="false" \
    LOG_LEVEL="debug" \
    INDEXER_NAME="stablecoin"

ENTRYPOINT ["/app/app"]