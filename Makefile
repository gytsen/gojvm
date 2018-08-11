PKG := github.com/gytsen/gojvm
OUTDIR := output


all: ijvm

$(OUTDIR):
	mkdir $(OUTDIR)

ijvm: $(OUTDIR)
	go build -o $(OUTDIR)/$@ $(PKG)/cmd/$@ 

test-stack: 
	go test $(PKG)/pkg/stack

.PHONY: clean

clean:
	rm -f output/*
	rmdir output
