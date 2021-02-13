// Copyright (c) 2021 Pierce Bartine. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/pbar1/pkill-go"
)

func main() {
	pids, err := pkill.Pkill(os.Args[1], syscall.SIGKILL)
	if err != nil {
		fmt.Println(err)
	}
	for _, pid := range pids {
		fmt.Println(pid)
	}
}
