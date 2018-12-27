MAIN_PACKAGE = $(dir $(shell grep -ir -l --exclude-dir vendor --exclude Makefile "func main()" ./*))
REPO = $(notdir $(CURDIR))
VERSION = $(shell grep 'Version string' $(MAIN_PACKAGE)/version.go | sed -E 's/.*"(.+)"$$/\1/')
COMMIT = $(shell git describe --always)
PACKAGES = $(shell go list ./... | grep -v '/vendor/')

GOX_OS = darwin linux windows
GOX_ARCH = amd64 386

default: test

build: go-generate
	@cd $(MAIN_PACKAGE) ; \
	gox \
	  -ldflags "-X main.GitCommit=$(COMMIT)" \
	  -os="$(firstword $(GOX_OS))" \
	  -arch="$(firstword $(GOX_ARCH))" \
	  -output="$(CURDIR)/pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

test: go-generate
	@echo "" > coverage.txt
	@for d in ${PACKAGES} ; do \
	  go test ${VERBOSE} -race -coverprofile=profile.out -covermode=atomic $$d ; \
	done
	@if [ -f profile.out ]; then \
	  cat profile.out >> coverage.txt ; \
	  rm profile.out ; \
	fi

test-cpu: go-generate
	go test ${VERBOSE} -cpu 1,2,4 ${PACKAGES}

test-race: go-generate
	go test ${VERBOSE} -race ${PACKAGES}

vet: go-generate
	go vet ${PACKAGES}

clean:
	@rm -fr ./pkg
	@rm -fr ./dist/$(VERSION)

install: clean build
	cp "$(CURDIR)/pkg/$(firstword $(GOX_OS))_$(firstword $(GOX_ARCH))/$(REPO)" "${GOPATH}/bin"

package: clean go-generate
	@cd $(MAIN_PACKAGE) ; \
	gox \
	  -ldflags "-X main.GitCommit=$(COMMIT)" \
	  -parallel=3 \
	  -os="$(GOX_OS)" \
	  -arch="$(GOX_ARCH)" \
	  -output="$(CURDIR)/pkg/{{.OS}}_{{.Arch}}/{{.Dir}}"

	@mkdir -p ./dist/$(VERSION)

	@for platform in $(foreach os,$(GOX_OS),$(foreach arch,$(GOX_ARCH),$(os)_$(arch))) ; do \
	  echo "zip ../../dist/$(VERSION)/$(REPO)_$(VERSION)_$$platform.zip ./*" ; \
	  (cd ./pkg/$$platform && zip ../../dist/$(VERSION)/$(REPO)_$(VERSION)_$$platform.zip ./*) ; \
	done

	@cd ./dist/$(VERSION) ; \
	echo "shasum -a 256 * > ./$(VERSION)_SHASUMS" ; \
	shasum -a 256 * > ./$(VERSION)_SHASUMS

release:
	ghr $(VERSION) ./dist/$(VERSION)

fmt:
	gofmt -w .

dep:
	dep ensure

go-generate:
	go generate ${VERBOSE} ${PACKAGES}

certificate:
	@rm -fr ./cert
	@mkdir -p ./cert
	openssl genrsa -out ./cert/server.key 2048
	openssl req -new -x509 -sha256 \
	  -key ./cert/server.key -out ./cert/server.crt -days 3650
	sudo security add-trusted-cert -d \
	  -r trustRoot -k /Library/Keychains/System.keychain \
	  ./cert/server.crt

.PHONY: build test test-cpu test-race vet clean install package release fmt dep go-generate certificate
