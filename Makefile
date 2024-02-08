GO ?= go
GOTAGS ?=

.PHONY: all
all:
	$(GO) generate -v -x ./...
	$(GO) mod tidy
	$(GO) build -v --tags=$(GOTAGS)

COVERAGE ?= /tmp/io.github.black-desk.busagent-test/coverage.out

.PHONY: test
test:
	mkdir -p $(shell dirname $(COVERAGE))
	$(GO) test ./... --tags=$(GOTAGS) -v --ginkgo.vv -coverprofile=$(COVERAGE) \

PREFIX ?= /usr/local
DESTDIR ?=

.PHONY: install
install:
	install -m755 -D busagent \
		$(DESTDIR)$(PREFIX)/bin/busagent

COVERAGE_REPORT ?= /tmp/io.github.black-desk.busagent-test/coverage.txt

.PHONY: test-coverage
test-coverage:
	go tool cover -func=$(COVERAGE) -o=$(COVERAGE_REPORT)
