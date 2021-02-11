FROM golang:1.15-buster AS builder
WORKDIR /gtg
COPY . /gtg
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /gtg/

FROM alpine:3.13
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /gtg/go-tesla-go /
ENTRYPOINT ["/go-tesla-go"]
