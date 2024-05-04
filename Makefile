GOX := $(shell which go)
BIN := crv
OUTDIR := dist

crv:
	mkdir -p $(OUTDIR)
	$(GOX) build \
		-v \
		-x \
		-o $(OUTDIR)/$(BIN)

clean:
	@rm -rf $(OUTDIR)
