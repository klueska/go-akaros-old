// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// +build akaros

package parlib

/*
#include <parlib.h>
#include <uthread.h>
#include <vcore.h>
#include <mcs.h>
#include <futex.h>
*/
import "C"
import (
	"unsafe"
	"errors"
)

var Procinfo *ProcinfoType = (*ProcinfoType)(unsafe.Pointer(uintptr(C.UINFO)))
func (p *ProcinfoType) Pid() int { return int(p.pid) }

func Futex(uaddr *int32, op int32, val int32,
           timeout *Timespec, uaddr2 *int32, val3 int32) (ret int32) {
	// For now, akaros futexes don't support uaddr2 or val3, so we
	// just 0 them out.
	uaddr2 = nil;
	val3 = 0;
	// Also, the minimum timout is 1ms, so up it to that if it's too small
	if (timeout != nil) {
		if (timeout.tv_sec == 0) {
			if (timeout.tv_nsec < 1000000) {
				timeout.tv_nsec = 1000000;
			}
		}
	}
	return int32(C.futex((*C.int)(unsafe.Pointer(uaddr)),
	                     C.int(op), C.int(val),
	                     (*C.struct_timespec)(unsafe.Pointer(timeout)),
	                     (*C.int)(unsafe.Pointer(uaddr2)), C.int(val3)))
}

func ProcinfoPackArgs(argv []*byte, envp []*byte) (pi ProcinfoType, err error) {
	p_pi := (*_Ctype_procinfo_t)(unsafe.Pointer(&pi))
    p_argv := (**_Ctype_char)(unsafe.Pointer(&argv[0]))
    p_envp := (**_Ctype_char)(unsafe.Pointer(&envp[0]))

	__err := C.procinfo_pack_args(p_pi, p_argv, p_envp)
	if __err == -1 {
		err = nil
	} else {
		err = errors.New("ProcinfoPackArgs: error packing argv and envp")
	}
	return pi, err
}

