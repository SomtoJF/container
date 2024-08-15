.PHONY: build_docker run_container run

run: run_container

build_docker:
	docker build -t gocker .

run_container:
	docker run -it --privileged gocker
