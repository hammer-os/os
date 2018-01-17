FROM alpine:3.7 as overlay
LABEL maintainer="Markus Sonderegger <marraison@gmail.com>"

RUN apk add --no-cache curl

ARG ARCHIVE="http://dl-cdn.alpinelinux.org/alpine/v3.7/main"
ARG ARCH="x86_64"

# TODO: automate this!

ARG OPENSSH_DEPS=" \
	libressl2.6-libcrypto-2.6.3-r0 \
	zlib-1.2.11-r1 \
	musl-1.1.18-r2 \
"

ARG OPENSSH=" \
	openssh-server-common-7.5_p1-r8 \
	openssh-server-7.5_p1-r8 \
	openssh-keygen-7.5_p1-r8 \
	openssh-client-7.5_p1-r8 \
	openssh-sftp-server-7.5_p1-r8 \
"

ARG BTRFS_DEPS=" \
	libuuid-2.31-r0 \
	libblkid-2.31-r0 \
	lzo-2.10-r2 \
"

ARG BTRFS_PROGS="btrfs-progs-4.13.2-r0"

RUN mkdir /rootfs \
	&& for pkg in ${OPENSSH_DEPS} ${OPENSSH}; do \
		curl -fL -o ${pkg}.apk ${ARCHIVE}/${ARCH}/${pkg}.apk; \
	done \
	&& for pkg in ${BTRFS_DEPS} ${BTRFS_PROGS}; do \
		curl -fL -o ${pkg}.apk ${ARCHIVE}/${ARCH}/${pkg}.apk; \
	done \
	&& for pkg in *.apk; do \
		tar xzf ${pkg} -C /rootfs; \
	done \
	&& mv /rootfs/sbin/* /rootfs/usr/sbin && rm -r /rootfs/sbin \
	&& rm /rootfs/.SIGN* /rootfs/.PKGINFO \
	&& rm -r /rootfs/etc/conf.d /rootfs/etc/init.d \
	&& rm -f /rootfs/usr/bin/findssl.sh


FROM scratch
COPY --from=overlay /rootfs /
