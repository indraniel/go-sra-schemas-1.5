.PHONY: check-env prepare download
SOURCES=$(*.xsd)

all: gofiles

gofiles: export PATH := $(GOPATH)/bin:$(PATH)
gofiles: $(SOURCES) check-env
	bash -x get-sra-schemas.sh

test:
	@go test ./keys ./block ./transaction ./db ./git

prepare: check-env
	@echo "GOPATH is: ${GOPATH}"
	@echo "GOROOT is: ${GOROOT}"
	go get github.com/tools/godep
	godep restore
	cd $(GOPATH)/src/github.com/metaleap/go-xsd/xsd-makepkg && go install

check-env:
ifndef GOROOT
    $(error environment variable GOROOT is undefined)
endif

ifndef GOPATH
    $(error environment variable GOPATH is undefined)
endif
