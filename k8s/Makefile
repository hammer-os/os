include ../build.mk

build: Dockerfile APKBUILD
	@docker build --tag=$(REPONAME)/$(NAME)-kubernetes:$(VERSION) \
		--build-arg=ALPINE_VERSION=$(ALPINE_VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-kubernetes:$(VERSION)

clean:
	@echo -n
