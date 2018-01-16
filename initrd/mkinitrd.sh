#!/bin/sh

case ${1} in
	# TODO: enable XZ kernel configuration
	initrd)
		find . | cpio -R root:root -H newc -o | gzip -9 ;;
	
	kernel)
		cat /kernel/kernel ;;

	iso)
		genisoimage -o /tmp/hammer.iso -l -J -R \
			-c boot/isolinux/boot.cat  \
			-b boot/isolinux/isolinux.bin \
			-no-emul-boot \
			-boot-load-size 4 \
			-boot-info-table \
			-joliet-long \
			-input-charset utf8 \
			-hide-rr-moved \
			-V HammerOS .

		isohybrid /tmp/hammer.iso
		cat /tmp/hammer.iso ;;


	*)
		echo "usage: mkinitrd [initrd|kernel|iso]" 1>&2;;
esac
