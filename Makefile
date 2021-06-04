.PHONY: build clean

build:
	@go mod download && \
 	go build -o candy-fight main.go && \
    chmod +x candy-fight && \
    echo "candy-fight has been built."

clean:
	@rm -f candy-fight

setup:
	go install honnef.co/go/tools/cmd/staticcheck github.com/swaggo/swag/cmd/swag
	git config core.hooksPath .githooks