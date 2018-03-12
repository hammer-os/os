package main

import (
	"log"
	"os"

	"github.com/hammer-os/os/initrd/init/linux"
)

const (
	nodev    = linux.NoDevice // do not allow access to devices (special files)
	noexec   = linux.NoExec   // do not allow programs to be executed
	nosuid   = linux.NoSuid   // do not honor set-user-ID and set-group-ID bits
	readonly = linux.Readonly // dount filesystem read-only

	relatime = linux.Relatime
	remount  = linux.Remount
	shared   = linux.Shared
)

func mount(src, dst string, fstype string, flags uintptr, data string) {
	if err := linux.Mount(src, dst, fstype, flags, data); err != nil {
		log.Printf("mounting %s to %s: %v", src, dst, err)
	}
}

func mkdir(path string, perm os.FileMode) {
	if err := os.MkdirAll(path, perm); err != nil {
		log.Printf("making directory %s: %v", path, err)
	}
}

func symlink(src, dst string) {
	if err := os.Symlink(src, dst); err != nil {
		log.Printf("making symlink %s: %v", dst, err)
	}
}

func mkchar(path string, mode, major, minor uint32) {
	if err := linux.Mkchar(path, mode, major, minor); err != nil {
		log.Printf("making device %s: %v", path, err)
	}
}

func touch(path string, perm os.FileMode) {
	if err := linux.Touch(path, perm); err != nil {
		log.Printf("creating file %s: %v", path, err)
	}
}

func write(path string, data []byte, perm os.FileMode) {
	if err := linux.Write(path, data, perm); err != nil {
		log.Printf("writing %s: %v", path, err)
	}
}

func run(name string, args ...string) {
	if err := linux.Run(name, args...); err != nil {
		log.Printf("running %s %v: %v", name, args, err)
	}
}

func exec(name string, args ...string) {
	if err := linux.Exec(name, args...); err != nil {
		log.Printf("running %s %v: %v", name, args, err)
	}
}

func doMount() {
	mount("dev", "/dev", "devtmpfs", nosuid|noexec|relatime, "size=10m,nr_inodes=248418,mode=755")

	mount("proc", "/proc", "proc", nodev|nosuid|noexec|relatime, "")
	mount("sysfs", "/sys", "sysfs", noexec|nosuid|nodev, "")
	mount("tmpfs", "/run", "tmpfs", nodev|nosuid|noexec|relatime, "size=10%,mode=755")
	mount("tmpfs", "/tmp", "tmpfs", nodev|nosuid|noexec|relatime, "size=10%,mode=1777")

	// see http://www.linuxfromscratch.org/lfs/view/6.1/chapter06/devices.html
	mkchar("/dev/console", 0600, 5, 1)
	mkchar("/dev/null", 0666, 1, 3)
	mkchar("/dev/zero", 0666, 1, 5)
	mkchar("/dev/ptmx", 0666, 5, 1)
	mkchar("/dev/tty", 0666, 5, 0)
	//mkchar("/dev/tty1", 0620, 4, 1)
	//mkchar("/dev/kmsg", 0660, 1, 11)

	symlink("/proc/self/fd", "/dev/fd")
	symlink("/proc/self/fd/0", "/dev/stdin")
	symlink("/proc/self/fd/1", "/dev/stdout")
	symlink("/proc/self/fd/2", "/dev/stderr")
	symlink("/proc/kcore", "/dev/kcore")

	mkdir("/dev/mqueue", 01777)
	mkdir("/dev/shm", 01777)
	mkdir("/dev/pts", 0755)
	mount("mqueue", "/dev/mqueue", "mqueue", noexec|nosuid|nodev, "")
	mount("shm", "/dev/shm", "tmpfs", noexec|nosuid|nodev, "mode=1777")
	mount("devpts", "/dev/pts", "devpts", noexec|nosuid, "gid=5,mode=0620")

	// mount cgroup root tmpfs
	mount("cgroup_root", "/sys/fs/cgroup", "tmpfs", nodev|noexec|nosuid, "mode=755,size=10m")
	linux.MountCgroupSubsystems()

	linux.MountSubsystems()

	mount("tmpfs", "/var", "tmpfs", nodev|nosuid|noexec|relatime, "size=50%,mode=755")
	mkdir("/var/cache", 0755)
	mkdir("/var/empty", 0555)
	mkdir("/var/lib", 0755)
	mkdir("/var/lib/udhcpd", 0755)
	mkdir("/var/local", 0755)
	mkdir("/var/lock", 0755)
	mkdir("/var/log", 0755)
	mkdir("/var/opt", 0755)
	mkdir("/var/spool", 0755)
	mkdir("/var/tmp", 01777)
	symlink("/run", "/var/run")

	touch("/run/resolv.conf", 0600)

	// Hide all kernel messages. Only kernel panics will be displayed.
	write("/proc/sys/kernel/printk", []byte("1"), 0644)
}

func doNetwork() {
	run("/sbin/ip", "addr", "add", "127.0.0.1/8", "dev", "lo", "brd", "+", "scope", "host")
	run("/sbin/ip", "route", "add", "127.0.0.0/8", "dev", "lo", "scope", "host")
	run("/sbin/ip", "link", "set", "lo", "up")

	run("/sbin/ip", "link", "set", "eth0", "up")
	// TODO:
	run("/sbin/udhcpc", "-s", "/etc/rc.dhcp", "-i", "eth0", "-v")
}

func doClock() { run("/sbin/hwclock", "--hctosys", "--utc") }

func doHotplug() {
	write("/proc/sys/kernel/hotplug", []byte("/sbin/mdev"), 0644)
	run("/sbin/mdev", "-s")
}

// http://www.linuxfromscratch.org/lfs/view/6.1/part3.html
func main() {
	doMount()
	doClock()
	doHotplug()
	doNetwork()
}
