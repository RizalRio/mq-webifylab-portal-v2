# ─── Run stage (~25MB final image) ───
FROM alpine:3.20
RUN apk add --no-cache ca-certificates tzdata   # wajib: HTTPS ke Telegram API
COPY webifylab-api /server
EXPOSE 8080
USER 1000:1000
ENTRYPOINT ["/server"]