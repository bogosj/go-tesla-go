build: check-version
	docker build -t ghcr.io/bogosj/go-tesla-go:latest .
	docker tag ghcr.io/bogosj/go-tesla-go:latest ghcr.io/bogosj/go-tesla-go:${GTG_VERSION}

deploy: build
	docker push ghcr.io/bogosj/go-tesla-go:latest
	docker push ghcr.io/bogosj/go-tesla-go:${GTG_VERSION}

check-version:
ifndef GTG_VERSION
	$(error GTG_VERSION is undefined)
endif