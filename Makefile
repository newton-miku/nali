NAME=nali
BINDIR=bin
VERSION=$(shell git describe --tags || echo "unknown version")
GOBUILD=CGO_ENABLED=0 go build -trimpath -ldflags '-X "github.com/zu1k/nali/internal/constant.Version=$(VERSION)" -w -s'

	$(GOBUILD) -o $(BINDIR)/$(NAME)-$@
	rm $(BINDIR)/*
