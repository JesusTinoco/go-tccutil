FROM golang:1.6
# Creating "sw_vers" script to fake the sw_vers of OS X. It will return a OS X version (10.11)
RUN echo '#!/bin/bash\n\
 echo "10.10.5"' > /usr/bin/sw_vers && chmod +x /usr/bin/sw_vers
RUN go get github.com/axw/gocov/gocov && \
	go get github.com/golang/lint/golint && \
	go get github.com/hashicorp/go-version && \
	go get github.com/mattn/go-sqlite3 && \
	go get github.com/spf13/pflag && \
	go get github.com/spf13/cobra
WORKDIR /usr/src/tccutil2