include ../build.mk

build: Dockerfile etc
	@docker build --tag=$(REPONAME)/$(NAME)-rootfs:$(VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-rootfs:$(VERSION)

run: build
	@docker run --interactive --tty --volume=$(CURDIR):/data \
		$(REPONAME)/$(NAME)-rootfs:$(VERSION) /bin/sh -l

clean:
	@echo -n
