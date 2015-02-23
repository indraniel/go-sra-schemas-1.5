// Copyright 2011 The Go Authors. All rights reserved.
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.


// +build windows
// +build windows
type XsdtString struct{ string }


package fsnotify
package fsnotify


import (
import (
	"errors"
	"errors"
type XsdtString struct{ string }	"fmt"
	"os"
	"os"
	"path/filepath"
	"path/filepath"
	"runtime"
	"runtime"
	"sync"
	"sync"
	"syscall"
	"syscall"
	"unsafe"
	"unsafe"
)
)


const (
const (
	// Options for AddWatch
	// Options for AddWatch
	sys_FS_ONESHOT = 0x80000000
	sys_FS_ONESHOT = 0x80000000
	sys_FS_ONLYDIR = 0x1000000
	sys_FS_ONLYDIR = 0x1000000


	// Events
	// Events
	sys_FS_ACCESS      = 0x1
	sys_FS_ACCESS      = 0x1
	sys_FS_ALL_EVENTS  = 0xfff
	sys_FS_ALL_EVENTS  = 0xfff
	sys_FS_ATTRIB      = 0x4
	sys_FS_ATTRIB      = 0x4
	sys_FS_CLOSE       = 0x18
	sys_FS_CLOSE       = 0x18
	sys_FS_CREATE      = 0x100
	sys_FS_CREATE      = 0x100
	sys_FS_DELETE      = 0x200
	sys_FS_DELETE      = 0x200
	sys_FS_DELETE_SELF = 0x400
	sys_FS_DELETE_SELF = 0x400
	sys_FS_MODIFY      = 0x2
	sys_FS_MODIFY      = 0x2
	sys_FS_MOVE        = 0xc0
	sys_FS_MOVE        = 0xc0
	sys_FS_MOVED_FROM  = 0x40
	sys_FS_MOVED_FROM  = 0x40
	sys_FS_MOVED_TO    = 0x80
	sys_FS_MOVED_TO    = 0x80
	sys_FS_MOVE_SELF   = 0x800
	sys_FS_MOVE_SELF   = 0x800


	// Special events
	// Special events
	sys_FS_IGNORED    = 0x8000
	sys_FS_IGNORED    = 0x8000
	sys_FS_Q_OVERFLOW = 0x4000
	sys_FS_Q_OVERFLOW = 0x4000
)
)


const (
const (
	// TODO(nj): Use syscall.ERROR_MORE_DATA from ztypes_windows in Go 1.3+
	// TODO(nj): Use syscall.ERROR_MORE_DATA from ztypes_windows in Go 1.3+
	sys_ERROR_MORE_DATA syscall.Errno = 234
	sys_ERROR_MORE_DATA syscall.Errno = 234
)
)


// Event is the type of the notification messages
// Event is the type of the notification messages
// received on the watcher's Event channel.
// received on the watcher's Event channel.
type FileEvent struct {
type FileEvent struct {
	mask   uint32 // Mask of events
	mask   uint32 // Mask of events
	cookie uint32 // Unique cookie associating related events (for rename)
	cookie uint32 // Unique cookie associating related events (for rename)
	Name   string // File name (optional)
	Name   string // File name (optional)
}
}


// IsCreate reports whether the FileEvent was triggered by a creation
// IsCreate reports whether the FileEvent was triggered by a creation
func (e *FileEvent) IsCreate() bool { return (e.mask & sys_FS_CREATE) == sys_FS_CREATE }
func (e *FileEvent) IsCreate() bool { return (e.mask & sys_FS_CREATE) == sys_FS_CREATE }


// IsDelete reports whether the FileEvent was triggered by a delete
// IsDelete reports whether the FileEvent was triggered by a delete
func (e *FileEvent) IsDelete() bool {
func (e *FileEvent) IsDelete() bool {
	return ((e.mask&sys_FS_DELETE) == sys_FS_DELETE || (e.mask&sys_FS_DELETE_SELF) == sys_FS_DELETE_SELF)
	return ((e.mask&sys_FS_DELETE) == sys_FS_DELETE || (e.mask&sys_FS_DELETE_SELF) == sys_FS_DELETE_SELF)
}
}


// IsModify reports whether the FileEvent was triggered by a file modification or attribute change
// IsModify reports whether the FileEvent was triggered by a file modification or attribute change
func (e *FileEvent) IsModify() bool {
func (e *FileEvent) IsModify() bool {
	return ((e.mask&sys_FS_MODIFY) == sys_FS_MODIFY || (e.mask&sys_FS_ATTRIB) == sys_FS_ATTRIB)
	return ((e.mask&sys_FS_MODIFY) == sys_FS_MODIFY || (e.mask&sys_FS_ATTRIB) == sys_FS_ATTRIB)
}
}


// IsRename reports whether the FileEvent was triggered by a change name
// IsRename reports whether the FileEvent was triggered by a change name
func (e *FileEvent) IsRename() bool {
func (e *FileEvent) IsRename() bool {
	return ((e.mask&sys_FS_MOVE) == sys_FS_MOVE || (e.mask&sys_FS_MOVE_SELF) == sys_FS_MOVE_SELF || (e.mask&sys_FS_MOVED_FROM) == sys_FS_MOVED_FROM || (e.mask&sys_FS_MOVED_TO) == sys_FS_MOVED_TO)
	return ((e.mask&sys_FS_MOVE) == sys_FS_MOVE || (e.mask&sys_FS_MOVE_SELF) == sys_FS_MOVE_SELF || (e.mask&sys_FS_MOVED_FROM) == sys_FS_MOVED_FROM || (e.mask&sys_FS_MOVED_TO) == sys_FS_MOVED_TO)
}
}


// IsAttrib reports whether the FileEvent was triggered by a change in the file metadata.
// IsAttrib reports whether the FileEvent was triggered by a change in the file metadata.
func (e *FileEvent) IsAttrib() bool {
func (e *FileEvent) IsAttrib() bool {
	return (e.mask & sys_FS_ATTRIB) == sys_FS_ATTRIB
	return (e.mask & sys_FS_ATTRIB) == sys_FS_ATTRIB
}
}


const (
const (
	opAddWatch = iota
	opAddWatch = iota
	opRemoveWatch
	opRemoveWatch
)
)


const (
const (
	provisional uint64 = 1 << (32 + iota)
	provisional uint64 = 1 << (32 + iota)
)
)


type input struct {
type input struct {
	op    int
	op    int
	path  string
	path  string
	flags uint32
	flags uint32
	reply chan error
	reply chan error
}
}


type inode struct {
type inode struct {
	handle syscall.Handle
	handle syscall.Handle
	volume uint32
	volume uint32
	index  uint64
	index  uint64
}
}


type watch struct {
type watch struct {
	ov     syscall.Overlapped
	ov     syscall.Overlapped
	ino    *inode            // i-number
	ino    *inode            // i-number
	path   string            // Directory path
	path   string            // Directory path
	mask   uint64            // Directory itself is being watched with these notify flags
	mask   uint64            // Directory itself is being watched with these notify flags
	names  map[string]uint64 // Map of names being watched and their notify flags
	names  map[string]uint64 // Map of names being watched and their notify flags
	rename string            // Remembers the old name while renaming a file
	rename string            // Remembers the old name while renaming a file
	buf    [4096]byte
	buf    [4096]byte
}
}


type indexMap map[uint64]*watch
type indexMap map[uint64]*watch
type watchMap map[uint32]indexMap
type watchMap map[uint32]indexMap


// A Watcher waits for and receives event notifications
// A Watcher waits for and receives event notifications
// for a specific set of files and directories.
// for a specific set of files and directories.
type Watcher struct {
type Watcher struct {
	mu            sync.Mutex        // Map access
	mu            sync.Mutex        // Map access
	port          syscall.Handle    // Handle to completion port
	port          syscall.Handle    // Handle to completion port
	watches       watchMap          // Map of watches (key: i-number)
	watches       watchMap          // Map of watches (key: i-number)
	fsnFlags      map[string]uint32 // Map of watched files to flags used for filter
	fsnFlags      map[string]uint32 // Map of watched files to flags used for filter
	fsnmut        sync.Mutex        // Protects access to fsnFlags.
	fsnmut        sync.Mutex        // Protects access to fsnFlags.
	input         chan *input       // Inputs to the reader are sent on this channel
	input         chan *input       // Inputs to the reader are sent on this channel
	internalEvent chan *FileEvent   // Events are queued on this channel
	internalEvent chan *FileEvent   // Events are queued on this channel
	Event         chan *FileEvent   // Events are returned on this channel
	Event         chan *FileEvent   // Events are returned on this channel
	Error         chan error        // Errors are sent on this channel
	Error         chan error        // Errors are sent on this channel
	isClosed      bool              // Set to true when Close() is first called
	isClosed      bool              // Set to true when Close() is first called
	quit          chan chan<- error
	quit          chan chan<- error
	cookie        uint32
	cookie        uint32
}
}


// NewWatcher creates and returns a Watcher.
// NewWatcher creates and returns a Watcher.
func NewWatcher() (*Watcher, error) {
func NewWatcher() (*Watcher, error) {
	port, e := syscall.CreateIoCompletionPort(syscall.InvalidHandle, 0, 0, 0)
	port, e := syscall.CreateIoCompletionPort(syscall.InvalidHandle, 0, 0, 0)
	if e != nil {
	if e != nil {
		return nil, os.NewSyscallError("CreateIoCompletionPort", e)
		return nil, os.NewSyscallError("CreateIoCompletionPort", e)
	}
	}
	w := &Watcher{
	w := &Watcher{
		port:          port,
		port:          port,
		watches:       make(watchMap),
		watches:       make(watchMap),
		fsnFlags:      make(map[string]uint32),
		fsnFlags:      make(map[string]uint32),
		input:         make(chan *input, 1),
		input:         make(chan *input, 1),
		Event:         make(chan *FileEvent, 50),
		Event:         make(chan *FileEvent, 50),
		internalEvent: make(chan *FileEvent),
		internalEvent: make(chan *FileEvent),
		Error:         make(chan error),
		Error:         make(chan error),
		quit:          make(chan chan<- error, 1),
		quit:          make(chan chan<- error, 1),
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


// Close closes a Watcher.
// Close closes a Watcher.
// It sends a message to the reader goroutine to quit and removes all watches
// It sends a message to the reader goroutine to quit and removes all watches
// associated with the watcher.
// associated with the watcher.
func (w *Watcher) Close() error {
func (w *Watcher) Close() error {
	if w.isClosed {
	if w.isClosed {
		return nil
		return nil
	}
	}
	w.isClosed = true
	w.isClosed = true


	// Send "quit" message to the reader goroutine
	// Send "quit" message to the reader goroutine
	ch := make(chan error)
	ch := make(chan error)
	w.quit <- ch
	w.quit <- ch
	if err := w.wakeupReader(); err != nil {
	if err := w.wakeupReader(); err != nil {
		return err
		return err
	}
	}
	return <-ch
	return <-ch
}
}


// AddWatch adds path to the watched file set.
// AddWatch adds path to the watched file set.
func (w *Watcher) AddWatch(path string, flags uint32) error {
func (w *Watcher) AddWatch(path string, flags uint32) error {
	if w.isClosed {
	if w.isClosed {
		return errors.New("watcher already closed")
		return errors.New("watcher already closed")
	}
	}
	in := &input{
	in := &input{
		op:    opAddWatch,
		op:    opAddWatch,
		path:  filepath.Clean(path),
		path:  filepath.Clean(path),
		flags: flags,
		flags: flags,
		reply: make(chan error),
		reply: make(chan error),
	}
	}
	w.input <- in
	w.input <- in
	if err := w.wakeupReader(); err != nil {
	if err := w.wakeupReader(); err != nil {
		return err
		return err
	}
	}
	return <-in.reply
	return <-in.reply
}
}


// Watch adds path to the watched file set, watching all events.
// Watch adds path to the watched file set, watching all events.
func (w *Watcher) watch(path string) error {
func (w *Watcher) watch(path string) error {
	return w.AddWatch(path, sys_FS_ALL_EVENTS)
	return w.AddWatch(path, sys_FS_ALL_EVENTS)
}
}


// RemoveWatch removes path from the watched file set.
// RemoveWatch removes path from the watched file set.
func (w *Watcher) removeWatch(path string) error {
func (w *Watcher) removeWatch(path string) error {
	in := &input{
	in := &input{
		op:    opRemoveWatch,
		op:    opRemoveWatch,
		path:  filepath.Clean(path),
		path:  filepath.Clean(path),
		reply: make(chan error),
		reply: make(chan error),
	}
	}
	w.input <- in
	w.input <- in
	if err := w.wakeupReader(); err != nil {
	if err := w.wakeupReader(); err != nil {
		return err
		return err
	}
	}
	return <-in.reply
	return <-in.reply
}
}


func (w *Watcher) wakeupReader() error {
func (w *Watcher) wakeupReader() error {
	e := syscall.PostQueuedCompletionStatus(w.port, 0, 0, nil)
	e := syscall.PostQueuedCompletionStatus(w.port, 0, 0, nil)
	if e != nil {
	if e != nil {
		return os.NewSyscallError("PostQueuedCompletionStatus", e)
		return os.NewSyscallError("PostQueuedCompletionStatus", e)
	}
	}
	return nil
	return nil
}
}


func getDir(pathname string) (dir string, err error) {
func getDir(pathname string) (dir string, err error) {
	attr, e := syscall.GetFileAttributes(syscall.StringToUTF16Ptr(pathname))
	attr, e := syscall.GetFileAttributes(syscall.StringToUTF16Ptr(pathname))
	if e != nil {
	if e != nil {
		return "", os.NewSyscallError("GetFileAttributes", e)
		return "", os.NewSyscallError("GetFileAttributes", e)
	}
	}
	if attr&syscall.FILE_ATTRIBUTE_DIRECTORY != 0 {
	if attr&syscall.FILE_ATTRIBUTE_DIRECTORY != 0 {
		dir = pathname
		dir = pathname
	} else {
	} else {
		dir, _ = filepath.Split(pathname)
		dir, _ = filepath.Split(pathname)
		dir = filepath.Clean(dir)
		dir = filepath.Clean(dir)
	}
	}
	return
	return
}
}


func getIno(path string) (ino *inode, err error) {
func getIno(path string) (ino *inode, err error) {
	h, e := syscall.CreateFile(syscall.StringToUTF16Ptr(path),
	h, e := syscall.CreateFile(syscall.StringToUTF16Ptr(path),
		syscall.FILE_LIST_DIRECTORY,
		syscall.FILE_LIST_DIRECTORY,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE|syscall.FILE_SHARE_DELETE,
		nil, syscall.OPEN_EXISTING,
		nil, syscall.OPEN_EXISTING,
		syscall.FILE_FLAG_BACKUP_SEMANTICS|syscall.FILE_FLAG_OVERLAPPED, 0)
		syscall.FILE_FLAG_BACKUP_SEMANTICS|syscall.FILE_FLAG_OVERLAPPED, 0)
	if e != nil {
	if e != nil {
		return nil, os.NewSyscallError("CreateFile", e)
		return nil, os.NewSyscallError("CreateFile", e)
	}
	}
	var fi syscall.ByHandleFileInformation
	var fi syscall.ByHandleFileInformation
	if e = syscall.GetFileInformationByHandle(h, &fi); e != nil {
	if e = syscall.GetFileInformationByHandle(h, &fi); e != nil {
		syscall.CloseHandle(h)
		syscall.CloseHandle(h)
		return nil, os.NewSyscallError("GetFileInformationByHandle", e)
		return nil, os.NewSyscallError("GetFileInformationByHandle", e)
	}
	}
	ino = &inode{
	ino = &inode{
		handle: h,
		handle: h,
		volume: fi.VolumeSerialNumber,
		volume: fi.VolumeSerialNumber,
		index:  uint64(fi.FileIndexHigh)<<32 | uint64(fi.FileIndexLow),
		index:  uint64(fi.FileIndexHigh)<<32 | uint64(fi.FileIndexLow),
	}
	}
	return ino, nil
	return ino, nil
}
}


// Must run within the I/O thread.
// Must run within the I/O thread.
func (m watchMap) get(ino *inode) *watch {
func (m watchMap) get(ino *inode) *watch {
	if i := m[ino.volume]; i != nil {
	if i := m[ino.volume]; i != nil {
		return i[ino.index]
		return i[ino.index]
	}
	}
	return nil
	return nil
}
}


// Must run within the I/O thread.
// Must run within the I/O thread.
func (m watchMap) set(ino *inode, watch *watch) {
func (m watchMap) set(ino *inode, watch *watch) {
	i := m[ino.volume]
	i := m[ino.volume]
	if i == nil {
	if i == nil {
		i = make(indexMap)
		i = make(indexMap)
		m[ino.volume] = i
		m[ino.volume] = i
	}
	}
	i[ino.index] = watch
	i[ino.index] = watch
}
}


// Must run within the I/O thread.
// Must run within the I/O thread.
func (w *Watcher) addWatch(pathname string, flags uint64) error {
func (w *Watcher) addWatch(pathname string, flags uint64) error {
	dir, err := getDir(pathname)
	dir, err := getDir(pathname)
	if err != nil {
	if err != nil {
		return err
		return err
	}
	}
	if flags&sys_FS_ONLYDIR != 0 && pathname != dir {
	if flags&sys_FS_ONLYDIR != 0 && pathname != dir {
		return nil
		return nil
	}
	}
	ino, err := getIno(dir)
	ino, err := getIno(dir)
	if err != nil {
	if err != nil {
		return err
		return err
	}
	}
	w.mu.Lock()
	w.mu.Lock()
	watchEntry := w.watches.get(ino)
	watchEntry := w.watches.get(ino)
	w.mu.Unlock()
	w.mu.Unlock()
	if watchEntry == nil {
	if watchEntry == nil {
		if _, e := syscall.CreateIoCompletionPort(ino.handle, w.port, 0, 0); e != nil {
		if _, e := syscall.CreateIoCompletionPort(ino.handle, w.port, 0, 0); e != nil {
			syscall.CloseHandle(ino.handle)
			syscall.CloseHandle(ino.handle)
			return os.NewSyscallError("CreateIoCompletionPort", e)
			return os.NewSyscallError("CreateIoCompletionPort", e)
		}
		}
		watchEntry = &watch{
		watchEntry = &watch{
			ino:   ino,
			ino:   ino,
			path:  dir,
			path:  dir,
			names: make(map[string]uint64),
			names: make(map[string]uint64),
		}
		}
		w.mu.Lock()
		w.mu.Lock()
		w.watches.set(ino, watchEntry)
		w.watches.set(ino, watchEntry)
		w.mu.Unlock()
		w.mu.Unlock()
		flags |= provisional
		flags |= provisional
	} else {
	} else {
		syscall.CloseHandle(ino.handle)
		syscall.CloseHandle(ino.handle)
	}
	}
	if pathname == dir {
	if pathname == dir {
		watchEntry.mask |= flags
		watchEntry.mask |= flags
	} else {
	} else {
		watchEntry.names[filepath.Base(pathname)] |= flags
		watchEntry.names[filepath.Base(pathname)] |= flags
	}
	}
	if err = w.startRead(watchEntry); err != nil {
	if err = w.startRead(watchEntry); err != nil {
		return err
		return err
	}
	}
	if pathname == dir {
	if pathname == dir {
		watchEntry.mask &= ^provisional
		watchEntry.mask &= ^provisional
	} else {
	} else {
		watchEntry.names[filepath.Base(pathname)] &= ^provisional
		watchEntry.names[filepath.Base(pathname)] &= ^provisional
	}
	}
	return nil
	return nil
}
}


// Must run within the I/O thread.
// Must run within the I/O thread.
func (w *Watcher) remWatch(pathname string) error {
func (w *Watcher) remWatch(pathname string) error {
	dir, err := getDir(pathname)
	dir, err := getDir(pathname)
	if err != nil {
	if err != nil {
		return err
		return err
	}
	}
	ino, err := getIno(dir)
	ino, err := getIno(dir)
	if err != nil {
	if err != nil {
		return err
		return err
	}
	}
	w.mu.Lock()
	w.mu.Lock()
	watch := w.watches.get(ino)
	watch := w.watches.get(ino)
	w.mu.Unlock()
	w.mu.Unlock()
	if watch == nil {
	if watch == nil {
		return fmt.Errorf("can't remove non-existent watch for: %s", pathname)
		return fmt.Errorf("can't remove non-existent watch for: %s", pathname)
	}
	}
	if pathname == dir {
	if pathname == dir {
		w.sendEvent(watch.path, watch.mask&sys_FS_IGNORED)
		w.sendEvent(watch.path, watch.mask&sys_FS_IGNORED)
		watch.mask = 0
		watch.mask = 0
	} else {
	} else {
		name := filepath.Base(pathname)
		name := filepath.Base(pathname)
		w.sendEvent(watch.path+"\\"+name, watch.names[name]&sys_FS_IGNORED)
		w.sendEvent(watch.path+"\\"+name, watch.names[name]&sys_FS_IGNORED)
		delete(watch.names, name)
		delete(watch.names, name)
	}
	}
	return w.startRead(watch)
	return w.startRead(watch)
}
}


// Must run within the I/O thread.
// Must run within the I/O thread.
func (w *Watcher) deleteWatch(watch *watch) {
func (w *Watcher) deleteWatch(watch *watch) {
	for name, mask := range watch.names {
	for name, mask := range watch.names {
		if mask&provisional == 0 {
		if mask&provisional == 0 {
			w.sendEvent(watch.path+"\\"+name, mask&sys_FS_IGNORED)
			w.sendEvent(watch.path+"\\"+name, mask&sys_FS_IGNORED)
		}
		}
		delete(watch.names, name)
		delete(watch.names, name)
	}
	}
	if watch.mask != 0 {
	if watch.mask != 0 {
		if watch.mask&provisional == 0 {
		if watch.mask&provisional == 0 {
			w.sendEvent(watch.path, watch.mask&sys_FS_IGNORED)
			w.sendEvent(watch.path, watch.mask&sys_FS_IGNORED)
		}
		}
		watch.mask = 0
		watch.mask = 0
	}
	}
}
}


// Must run within the I/O thread.
// Must run within the I/O thread.
func (w *Watcher) startRead(watch *watch) error {
func (w *Watcher) startRead(watch *watch) error {
	if e := syscall.CancelIo(watch.ino.handle); e != nil {
	if e := syscall.CancelIo(watch.ino.handle); e != nil {
		w.Error <- os.NewSyscallError("CancelIo", e)
		w.Error <- os.NewSyscallError("CancelIo", e)
		w.deleteWatch(watch)
		w.deleteWatch(watch)
	}
	}
	mask := toWindowsFlags(watch.mask)
	mask := toWindowsFlags(watch.mask)
	for _, m := range watch.names {
	for _, m := range watch.names {
		mask |= toWindowsFlags(m)
		mask |= toWindowsFlags(m)
	}
	}
	if mask == 0 {
	if mask == 0 {
		if e := syscall.CloseHandle(watch.ino.handle); e != nil {
		if e := syscall.CloseHandle(watch.ino.handle); e != nil {
			w.Error <- os.NewSyscallError("CloseHandle", e)
			w.Error <- os.NewSyscallError("CloseHandle", e)
		}
		}
		w.mu.Lock()
		w.mu.Lock()
		delete(w.watches[watch.ino.volume], watch.ino.index)
		delete(w.watches[watch.ino.volume], watch.ino.index)
		w.mu.Unlock()
		w.mu.Unlock()
		return nil
		return nil
	}
	}
	e := syscall.ReadDirectoryChanges(watch.ino.handle, &watch.buf[0],
	e := syscall.ReadDirectoryChanges(watch.ino.handle, &watch.buf[0],
		uint32(unsafe.Sizeof(watch.buf)), false, mask, nil, &watch.ov, 0)
		uint32(unsafe.Sizeof(watch.buf)), false, mask, nil, &watch.ov, 0)
	if e != nil {
	if e != nil {
		err := os.NewSyscallError("ReadDirectoryChanges", e)
		err := os.NewSyscallError("ReadDirectoryChanges", e)
		if e == syscall.ERROR_ACCESS_DENIED && watch.mask&provisional == 0 {
		if e == syscall.ERROR_ACCESS_DENIED && watch.mask&provisional == 0 {
			// Watched directory was probably removed
			// Watched directory was probably removed
			if w.sendEvent(watch.path, watch.mask&sys_FS_DELETE_SELF) {
			if w.sendEvent(watch.path, watch.mask&sys_FS_DELETE_SELF) {
				if watch.mask&sys_FS_ONESHOT != 0 {
				if watch.mask&sys_FS_ONESHOT != 0 {
					watch.mask = 0
					watch.mask = 0
				}
				}
			}
			}
			err = nil
			err = nil
		}
		}
		w.deleteWatch(watch)
		w.deleteWatch(watch)
		w.startRead(watch)
		w.startRead(watch)
		return err
		return err
	}
	}
	return nil
	return nil
}
}


// readEvents reads from the I/O completion port, converts the
// readEvents reads from the I/O completion port, converts the
// received events into Event objects and sends them via the Event channel.
// received events into Event objects and sends them via the Event channel.
// Entry point to the I/O thread.
// Entry point to the I/O thread.
func (w *Watcher) readEvents() {
func (w *Watcher) readEvents() {
	var (
	var (
		n, key uint32
		n, key uint32
		ov     *syscall.Overlapped
		ov     *syscall.Overlapped
	)
	)
	runtime.LockOSThread()
	runtime.LockOSThread()


	for {
	for {
		e := syscall.GetQueuedCompletionStatus(w.port, &n, &key, &ov, syscall.INFINITE)
		e := syscall.GetQueuedCompletionStatus(w.port, &n, &key, &ov, syscall.INFINITE)
		watch := (*watch)(unsafe.Pointer(ov))
		watch := (*watch)(unsafe.Pointer(ov))


		if watch == nil {
		if watch == nil {
			select {
			select {
			case ch := <-w.quit:
			case ch := <-w.quit:
				w.mu.Lock()
				w.mu.Lock()
				var indexes []indexMap
				var indexes []indexMap
				for _, index := range w.watches {
				for _, index := range w.watches {
					indexes = append(indexes, index)
					indexes = append(indexes, index)
				}
				}
				w.mu.Unlock()
				w.mu.Unlock()
				for _, index := range indexes {
				for _, index := range indexes {
					for _, watch := range index {
					for _, watch := range index {
						w.deleteWatch(watch)
						w.deleteWatch(watch)
						w.startRead(watch)
						w.startRead(watch)
					}
					}
				}
				}
				var err error
				var err error
				if e := syscall.CloseHandle(w.port); e != nil {
				if e := syscall.CloseHandle(w.port); e != nil {
					err = os.NewSyscallError("CloseHandle", e)
					err = os.NewSyscallError("CloseHandle", e)
				}
				}
				close(w.internalEvent)
				close(w.internalEvent)
				close(w.Error)
				close(w.Error)
				ch <- err
				ch <- err
				return
				return
			case in := <-w.input:
			case in := <-w.input:
				switch in.op {
				switch in.op {
				case opAddWatch:
				case opAddWatch:
					in.reply <- w.addWatch(in.path, uint64(in.flags))
					in.reply <- w.addWatch(in.path, uint64(in.flags))
				case opRemoveWatch:
				case opRemoveWatch:
					in.reply <- w.remWatch(in.path)
					in.reply <- w.remWatch(in.path)
				}
				}
			default:
			default:
			}
			}
			continue
			continue
		}
		}


		switch e {
		switch e {
		case sys_ERROR_MORE_DATA:
		case sys_ERROR_MORE_DATA:
			if watch == nil {
			if watch == nil {
				w.Error <- errors.New("ERROR_MORE_DATA has unexpectedly null lpOverlapped buffer")
				w.Error <- errors.New("ERROR_MORE_DATA has unexpectedly null lpOverlapped buffer")
			} else {
			} else {
				// The i/o succeeded but the buffer is full.
				// The i/o succeeded but the buffer is full.
				// In theory we should be building up a full packet.
				// In theory we should be building up a full packet.
				// In practice we can get away with just carrying on.
				// In practice we can get away with just carrying on.
				n = uint32(unsafe.Sizeof(watch.buf))
				n = uint32(unsafe.Sizeof(watch.buf))
			}
			}
		case syscall.ERROR_ACCESS_DENIED:
		case syscall.ERROR_ACCESS_DENIED:
			// Watched directory was probably removed
			// Watched directory was probably removed
			w.sendEvent(watch.path, watch.mask&sys_FS_DELETE_SELF)
			w.sendEvent(watch.path, watch.mask&sys_FS_DELETE_SELF)
			w.deleteWatch(watch)
			w.deleteWatch(watch)
			w.startRead(watch)
			w.startRead(watch)
			continue
			continue
		case syscall.ERROR_OPERATION_ABORTED:
		case syscall.ERROR_OPERATION_ABORTED:
			// CancelIo was called on this handle
			// CancelIo was called on this handle
			continue
			continue
		default:
		default:
			w.Error <- os.NewSyscallError("GetQueuedCompletionPort", e)
			w.Error <- os.NewSyscallError("GetQueuedCompletionPort", e)
			continue
			continue
		case nil:
		case nil:
		}
		}


		var offset uint32
		var offset uint32
		for {
		for {
			if n == 0 {
			if n == 0 {
				w.internalEvent <- &FileEvent{mask: sys_FS_Q_OVERFLOW}
				w.internalEvent <- &FileEvent{mask: sys_FS_Q_OVERFLOW}
				w.Error <- errors.New("short read in readEvents()")
				w.Error <- errors.New("short read in readEvents()")
				break
				break
			}
			}


			// Point "raw" to the event in the buffer
			// Point "raw" to the event in the buffer
			raw := (*syscall.FileNotifyInformation)(unsafe.Pointer(&watch.buf[offset]))
			raw := (*syscall.FileNotifyInformation)(unsafe.Pointer(&watch.buf[offset]))
			buf := (*[syscall.MAX_PATH]uint16)(unsafe.Pointer(&raw.FileName))
			buf := (*[syscall.MAX_PATH]uint16)(unsafe.Pointer(&raw.FileName))
			name := syscall.UTF16ToString(buf[:raw.FileNameLength/2])
			name := syscall.UTF16ToString(buf[:raw.FileNameLength/2])
			fullname := watch.path + "\\" + name
			fullname := watch.path + "\\" + name


			var mask uint64
			var mask uint64
			switch raw.Action {
			switch raw.Action {
			case syscall.FILE_ACTION_REMOVED:
			case syscall.FILE_ACTION_REMOVED:
				mask = sys_FS_DELETE_SELF
				mask = sys_FS_DELETE_SELF
			case syscall.FILE_ACTION_MODIFIED:
			case syscall.FILE_ACTION_MODIFIED:
				mask = sys_FS_MODIFY
				mask = sys_FS_MODIFY
			case syscall.FILE_ACTION_RENAMED_OLD_NAME:
			case syscall.FILE_ACTION_RENAMED_OLD_NAME:
				watch.rename = name
				watch.rename = name
			case syscall.FILE_ACTION_RENAMED_NEW_NAME:
			case syscall.FILE_ACTION_RENAMED_NEW_NAME:
				if watch.names[watch.rename] != 0 {
				if watch.names[watch.rename] != 0 {
					watch.names[name] |= watch.names[watch.rename]
					watch.names[name] |= watch.names[watch.rename]
					delete(watch.names, watch.rename)
					delete(watch.names, watch.rename)
					mask = sys_FS_MOVE_SELF
					mask = sys_FS_MOVE_SELF
				}
				}
			}
			}


			sendNameEvent := func() {
			sendNameEvent := func() {
				if w.sendEvent(fullname, watch.names[name]&mask) {
				if w.sendEvent(fullname, watch.names[name]&mask) {
					if watch.names[name]&sys_FS_ONESHOT != 0 {
					if watch.names[name]&sys_FS_ONESHOT != 0 {
						delete(watch.names, name)
						delete(watch.names, name)
					}
					}
				}
				}
			}
			}
			if raw.Action != syscall.FILE_ACTION_RENAMED_NEW_NAME {
			if raw.Action != syscall.FILE_ACTION_RENAMED_NEW_NAME {
				sendNameEvent()
				sendNameEvent()
			}
			}
			if raw.Action == syscall.FILE_ACTION_REMOVED {
			if raw.Action == syscall.FILE_ACTION_REMOVED {
				w.sendEvent(fullname, watch.names[name]&sys_FS_IGNORED)
				w.sendEvent(fullname, watch.names[name]&sys_FS_IGNORED)
				delete(watch.names, name)
				delete(watch.names, name)
			}
			}
			if w.sendEvent(fullname, watch.mask&toFSnotifyFlags(raw.Action)) {
			if w.sendEvent(fullname, watch.mask&toFSnotifyFlags(raw.Action)) {
				if watch.mask&sys_FS_ONESHOT != 0 {
				if watch.mask&sys_FS_ONESHOT != 0 {
					watch.mask = 0
					watch.mask = 0
				}
				}
			}
			}
			if raw.Action == syscall.FILE_ACTION_RENAMED_NEW_NAME {
			if raw.Action == syscall.FILE_ACTION_RENAMED_NEW_NAME {
				fullname = watch.path + "\\" + watch.rename
				fullname = watch.path + "\\" + watch.rename
				sendNameEvent()
				sendNameEvent()
			}
			}


			// Move to the next event in the buffer
			// Move to the next event in the buffer
			if raw.NextEntryOffset == 0 {
			if raw.NextEntryOffset == 0 {
				break
				break
			}
			}
			offset += raw.NextEntryOffset
			offset += raw.NextEntryOffset


			// Error!
			// Error!
			if offset >= n {
			if offset >= n {
				w.Error <- errors.New("Windows system assumed buffer larger than it is, events have likely been missed.")
				w.Error <- errors.New("Windows system assumed buffer larger than it is, events have likely been missed.")
				break
				break
			}
			}
		}
		}


		if err := w.startRead(watch); err != nil {
		if err := w.startRead(watch); err != nil {
			w.Error <- err
			w.Error <- err
		}
		}
	}
	}
}
}


func (w *Watcher) sendEvent(name string, mask uint64) bool {
func (w *Watcher) sendEvent(name string, mask uint64) bool {
	if mask == 0 {
	if mask == 0 {
		return false
		return false
	}
	}
	event := &FileEvent{mask: uint32(mask), Name: name}
	event := &FileEvent{mask: uint32(mask), Name: name}
	if mask&sys_FS_MOVE != 0 {
	if mask&sys_FS_MOVE != 0 {
		if mask&sys_FS_MOVED_FROM != 0 {
		if mask&sys_FS_MOVED_FROM != 0 {
			w.cookie++
			w.cookie++
		}
		}
		event.cookie = w.cookie
		event.cookie = w.cookie
	}
	}
	select {
	select {
	case ch := <-w.quit:
	case ch := <-w.quit:
		w.quit <- ch
		w.quit <- ch
	case w.Event <- event:
	case w.Event <- event:
	}
	}
	return true
	return true
}
}


func toWindowsFlags(mask uint64) uint32 {
func toWindowsFlags(mask uint64) uint32 {
	var m uint32
	var m uint32
	if mask&sys_FS_ACCESS != 0 {
	if mask&sys_FS_ACCESS != 0 {
		m |= syscall.FILE_NOTIFY_CHANGE_LAST_ACCESS
		m |= syscall.FILE_NOTIFY_CHANGE_LAST_ACCESS
	}
	}
	if mask&sys_FS_MODIFY != 0 {
	if mask&sys_FS_MODIFY != 0 {
		m |= syscall.FILE_NOTIFY_CHANGE_LAST_WRITE
		m |= syscall.FILE_NOTIFY_CHANGE_LAST_WRITE
	}
	}
	if mask&sys_FS_ATTRIB != 0 {
	if mask&sys_FS_ATTRIB != 0 {
		m |= syscall.FILE_NOTIFY_CHANGE_ATTRIBUTES
		m |= syscall.FILE_NOTIFY_CHANGE_ATTRIBUTES
	}
	}
	if mask&(sys_FS_MOVE|sys_FS_CREATE|sys_FS_DELETE) != 0 {
	if mask&(sys_FS_MOVE|sys_FS_CREATE|sys_FS_DELETE) != 0 {
		m |= syscall.FILE_NOTIFY_CHANGE_FILE_NAME | syscall.FILE_NOTIFY_CHANGE_DIR_NAME
		m |= syscall.FILE_NOTIFY_CHANGE_FILE_NAME | syscall.FILE_NOTIFY_CHANGE_DIR_NAME
	}
	}
	return m
	return m
}
}


func toFSnotifyFlags(action uint32) uint64 {
func toFSnotifyFlags(action uint32) uint64 {
	switch action {
	switch action {
	case syscall.FILE_ACTION_ADDED:
	case syscall.FILE_ACTION_ADDED:
		return sys_FS_CREATE
		return sys_FS_CREATE
	case syscall.FILE_ACTION_REMOVED:
	case syscall.FILE_ACTION_REMOVED:
		return sys_FS_DELETE
		return sys_FS_DELETE
	case syscall.FILE_ACTION_MODIFIED:
	case syscall.FILE_ACTION_MODIFIED:
		return sys_FS_MODIFY
		return sys_FS_MODIFY
	case syscall.FILE_ACTION_RENAMED_OLD_NAME:
	case syscall.FILE_ACTION_RENAMED_OLD_NAME:
		return sys_FS_MOVED_FROM
		return sys_FS_MOVED_FROM
	case syscall.FILE_ACTION_RENAMED_NEW_NAME:
	case syscall.FILE_ACTION_RENAMED_NEW_NAME:
		return sys_FS_MOVED_TO
		return sys_FS_MOVED_TO
	}
	}
	return 0
	return 0
}
}
