.PHONY: build_docker run_container run

build_docker:
	docker build -t gocker .

run_container:
	docker run --privileged -it gocker run $(ARGS)

run: build_docker run_container