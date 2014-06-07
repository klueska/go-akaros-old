// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "parlib/akaros.h"

#define SS_DISABLE 2

// Akaros-specific system calls
int32	runtime·futex(uint32*, int32, uint32, Timespec*, uint32*, uint32);
int32	runtime·clone(int32, void*, M*, G*, void(*)(void));

struct Sigaction;
int32	runtime·sigaction(int32, struct Sigaction*, struct Sigaction*);
void	runtime·sigpanic(void);
void runtime·setitimer(int32, Itimerval*, Itimerval*);

int32	runtime·sigprocmask(int32, Sigset*, Sigset*);
void	runtime·unblocksignals(void);
#define SIG_SETMASK 2

#define RLIMIT_AS 9
typedef struct Rlimit Rlimit;
struct Rlimit {
	uintptr	rlim_cur;
	uintptr	rlim_max;
};
int32	runtime·getrlimit(int32, Rlimit*);
