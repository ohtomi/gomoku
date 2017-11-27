MAIN_PACKAGE = $(dir $(shell grep -ir -l --exclude-dir vendor --exclude Makefile "func main()" ./*))
REPO = $(notdir $(CURDIR))
VERSION = $(shell grep 'Version string' $(MAIN_PACKAGE)/version.go | sed -E 's/.*"(.+)"$$/\1/')
COMMIT = $(shell git describe --always)
PACKAGES = $(shell go list ./... | grep -v '/vendor/')

GOX_OS = darwin linux windows
GOX_ARCH = amd64 386

default: test

build: go-assets-builder
	@cd $(MAIN_PACKAGE) ; \
	gox \
	  -ldflags "-X main.GitCommit=$(COMMIT)" \
	  -os="$(firstword $(GOX_OS))" \
	  -arch="$(firstword $(GOX_ARCH))" \
	  -output="$(CURDIR)/pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

test:
	go test -v -parallel=4 ${PACKAGES}

test-race:
	go test -v -race ${PACKAGES}

vet:
	go vet ${PACKAGES}

clean:
	@rm -fr ./pkg
	@rm -fr ./dist/$(VERSION)

package: clean go-assets-builder
	@cd $(MAIN_PACKAGE) ; \
	gox \
	  -ldflags "-X main.GitCommit=$(COMMIT)" \
	  -parallel=3 \
	  -os="$(GOX_OS)" \
	  -arch="$(GOX_ARCH)" \
	  -output="$(CURDIR)/pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

	@mkdir -p ./dist/$(VERSION)

	@for platform in $(foreach os,$(GOX_OS),$(foreach arch,$(GOX_ARCH),$(os)_$(arch))) ; do \
	  echo "zip ../../dist/$(VERSION)/xxx_$(VERSION)_$$platform.zip ./*" ; \
	  (cd ./pkg/$$platform && zip ../../dist/$(VERSION)/$(REPO)_$(VERSION)_$$platform.zip ./*) ; \
	done

	@cd ./dist/$(VERSION) ; \
	echo "shasum -a 256 * > ./$(VERSION)_SHASUMS" ; \
	shasum -a 256 * > ./$(VERSION)_SHASUMS

release: package
	ghr $(VERSION) ./dist/$(VERSION)

fmt:
	gofmt -w .

go-assets-builder:
	+$(MAKE) -C server

.PHONY: build test test-race vet clean package release fmt go-assets-builder