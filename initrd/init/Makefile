build: rc.init

rc.init: main.go
	GOOS=linux GOARCH=amd64 go build -o ../rc.init

clean:
	@rm -f ../rc.init rc.init init
