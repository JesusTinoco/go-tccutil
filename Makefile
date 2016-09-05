init:
		go get -t -v ./...
		go get github.com/golang/lint/golint

get-deps: init
		go get github.com/axw/gocov/gocov

get-travis-deps: init
		go get github.com/mattn/goveralls
		go get golang.org/x/tools/cmd/cover

check: vet test lint coverage

travis-check: vet test lint

vet:
		go vet ./...

test:
		go test -v -race ./...

lint:
		golint -set_exit_status ./...

coverage:
		gocov test ./... | gocov report

goveralls:
		cd ./stacksmith && $(HOME)/gopath/bin/goveralls -service=travis-ci%
