package main

import (
	"log"
	"os"
	"syscall"
)

func main() {
	// get pid of process
	log.Printf("Pid of process: %d", os.Getpid())

	// get parent id
	log.Printf("Parent id: %d", os.Getppid())

	// real user id (uid) & group id (gid)
	log.Printf("Real User Id: %d, Group Id: %d", os.Getuid(), os.Getgid())

	// effective user id (uid) & group id (gid)
	// will change effective ids to 1000 (parikshit)
	syscall.Setegid(1000)
	syscall.Seteuid(1000)
	log.Printf("Effective User Id: %d, Group Id: %d", os.Geteuid(), os.Getegid())
}
