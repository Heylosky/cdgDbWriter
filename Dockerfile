FROM golang:1.19 AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build
COPY . .
RUN go build -o dbWriter .

FROM scratch
COPY --from=builder /build/dbWriter /
COPY --from=builder /build/dbWriter.log /

ENTRYPOINT ["/dbWriter"]