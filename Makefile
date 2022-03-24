build:
	env GOOS=linux GOARCH=arm go build

run-create:
	go run main.go create
