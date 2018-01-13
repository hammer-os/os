FROM alpine:3.7
LABEL maintainer="Markus Sonderegger <marraison@gmail.com>"

RUN apk add --no-cache \
    argp-standalone \
    automake \
    bash \
    bc \
    binutils-dev \
    bison \
    build-base \
    curl \
    diffutils \
    flex \
    git \
    gmp-dev \
    gnupg \
    installkernel \
    kmod \
    libelf-dev \
    libressl-dev \
    linux-headers \
	libunwind-dev \
    mpc1-dev \
    mpfr-dev \
    ncurses-dev \
    sed \
    squashfs-tools \
    tar \
    xz \
    xz-dev \
    zlib-dev

ARG KERNEL_VERSION

ARG KERNEL_SOURCE=https://www.kernel.org/pub/linux/kernel/v4.x/linux-${KERNEL_VERSION}.tar.xz
ARG KERNEL_SHA256_SUMS=https://www.kernel.org/pub/linux/kernel/v4.x/sha256sums.asc
ARG KERNEL_PGP2_SIGN=https://www.kernel.org/pub/linux/kernel/v4.x/linux-${KERNEL_VERSION}.tar.sign

ARG KERNEL_DEF_AMD64_CONFIG=/linux/arch/x86/configs/x86_64_defconfig

ADD keys.asc /keys.asc

RUN curl -fsSLO ${KERNEL_SHA256_SUMS} \
    && gpg -q --import keys.asc \
    && gpg --verify sha256sums.asc \
    && KERNEL_SHA256=$(grep linux-${KERNEL_VERSION}.tar.xz sha256sums.asc | cut -d ' ' -f 1) \
    && [ -f linux-${KERNEL_VERSION}.tar.xz ] || curl -fsSLO ${KERNEL_SOURCE} \
    && echo "${KERNEL_SHA256}  linux-${KERNEL_VERSION}.tar.xz" | sha256sum -c - \
    && xz -d linux-${KERNEL_VERSION}.tar.xz \
    && curl -fsSLO ${KERNEL_PGP2_SIGN} \
    && gpg --verify linux-${KERNEL_VERSION}.tar.sign linux-${KERNEL_VERSION}.tar \
    && cat linux-${KERNEL_VERSION}.tar | tar --absolute-names -x \
	&& mv /linux-${KERNEL_VERSION} /linux

ADD config-amd64 ${KERNEL_DEF_AMD64_CONFIG}
WORKDIR /linux

RUN make defconfig && make oldconfig \
	&& make -j "$(getconf _NPROCESSORS_ONLN)" KCFLAGS="-fno-pie"

RUN make -j "$(getconf _NPROCESSORS_ONLN)" modules

RUN mkdir /kernel && cp arch/x86_64/boot/bzImage /kernel/bzImage

# kernel headers
RUN mkdir -p /tmp/kernel-headers/usr \
	&& make INSTALL_HDR_PATH=/tmp/kernel-headers/usr headers_install \
	&& (cd /tmp/kernel-headers && tar cf /kernel/kernel-headers.tar usr)

RUN DVER=$(basename $(find /tmp/kernel-modules/lib/modules/ -mindepth 1 -maxdepth 1)) \
	&& dir=/tmp/usr/src/linux-headers-$DVER \
	&& mkdir -p $dir \
	&& cp /linux/.config $dir \
	&& cp /linux/Module.symvers $dir \
	&& find . -path './include/*' -prune -o \
		-path './arch/*/include' -prune -o \
		-path './scripts/*' -prune -o \
		-type f \( -name 'Makefile*' -o -name 'Kconfig*' -o -name 'Kbuild*' -o \
		-name '*.lds' -o -name '*.pl' -o -name '*.sh' -o \
		-name 'objtool' -o -name 'fixdep' -o -name 'randomize_layout_seed.h' \) | \
		tar cf - -T - | (cd $dir; tar xf -) \
	&& (cd /tmp && tar cf /kernel/kernel-dev.tar usr/src)
