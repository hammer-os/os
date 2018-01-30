package linux

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"syscall"

	"golang.org/x/sys/unix"
)

const (
	NoDevice = unix.MS_NODEV  // do not allow access to devices (special files)
	NoExec   = unix.MS_NOEXEC // do not allow programs to be executed
	NoSuid   = unix.MS_NOSUID // do not honor set-user-ID and set-group-ID bits
	Readonly = unix.MS_RDONLY // dount filesystem read-only

	Relatime = unix.MS_RELATIME
	Remount  = unix.MS_REMOUNT
	Shared   = unix.MS_SHARED
)

func Mount(src, dst, fstype string, flags uintptr, data string) error {
	if err := unix.Mount(src, dst, fstype, flags, data); err != nil {
		return &os.PathError{"mount", dst, err}
	}
	return nil
}

func Mkdir(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func open(path string, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func touch(path string, perm os.FileMode) (*os.File, error) {
	// Fast path: if we can tell whether path is a directory or file, stop
	// with success or error.
	fi, err := os.Lstat(path)
	if err == nil {
		if !fi.IsDir() {
			return open(path, perm)
		}
		return nil, &os.PathError{"touch", path, unix.EISDIR}
	}

	i := len(path)
	for i > 0 && os.IsPathSeparator(path[i-1]) {
		i--
	}

	j := i
	for j > 0 && !os.IsPathSeparator(path[j-1]) {
		j--
	}

	if j > 1 {
		if err = os.MkdirAll(path[0:j-1], 0755); err != nil {
			return nil, err
		}
	}
	return open(path, perm)
}

func Touch(path string, perm os.FileMode) error {
	f, err := touch(path, perm)
	if err != nil {
		return err
	}
	return f.Close()
}

func Mkchar(path string, mode, major, minor uint32) error {
	_, err := os.Lstat(path) // character device already exists
	if err == nil {
		return nil
	}

	dev := int(unix.Mkdev(major, minor))
	if err = unix.Mknod(path, mode, dev); err != nil {
		return &os.PathError{"mknod", path, err}
	}
	return nil
}

func Write(path string, data []byte, perm os.FileMode) error {
	f, err := touch(path, perm)
	if err != nil {
		return err
	}

	n, err := f.Write(data)
	if err == nil && n < len(data) {
		err = io.ErrShortWrite
	}
	if err != nil {
		f.Close()
		return err
	}
	return f.Close()
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
	if err := s.Err(); err != nil {
		return nil, err
	}
	return sub, nil
}

const (
	subFlags    = unix.MS_NODEV | unix.MS_NOSUID | unix.MS_NOEXEC
	cgroupsPath = "/proc/cgroups"
)

func MountCgroupSubsystems() error {
	f, err := os.Open(cgroupsPath)
	if err != nil {
		return err
	}
	defer f.Close()

	subsystems, err := cgroupSubsystems(f)
	if err != nil {
		return &os.PathError{"scan", cgroupsPath, err}
	}

	for _, name := range subsystems {
		path := "/sys/fs/cgroup/" + name
		if err = os.Mkdir(path, 0755); err != nil {
			return err
		}
		err = unix.Mount(name, path, "cgroup", subFlags, name)
		if err != nil {
			return err
		}
	}
	return nil
}

// some of the subsystems may not exist -> ignore all errors
func MountSubsystems() {
	unix.Mount("securityfs", "/sys/kernel/security", "securityfs", subFlags, "")
	unix.Mount("debugfs", "/sys/kernel/debug", "debugfs", subFlags, "")
	unix.Mount("configfs", "/sys/kernel/config", "configfs", subFlags, "")
	unix.Mount("fusectl", "/sys/fs/fuse/connections", "fusectl", subFlags, "")
	unix.Mount("selinuxfs", "/sys/fs/selinux", "selinuxfs", subFlags, "")
	unix.Mount("pstore", "/sys/fs/pstore", "pstore", subFlags, "")
	unix.Mount("efivarfs", "/sys/firmware/efi/efivars", "efivarfs", subFlags, "")
}

func Run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			return e // TODO
		}
		return err
	}
	return nil
}

// Exec invokes the execve(2) system call.
func Exec(name string, args ...string) error {
	return syscall.Exec(name, args, os.Environ())
}
