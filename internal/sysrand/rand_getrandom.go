// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || freebsd || linux || solaris

package sysrand

import (
	"errors"
	"math"
	"runtime"
	"syscall"
	"unsafe"
)

const sysGetrandom = 318 // SYS_GETRANDOM on linux/amd64

func getrandom(buf []byte, flags uintptr) (int, error) {
	if len(buf) == 0 {
		return 0, nil
	}
	r1, _, errno := syscall.Syscall(sysGetrandom, uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), flags)
	if errno != 0 {
		return int(r1), errno
	}
	return int(r1), nil
}

func read(b []byte) error {
	maxSize := math.MaxInt32
	if runtime.GOOS == "solaris" {
		maxSize = 133120
	}

	for len(b) > 0 {
		size := len(b)
		if size > maxSize {
			size = maxSize
		}
		n, err := getrandom(b[:size], 0)
		if errors.Is(err, syscall.ENOSYS) {
			return urandomRead(b)
		}
		if errors.Is(err, syscall.EINTR) {
			continue
		}
		if err != nil {
			return err
		}
		b = b[n:]
	}
	return nil
}
