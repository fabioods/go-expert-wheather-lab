PKG := $(shell go list ./...)
COVERAGE_FILE := coverage.out
COVERAGE_HTML := coverage.html

test:
	go test -v -cover -coverprofile=$(COVERAGE_FILE) $(PKG)

coverage:
	go tool cover -func=$(COVERAGE_FILE)

coverage-html: test
	go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	open $(COVERAGE_HTML)

clean:
	rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
