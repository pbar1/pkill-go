package main

import (
	"fmt"
	"syscall"

	"github.com/pbar1/pkill-go"
)

func main() {
	// terminate ourself
	pids, err := pkill.Pkill("^go$", syscall.SIGTERM)
	if err != nil {
		fmt.Println(err)
	}
	for _, pid := range pids {
		fmt.Println(pid)
	}
}
