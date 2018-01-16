include ../build.mk

build: Dockerfile init-amd64
	@docker build --tag=$(NAME)-initrd:$(VERSION) .

iso: build
	@docker run $(NAME)-initrd:$(VERSION) iso > hammer-amd64.iso

kernel: build
	@docker run $(NAME)-initrd:$(VERSION) kernel > kernel

initrd: build
	@docker run $(NAME)-initrd:$(VERSION) initrd > initrd


#	@docker run --interactive --tty --volume=$(CURDIR):/data \
#		$(NAME)-initrd:$(VERSION) /bin/sh -l

init-amd64:
	GOOS=linux GOARCH=amd64 go build -o rc.init rc.init.go

clean:
	@go clean ./...
