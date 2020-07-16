run:
	./bin/bench -remotes=127.0.0.1 -output results/result_`date +"%Y%m%d_%H%M%S"`.json

all: build

.PHONY: build
build:
	GOPATH=`pwd`:`pwd`/vendor go install ./src/cmd/...

.PHONY: race
race:
	GOPATH=`pwd`:`pwd`/vendor go install -race ./src/cmd/...
