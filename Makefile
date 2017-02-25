PREFIX:=$(shell pwd)
BINDIR:=$(PREFIX)/bin
EXE=$(BINDIR)/rudbeckia
GOSRC=$(shell find . -name "*.go")

build: $(EXE)

$(EXE): $(GOSRC)
	@mkdir -p $(BINDIR)
	go build -o $(EXE) ./cli
