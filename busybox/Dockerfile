FROM alpine:3.7 AS busybox-build
LABEL maintainer="Markus Sonderegger <marraison@gmail.com>"

RUN apk add --no-cache \
		bzip2 \
		coreutils \
		curl \
		gcc \
		gnupg \
		linux-headers \
		make \
		musl-dev \
		tzdata

ARG HAMMER_VERSION
ARG BUSYBOX_VERSION 

COPY keys.asc /

RUN gpg -q --import keys.asc \
	&&tarball="busybox-${BUSYBOX_VERSION}.tar.bz2" \
	&& curl -fL -o busybox.tar.bz2 "https://busybox.net/downloads/$tarball" \
	&& curl -fL -o busybox.tar.bz2.sign "https://busybox.net/downloads/$tarball.sign" \
	&& gpg --batch --decrypt --output busybox.tar.bz2.txt busybox.tar.bz2.sign \
	&& awk '$1 == "SHA1:" && $2 ~ /^[0-9a-f]+$/ && $3 == "'"$tarball"'" \
		{ print $2, "*busybox.tar.bz2" }' busybox.tar.bz2.txt > busybox.tar.bz2.sha1 \
	&& test -s busybox.tar.bz2.sha1 \
	&& sha1sum -c busybox.tar.bz2.sha1 \
	&& mkdir -p /busybox \
	&& tar -xf busybox.tar.bz2 -C /busybox --strip-components 1 \
	&& rm busybox.tar.bz2*

ADD config /busybox/.config

COPY etc /busybox/rootfs/etc
WORKDIR /busybox

RUN chmod 755 /busybox/rootfs/etc && chmod 644 /busybox/rootfs/etc/*

# https://www.mail-archive.com/toybox@lists.landley.net/msg02528.html
# https://www.mail-archive.com/toybox@lists.landley.net/msg02526.html
RUN sed -i 's/^struct kconf_id \*$/static &/g' scripts/kconfig/zconf.hash.c_shipped

#	&& mkdir rootfs/usr && chmod 755 rootfs/usr \
#	&& mkdir 	rootfs/usr/sbin && chmod 755 rootfs/usr/sbin \
#	&& mkdir 	rootfs/usr/bin && chmod 755 rootfs/usr/bin \
#	&& mkdir rootfs/boot && chmod 755 rootfs/boot \
RUN make oldconfig && make busybox \
	&& make CONFIG_PREFIX=/busybox/rootfs install \
	&& mkdir rootfs/tmp && chmod 1777 rootfs/tmp \
	&& mkdir rootfs/root && chmod 700 rootfs/root \
	&& mkdir rootfs/home && chmod 755 rootfs/home \
	&&	mkdir rootfs/etc/profile.d && chmod 755 rootfs/etc/profile.d \
	&& mkdir rootfs/dev && chmod 755 rootfs/dev \
	&& mkdir rootfs/mnt && chmod 755 rootfs/mnt \
	&& mkdir rootfs/var && chmod 755 rootfs/var \
	&& mkdir rootfs/run && chmod 755 rootfs/run \
	&& mkdir rootfs/sys && chmod 555 rootfs/sys \
	&& mkdir rootfs/proc && chmod 555 rootfs/proc


FROM scratch
COPY --from=busybox-build /busybox/rootfs /
