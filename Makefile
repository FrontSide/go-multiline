VERSION := $(shell grep -Po "\d+[.]\d+[.]\d+" version.go)
PACKAGE := "github.com/frontside/multiline@v$(VERSION)"

release:
	@echo Releasing $(PACKAGE)
	-git tag v$(VERSION)
	-git push origin v$(VERSION)
	GOPROXY=proxy.golang.org go list -m $(PACKAGE)

