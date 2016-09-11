docker-test:
		docker build -t docker-tccutil .
		docker run --rm docker-tccutil go vet ./... && golint -set_exit_status ./... && go test -v -race ./... && gocov test ./... | gocov report
		docker rmi docker-tccutil
vet:
		go vet ./...

lint:
		golint -set_exit_status ./...