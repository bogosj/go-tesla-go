FROM golang:1.21-bookworm AS builder
WORKDIR /gtg
COPY . /gtg
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /gtg/

FROM alpine:3.18
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /gtg/go-tesla-go /
ENTRYPOINT ["/go-tesla-go"]
