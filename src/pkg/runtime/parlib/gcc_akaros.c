// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <sys/syscall.h>
#include <futex.h>
#include "gcc_akaros.h"

// Akaros syscalls
typedef void (*gcc_call_t)(void *arg);
const gcc_call_t gcc_syscall = (gcc_call_t)ros_syscall_sync;

// Akaros style futexes
static void __gcc_futex(void *__arg)
{
	// For now, akaros futexes don't support timeout, uaddr2 or val3, so we
	// just 0 them out.
	gcc_futex_arg_t *a = (gcc_futex_arg_t*)__arg;
	a->timeout = NULL;
	a->uaddr2 = NULL;
	a->val3 = 0;
	a->retval = futex(a->uaddr, a->op, a->val, a->timeout, a->uaddr2, a->val3);
}
const gcc_call_t gcc_futex = __gcc_futex;

