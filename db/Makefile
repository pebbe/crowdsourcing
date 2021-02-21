
include ../bin/Makefile.cfg

% : %.go
	$(GO) build $(OPTS) $<

data.sqlite: makedb questions.csv
	rm -f data.sqlite
	./makedb
