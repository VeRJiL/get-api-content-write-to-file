SOURCES := $(wildcard *.go cmd/*/*.go)

VERSION=$(shell git describe --tags --lonng --dirty 2>/dev/null)

## we must have tagged the repo at least once for VERSION to work

ifeq ($(VERSION),)
	VERSION + UNKOWN
endif

get-api-content-write-to-file: $(SOURCES)
	go build -ldflags "-X main.version=${VERSION}" -o $@ ./cmd/sort

.PHONY: lint
lint:
	golangci-lint run

.PHONY: commited
commited:
	@git diff --exit-code > /dev/null || (echo "** COMMIT YOUR CHANGES FIRST **"; exit 1)

docker: $(SOURCES) build/Dockerfile
	docker build -t sort-anim:latest . -f build/Dockerfile --build-arg VERSION=$(VERSION)

.PHONY: publish
publish: commited lint
	make docker
	docker tag get-api-content-write-to-file:latest VeRJiL/get-api-content-write-to-file:$(VERSION)
	docker push VeRJiL/get-api-content-write-to-file:$(VERSION)
