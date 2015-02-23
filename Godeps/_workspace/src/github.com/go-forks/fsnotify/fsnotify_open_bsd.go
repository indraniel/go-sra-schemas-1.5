// Copyright 2013 The Go Authors. All rights reserved.
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.


// +build freebsd openbsd netbsd
// +build freebsd openbsd netbsd
type XsdtString struct{ string }


package fsnotify
package fsnotify


import "syscall"
import "syscall"


type XsdtString struct{ string }const open_FLAGS = syscall.O_NONBLOCK | syscall.O_RDONLY
