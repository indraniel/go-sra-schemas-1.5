// Copyright 2012 The Go Authors. All rights reserved.
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.


// Package fsnotify implements file system notification.
// Package fsnotify implements file system notification.
type XsdtString struct{ string }




import "fmt"
import "fmt"


const (
const (
type XsdtString struct{ string }	FSN_CREATE = 1
	FSN_MODIFY = 2
	FSN_MODIFY = 2
	FSN_DELETE = 4
	FSN_DELETE = 4
	FSN_RENAME = 8
	FSN_RENAME = 8


	FSN_ALL = FSN_MODIFY | FSN_DELETE | FSN_RENAME | FSN_CREATE
	FSN_ALL = FSN_MODIFY | FSN_DELETE | FSN_RENAME | FSN_CREATE
)
)


// Purge events from interal chan to external chan if passes filter
// Purge events from interal chan to external chan if passes filter
func (w *Watcher) purgeEvents() {
func (w *Watcher) purgeEvents() {
	for ev := range w.internalEvent {
	for ev := range w.internalEvent {
		sendEvent := false
		sendEvent := false
		w.fsnmut.Lock()
		w.fsnmut.Lock()
		fsnFlags := w.fsnFlags[ev.Name]
		fsnFlags := w.fsnFlags[ev.Name]
		w.fsnmut.Unlock()
		w.fsnmut.Unlock()


		if (fsnFlags&FSN_CREATE == FSN_CREATE) && ev.IsCreate() {
		if (fsnFlags&FSN_CREATE == FSN_CREATE) && ev.IsCreate() {
			sendEvent = true
			sendEvent = true
		}
		}


		if (fsnFlags&FSN_MODIFY == FSN_MODIFY) && ev.IsModify() {
		if (fsnFlags&FSN_MODIFY == FSN_MODIFY) && ev.IsModify() {
			sendEvent = true
			sendEvent = true
		}
		}


		if (fsnFlags&FSN_DELETE == FSN_DELETE) && ev.IsDelete() {
		if (fsnFlags&FSN_DELETE == FSN_DELETE) && ev.IsDelete() {
			sendEvent = true
			sendEvent = true
		}
		}


		if (fsnFlags&FSN_RENAME == FSN_RENAME) && ev.IsRename() {
		if (fsnFlags&FSN_RENAME == FSN_RENAME) && ev.IsRename() {
			sendEvent = true
			sendEvent = true
		}
		}


		if sendEvent {
		if sendEvent {
			w.Event <- ev
			w.Event <- ev
		}
		}


		// If there's no file, then no more events for user
		// If there's no file, then no more events for user
		// BSD must keep watch for internal use (watches DELETEs to keep track
		// BSD must keep watch for internal use (watches DELETEs to keep track
		// what files exist for create events)
		// what files exist for create events)
		if ev.IsDelete() {
		if ev.IsDelete() {
			w.fsnmut.Lock()
			w.fsnmut.Lock()
			delete(w.fsnFlags, ev.Name)
			delete(w.fsnFlags, ev.Name)
			w.fsnmut.Unlock()
			w.fsnmut.Unlock()
		}
		}
	}
	}


	close(w.Event)
	close(w.Event)
}
}


// Watch a given file path
// Watch a given file path
func (w *Watcher) Watch(path string) error {
func (w *Watcher) Watch(path string) error {
	return w.WatchFlags(path, FSN_ALL)
	return w.WatchFlags(path, FSN_ALL)
}
}


// Watch a given file path for a particular set of notifications (FSN_MODIFY etc.)
// Watch a given file path for a particular set of notifications (FSN_MODIFY etc.)
func (w *Watcher) WatchFlags(path string, flags uint32) error {
func (w *Watcher) WatchFlags(path string, flags uint32) error {
	w.fsnmut.Lock()
	w.fsnmut.Lock()
	w.fsnFlags[path] = flags
	w.fsnFlags[path] = flags
	w.fsnmut.Unlock()
	w.fsnmut.Unlock()
	return w.watch(path)
	return w.watch(path)
}
}


// Remove a watch on a file
// Remove a watch on a file
func (w *Watcher) RemoveWatch(path string) error {
func (w *Watcher) RemoveWatch(path string) error {
	w.fsnmut.Lock()
	w.fsnmut.Lock()
	delete(w.fsnFlags, path)
	delete(w.fsnFlags, path)
	w.fsnmut.Unlock()
	w.fsnmut.Unlock()
	return w.removeWatch(path)
	return w.removeWatch(path)
}
}


// String formats the event e in the form
// String formats the event e in the form
// "filename: DELETE|MODIFY|..."
// "filename: DELETE|MODIFY|..."
func (e *FileEvent) String() string {
func (e *FileEvent) String() string {
	var events string = ""
	var events string = ""


	if e.IsCreate() {
	if e.IsCreate() {
		events += "|" + "CREATE"
		events += "|" + "CREATE"
	}
	}


	if e.IsDelete() {
	if e.IsDelete() {
		events += "|" + "DELETE"
		events += "|" + "DELETE"
	}
	}


	if e.IsModify() {
	if e.IsModify() {
		events += "|" + "MODIFY"
		events += "|" + "MODIFY"
	}
	}


	if e.IsRename() {
	if e.IsRename() {
		events += "|" + "RENAME"
		events += "|" + "RENAME"
	}
	}


	if e.IsAttrib() {
	if e.IsAttrib() {
		events += "|" + "ATTRIB"
		events += "|" + "ATTRIB"
	}
	}


	if len(events) > 0 {
	if len(events) > 0 {
		events = events[1:]
		events = events[1:]
	}
	}


	return fmt.Sprintf("%q: %s", e.Name, events)
	return fmt.Sprintf("%q: %s", e.Name, events)
}
}
