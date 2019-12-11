PACKAGES = ./channel ./context ./syntax ./testing ./types ./tips...

.PHONY: all
all: lib fmt test
	@echo "All done."

test:
	@echo "testing..."
	@go test -race -cover $(PACKAGES)

fmt:
	@./fmt.sh

lint:
	@golint $(PACKAGES)

lib:
	@go mod download