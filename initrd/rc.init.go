package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/sys/unix"
)

const (
	nodev    = unix.MS_NODEV  // do not allow access to devices (special files)
	noexec   = unix.MS_NOEXEC // do not allow programs to be executed
	nosuid   = unix.MS_NOSUID // do not honor set-user-ID and set-group-ID bits
	readonly = unix.MS_RDONLY // dount filesystem read-only

	relatime = unix.MS_RELATIME
	remount  = unix.MS_REMOUNT
	shared   = unix.MS_SHARED
)

func mount(source, target string, fstype string, flags uintptr, data string) {
	if err := unix.Mount(source, target, fstype, flags, data); err != nil {
		log.Printf("error mounting %s to %s: %v", source, target, err)
	}
}

func mkdir(path string, perm os.FileMode) {
	if err := os.MkdirAll(path, perm); err != nil {
		log.Printf("error making directory %s: %v", path, err)
	}
}

func symlink(oldpath string, newpath string) {
	if err := os.Symlink(oldpath, newpath); err != nil {
		log.Printf("error making symlink %s: %v", newpath, err)
	}
}

func mkchar(path string, mode, major, minor uint32) {
	_, err := os.Lstat(path) // character device already exists
	if err == nil {
		return
	}

	dev := int(unix.Mkdev(major, minor))
	if err = unix.Mknod(path, mode, dev); err != nil {
		log.Printf("error making device %s: %v", path, err)
	}
}

// TODO
func run(name string, args ...string) {
	cmd := exec.Command(name, args...)
	if err := cmd.Run(); err != nil {
		log.Printf("error running %s %v: %v", name, args, err)
	}
}

// TOOD
//func start(name string, args ...string) {
//	cmd := exec.Command(name, args...)
//	if err := cmd.Start(); err != nil {
//		log.Printf("error starting %s %v: %v", name, args, err)
//	}
//}

func write(path string, data string) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Printf("error opening %s: %v", path, err)
		return
	}

	n, err := f.Write([]byte(data))
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err != nil {
		log.Printf("error writing to %s: %v", path, err)
		return
	}
	if err := f.Close(); err != nil {
		log.Printf("error closing %s: %v", path, err)
	}
}

func touch(path string, perm os.FileMode) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		log.Printf("error creating %s: %v", path, err)
		return
	}
	if err = f.Close(); err != nil {
		log.Printf("error closint %s: %v", path, err)
	}
}

func cgroupSubsystems(r io.Reader) ([]string, error) {
	sub, name := []string{}, ""
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanWords)
	for n := 0; s.Scan(); {
		switch n {
		case 0:
			name = s.Text()
			n++
		case 1, 2:
			n++
		case 3:
			if len(name) > 0 && name[0] != '#' {
				if s.Text() == "1" {
					sub = append(sub, name)
				}
			}
			name, n = "", 0
		}
	}
	return sub, s.Err()
}

func doCgroupMount() {
	f, err := os.Open("/proc/cgroups")
	if err != nil {
		log.Printf("error opening /proc/cgroups: %v", err)
		return
	}
	defer f.Close()

	subsystems, err := cgroupSubsystems(f)
	if err != nil {
		log.Printf("error finding cgroup subsystems: %v", err)
		return
	}

	for _, name := range subsystems {
		path := filepath.Join("/sys/fs/cgroup", name)
		mkdir(path, 0555)
		mount(name, path, "cgroup", noexec|nosuid|nodev, name)
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
	doCgroupMount()

	// TODO: some of the subsystems may not exist
	mount("securityfs", "/sys/kernel/security", "securityfs", noexec|nosuid|nodev, "")
	mount("debugfs", "/sys/kernel/debug", "debugfs", noexec|nosuid|nodev, "")
	mount("configfs", "/sys/kernel/config", "configfs", noexec|nosuid|nodev, "")
	mount("fusectl", "/sys/fs/fuse/connections", "fusectl", noexec|nosuid|nodev, "")
	mount("selinuxfs", "/sys/fs/selinux", "selinuxfs", noexec|nosuid, "")
	mount("pstore", "/sys/fs/pstore", "pstore", noexec|nosuid|nodev, "")
	mount("efivarfs", "/sys/firmware/efi/efivars", "efivarfs", noexec|nosuid|nodev, "")

	mount("tmpfs", "/var", "tmpfs", nodev|nosuid|noexec|relatime, "size=50%,mode=755")
	mkdir("/var/cache", 0755)
	mkdir("/var/empty", 0555)
	mkdir("/var/lib", 0755)
	mkdir("/var/lib/containerd", 0755)
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
	write("/proc/sys/kernel/printk", "1")
}

func doContainerd() {
	mkdir("/run/rw/containerd", 0755)
	//start("/usr/bin/containerd")
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
	write("/proc/sys/kernel/hotplug", "/sbin/mdev")
	run("/sbin/mdev", "-s")
}

// http://www.linuxfromscratch.org/lfs/view/6.1/part3.html
func main() {
	doMount()
	doClock()
	doHotplug()
	doNetwork()
	doContainerd()
}
