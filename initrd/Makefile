include ../build.mk

build: Dockerfile init-amd64
	@docker build --tag=$(REPONAME)/$(NAME)-initrd:$(VERSION) .

push: build
	@docker push $(REPONAME)/$(NAME)-initrd:$(VERSION)

iso: build
	@docker run $(REPONAME)/$(NAME)-initrd:$(VERSION) \
		iso > hammer-amd64.iso

kernel: build
	@docker run $(REPONAME)/$(NAME)-initrd:$(VERSION) \
		kernel > kernel

initrd: build
	@docker run $(REPONAME)/$(NAME)-initrd:$(VERSION) \
		initrd > initrd

init-amd64:
	GOOS=linux GOARCH=amd64 go build -o rc.init rc.init.go

clean:
	@go clean ./...
