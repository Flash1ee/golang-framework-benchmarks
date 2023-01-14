ROOT = $(PWD)
TIME = $(shell time date +"%Y_%m_%d-%H:%M:%S")

.PHONY: time
time:
	@echo ${TIME}

.PHONY: build-docker
build-docker:
	docker build --build-arg FILENAME=${TIME} -t go-bench-framework . -f ./benchmarks/Dockerfile

.PHONY: run-docker
run-docker:
	docker run  -v  ${PWD}/benchmarks:/app --memory="4g" --cpus="4.0" go-bench-framework