ROOT = $(PWD)
TIME = $(shell time date +"%Y_%m_%d-%H:%M:%S")

.PHONY: time
time:
	@echo ${TIME}

.PHONY: build-docker
build-docker:
	docker build --build-arg FILENAME=${TIME} -t go-bench-framework . -f ./benchmarks/Dockerfile

.PHONY: build-bench-docker
build-bench-docker:
	docker build --build-arg FILENAME=${TIME} -t go-ab-bench-framework . -f ./benchmarks/ab.Dockerfile

.PHONY: run-docker
run-docker:
	docker run  -v  ${PWD}:/app --memory="4g" --cpus="4.0" go-bench-framework

.PHONY: run-ab-get
run-ab-get:
	ab -n 1000 -g $(dir)/echo-$(num).csv -c 100 http://127.0.0.1:3001/param/2
	ab -n 1000 -g $(dir)/fiber-$(num).csv -c 100 http://127.0.0.1:3002/param/2
	ab -n 1000 -g $(dir)/gin-$(num).csv -c 100 http://127.0.0.1:3003/param/2
	ab -n 1000 -g $(dir)/http-$(num).csv -c 100 http://127.0.0.1:3004/param/2

run-ab-post: run-ab-post
	ab -p ./data/post.txt -T application/json -g ./$(dir)/echo-p-$(num).csv -n 1000 -c 100 http://127.0.0.1:3001/
	ab -p ./data/post.txt -T application/json  -g ./$(dir)/fiber-p-$(num).csv -n 1000 -c 100 http://127.0.0.1:3002/
	ab -p ./data/post.txt -T application/json -g ./$(dir)/gin-p-$(num).csv -n 1000 -c 100 http://127.0.0.1:3003/
	ab -p ./data/post.txt -T application/json -g ./$(dir)/http-p-$(num).csv -n 1000 -c 100 http://127.0.0.1:3004/

