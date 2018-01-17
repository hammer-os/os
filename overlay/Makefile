include ../build.mk

build: Dockerfile
	@docker build --tag=$(REPONAME)/$(NAME)-overlay:$(VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-overlay:$(VERSION)

run: build
	@docker run --interactive --tty --volume=$(CURDIR):/data \
		$(REPONAME)/$(NAME)-overlay:$(VERSION) /bin/sh -l

clean:
	@echo -n
