BUILDPATH := $(CURDIR)/build
PKGNAME := dummy-backend

.PHONY: build
build:
	@mkdir -p $(BUILDPATH)
	@CGO_ENABLED=0 go build -mod=vendor -ldflags -s -o $(BUILDPATH)/$(PKGNAME) .