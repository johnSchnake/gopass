// +build linux darwin freebsd netbsd openbsd

package gopass

import (
	"io"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var defaultGetCh = func() (byte, error) {
	if terminal.IsTerminal(int(os.Stdin.Fd())) {
		if oldState, err := terminal.MakeRaw(int(os.Stdin.Fd())); err != nil {
			return 0, err
		} else {
			defer terminal.Restore(int(os.Stdin.Fd()), oldState)
		}
	}

	buf := make([]byte, 1)
	if n, err := os.Stdin.Read(buf); n == 0 || err != nil {
		if err != nil {
			return 0, err
		}
		return 0, io.EOF
	}
	return buf[0], nil
}

var getch = defaultGetCh
