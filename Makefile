build:
	env GOOS=linux GOARCH=arm go build

run-create:
	go run main.go create

run-details:
	go run main.go details
