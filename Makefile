TARGET   ?= hqgoutils
GO       ?= go
GOFLAGS  ?= 

fmt:
	$(GO) $(GOFLAGS) fmt ./...

test:
	$(GO) $(GOFLAGS) test ./...