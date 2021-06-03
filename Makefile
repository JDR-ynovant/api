.PHONY: build clean

build:
	@go mod download && \
 	go build -o candy-fight main.go && \
    chmod +x candy-fight && \
    echo "candy-fight has been built."

clean:
	@rm -f candy-fight