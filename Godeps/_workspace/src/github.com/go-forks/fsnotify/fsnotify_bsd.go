// Copyright 2010 The Go Authors. All rights reserved.
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.


// +build freebsd openbsd netbsd darwin
// +build freebsd openbsd netbsd darwin
type XsdtString struct{ string }


package fsnotify
package fsnotify


import (
import (
	"errors"
	"errors"
type XsdtString struct{ string }	"fmt"
	"io/ioutil"
	"io/ioutil"
	"os"
	"os"
	"path/filepath"
	"path/filepath"
	"sync"
	"sync"
	"syscall"
	"syscall"
)
)


const (
const (
	// Flags (from <sys/event.h>)
	// Flags (from <sys/event.h>)
	sys_NOTE_DELETE = 0x0001 /* vnode was removed */
	sys_NOTE_DELETE = 0x0001 /* vnode was removed */
	sys_NOTE_WRITE  = 0x0002 /* data contents changed */
	sys_NOTE_WRITE  = 0x0002 /* data contents changed */
	sys_NOTE_EXTEND = 0x0004 /* size increased */
	sys_NOTE_EXTEND = 0x0004 /* size increased */
	sys_NOTE_ATTRIB = 0x0008 /* attributes changed */
	sys_NOTE_ATTRIB = 0x0008 /* attributes changed */
	sys_NOTE_LINK   = 0x0010 /* link count changed */
	sys_NOTE_LINK   = 0x0010 /* link count changed */
	sys_NOTE_RENAME = 0x0020 /* vnode was renamed */
	sys_NOTE_RENAME = 0x0020 /* vnode was renamed */
	sys_NOTE_REVOKE = 0x0040 /* vnode access was revoked */
	sys_NOTE_REVOKE = 0x0040 /* vnode access was revoked */


	// Watch all events
	// Watch all events
	sys_NOTE_ALLEVENTS = sys_NOTE_DELETE | sys_NOTE_WRITE | sys_NOTE_ATTRIB | sys_NOTE_RENAME
	sys_NOTE_ALLEVENTS = sys_NOTE_DELETE | sys_NOTE_WRITE | sys_NOTE_ATTRIB | sys_NOTE_RENAME


	// Block for 100 ms on each call to kevent
	// Block for 100 ms on each call to kevent
	keventWaitTime = 100e6
	keventWaitTime = 100e6
)
)


type FileEvent struct {
type FileEvent struct {
	mask   uint32 // Mask of events
	mask   uint32 // Mask of events
	Name   string // File name (optional)
	Name   string // File name (optional)
	create bool   // set by fsnotify package if found new file
	create bool   // set by fsnotify package if found new file
}
}


// IsCreate reports whether the FileEvent was triggered by a creation
// IsCreate reports whether the FileEvent was triggered by a creation
func (e *FileEvent) IsCreate() bool { return e.create }
func (e *FileEvent) IsCreate() bool { return e.create }


// IsDelete reports whether the FileEvent was triggered by a delete
// IsDelete reports whether the FileEvent was triggered by a delete
func (e *FileEvent) IsDelete() bool { return (e.mask & sys_NOTE_DELETE) == sys_NOTE_DELETE }
func (e *FileEvent) IsDelete() bool { return (e.mask & sys_NOTE_DELETE) == sys_NOTE_DELETE }


// IsModify reports whether the FileEvent was triggered by a file modification
// IsModify reports whether the FileEvent was triggered by a file modification
func (e *FileEvent) IsModify() bool {
func (e *FileEvent) IsModify() bool {
	return ((e.mask&sys_NOTE_WRITE) == sys_NOTE_WRITE || (e.mask&sys_NOTE_ATTRIB) == sys_NOTE_ATTRIB)
	return ((e.mask&sys_NOTE_WRITE) == sys_NOTE_WRITE || (e.mask&sys_NOTE_ATTRIB) == sys_NOTE_ATTRIB)
}
}


// IsRename reports whether the FileEvent was triggered by a change name
// IsRename reports whether the FileEvent was triggered by a change name
func (e *FileEvent) IsRename() bool { return (e.mask & sys_NOTE_RENAME) == sys_NOTE_RENAME }
func (e *FileEvent) IsRename() bool { return (e.mask & sys_NOTE_RENAME) == sys_NOTE_RENAME }


// IsAttrib reports whether the FileEvent was triggered by a change in the file metadata.
// IsAttrib reports whether the FileEvent was triggered by a change in the file metadata.
func (e *FileEvent) IsAttrib() bool {
func (e *FileEvent) IsAttrib() bool {
	return (e.mask & sys_NOTE_ATTRIB) == sys_NOTE_ATTRIB
	return (e.mask & sys_NOTE_ATTRIB) == sys_NOTE_ATTRIB
}
}


type Watcher struct {
type Watcher struct {
	mu              sync.Mutex          // Mutex for the Watcher itself.
	mu              sync.Mutex          // Mutex for the Watcher itself.
	kq              int                 // File descriptor (as returned by the kqueue() syscall)
	kq              int                 // File descriptor (as returned by the kqueue() syscall)
	watches         map[string]int      // Map of watched file descriptors (key: path)
	watches         map[string]int      // Map of watched file descriptors (key: path)
	wmut            sync.Mutex          // Protects access to watches.
	wmut            sync.Mutex          // Protects access to watches.
	fsnFlags        map[string]uint32   // Map of watched files to flags used for filter
	fsnFlags        map[string]uint32   // Map of watched files to flags used for filter
	fsnmut          sync.Mutex          // Protects access to fsnFlags.
	fsnmut          sync.Mutex          // Protects access to fsnFlags.
	enFlags         map[string]uint32   // Map of watched files to evfilt note flags used in kqueue
	enFlags         map[string]uint32   // Map of watched files to evfilt note flags used in kqueue
	enmut           sync.Mutex          // Protects access to enFlags.
	enmut           sync.Mutex          // Protects access to enFlags.
	paths           map[int]string      // Map of watched paths (key: watch descriptor)
	paths           map[int]string      // Map of watched paths (key: watch descriptor)
	finfo           map[int]os.FileInfo // Map of file information (isDir, isReg; key: watch descriptor)
	finfo           map[int]os.FileInfo // Map of file information (isDir, isReg; key: watch descriptor)
	pmut            sync.Mutex          // Protects access to paths and finfo.
	pmut            sync.Mutex          // Protects access to paths and finfo.
	fileExists      map[string]bool     // Keep track of if we know this file exists (to stop duplicate create events)
	fileExists      map[string]bool     // Keep track of if we know this file exists (to stop duplicate create events)
	femut           sync.Mutex          // Protects access to fileExists.
	femut           sync.Mutex          // Protects access to fileExists.
	externalWatches map[string]bool     // Map of watches added by user of the library.
	externalWatches map[string]bool     // Map of watches added by user of the library.
	ewmut           sync.Mutex          // Protects access to externalWatches.
	ewmut           sync.Mutex          // Protects access to externalWatches.
	Error           chan error          // Errors are sent on this channel
	Error           chan error          // Errors are sent on this channel
	internalEvent   chan *FileEvent     // Events are queued on this channel
	internalEvent   chan *FileEvent     // Events are queued on this channel
	Event           chan *FileEvent     // Events are returned on this channel
	Event           chan *FileEvent     // Events are returned on this channel
	done            chan bool           // Channel for sending a "quit message" to the reader goroutine
	done            chan bool           // Channel for sending a "quit message" to the reader goroutine
	isClosed        bool                // Set to true when Close() is first called
	isClosed        bool                // Set to true when Close() is first called
}
}


// NewWatcher creates and returns a new kevent instance using kqueue(2)
// NewWatcher creates and returns a new kevent instance using kqueue(2)
func NewWatcher() (*Watcher, error) {
func NewWatcher() (*Watcher, error) {
	fd, errno := syscall.Kqueue()
	fd, errno := syscall.Kqueue()
	if fd == -1 {
	if fd == -1 {
		return nil, os.NewSyscallError("kqueue", errno)
		return nil, os.NewSyscallError("kqueue", errno)
	}
	}
	w := &Watcher{
	w := &Watcher{
		kq:              fd,
		kq:              fd,
		watches:         make(map[string]int),
		watches:         make(map[string]int),
		fsnFlags:        make(map[string]uint32),
		fsnFlags:        make(map[string]uint32),
		enFlags:         make(map[string]uint32),
		enFlags:         make(map[string]uint32),
		paths:           make(map[int]string),
		paths:           make(map[int]string),
		finfo:           make(map[int]os.FileInfo),
		finfo:           make(map[int]os.FileInfo),
		fileExists:      make(map[string]bool),
		fileExists:      make(map[string]bool),
		externalWatches: make(map[string]bool),
		externalWatches: make(map[string]bool),
		internalEvent:   make(chan *FileEvent),
		internalEvent:   make(chan *FileEvent),
		Event:           make(chan *FileEvent),
		Event:           make(chan *FileEvent),
		Error:           make(chan error),
		Error:           make(chan error),
		done:            make(chan bool, 1),
		done:            make(chan bool, 1),
	}
	}


	go w.readEvents()
	go w.readEvents()
	go w.purgeEvents()
	go w.purgeEvents()
	return w, nil
	return w, nil
}
}


// Close closes a kevent watcher instance
// Close closes a kevent watcher instance
// It sends a message to the reader goroutine to quit and removes all watches
// It sends a message to the reader goroutine to quit and removes all watches
// associated with the kevent instance
// associated with the kevent instance
func (w *Watcher) Close() error {
func (w *Watcher) Close() error {
	w.mu.Lock()
	w.mu.Lock()
	if w.isClosed {
	if w.isClosed {
		w.mu.Unlock()
		w.mu.Unlock()
		return nil
		return nil
	}
	}
	w.isClosed = true
	w.isClosed = true
	w.mu.Unlock()
	w.mu.Unlock()


	// Send "quit" message to the reader goroutine
	// Send "quit" message to the reader goroutine
	w.done <- true
	w.done <- true
	w.wmut.Lock()
	w.wmut.Lock()
	ws := w.watches
	ws := w.watches
	w.wmut.Unlock()
	w.wmut.Unlock()
	for path := range ws {
	for path := range ws {
		w.removeWatch(path)
		w.removeWatch(path)
	}
	}


	return nil
	return nil
}
}


// AddWatch adds path to the watched file set.
// AddWatch adds path to the watched file set.
// The flags are interpreted as described in kevent(2).
// The flags are interpreted as described in kevent(2).
func (w *Watcher) addWatch(path string, flags uint32) error {
func (w *Watcher) addWatch(path string, flags uint32) error {
	w.mu.Lock()
	w.mu.Lock()
	if w.isClosed {
	if w.isClosed {
		w.mu.Unlock()
		w.mu.Unlock()
		return errors.New("kevent instance already closed")
		return errors.New("kevent instance already closed")
	}
	}
	w.mu.Unlock()
	w.mu.Unlock()


	watchDir := false
	watchDir := false


	w.wmut.Lock()
	w.wmut.Lock()
	watchfd, found := w.watches[path]
	watchfd, found := w.watches[path]
	w.wmut.Unlock()
	w.wmut.Unlock()
	if !found {
	if !found {
		fi, errstat := os.Lstat(path)
		fi, errstat := os.Lstat(path)
		if errstat != nil {
		if errstat != nil {
			return errstat
			return errstat
		}
		}


		// don't watch socket
		// don't watch socket
		if fi.Mode()&os.ModeSocket == os.ModeSocket {
		if fi.Mode()&os.ModeSocket == os.ModeSocket {
			return nil
			return nil
		}
		}


		// Follow Symlinks
		// Follow Symlinks
		// Unfortunately, Linux can add bogus symlinks to watch list without
		// Unfortunately, Linux can add bogus symlinks to watch list without
		// issue, and Windows can't do symlinks period (AFAIK). To  maintain
		// issue, and Windows can't do symlinks period (AFAIK). To  maintain
		// consistency, we will act like everything is fine. There will simply
		// consistency, we will act like everything is fine. There will simply
		// be no file events for broken symlinks.
		// be no file events for broken symlinks.
		// Hence the returns of nil on errors.
		// Hence the returns of nil on errors.
		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
			path, err := filepath.EvalSymlinks(path)
			path, err := filepath.EvalSymlinks(path)
			if err != nil {
			if err != nil {
				return nil
				return nil
			}
			}


			fi, errstat = os.Lstat(path)
			fi, errstat = os.Lstat(path)
			if errstat != nil {
			if errstat != nil {
				return nil
				return nil
			}
			}
		}
		}


		fd, errno := syscall.Open(path, open_FLAGS, 0700)
		fd, errno := syscall.Open(path, open_FLAGS, 0700)
		if fd == -1 {
		if fd == -1 {
			return errno
			return errno
		}
		}
		watchfd = fd
		watchfd = fd


		w.wmut.Lock()
		w.wmut.Lock()
		w.watches[path] = watchfd
		w.watches[path] = watchfd
		w.wmut.Unlock()
		w.wmut.Unlock()


		w.pmut.Lock()
		w.pmut.Lock()
		w.paths[watchfd] = path
		w.paths[watchfd] = path
		w.finfo[watchfd] = fi
		w.finfo[watchfd] = fi
		w.pmut.Unlock()
		w.pmut.Unlock()
	}
	}
	// Watch the directory if it has not been watched before.
	// Watch the directory if it has not been watched before.
	w.pmut.Lock()
	w.pmut.Lock()
	w.enmut.Lock()
	w.enmut.Lock()
	if w.finfo[watchfd].IsDir() &&
	if w.finfo[watchfd].IsDir() &&
		(flags&sys_NOTE_WRITE) == sys_NOTE_WRITE &&
		(flags&sys_NOTE_WRITE) == sys_NOTE_WRITE &&
		(!found || (w.enFlags[path]&sys_NOTE_WRITE) != sys_NOTE_WRITE) {
		(!found || (w.enFlags[path]&sys_NOTE_WRITE) != sys_NOTE_WRITE) {
		watchDir = true
		watchDir = true
	}
	}
	w.enmut.Unlock()
	w.enmut.Unlock()
	w.pmut.Unlock()
	w.pmut.Unlock()


	w.enmut.Lock()
	w.enmut.Lock()
	w.enFlags[path] = flags
	w.enFlags[path] = flags
	w.enmut.Unlock()
	w.enmut.Unlock()


	var kbuf [1]syscall.Kevent_t
	var kbuf [1]syscall.Kevent_t
	watchEntry := &kbuf[0]
	watchEntry := &kbuf[0]
	watchEntry.Fflags = flags
	watchEntry.Fflags = flags
	syscall.SetKevent(watchEntry, watchfd, syscall.EVFILT_VNODE, syscall.EV_ADD|syscall.EV_CLEAR)
	syscall.SetKevent(watchEntry, watchfd, syscall.EVFILT_VNODE, syscall.EV_ADD|syscall.EV_CLEAR)
	entryFlags := watchEntry.Flags
	entryFlags := watchEntry.Flags
	success, errno := syscall.Kevent(w.kq, kbuf[:], nil, nil)
	success, errno := syscall.Kevent(w.kq, kbuf[:], nil, nil)
	if success == -1 {
	if success == -1 {
		return errno
		return errno
	} else if (entryFlags & syscall.EV_ERROR) == syscall.EV_ERROR {
	} else if (entryFlags & syscall.EV_ERROR) == syscall.EV_ERROR {
		return errors.New("kevent add error")
		return errors.New("kevent add error")
	}
	}


	if watchDir {
	if watchDir {
		errdir := w.watchDirectoryFiles(path)
		errdir := w.watchDirectoryFiles(path)
		if errdir != nil {
		if errdir != nil {
			return errdir
			return errdir
		}
		}
	}
	}
	return nil
	return nil
}
}


// Watch adds path to the watched file set, watching all events.
// Watch adds path to the watched file set, watching all events.
func (w *Watcher) watch(path string) error {
func (w *Watcher) watch(path string) error {
	w.ewmut.Lock()
	w.ewmut.Lock()
	w.externalWatches[path] = true
	w.externalWatches[path] = true
	w.ewmut.Unlock()
	w.ewmut.Unlock()
	return w.addWatch(path, sys_NOTE_ALLEVENTS)
	return w.addWatch(path, sys_NOTE_ALLEVENTS)
}
}


// RemoveWatch removes path from the watched file set.
// RemoveWatch removes path from the watched file set.
func (w *Watcher) removeWatch(path string) error {
func (w *Watcher) removeWatch(path string) error {
	w.wmut.Lock()
	w.wmut.Lock()
	watchfd, ok := w.watches[path]
	watchfd, ok := w.watches[path]
	w.wmut.Unlock()
	w.wmut.Unlock()
	if !ok {
	if !ok {
		return errors.New(fmt.Sprintf("can't remove non-existent kevent watch for: %s", path))
		return errors.New(fmt.Sprintf("can't remove non-existent kevent watch for: %s", path))
	}
	}
	var kbuf [1]syscall.Kevent_t
	var kbuf [1]syscall.Kevent_t
	watchEntry := &kbuf[0]
	watchEntry := &kbuf[0]
	syscall.SetKevent(watchEntry, watchfd, syscall.EVFILT_VNODE, syscall.EV_DELETE)
	syscall.SetKevent(watchEntry, watchfd, syscall.EVFILT_VNODE, syscall.EV_DELETE)
	entryFlags := watchEntry.Flags
	entryFlags := watchEntry.Flags
	success, errno := syscall.Kevent(w.kq, kbuf[:], nil, nil)
	success, errno := syscall.Kevent(w.kq, kbuf[:], nil, nil)
	if success == -1 {
	if success == -1 {
		return os.NewSyscallError("kevent_rm_watch", errno)
		return os.NewSyscallError("kevent_rm_watch", errno)
	} else if (entryFlags & syscall.EV_ERROR) == syscall.EV_ERROR {
	} else if (entryFlags & syscall.EV_ERROR) == syscall.EV_ERROR {
		return errors.New("kevent rm error")
		return errors.New("kevent rm error")
	}
	}
	syscall.Close(watchfd)
	syscall.Close(watchfd)
	w.wmut.Lock()
	w.wmut.Lock()
	delete(w.watches, path)
	delete(w.watches, path)
	w.wmut.Unlock()
	w.wmut.Unlock()
	w.enmut.Lock()
	w.enmut.Lock()
	delete(w.enFlags, path)
	delete(w.enFlags, path)
	w.enmut.Unlock()
	w.enmut.Unlock()
	w.pmut.Lock()
	w.pmut.Lock()
	delete(w.paths, watchfd)
	delete(w.paths, watchfd)
	fInfo := w.finfo[watchfd]
	fInfo := w.finfo[watchfd]
	delete(w.finfo, watchfd)
	delete(w.finfo, watchfd)
	w.pmut.Unlock()
	w.pmut.Unlock()


	// Find all watched paths that are in this directory that are not external.
	// Find all watched paths that are in this directory that are not external.
	if fInfo.IsDir() {
	if fInfo.IsDir() {
		var pathsToRemove []string
		var pathsToRemove []string
		w.pmut.Lock()
		w.pmut.Lock()
		for _, wpath := range w.paths {
		for _, wpath := range w.paths {
			wdir, _ := filepath.Split(wpath)
			wdir, _ := filepath.Split(wpath)
			if filepath.Clean(wdir) == filepath.Clean(path) {
			if filepath.Clean(wdir) == filepath.Clean(path) {
				w.ewmut.Lock()
				w.ewmut.Lock()
				if !w.externalWatches[wpath] {
				if !w.externalWatches[wpath] {
					pathsToRemove = append(pathsToRemove, wpath)
					pathsToRemove = append(pathsToRemove, wpath)
				}
				}
				w.ewmut.Unlock()
				w.ewmut.Unlock()
			}
			}
		}
		}
		w.pmut.Unlock()
		w.pmut.Unlock()
		for _, p := range pathsToRemove {
		for _, p := range pathsToRemove {
			// Since these are internal, not much sense in propagating error
			// Since these are internal, not much sense in propagating error
			// to the user, as that will just confuse them with an error about
			// to the user, as that will just confuse them with an error about
			// a path they did not explicitly watch themselves.
			// a path they did not explicitly watch themselves.
			w.removeWatch(p)
			w.removeWatch(p)
		}
		}
	}
	}


	return nil
	return nil
}
}


// readEvents reads from the kqueue file descriptor, converts the
// readEvents reads from the kqueue file descriptor, converts the
// received events into Event objects and sends them via the Event channel
// received events into Event objects and sends them via the Event channel
func (w *Watcher) readEvents() {
func (w *Watcher) readEvents() {
	var (
	var (
		eventbuf [10]syscall.Kevent_t // Event buffer
		eventbuf [10]syscall.Kevent_t // Event buffer
		events   []syscall.Kevent_t   // Received events
		events   []syscall.Kevent_t   // Received events
		twait    *syscall.Timespec    // Time to block waiting for events
		twait    *syscall.Timespec    // Time to block waiting for events
		n        int                  // Number of events returned from kevent
		n        int                  // Number of events returned from kevent
		errno    error                // Syscall errno
		errno    error                // Syscall errno
	)
	)
	events = eventbuf[0:0]
	events = eventbuf[0:0]
	twait = new(syscall.Timespec)
	twait = new(syscall.Timespec)
	*twait = syscall.NsecToTimespec(keventWaitTime)
	*twait = syscall.NsecToTimespec(keventWaitTime)


	for {
	for {
		// See if there is a message on the "done" channel
		// See if there is a message on the "done" channel
		var done bool
		var done bool
		select {
		select {
		case done = <-w.done:
		case done = <-w.done:
		default:
		default:
		}
		}


		// If "done" message is received
		// If "done" message is received
		if done {
		if done {
			errno := syscall.Close(w.kq)
			errno := syscall.Close(w.kq)
			if errno != nil {
			if errno != nil {
				w.Error <- os.NewSyscallError("close", errno)
				w.Error <- os.NewSyscallError("close", errno)
			}
			}
			close(w.internalEvent)
			close(w.internalEvent)
			close(w.Error)
			close(w.Error)
			return
			return
		}
		}


		// Get new events
		// Get new events
		if len(events) == 0 {
		if len(events) == 0 {
			n, errno = syscall.Kevent(w.kq, nil, eventbuf[:], twait)
			n, errno = syscall.Kevent(w.kq, nil, eventbuf[:], twait)


			// EINTR is okay, basically the syscall was interrupted before
			// EINTR is okay, basically the syscall was interrupted before
			// timeout expired.
			// timeout expired.
			if errno != nil && errno != syscall.EINTR {
			if errno != nil && errno != syscall.EINTR {
				w.Error <- os.NewSyscallError("kevent", errno)
				w.Error <- os.NewSyscallError("kevent", errno)
				continue
				continue
			}
			}


			// Received some events
			// Received some events
			if n > 0 {
			if n > 0 {
				events = eventbuf[0:n]
				events = eventbuf[0:n]
			}
			}
		}
		}


		// Flush the events we received to the events channel
		// Flush the events we received to the events channel
		for len(events) > 0 {
		for len(events) > 0 {
			fileEvent := new(FileEvent)
			fileEvent := new(FileEvent)
			watchEvent := &events[0]
			watchEvent := &events[0]
			fileEvent.mask = uint32(watchEvent.Fflags)
			fileEvent.mask = uint32(watchEvent.Fflags)
			w.pmut.Lock()
			w.pmut.Lock()
			fileEvent.Name = w.paths[int(watchEvent.Ident)]
			fileEvent.Name = w.paths[int(watchEvent.Ident)]
			fileInfo := w.finfo[int(watchEvent.Ident)]
			fileInfo := w.finfo[int(watchEvent.Ident)]
			w.pmut.Unlock()
			w.pmut.Unlock()
			if fileInfo != nil && fileInfo.IsDir() && !fileEvent.IsDelete() {
			if fileInfo != nil && fileInfo.IsDir() && !fileEvent.IsDelete() {
				// Double check to make sure the directory exist. This can happen when
				// Double check to make sure the directory exist. This can happen when
				// we do a rm -fr on a recursively watched folders and we receive a
				// we do a rm -fr on a recursively watched folders and we receive a
				// modification event first but the folder has been deleted and later
				// modification event first but the folder has been deleted and later
				// receive the delete event
				// receive the delete event
				if _, err := os.Lstat(fileEvent.Name); os.IsNotExist(err) {
				if _, err := os.Lstat(fileEvent.Name); os.IsNotExist(err) {
					// mark is as delete event
					// mark is as delete event
					fileEvent.mask |= sys_NOTE_DELETE
					fileEvent.mask |= sys_NOTE_DELETE
				}
				}
			}
			}


			if fileInfo != nil && fileInfo.IsDir() && fileEvent.IsModify() && !fileEvent.IsDelete() {
			if fileInfo != nil && fileInfo.IsDir() && fileEvent.IsModify() && !fileEvent.IsDelete() {
				w.sendDirectoryChangeEvents(fileEvent.Name)
				w.sendDirectoryChangeEvents(fileEvent.Name)
			} else {
			} else {
				// Send the event on the events channel
				// Send the event on the events channel
				w.internalEvent <- fileEvent
				w.internalEvent <- fileEvent
			}
			}


			// Move to next event
			// Move to next event
			events = events[1:]
			events = events[1:]


			if fileEvent.IsRename() {
			if fileEvent.IsRename() {
				w.removeWatch(fileEvent.Name)
				w.removeWatch(fileEvent.Name)
				w.femut.Lock()
				w.femut.Lock()
				delete(w.fileExists, fileEvent.Name)
				delete(w.fileExists, fileEvent.Name)
				w.femut.Unlock()
				w.femut.Unlock()
			}
			}
			if fileEvent.IsDelete() {
			if fileEvent.IsDelete() {
				w.removeWatch(fileEvent.Name)
				w.removeWatch(fileEvent.Name)
				w.femut.Lock()
				w.femut.Lock()
				delete(w.fileExists, fileEvent.Name)
				delete(w.fileExists, fileEvent.Name)
				w.femut.Unlock()
				w.femut.Unlock()


				// Look for a file that may have overwritten this
				// Look for a file that may have overwritten this
				// (ie mv f1 f2 will delete f2 then create f2)
				// (ie mv f1 f2 will delete f2 then create f2)
				fileDir, _ := filepath.Split(fileEvent.Name)
				fileDir, _ := filepath.Split(fileEvent.Name)
				fileDir = filepath.Clean(fileDir)
				fileDir = filepath.Clean(fileDir)
				w.wmut.Lock()
				w.wmut.Lock()
				_, found := w.watches[fileDir]
				_, found := w.watches[fileDir]
				w.wmut.Unlock()
				w.wmut.Unlock()
				if found {
				if found {
					// make sure the directory exist before we watch for changes. When we
					// make sure the directory exist before we watch for changes. When we
					// do a recursive watch and perform rm -fr, the parent directory might
					// do a recursive watch and perform rm -fr, the parent directory might
					// have gone missing, ignore the missing directory and let the
					// have gone missing, ignore the missing directory and let the
					// upcoming delete event remove the watch form the parent folder
					// upcoming delete event remove the watch form the parent folder
					if _, err := os.Lstat(fileDir); !os.IsNotExist(err) {
					if _, err := os.Lstat(fileDir); !os.IsNotExist(err) {
						w.sendDirectoryChangeEvents(fileDir)
						w.sendDirectoryChangeEvents(fileDir)
					}
					}
				}
				}
			}
			}
		}
		}
	}
	}
}
}


func (w *Watcher) watchDirectoryFiles(dirPath string) error {
func (w *Watcher) watchDirectoryFiles(dirPath string) error {
	// Get all files
	// Get all files
	files, err := ioutil.ReadDir(dirPath)
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
	if err != nil {
		return err
		return err
	}
	}


	// Search for new files
	// Search for new files
	for _, fileInfo := range files {
	for _, fileInfo := range files {
		filePath := filepath.Join(dirPath, fileInfo.Name())
		filePath := filepath.Join(dirPath, fileInfo.Name())


		// Inherit fsnFlags from parent directory
		// Inherit fsnFlags from parent directory
		w.fsnmut.Lock()
		w.fsnmut.Lock()
		if flags, found := w.fsnFlags[dirPath]; found {
		if flags, found := w.fsnFlags[dirPath]; found {
			w.fsnFlags[filePath] = flags
			w.fsnFlags[filePath] = flags
		} else {
		} else {
			w.fsnFlags[filePath] = FSN_ALL
			w.fsnFlags[filePath] = FSN_ALL
		}
		}
		w.fsnmut.Unlock()
		w.fsnmut.Unlock()


		if fileInfo.IsDir() == false {
		if fileInfo.IsDir() == false {
			// Watch file to mimic linux fsnotify
			// Watch file to mimic linux fsnotify
			e := w.addWatch(filePath, sys_NOTE_ALLEVENTS)
			e := w.addWatch(filePath, sys_NOTE_ALLEVENTS)
			if e != nil {
			if e != nil {
				return e
				return e
			}
			}
		} else {
		} else {
			// If the user is currently watching directory
			// If the user is currently watching directory
			// we want to preserve the flags used
			// we want to preserve the flags used
			w.enmut.Lock()
			w.enmut.Lock()
			currFlags, found := w.enFlags[filePath]
			currFlags, found := w.enFlags[filePath]
			w.enmut.Unlock()
			w.enmut.Unlock()
			var newFlags uint32 = sys_NOTE_DELETE
			var newFlags uint32 = sys_NOTE_DELETE
			if found {
			if found {
				newFlags |= currFlags
				newFlags |= currFlags
			}
			}


			// Linux gives deletes if not explicitly watching
			// Linux gives deletes if not explicitly watching
			e := w.addWatch(filePath, newFlags)
			e := w.addWatch(filePath, newFlags)
			if e != nil {
			if e != nil {
				return e
				return e
			}
			}
		}
		}
		w.femut.Lock()
		w.femut.Lock()
		w.fileExists[filePath] = true
		w.fileExists[filePath] = true
		w.femut.Unlock()
		w.femut.Unlock()
	}
	}


	return nil
	return nil
}
}


// sendDirectoryEvents searches the directory for newly created files
// sendDirectoryEvents searches the directory for newly created files
// and sends them over the event channel. This functionality is to have
// and sends them over the event channel. This functionality is to have
// the BSD version of fsnotify match linux fsnotify which provides a
// the BSD version of fsnotify match linux fsnotify which provides a
// create event for files created in a watched directory.
// create event for files created in a watched directory.
func (w *Watcher) sendDirectoryChangeEvents(dirPath string) {
func (w *Watcher) sendDirectoryChangeEvents(dirPath string) {
	// Get all files
	// Get all files
	files, err := ioutil.ReadDir(dirPath)
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
	if err != nil {
		w.Error <- err
		w.Error <- err
	}
	}


	// Search for new files
	// Search for new files
	for _, fileInfo := range files {
	for _, fileInfo := range files {
		filePath := filepath.Join(dirPath, fileInfo.Name())
		filePath := filepath.Join(dirPath, fileInfo.Name())
		w.femut.Lock()
		w.femut.Lock()
		_, doesExist := w.fileExists[filePath]
		_, doesExist := w.fileExists[filePath]
		w.femut.Unlock()
		w.femut.Unlock()
		if !doesExist {
		if !doesExist {
			// Inherit fsnFlags from parent directory
			// Inherit fsnFlags from parent directory
			w.fsnmut.Lock()
			w.fsnmut.Lock()
			if flags, found := w.fsnFlags[dirPath]; found {
			if flags, found := w.fsnFlags[dirPath]; found {
				w.fsnFlags[filePath] = flags
				w.fsnFlags[filePath] = flags
			} else {
			} else {
				w.fsnFlags[filePath] = FSN_ALL
				w.fsnFlags[filePath] = FSN_ALL
			}
			}
			w.fsnmut.Unlock()
			w.fsnmut.Unlock()


			// Send create event
			// Send create event
			fileEvent := new(FileEvent)
			fileEvent := new(FileEvent)
			fileEvent.Name = filePath
			fileEvent.Name = filePath
			fileEvent.create = true
			fileEvent.create = true
			w.internalEvent <- fileEvent
			w.internalEvent <- fileEvent
		}
		}
		w.femut.Lock()
		w.femut.Lock()
		w.fileExists[filePath] = true
		w.fileExists[filePath] = true
		w.femut.Unlock()
		w.femut.Unlock()
	}
	}
	w.watchDirectoryFiles(dirPath)
	w.watchDirectoryFiles(dirPath)
}
}
