test:
	cd analytics && go test .

all: test
	cd analytics && go build -o ../netlify/functions/analytics .

clean:
	rm -rf netlify
