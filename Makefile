build/container: stage/jadvisor Dockerfile
	docker build --no-cache -t jadvisor .
	touch build/container

build/jadvisor: *.go */*.go
	godep go build -o build/jadvisor

stage/jadvisor: build/jadvisor
	mkdir -p stage
	cp build/jadvisor stage/jadvisor

release:
	docker tag jadvisor fabric8/jadvisor
	docker push fabric8/jadvisor

.PHONY: clean
clean:
	rm -rf build
