package main

import (
	"fmt"
	"log"
	"os"
)

func osrelease(path string, version string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, `PRETTY_NAME="v%s"
NAME="HammerOS"
ID="hammer"
VERSION="v%s"
`, version, version)
	return nil
}

func issue(path string, version string) error {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	fmt.Fprintf(f, `

HammerOS v%s (Kernel \r on a \m (\l))


`, version)
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s version\n", os.Args[0])
		os.Exit(2)
	}
	version := os.Args[1]

	if err := osrelease("etc/os-release", version); err != nil {
		log.Fatalf("error generating etc/os-release: %v", err)
	}
	if err := issue("etc/issue", version); err != nil {
		log.Fatalf("error generating etc/issue: %v", err)
	}
}
