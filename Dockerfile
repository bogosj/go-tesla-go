FROM golang:1.15.4-buster AS builder
RUN go get github.com/bogosj/go-tesla-go
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go-tesla-go github.com/bogosj/go-tesla-go

FROM alpine:3.12.1
RUN apk --no-cache add ca-certificates tzdata
COPY --from=builder /go-tesla-go /
ENTRYPOINT ["/go-tesla-go"]
