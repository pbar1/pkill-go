// Copyright (c) 2021 Pierce Bartine. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

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
