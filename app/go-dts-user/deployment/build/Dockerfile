FROM golang:1.20 AS builder
WORKDIR /app/go
COPY . .
RUN ls && go mod tidy
RUN CGO_ENABLED=0 go build -o go-dts-user

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/go/go-dts-user ./
COPY --from=builder /app/go/config/local.yaml ./config/local.yaml
CMD ["./go-dts-user"]  