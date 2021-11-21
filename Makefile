CMD=go
OUTPUT=lepus
BUILD=$(CMD) build
GOARCH=386

run:
	$(CMD) run .

deps:
	$(CMD) mod download

test:
	$(CMD) test .

clean:
	$(CMD) -rf ./bin

compile:
	$(CMD) build -o bin/lepus .

compile-all:
	echo "Compiling for every OS and Platform"
	GOOS=freebsd GOARCH=$(GOARCH) $(BUILD) -o bin/$(OUTPUT)-freebsd-$(GOARCH) .
	GOOS=linux GOARCH=$(GOARCH) $(BUILD) -o bin/$(OUTPUT)-linux-$(GOARCH) .
	GOOS=windows GOARCH=$(GOARCH) $(BUILD) -o bin/$(OUTPUT)-windows-$(GOARCH) .