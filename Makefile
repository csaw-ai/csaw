GO ?= go

.PHONY: fmt test vet docs-check check

fmt:
	$(GO) fmt ./...

test:
	$(GO) test ./...

vet:
	$(GO) vet ./...

docs-check:
	./scripts/check_docs.sh

check: fmt test vet docs-check
