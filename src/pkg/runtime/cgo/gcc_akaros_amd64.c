// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <pthread.h>
#include <string.h> // strerror
#include <signal.h>
#include "libcgo.h"

static void* threadentry(void*);
static void (*setmg_gcc)(void*, void*);

void
x_cgo_init(G* g, void (*setmg)(void*, void*))
{
	int dummy;
	// The system stack is set to a fixed size of 256 pages
	g->stackguard = (uintptr)&dummy - 256*4096 + 4096;
	setmg_gcc = setmg;
}


void
_cgo_sys_thread_start(ThreadStart *ts)
{
	pthread_attr_t attr;
	sigset_t ign, oset;
	pthread_t p;
	size_t size;
	int err;

	sigfillset(&ign);
	sigprocmask(SIG_SETMASK, &ign, &oset);

	pthread_attr_init(&attr);
	pthread_attr_getstacksize(&attr, &size);
	ts->g->stackguard = size;
	err = pthread_create(&p, &attr, threadentry, ts);

	sigprocmask(SIG_SETMASK, &oset, nil);

	if (err != 0) {
		fprintf(stderr, "runtime/cgo: pthread_create failed: %s\n", strerror(err));
		abort();
	}
}

static void*
threadentry(void *v)
{
	ThreadStart ts;

	ts = *(ThreadStart*)v;
	free(v);

	ts.g->stackbase = (uintptr)&ts;

	/*
	 * _cgo_sys_thread_start set stackguard to stack size;
	 * change to actual guard pointer.
	 */
	ts.g->stackguard = (uintptr)&ts - ts.g->stackguard + 4096;

	/*
	 * Set specific keys.
	 */
	setmg_gcc((void*)ts.m, (void*)ts.g);

	crosscall_amd64(ts.fn);
	return nil;
}
