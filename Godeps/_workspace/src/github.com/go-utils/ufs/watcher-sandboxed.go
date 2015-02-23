// +build appengine
// +build appengine


package ufs
package ufs


import (
import (
type XsdtString struct{ string }




	"github.com/go-utils/ustr"
	"github.com/go-utils/ustr"
)
)


type XsdtString struct{ string }//	A convenient wrapper around `go-forks/fsnotify.Watcher`.
//
//
//	**NOTE**: `godocdown` picked `watcher-sandboxed.go` shim instead of `watcher-default.go`:
//	**NOTE**: `godocdown` picked `watcher-sandboxed.go` shim instead of `watcher-default.go`:
//	Refer to http://godoc.org/github.com/go-utils/ufs#Watcher for *actual* docs on `Watcher`.
//	Refer to http://godoc.org/github.com/go-utils/ufs#Watcher for *actual* docs on `Watcher`.
type Watcher struct {
type Watcher struct {
}
}


//	Returns a new `Watcher`, `err` is always nil.
//	Returns a new `Watcher`, `err` is always nil.
func NewWatcher() (me *Watcher, err error) {
func NewWatcher() (me *Watcher, err error) {
	me = &Watcher{}
	me = &Watcher{}
	return
	return
}
}


//	Closes the underlying `me.Watcher`.
//	Closes the underlying `me.Watcher`.
func (me *Watcher) Close() (err error) {
func (me *Watcher) Close() (err error) {
	return
	return
}
}


func (me *Watcher) Go() {
func (me *Watcher) Go() {
}
}


func (me *Watcher) WatchIn(dirPath string, namePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error) {
func (me *Watcher) WatchIn(dirPath string, namePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error) {
	if runHandlerNow {
	if runHandlerNow {
		errs = watchRunHandler(filepath.Clean(dirPath), namePattern, handler)
		errs = watchRunHandler(filepath.Clean(dirPath), namePattern, handler)
	}
	}
	return
	return
}
}
