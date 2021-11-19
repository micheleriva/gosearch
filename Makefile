run:
	go run .

clean:
	rm -rf ./bin

compile:
	go build -o bin/lepus .

compile-all:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=386 go build -o bin/lepus-freebsd-386 .
	GOOS=linux GOARCH=386 go build -o bin/lepus-linux-386 .
	GOOS=windows GOARCH=386 go build -o bin/lepus-windows-386 .