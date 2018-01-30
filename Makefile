include build.mk

build:
	@make -C busybox build
	@make -C rootfs build
	@make -C initrd build

iso: build
	@make -C initrd iso

push:
	@make -C busybox push
	@make -C rootfs push
	@make -C initrd push

clean:
	@make -C busybox clean
	@make -C rootfs clean
	@make -C kernel clean
	@make -C initrd clean
	
