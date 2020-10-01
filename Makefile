test:
	go test -coverprofile=cover.out
cover: test
	go tool cover -html=cover.out -o cover.html
