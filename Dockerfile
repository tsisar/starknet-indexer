# Minimal runtime image
FROM alpine:latest

ARG BIN_PATH=bin/indexer

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

COPY ${BIN_PATH} /app/app

EXPOSE 8040
EXPOSE 8080

ENV METRICS_ENABLED="true" \
    METRICS_PORT="8040" \
    SAVE_LOGS="false" \
    LOG_LEVEL="debug"

ENTRYPOINT ["./app"]