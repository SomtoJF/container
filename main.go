//go:build linux
// +build linux

package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		panic("no command provided")
	}

	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("unsupported command")
	}
}

func run() {
	fmt.Printf("Running %v\n", os.Args[2:])

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running child: %v\n", err)
	}
}

func child() {
	fmt.Println("This is the child function")
	fmt.Printf("Running %v\n", os.Args[2:])

	// Set the hostname
	syscall.Sethostname([]byte("container"))

	// Create a new directory for pivoting
	must(os.MkdirAll("/rootfs", 0700))

	// Bind mount root to /rootfs
	must(syscall.Mount("/", "/rootfs", "", syscall.MS_BIND, ""))
	must(os.MkdirAll("/rootfs/oldroot", 0700))
	must(syscall.PivotRoot("/rootfs", "/rootfs/oldroot"))
	must(os.Chdir("/"))

	// Run the provided command
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Printf("Error running command: %v\n", err)
	}
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
