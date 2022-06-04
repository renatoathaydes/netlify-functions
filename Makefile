## To use a different version of Go, run like this:
##    make all GO=go1.16.5
GO:=go

.PHONY: setup
setup:
	cd analytics && rm go.mod go.sum && $(GO) mod init renato/analytics && $(GO) mod tidy

.PHONY: test
test:
	cd analytics && $(GO) test .

../netlify/functions/analytics: test
	cd analytics && $(GO) build -o ../netlify/functions/analytics .

all: ../netlify/functions/analytics

.PHONY: clean
clean:
	rm -rf netlify
