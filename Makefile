PACKAGES = ./channel ./context ./syntax ./testing ./types ./tips...
OUTPUT = ./output

.PHONY: all
all: check-dir lib fmt test
	@echo "All done."

test:
	@echo "testing..."
	@mkdir -p output
	@go test -race -coverprofile ./output/coverage.out $(PACKAGES)/...

fmt:
	@./fmt.sh

lint:
	@golint $(PACKAGES)

lib:
	@go mod download

check-dir:
	@if [ ! -d $(OUTPUT) ]; then mkdir $(OUTPUT); fi

clean:
	@rm -rf $(OUTPUT)