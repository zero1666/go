// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
)

var is []func() int

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if i == j { // 10 skips
				continue
			}
			sum++
		}
		if i&1 == 0 {
			is = append(is, func() int {
				if i%17 == 15 {
					i++
				}
				return i
			})
		}
	}

	bug := false
	if sum != 100-10 {
		fmt.Printf("wrong sum, expected %d, saw %d\n", 90, sum)
		bug = true
	}
	sum = 0
	for _, f := range is {
		sum += f()
	}
	if sum != 2+4+6+8 {
		fmt.Printf("wrong sum, expected %d, saw %d\n", 20, sum)
		bug = true
	}
	if !bug {
		fmt.Printf("PASS\n")
	} else {
		os.Exit(11)
	}
}
