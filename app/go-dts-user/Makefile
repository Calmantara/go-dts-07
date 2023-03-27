run:
	go run main.go -configFile=local -source=GORM
build-image:
	docker build -t go-dts-user -f deployment/build/Dockerfile .
mock:
	echo "mock"
swagger:
	echo "swagger"
test:
	go test -race -short -coverprofile=./cov.out ./...
