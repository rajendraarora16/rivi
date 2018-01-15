// Copyright 2014 Canonical Ltd.
// Copyright 2014 Cloudbase Solutions SRL
// Licensed under the LGPLv3, see LICENCE file for details.

// mksyscall_windows.pl -l32 uptime_windows.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package uptime

import "syscall"

var (
	modkernel32 = syscall.NewLazyDLL("kernel32.dll")

	procGetTickCount64 = modkernel32.NewProc("GetTickCount64")
)

func getTickCount64() (uptime uint64, err error) {
	r0, _, e1 := syscall.Syscall(procGetTickCount64.Addr(), 0, 0, 0, 0)
	uptime = uint64(r0)
	if uptime == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}
