######################################################################
# @author      : annika (annika@berlin.ccc.de)
# @file        : Makefile
# @created     : Sunday Aug 16, 2020 19:24:54 CEST
######################################################################

all: pantraserver

pantraserver: build

build: builddocs
	cd cmd/pantra_server && go build
	mkdir -p ./bin
	mv cmd/pantra_server/pantra_server ./bin
	chmod 755 bin/pantra_server

builddocs:
	cd cmd/pantra_server && swag init

.PHONY:
	clean test

run:
	go run cmd/pantra_server/main.go

runbin:
	./bin/pantra_server

test:
	cd pkg/pantra_server/model/expkey && go test -v
	cd pkg/pantra_server/expkeyservice && go test -v

gitupdate:
	git pull

clean:
	rm -f bin/*
	rm -f cmd/pantra_server/docs/*


refreshrun: clean gitupdate build runbin
