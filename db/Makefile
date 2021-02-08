
include ../bin/Makefile.cfg

% : %.go
	$(GO) build $<

data.sqlite: makedb questions.csv
	rm -f data.sqlite
	./makedb
