build:
	cd gui && npm run build
	go build -o dist/app main.go
