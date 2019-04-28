GOCMD=go
GOTEST=$(GOCMD) test
TESTOPT=

test:
	$(GOTEST) $(TESTOPT) hello_test.go hello.go
	$(GOTEST) $(TESTOPT) packages_test.go packages.go
	$(GOTEST) $(TESTOPT) imports_test.go imports.go
	$(GOTEST) $(TESTOPT) functions_test.go functions.go
