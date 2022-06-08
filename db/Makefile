
include ../bin/Makefile.cfg

all:
	@echo
	@echo Usage: make new
	@echo
	@echo WARNING: this will destroy all previously collected data
	@echo

new: data.sqlite

data.sqlite: makedb.go questions.csv
	rm -f data.sqlite
	$(GO) fmt .
	$(GO) run -v $<
