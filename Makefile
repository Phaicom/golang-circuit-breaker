NAME     := golang-circuit-breaker
VERSION  := v0.1.0

.PHONY: clean
clean:
	@echo "Cleaning up..."
	cd ./downstream && \
	rm -rf bin/* \
	rm -rf vendor/*
	cd ./upstream && \
	rm -rf bin/* \
	rm -rf vendor/*

.PHONY: dep
dep:
	@echo "Checking dep..."
ifeq ($(shell command -v dep 2> /dev/null),)
	@echo "Install dep..."
	go get -u -v github.com/golang/dep/cmd/dep
endif

.PHONY: deps
deps: dep
	@echo "Install dependencie..."
	cd ./downstream && \
	dep ensure -v
	cd ./upstream && \
	dep ensure -v