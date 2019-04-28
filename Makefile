GOCMD=go
GOTEST=$(GOCMD) test

test:
	$(GOTEST) -v hello_test.go hello.go
	$(GOTEST) -v packages_test.go packages.go
