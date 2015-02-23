.PHONY: check-env prepare download clean tidy

SOURCES=$(*.xsd)
GODEP := $(GOPATH)/bin/godep

XSD_DIR := xsd
TOOLS_DIR := tools

all: gofiles

gofiles: export PATH := $(GOPATH)/bin:$(PATH)
gofiles: $(SOURCES) check-env
	bash -x $(TOOLS_DIR)/get-sra-schemas.sh

prepare: check-env
	@echo "GOPATH is: ${GOPATH}"
	@echo "GOROOT is: ${GOROOT}"
	go get github.com/tools/godep
	$(GODEP) restore
	cd $(GOPATH)/src/github.com/metaleap/go-xsd/xsd-makepkg && go install

check-env:
ifndef GOROOT
    $(error environment variable GOROOT is undefined)
endif

ifndef GOPATH
    $(error environment variable GOPATH is undefined)
endif

clean:
	rm *.xsd 2>/dev/null
	rm -rf SRA.*

tidy:
	test -d $(XSD_DIR) || mkdir -p $(XSD_DIR)
	mv *.xsd $(XSD_DIR)
