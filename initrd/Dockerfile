ARG ALPINE_VERSION
ARG VERSION

FROM marraison/hammer-rootfs:${VERSION} as rootfs
FROM marraison/hammer-kernel:${VERSION} as kernel
FROM alpine:${ALPINE_VERSION} as syslinux

ARG REPO_URL="https://github.com/hammer-os/os"

RUN apk add --no-cache libarchive-tools cdrkit syslinux xz

#COPY --from=overlay / /rootfs
COPY --from=rootfs / /rootfs
COPY --from=kernel / /kernel

# install syslinux bootloader
RUN mkdir -p /rootfs/boot/isolinux && chmod 755 /rootfs/boot/isolinux \
	&& cp /usr/share/syslinux/isolinux.bin /rootfs/boot/isolinux/isolinux.bin \
	&& cp /usr/share/syslinux/ldlinux.c32  /rootfs/boot/isolinux/ldlinux.c32
ADD isolinux.cfg /rootfs/boot/isolinux/isolinux.cfg

# TODO: fix dead syslink at rootfs/lib/modules/4.14.13-hammer/build 
# install linux kernel and kernel modules
RUN mkdir -p /rootfs/boot && chmod 755 /rootfs/boot \
	&& cp /kernel/kernel /rootfs/boot/kernel \
	&& cp /kernel/System.map /rootfs/boot/System.map

# TODO: kernel headers and development files
#	&& tar xf /kernel/kernel-headers.tar -C /rootfs \
#	&& tar xf /kernel/kernel-dev.tar -C /rootfs \
RUN tar xf /kernel/kernel.tar -C /rootfs \
	&& chmod 755 /rootfs/lib \
	&& rm -f /rootfs/lib/modules/*/build


# install system specific configuration and create /etc/os-release
COPY etc /rootfs/etc
# TODO
RUN	echo PRETTY_NAME=\"HammerOS v0.0.2\" > /rootfs/etc/os-release \
	&& echo NAME=\"HammerOS\" >> /rootfs/etc/os-release \
	&& echo ID=\"hammer\" >> /rootfs/etc/os-release \
	&& echo VERSION=\"v0.0.2\" >> /rootfs/etc/os-release
#RUN	echo PRETTY_NAME=\"HammerOS v${VERSION}\" > /rootfs/etc/os-release \
#	&& echo NAME=\"HammerOS\" >> /rootfs/etc/os-release \
#	&& echo ID=\"hammer\" >> /rootfs/etc/os-release \
#	&& echo VERSION=\"v${VERSION}\" >> /rootfs/etc/os-release \
#	&& echo HOME_URL=\"${REPO_URL}\" >> /rootfs/etc/os-release

# install rc.(start|shutdown|...) handler (/sbin/init)
COPY rc.init /rootfs/etc/rc.init
COPY rc.dhcp /rootfs/etc/rc.dhcp
RUN chmod 755 /rootfs/etc/rc.init /rootfs/etc/rc.dhcp

# install /etc/resolv.conf symlink and /etc/docker configuration
# directory (both must be writable), but /etc is read-only.
RUN ln -s /run/hammer/resolv.conf /rootfs/etc/resolv.conf
RUN ln -s /run/hammer/docker /rootfs/etc/docker

# create /etc/mtab symlink
RUN ln -s /proc/mounts /rootfs/etc/mtab

COPY mkinitrd.sh /bin/mkinitrd
RUN chmod 755 /bin/mkinitrd

WORKDIR /rootfs
ENTRYPOINT ["/bin/mkinitrd"]
CMD ["help"]
