// This file has been manually generated by trying to build particular GO
// packages and realizing that we don't yet have definitions for particular
// system call numbers as originally required by the linux port we derive
// outselves from. As we continue porting the contants in here will start to
// disappear. The contants are copied out of zsysnum_linux_amd64.go

package syscall

const (
	SYS_PRCTL                        = 300
	SYS_GETPPID                      = 301
	SYS_KILL                         = 302
	SYS_PTRACE                       = 303
	SYS_SETSID                       = 304
	SYS_SETPGID                      = 305
	SYS_CHROOT                       = 306
	SYS_SETGROUPS                    = 307
	SYS_DUP2                         = 308
	SYS_IOCTL                        = 309
	SYS_EXECVE                       = 310
	SYS_EXIT                         = 311
	SYS_OPENAT                       = 312
	SYS_PIPE                         = 313
	SYS_PIPE2                        = 314
	SYS_UTIMES                       = 315
	SYS_UTIMENSAT                    = 316
	SYS_FUTIMESAT                    = 317
	SYS_WAIT4                        = 318
	SYS_MOUNT                        = 319
	SYS_ACCT                         = 320
	SYS_ADJTIMEX                     = 321
	SYS_CREAT                        = 322
	SYS_DUP                          = 323
	SYS_EPOLL_CREATE                 = 324
	SYS_EPOLL_CREATE1                = 325
	SYS_EPOLL_CTL                    = 326
	SYS_EPOLL_WAIT                   = 327
	SYS_EXIT_GROUP                   = 328
	SYS_FACCESSAT                    = 329
	SYS_FALLOCATE                    = 330
	SYS_FCHDIR                       = 331
	SYS_FCHMOD                       = 332
	SYS_FCHMODAT                     = 333
	SYS_FCHOWNAT                     = 334
	SYS_FDATASYNC                    = 335
	SYS_FLOCK                        = 336
	SYS_FSYNC                        = 337
	SYS_GETDENTS64                   = 338
	SYS_GETPGID                      = 339
	SYS_GETPGRP                      = 340
	SYS_GETPRIORITY                  = 341
	SYS_GETRUSAGE                    = 342
	SYS_GETTID                       = 343
	SYS_GETXATTR                     = 344
	SYS_INOTIFY_ADD_WATCH            = 345
	SYS_INOTIFY_INIT                 = 346
	SYS_INOTIFY_INIT1                = 347
	SYS_INOTIFY_RM_WATCH             = 348
	SYS_SYSLOG                       = 349
	SYS_LISTXATTR                    = 350
	SYS_MKDIRAT                      = 351
	SYS_MKNOD                        = 352
	SYS_MKNODAT                      = 353
	SYS_NANOSLEEP                    = 354
	SYS_PAUSE                        = 355
	SYS_PIVOT_ROOT                   = 356
	SYS_PRLIMIT64                    = 357
	SYS_REMOVEXATTR                  = 358
	SYS_RENAME                       = 359
	SYS_RENAMEAT                     = 360
	SYS_SETDOMAINNAME                = 361
	SYS_SETHOSTNAME                  = 362
	SYS_SETTIMEOFDAY                 = 363
	SYS_SETPRIORITY                  = 364
	SYS_SETXATTR                     = 365
	SYS_SYNC                         = 366
	SYS_SYSINFO                      = 367
	SYS_TEE                          = 368
	SYS_TGKILL                       = 369
	SYS_TIMES                        = 370
	SYS_UNAME                        = 371
	SYS_UNLINKAT                     = 372
	SYS_UMOUNT2                      = 373
	SYS_UNSHARE                      = 374
	SYS_USTAT                        = 375
	SYS_UTIME                        = 376
	SYS_MADVISE                      = 377
	SYS_MLOCK                        = 378
	SYS_MUNLOCK                      = 379
	SYS_MLOCKALL                     = 380
	SYS_MUNLOCKALL                   = 381
	SYS_CHOWN                        = 382
	SYS_FCHOWN                       = 383
	SYS_FSTATFS                      = 384
	SYS_FTRUNCATE                    = 385
	SYS_GETEGID                      = 386
	SYS_GETEUID                      = 387
	SYS_GETGID                       = 388
	SYS_GETRLIMIT                    = 389
	SYS_GETUID                       = 390
	SYS_IOPERM                       = 391
	SYS_IOPL                         = 392
	SYS_LCHOWN                       = 393
	SYS_PREAD64                      = 394
	SYS_PWRITE64                     = 395
	SYS_SENDFILE                     = 396
	SYS_SETFSGID                     = 397
	SYS_SETFSUID                     = 398
	SYS_SETREGID                     = 399
	SYS_SETRESGID                    = 400
	SYS_SETRESUID                    = 401
	SYS_SETRLIMIT                    = 402
	SYS_SETREUID                     = 403
	SYS_SHUTDOWN                     = 404
	SYS_SPLICE                       = 405
	SYS_STATFS                       = 406
	SYS_SYNC_FILE_RANGE              = 407
	SYS_TRUNCATE                     = 408
	SYS_ACCEPT4                      = 409
	SYS_GETGROUPS                    = 410
	SYS_GETSOCKOPT                   = 411
	SYS_SETSOCKOPT                   = 412
	SYS_SOCKETPAIR                   = 413
	SYS_GETPEERNAME                  = 414
	SYS_GETSOCKNAME                  = 415
	SYS_RECVMSG                      = 416
	SYS_SENDMSG                      = 417
)
