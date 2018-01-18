ARG ALPINE_VERSION

FROM alpine:${ALPINE_VERSION} as overlay

# create rootfs namespace and install Alpine Linux binary packages
RUN mkdir -p /rootfs/etc/apk \
	&& cp -r /etc/apk/* /rootfs/etc/apk/ \
	&& apk add --no-cache --initdb --root /rootfs \
		busybox ca-certificates \
		openssh-sftp-server \
		openssh-server \
		openssh-client \
		btrfs-progs \
		iptables \
		device-mapper-libs \
		libltdl \
		libseccomp \
		docker \
		musl \
	&& true

# remove directories/files with no use
ARG NOT_WANT=" \
	/rootfs/etc/apk /rootfs/lib/apk \
	/rootfs/etc/conf.d \
	/rootfs/etc/init.d \
	/rootfs/etc/network \
	/rootfs/etc/securetty \
	/rootfs/usr/bin/findssl.sh \
	/rootfs/usr/bin/iptables-xml \
	/rootfs/usr/local \
	/rootfs/var \
"
RUN rm -rf ${NOT_WANT}

# create system and distribution relevant directories/files
ARG SYSTEM_MOUNTPOINTS="proc sys"
ARG MOUNTPOINTS="dev run tmp var"
RUN mkdir -p /rootfs/home/adm && chmod 750 /rootfs/home/adm \
	&& chown adm:adm /rootfs/home/adm \
	&& mkdir -p /rootfs/root && chmod 750 /rootfs/root \
	&& mkdir -p /rootfs/tmp && chmod 1777 /rootfs/tmp \
	&& for d in ${SYSTEM_MOUNTPOINTS}; do \
		mkdir -p /rootfs/${d} && chmod 555 /rootfs/${d}; \
	done \
	&& for d in ${MOUNTPOINTS}; do \
		mkdir -p /rootfs/${d} && chmod 755 /rootfs/${d}; \
	done

# install distribution configuration files
ADD etc /rootfs/etc

COPY service.c service.c
RUN apk add --no-cache \
		bsd-compat-headers \
		linux-headers \
		musl-dev \
		gcc \
	&& gcc -o /rootfs/sbin/service service.c

FROM scratch
COPY --from=overlay /rootfs /
LABEL maintainer="Markus Sonderegger <marraison@gmail.com>"
