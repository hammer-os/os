include build.mk

build:
	@make -C kernel build
	@make -C rootfs build
	@make -C initrd build

clean:
	@make -C rootfs clean
	@make -C kernel clean
	@make -C initrd clean
	
