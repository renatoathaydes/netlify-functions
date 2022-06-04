GO:=go

.PHONY: tidy
tidy:
	cd analytics && $(GO) mod tidy

.PHONY: test
test: tidy
	cd analytics && $(GO) test .

../netlify/functions/analytics: test
	cd analytics && $(GO) build -o ../netlify/functions/analytics .

all: ../netlify/functions/analytics

.PHONY: clean
clean:
	rm -rf netlify
