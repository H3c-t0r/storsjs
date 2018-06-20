.PHONY: test lint proto check-copyrights build-dev-deps

TAG    := $$(git rev-parse --short HEAD)

lint: check-copyrights
	@echo "Running ${@}"
	@gometalinter \
	--deadline=170s \
	--disable-all \
	--enable=golint \
	--enable=goimports \
	--enable=vet \
	--enable=deadcode \
	--enable=goconst \
	--exclude=.*\.pb\.go \
	--exclude=.*_test.go \
	./...

check-copyrights:
	@echo "Running ${@}"
	@./scripts/check-for-header.sh

# Applies goimports to every go file (excluding vendored files)
goimports-fix:
	goimports -w $$(find . -type f -name '*.go' -not -path "*/vendor/*")

proto:
	@echo "Running ${@}"
	./scripts/build-protos.sh

build-dev-deps:
	go get github.com/golang/protobuf/protoc-gen-go
	go get github.com/alecthomas/gometalinter
	gometalinter --install --force

test: lint
	go install -v ./...
	go test ./...
	@echo done

build-binaries:
	docker build -t overlay .

run-overlay:
	docker network create test-net

	docker run -d \
		--name redis \
		--network test-net \
		-p 127.0.0.1:6379:6379 \
		redis

	docker run -d \
		--name=overlay \
		--network test-net \
		-e REDIS_ADDRESS=redis:6379 \
		-e REDIS_PASSWORD="" \
		-e REDIS_DB=1 \
		-e OVERLAY_PORT=8080 \
		overlay

clean-local:
	# cleanup overlay
	docker stop overlay || true
	docker rm overlay || true
	# cleanup redis
	docker stop redis || true
	docker rm redis || true
	# cleanup docker network
	docker network rm test-net || true

images:
	docker build -t overlay:${TAG} .
	docker tag overlay:${TAG} overlay:latest

push-images:
	docker push storjlabs/overlay