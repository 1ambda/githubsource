.PHONY: install

VERSION=0.0.1
BUILD_TIME=`date +%FT%T%z`

install:
	curl https://glide.sh/get | sh
	glide install
	go get github.com/onsi/ginkgo/ginkgo
	go get github.com/onsi/gomega

test:
	ginkgo


