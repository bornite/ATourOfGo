GOCMD=go
GOTEST=$(GOCMD) test
TESTOPT=

test:
	$(GOTEST) $(TESTOPT) hello_test.go hello.go
	$(GOTEST) $(TESTOPT) packages_test.go packages.go
	$(GOTEST) $(TESTOPT) imports_test.go imports.go
	$(GOTEST) $(TESTOPT) functions_test.go functions.go
	$(GOTEST) $(TESTOPT) multiple-results_test.go multiple-results.go
	$(GOTEST) $(TESTOPT) named-results_test.go named-results.go
	$(GOTEST) $(TESTOPT) variables_test.go variables.go
	$(GOTEST) $(TESTOPT) variables-with-initializers_test.go variables-with-initializers.go
	$(GOTEST) $(TESTOPT) short-variable-declarations_test.go short-variable-declarations.go
	$(GOTEST) $(TESTOPT) basic-types_test.go basic-types.go
	$(GOTEST) $(TESTOPT) zero_test.go zero.go
