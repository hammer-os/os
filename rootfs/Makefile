include ../build.mk

build: Dockerfile etc
	@echo "HammerOS v$(VERSION) (Kernel \r on an \m (\l))" > etc/issue
	@echo >> etc/issue
	@docker build --tag=$(REPONAME)/$(NAME)-rootfs:$(VERSION) \
		--build-arg=CONTAINERD_VERSION=$(CONTAINERD_VERSION) \
		--build-arg=ALPINE_VERSION=$(ALPINE_VERSION) \
		--build-arg=VERSION=$(VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-rootfs:$(VERSION)

run: build
	@docker run --interactive --tty --volume=$(CURDIR):/data \
		$(REPONAME)/$(NAME)-rootfs:$(VERSION) /bin/sh -l

clean:
	@echo -n
