ARG ALPINE_VERSION
ARG VERSION
FROM marraison/hammer-busybox:${VERSION} as busybox
FROM alpine:${ALPINE_VERSION} as packages


# create rootfs namespace and install Alpine Linux binary packages
COPY --from=busybox / /pkg
RUN mkdir -p /rootfs/etc/apk && cp -r /etc/apk/* /rootfs/etc/apk/ \
	&& apk add --no-cache --initdb --allow-untrusted -p /rootfs \
		/pkg/x86_64/busybox-1.28.0-r0.apk \
	&& apk add --no-cache -p /rootfs \
		ca-certificates \
		openssh-server \
		btrfs-progs \
		musl \
	&& true

# remove directories/files with no use
ARG NOT_WANT=" \
	/rootfs/etc/apk /rootfs/lib/apk \
	/rootfs/etc/ca-certificates.conf \
	/rootfs/etc/ca-certificates \
	/rootfs/var \
	/rootfs/usr/bin/c_rehash \
	/rootfs/usr/sbin/update-ca-certificates \
	/rootfs/usr/local \
	/rootfs/etc/conf.d \
	/rootfs/etc/init.d \
	/rootfs/etc/network \
	/rootfs/etc/securetty \
	/rootfs/usr/bin/findssl.sh \
	/rootfs/usr/bin/iptables-xml \
"
RUN rm -rf ${NOT_WANT}

# alpine linux v3.7.x does not include containerd. build and install ...
ENV PATH=${PATH}:/go/bin
ENV GOPATH=/go
ARG CONTAINERD_VERSION

RUN apk update && apk upgrade -a && apk add \
		git gcc go make curl \
		musl-dev \
		libc-dev \
		btrfs-progs-dev \
		linux-headers \
	&& true

RUN mkdir -p ${GOPATH}/src/github.com/containerd \
	&& cd ${GOPATH}/src/github.com/containerd \
	&& git clone https://github.com/containerd/containerd.git \
	&& cd ${GOPATH}/src/github.com/containerd/containerd \
	&& git checkout v${CONTAINERD_VERSION} \
	&& make binaries \
		EXTRA_FLAGS="-buildmode pie" \
		EXTRA_LDFLAGS='-extldflags "-fno-PIC -static"' \
		BUILD_TAGS="static_build" \
	&& cp bin/containerd /rootfs/sbin/containerd \
	&& cp bin/containerd-shim /rootfs/sbin/containerd-shim \
	&& cp bin/ctr /rootfs/sbin/ctr

# create system and distribution relevant directories/files
ARG SYSTEM_MOUNTPOINTS="proc sys"
ARG MOUNTPOINTS="dev run tmp var"
#RUN mkdir -p /rootfs/home/adm && chmod 750 /rootfs/home/adm \
#	&& chown adm:adm /rootfs/home/adm \
RUN mkdir -p /rootfs/root && chmod 750 /rootfs/root \
	&& mkdir -p /rootfs/tmp && chmod 1777 /rootfs/tmp \
	&& for d in ${SYSTEM_MOUNTPOINTS}; do \
		mkdir -p /rootfs/${d} && chmod 555 /rootfs/${d}; \
	done \
	&& for d in ${MOUNTPOINTS}; do \
		mkdir -p /rootfs/${d} && chmod 755 /rootfs/${d}; \
	done

# install distribution configuration files
COPY etc /rootfs/etc

FROM scratch
COPY --from=packages /rootfs /
