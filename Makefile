all: lint vet test coverage

test:
		go test -v -race ./...

coverage:
		gocov test ./... | gocov report

vet:
		go vet ./...

lint:
		golint -set_exit_status ./...

build:
		go build -v ./...

docker-all: docker-build docker-tccutil-test docker-tccutil-build docker-rmi

docker-build:
		docker build --no-cache -t docker-tccutil .

docker-tccutil-test:
		docker run --rm -v ${PWD}:/usr/src/tccutil2 docker-tccutil bash -c "go get -d && make all"

docker-tccutil-build:
		docker run --rm -e GOOS=darwin -e GOARCH=amd64 -v ${PWD}:/usr/src/tccutil2 docker-tccutil bash -c "go get -d && go build -gccgoflags"
 
docker-rmi:
		docker rmi docker-tccutil 