#!/bin/sh

ENABLE_KVM=""
[ -c /dev/kvm ] && ENABLE_KVM=-enable-kvm

qemu-system-x86_64 ${ENABLE_KVM} -boot d -cdrom hammer-amd64.iso \
	-device virtio-net-pci,netdev=t0,mac=b6:d3:e7:05:5c:1b \
	-netdev user,id=t0 \
	-object rng-random,id=rng0,filename=/dev/urandom \
	-device virtio-rng-pci,rng=rng0 \
	-nographic \
	-smp 1 \
	-m 1024 \
	-machine q35
