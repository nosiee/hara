BINARY_NAME=hara

build:
	go build -o bin/${BINARY_NAME} cmd/hara/main.go
	go build -o bin/migrations cmd/migrations/migrations.go

run:
	bin/${BINARY_NAME} --config ../configs/api_dev.toml

test:
	go test -v -cover ./...

testscover:
	go test -coverprofile testdata/tests_cover.out ./...
	go tool cover -html=testdata/tests_cover.out

totalcover:
	go test -coverprofile testdata/tests_cover.out ./...
	go tool cover -func testdata/tests_cover.out

clean:
	go clean
	rm testdata/tests_cover.out
	rm bin/${BINARY_NAME}
	rm bin/migrations
