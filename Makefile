######################################################################
# @author      : annika (annika@berlin.ccc.de)
# @file        : Makefile
# @created     : Sunday Aug 16, 2020 19:24:54 CEST
######################################################################

all: pantraserver

pantraserver:
	cd cmd/pantra_server && go build

.PHONY:
	clean test

run:
	go run cmd/pantra_server/main.go

test:
	cd pkg/pantra_server/model/expkey && go test -v
	cd pkg/pantra_server/expkeyservice && go test -v

clean:
	rm -f cmd/b3scaled/b3scaled
	rm -f cmd/b3scalectl/b3scalectl

