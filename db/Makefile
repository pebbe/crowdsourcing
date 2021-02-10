
include ../bin/Makefile.cfg

% : %.go
	$(GO) build -v $<

data.sqlite: makedb questions.csv
	rm -f data.sqlite
	./makedb
