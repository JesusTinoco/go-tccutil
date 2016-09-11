FROM golang:1.6-onbuild
COPY . /go/src/app
RUN go get -d -v
RUN go get github.com/axw/gocov/gocov
RUN go get github.com/golang/lint/golint
# Creating "sw_vers" script to fake the sw_vers of OS X. It will return a OS X version (10.11)
RUN echo "10.11" > /usr/bin/sw_vers && chmod +x /usr/bin/sw_vers