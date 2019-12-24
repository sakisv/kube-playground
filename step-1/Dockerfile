FROM golang:1.13-alpine AS builder
COPY main.go .
RUN CGO_ENABLED=0 go build -o miniserver .

FROM alpine
COPY --from=builder /go/miniserver .
CMD ["./miniserver"]
