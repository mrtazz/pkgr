#
# some housekeeping tasks
#

# variable definitions
NAME := pkgr
DESC := FreeBSD pkg creation tool
PREFIX ?= usr/local
VERSION := $(shell git describe --tags --always --dirty)
GOVERSION := $(shell go version)
BUILDTIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
BUILDDATE := $(shell date -u +"%B %d, %Y")
BUILDER := $(shell echo "`git config user.name` <`git config user.email`>")
PKG_RELEASE ?= 1
PROJECT_URL := "https://github.com/mrtazz/$(NAME)"
LDFLAGS := -X 'main.version=$(VERSION)' \
           -X 'main.builder=$(BUILDER)' \
           -X 'main.goversion=$(GOVERSION)'
GOFLAGS := -mod=vendor

PKGNG_ARCH ?= ""

PACKAGES := $(shell find ./* -type d | grep -v vendor)

CMD_SOURCES := $(shell find cmd -name main.go)
TARGETS := $(patsubst cmd/%/main.go,%,$(CMD_SOURCES))
MAN_SOURCES := $(shell find man -name "*.troff")
MAN_TARGETS := $(patsubst man/man1/%.troff,%,$(MAN_SOURCES))

INSTALLED_TARGETS = $(addprefix $(PREFIX)/bin/, $(TARGETS))
INSTALLED_MAN_TARGETS = $(addprefix $(PREFIX)/man/man1/, $(MAN_TARGETS))

MANIFEST:
	echo '{ "name": "pkgr", "version": "$(VERSION)", "comment": "create pkgng packages from directory", "desc": "create pkgng packages from directory", "maintainer": "Daniel Schauenberg <d@unwiredcouch.com>", "www": "https://github.com/mrtazz/pkgr", "arch": "$(PKGNG_ARCH)" }' > MANIFEST

%: cmd/%/main.go
	go build -ldflags "$(LDFLAGS)" -o $@ $<

%.1: man/man1/%.1.troff
	sed "s/REPLACE_DATE/$(BUILDDATE)/" $< > $@

all: $(TARGETS) $(MAN_TARGETS)
.DEFAULT_GOAL:=all

# development tasks
test:
	go test $(GOFLAGS) -v ./...

coverage:
	@echo "mode: set" > cover.out
	@for package in $(PACKAGES); do \
		if ls $${package}/*.go &> /dev/null; then  \
		go test -coverprofile=$${package}/profile.out $${package}; fi; \
		if test -f $${package}/profile.out; then \
	 	cat $${package}/profile.out | grep -v "mode: set" >> cover.out; fi; \
	done
	@-go tool cover -html=cover.out -o cover.html

benchmark:
	@echo "Running tests..."
	@go test -bench=. ${NAME}

# install tasks
$(PREFIX)/bin/%: %
	install -d $$(dirname $@)
	install -m 755 $< $@

$(PREFIX)/man/man1/%: %
	install -d $$(dirname $@)
	install -m 644 $< $@

install: $(INSTALLED_TARGETS) $(INSTALLED_MAN_TARGETS)

local-install:
	$(MAKE) install PREFIX=usr/local

# packaging tasks
packages: local-install rpm deb

deploy-packages: packages
	package_cloud push mrtazz/$(NAME)/el/7 *.rpm
	package_cloud push mrtazz/$(NAME)/debian/wheezy *.deb
	package_cloud push mrtazz/$(NAME)/ubuntu/trusty *.deb


clean: clean-docs
	$(RM) -r ./usr
	$(RM) $(TARGETS)

clean-docs:
	$(RM) $(MAN_TARGETS)

.PHONY: all test rpm deb install local-install packages govendor coverage clean-deps clean clean-docs pizza
