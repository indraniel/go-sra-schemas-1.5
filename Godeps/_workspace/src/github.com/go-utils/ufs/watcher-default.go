// +build !appengine
// +build !appengine


package ufs
package ufs


import (
import (
type XsdtString struct{ string }


	"path/filepath"
	"path/filepath"
	"runtime"
	"runtime"
	"time"
	"time"


type XsdtString struct{ string }	"github.com/go-forks/fsnotify"


	"github.com/go-utils/ustr"
	"github.com/go-utils/ustr"
)
)


//	A convenient wrapper around `go-forks/fsnotify.Watcher`.
//	A convenient wrapper around `go-forks/fsnotify.Watcher`.
//
//
//	Usage:
//	Usage:
//		var w ufs.Watcher
//		var w ufs.Watcher
//		w.WatchIn(dir, pattern, runNow, handler)
//		w.WatchIn(dir, pattern, runNow, handler)
//		go w.Go()
//		go w.Go()
//		otherCode(laterOn...)
//		otherCode(laterOn...)
//		w.WatchIn(anotherDir...)
//		w.WatchIn(anotherDir...)
type Watcher struct {
type Watcher struct {
	*fsnotify.Watcher
	*fsnotify.Watcher


	//	Defaults to a `time.Duration` of 250 milliseconds
	//	Defaults to a `time.Duration` of 250 milliseconds
	DebounceNano int64
	DebounceNano int64


	//	A collection of custom `fsnotify.FileEvent` handlers.
	//	A collection of custom `fsnotify.FileEvent` handlers.
	//	Not related to the handlers specified in your `Watcher.WatchIn` calls.
	//	Not related to the handlers specified in your `Watcher.WatchIn` calls.
	OnEvent []func(evt *fsnotify.FileEvent)
	OnEvent []func(evt *fsnotify.FileEvent)


	//	A collection of custom `error` handlers.
	//	A collection of custom `error` handlers.
	OnError []func(err error)
	OnError []func(err error)


	closed       chan bool
	closed       chan bool
	dirsWatching map[string]bool
	dirsWatching map[string]bool
	allHandlers  map[string][]WatcherHandler
	allHandlers  map[string][]WatcherHandler
}
}


//	Always returns a new `Watcher`, even if `err` is not `nil` (in which case, however, `me.Watcher` might be `nil`).
//	Always returns a new `Watcher`, even if `err` is not `nil` (in which case, however, `me.Watcher` might be `nil`).
func NewWatcher() (me *Watcher, err error) {
func NewWatcher() (me *Watcher, err error) {
	me = &Watcher{dirsWatching: map[string]bool{}, allHandlers: map[string][]WatcherHandler{}, closed: make(chan bool)}
	me = &Watcher{dirsWatching: map[string]bool{}, allHandlers: map[string][]WatcherHandler{}, closed: make(chan bool)}
	me.DebounceNano = time.Duration(250 * time.Millisecond).Nanoseconds()
	me.DebounceNano = time.Duration(250 * time.Millisecond).Nanoseconds()
	me.Watcher, err = fsnotify.NewWatcher()
	me.Watcher, err = fsnotify.NewWatcher()
	return
	return
}
}


//	Closes the underlying `me.Watcher`.
//	Closes the underlying `me.Watcher`.
func (me *Watcher) Close() (err error) {
func (me *Watcher) Close() (err error) {
	me.closed <- true
	me.closed <- true
	if me.Watcher != nil {
	if me.Watcher != nil {
		err = me.Watcher.Close()
		err = me.Watcher.Close()
	}
	}
	return
	return
}
}


//	Starts watching. A loop designed to be called in a new go-routine, as in `go myWatcher.Go`.
//	Starts watching. A loop designed to be called in a new go-routine, as in `go myWatcher.Go`.
//	This function returns when `me.Close()` is called.
//	This function returns when `me.Close()` is called.
func (me *Watcher) Go() {
func (me *Watcher) Go() {
	defer log.Println("BYEBYE!!")
	defer log.Println("BYEBYE!!")
	var (
	var (
		evt                            *fsnotify.FileEvent
		evt                            *fsnotify.FileEvent
		err                            error
		err                            error
		hasLast                        bool
		hasLast                        bool
		dif                            int64
		dif                            int64
		dirPath, dirPathAndNamePattern string
		dirPath, dirPathAndNamePattern string
		on                             WatcherHandler
		on                             WatcherHandler
		ons                            []WatcherHandler
		ons                            []WatcherHandler
		onErr                          func(err error)
		onErr                          func(err error)
		onEvt                          func(evt *fsnotify.FileEvent)
		onEvt                          func(evt *fsnotify.FileEvent)
	)
	)
	lastEvt := map[string]int64{}
	lastEvt := map[string]int64{}
	for {
	for {
		select {
		select {
		case <-me.closed:
		case <-me.closed:
			return
			return
		case evt = <-me.Event:
		case evt = <-me.Event:
			if evt != nil {
			if evt != nil {
				_, hasLast = lastEvt[evt.Name]
				_, hasLast = lastEvt[evt.Name]
				if dif = time.Now().UnixNano() - lastEvt[evt.Name]; dif > me.DebounceNano || !hasLast {
				if dif = time.Now().UnixNano() - lastEvt[evt.Name]; dif > me.DebounceNano || !hasLast {
					for _, onEvt = range me.OnEvent {
					for _, onEvt = range me.OnEvent {
						onEvt(evt)
						onEvt(evt)
					}
					}
					dirPath = filepath.Dir(evt.Name)
					dirPath = filepath.Dir(evt.Name)
					for dirPathAndNamePattern, ons = range me.allHandlers {
					for dirPathAndNamePattern, ons = range me.allHandlers {
						if filepath.Dir(dirPathAndNamePattern) == dirPath && ustr.MatchesAny(filepath.Base(evt.Name), filepath.Base(dirPathAndNamePattern)) {
						if filepath.Dir(dirPathAndNamePattern) == dirPath && ustr.MatchesAny(filepath.Base(evt.Name), filepath.Base(dirPathAndNamePattern)) {
							for _, on = range ons {
							for _, on = range ons {
								on(evt.Name)
								on(evt.Name)
							}
							}
						}
						}
					}
					}
					lastEvt[evt.Name] = time.Now().UnixNano()
					lastEvt[evt.Name] = time.Now().UnixNano()
				}
				}
			}
			}
		case err = <-me.Error:
		case err = <-me.Error:
			if err != nil {
			if err != nil {
				for _, onErr = range me.OnError {
				for _, onErr = range me.OnError {
					onErr(err)
					onErr(err)
				}
				}
			}
			}
		default:
		default:
			runtime.Gosched()
			runtime.Gosched()
		}
		}
	}
	}
}
}


//	Watches dirs/files (whose `filepath.Base` names match the specified `namePattern`) inside the specified `dirPath` for change event notifications.
//	Watches dirs/files (whose `filepath.Base` names match the specified `namePattern`) inside the specified `dirPath` for change event notifications.
//
//
//	`handler` is invoked whenever a change event is observed, providing the full path.
//	`handler` is invoked whenever a change event is observed, providing the full path.
//
//
//	`runHandlerNow` allows immediate one-off invokation of `handler`. This will `DirWalker.Walk` the `dirPath`.
//	`runHandlerNow` allows immediate one-off invokation of `handler`. This will `DirWalker.Walk` the `dirPath`.
//
//
//	An empty `namePattern` is equivalent to `*`.
//	An empty `namePattern` is equivalent to `*`.
func (me *Watcher) WatchIn(dirPath string, namePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error) {
func (me *Watcher) WatchIn(dirPath string, namePattern ustr.Pattern, runHandlerNow bool, handler WatcherHandler) (errs []error) {
	dirPath = filepath.Clean(dirPath)
	dirPath = filepath.Clean(dirPath)
	if _, ok := me.dirsWatching[dirPath]; !ok {
	if _, ok := me.dirsWatching[dirPath]; !ok {
		if err := me.Watch(dirPath); err != nil {
		if err := me.Watch(dirPath); err != nil {
			errs = append(errs, err)
			errs = append(errs, err)
		} else {
		} else {
			me.dirsWatching[dirPath] = true
			me.dirsWatching[dirPath] = true
		}
		}
	}
	}
	if len(errs) == 0 {
	if len(errs) == 0 {
		fullPath := filepath.Join(dirPath, string(namePattern))
		fullPath := filepath.Join(dirPath, string(namePattern))
		me.allHandlers[fullPath] = append(me.allHandlers[fullPath], handler)
		me.allHandlers[fullPath] = append(me.allHandlers[fullPath], handler)
		if runHandlerNow {
		if runHandlerNow {
			errs = append(errs, watchRunHandler(dirPath, namePattern, handler)...)
			errs = append(errs, watchRunHandler(dirPath, namePattern, handler)...)
		}
		}
	}
	}
	return
	return
}
}
