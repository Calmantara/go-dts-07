run:
	go run main.go -configFile=local -source=GORM
build-image: mock test swagger
	./deployment/script/build.sh
mock:
	echo "mock"
swagger:
	swag init --parseDependency -q -g ./server/http.go
test:
	go test -race -short -coverprofile=./cov.out ./...
