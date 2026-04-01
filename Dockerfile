FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go mod download && CGO_ENABLED=0 go build -o mailbag ./cmd/mailbag/

FROM alpine:3.19
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/mailbag .
ENV PORT=9230 DATA_DIR=/data
EXPOSE 9230
CMD ["./mailbag"]
