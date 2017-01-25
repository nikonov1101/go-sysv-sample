default:
	go run main.go

linux:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s" -o simple_app_linux64
