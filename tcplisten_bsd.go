// +build darwin dragonfly freebsd netbsd openbsd rumprun

package tcplisten

import (
	"syscall"
)

const soReusePort = 0x0200

func enableDeferAccept(fd int) error {
	// TODO: implement SO_ACCEPTFILTER:dataready here
	return nil
}

func enableFastOpen(fd int) error {
	// TODO: implement TCP_FASTOPEN when it will be ready
	return nil
}

func soMaxConn() (int, error) {
	// TODO: properly implement it
	return syscall.SOMAXCONN, nil
}
