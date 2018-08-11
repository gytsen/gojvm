PKG := github.com/gytsen/gojvm
OUTDIR := output


all: ijvm

$(OUTDIR):
	mkdir $(OUTDIR)

ijvm: $(OUTDIR)
	go build -o $(OUTDIR)/$@ $(PKG)/cmd/ijvm 

.PHONY: clean

clean:
	rm -f output/*
	rmdir output
