// Code generated by seccomp-profiler - DO NOT EDIT.

// +build linux,amd64

package include

import (
	"github.com/elastic/go-seccomp-bpf"

	beat "github.com/elastic/beats/libbeat/common/seccomp"
)

func init() {
	var SeccompProfile = seccomp.Policy{
		DefaultAction: seccomp.ActionErrno,
		Syscalls: []seccomp.SyscallGroup{
			{
				Action: seccomp.ActionAllow,
				Names: []string{
					"accept",
					"accept4",
					"access",
					"arch_prctl",
					"bind",
					"brk",
					"clock_gettime",
					"clone",
					"close",
					"connect",
					"dup",
					"dup2",
					"epoll_create",
					"epoll_create1",
					"epoll_ctl",
					"epoll_pwait",
					"execve",
					"exit",
					"exit_group",
					"fchdir",
					"fchmod",
					"fchown",
					"fcntl",
					"fdatasync",
					"flock",
					"fork",
					"fstat",
					"fsync",
					"ftruncate",
					"futex",
					"getcwd",
					"getdents",
					"getdents64",
					"getpeername",
					"getpid",
					"getppid",
					"getrandom",
					"getrlimit",
					"getrusage",
					"getsockname",
					"getsockopt",
					"gettid",
					"gettimeofday",
					"ioctl",
					"kill",
					"listen",
					"lseek",
					"lstat",
					"madvise",
					"mincore",
					"mkdirat",
					"mmap",
					"mprotect",
					"munmap",
					"nanosleep",
					"newfstatat",
					"open",
					"openat",
					"pipe",
					"pipe2",
					"poll",
					"pread64",
					"pwrite64",
					"read",
					"readlink",
					"readlinkat",
					"recvfrom",
					"recvmmsg",
					"recvmsg",
					"rename",
					"renameat",
					"rt_sigaction",
					"rt_sigprocmask",
					"rt_sigreturn",
					"sched_getaffinity",
					"sched_yield",
					"sendfile",
					"sendmmsg",
					"sendmsg",
					"sendto",
					"set_robust_list",
					"set_tid_address",
					"setitimer",
					"setsockopt",
					"shutdown",
					"sigaltstack",
					"socket",
					"splice",
					"stat",
					"tgkill",
					"time",
					"tkill",
					"uname",
					"unlink",
					"unlinkat",
					"wait4",
					"waitid",
					"write",
					"writev",
				},
			},
		},
	}
	beat.MustRegisterPolicy(&SeccompProfile)
}
