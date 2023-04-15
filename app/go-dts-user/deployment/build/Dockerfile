FROM golang:1.20 AS builder
WORKDIR /app/go
COPY . .
RUN ls && go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go-dts-user

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder ./go-dts-user ./
COPY --from=builder ./config/local.yaml ./config/local.yaml
COPY --from=builder ./docs ./
CMD ["./go-dts-user","-configFile=local", "-source=GORM"]  