package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("help")
	}
}

func run() {
	fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())

	cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// namespace
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

func child() {
	fmt.Printf("Running %v as pid %d\n", os.Args[2:], os.Getpid())

	cg()

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	must(syscall.Sethostname([]byte("container")))
	must(syscall.Chroot("/home/parikshit/alpine/"))
	must(syscall.Chdir("/"))
	must(syscall.Mount("proc", "proc", "proc", 0, ""))

	must(cmd.Run())

	must(syscall.Unmount("proc", 0))
}

func cg() {
	cgroups := "/sys/fs/cgroup/"

	// must(os.Mkdir(filepath.Join(cgroups, "mycontainer"), 0755))
	must(os.WriteFile(filepath.Join(cgroups, "mycontainer/pids.max"), []byte("10"), 0700))

	must(os.WriteFile(filepath.Join(cgroups, "mycontainer/cgroup.procs"), []byte(strconv.Itoa(os.Getpid())), 0700))
}
