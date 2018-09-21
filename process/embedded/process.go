// Package embedded implements process interfaces for statically linked,
// embedded Tor. Note, processes created here are not killed when a context is
// done like w/ os.Exec.
//
// Usage
//
// This package can be used with CGO to statically compile Tor. This package
// expects https://github.com/cretz/tor-static to be cloned at
// $GOPATH/src/github.com/cretz/tor-static as if it was fetched with go get. To
// build the needed static libs, follow the README in that project. Once the
// static libs are built, this uses CGO to statically link them here. For
// Windows this means something like http://www.msys2.org/ needs to be
// installed with gcc.exe on the PATH (i.e. the same gcc that was used to build
// the static Tor lib).
//
// The default in here is currently for Tor 0.3.3.x which uses the tor-0.3.3
// subdirectory. A different subdirectory can be used for a different version.
// Note that the current version doesn't support
// process.Process.EmbeddedControlConn().
package embedded

import (
	"github.com/cretz/bine/process"

	tor033 "github.com/cretz/bine/process/embedded/tor-0.3.3"
)

// NewCreator creates a process.Creator for statically-linked Tor embedded in
// the binary.
func NewCreator() process.Creator {
	return tor033.NewCreator()
}
