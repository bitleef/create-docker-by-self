package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func UTS_Namespace() (err error) {
	cmd := exec.Command("sh")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		fmt.Println(err)
	}
	return err
}
