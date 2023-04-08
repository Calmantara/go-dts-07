# FROM golang:1.20 AS builder
# WORKDIR /app/go
# COPY . .
# RUN ls && go mod tidy
# RUN CGO_ENABLED=0 go build -o go-dts-user

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY ./go-dts-user ./
COPY ./config/local.yaml ./config/local.yaml
COPY ./docs ./
CMD ["./go-dts-user","-configFile=local", "-source=GORM"]  