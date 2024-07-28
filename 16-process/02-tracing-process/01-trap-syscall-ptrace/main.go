package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"

	sec "github.com/seccomp/libseccomp-golang"
)

func main() {
	fmt.Printf("Run %v\n", os.Args[1:])

	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Ptrace: true,
	}

	cmd.Start()
	err := cmd.Wait()
	if err != nil {
		fmt.Println(err)
	}

	pid := cmd.Process.Pid

	var regs syscall.PtraceRegs

	for {
		err := syscall.PtraceGetRegs(pid, &regs)
		if err != nil {
			fmt.Println(err)
			break
		}

		name, _ := sec.ScmpSyscall(regs.Orig_rax).GetName()
		fmt.Println(name)

		if name == "getcwd" {
			adr := uintptr(regs.Rdi)
			syscall.PtracePokeData(pid, adr, []byte("/Hi/"))
			// if _, _, errno := unix.RawSyscall6(unix.SYS_PTRACE, unix.PTRACE_SYSEMU, uintptr(pid), 0, 0, 0, 0); errno != 0 {
			// 	panic(fmt.Sprintf("ptrace sysemu failed: %v", errno))
			// }
		}

		err = syscall.PtraceSyscall(pid, 0)
		if err != nil {
			panic(err)
		}

		_, err = syscall.Wait4(pid, nil, 0, nil)
		if err != nil {
			panic(err)
		}
	}
}
