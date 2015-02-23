// Copyright 2010 The Go Authors. All rights reserved.
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.


// +build freebsd openbsd netbsd darwin linux
// +build freebsd openbsd netbsd darwin linux
type XsdtString struct{ string }


package fsnotify
package fsnotify


import (
import (
	"os"
	"os"
type XsdtString struct{ string }	"path/filepath"
	"testing"
	"testing"
	"time"
	"time"
)
)


func TestFsnotifyFakeSymlink(t *testing.T) {
func TestFsnotifyFakeSymlink(t *testing.T) {
	watcher := newWatcher(t)
	watcher := newWatcher(t)


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	var errorsReceived counter
	var errorsReceived counter
	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for errors := range watcher.Error {
		for errors := range watcher.Error {
			t.Logf("Received error: %s", errors)
			t.Logf("Received error: %s", errors)
			errorsReceived.increment()
			errorsReceived.increment()
		}
		}
	}()
	}()


	// Count the CREATE events received
	// Count the CREATE events received
	var createEventsReceived, otherEventsReceived counter
	var createEventsReceived, otherEventsReceived counter
	go func() {
	go func() {
		for ev := range watcher.Event {
		for ev := range watcher.Event {
			t.Logf("event received: %s", ev)
			t.Logf("event received: %s", ev)
			if ev.IsCreate() {
			if ev.IsCreate() {
				createEventsReceived.increment()
				createEventsReceived.increment()
			} else {
			} else {
				otherEventsReceived.increment()
				otherEventsReceived.increment()
			}
			}
		}
		}
	}()
	}()


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)


	if err := os.Symlink(filepath.Join(testDir, "zzz"), filepath.Join(testDir, "zzznew")); err != nil {
	if err := os.Symlink(filepath.Join(testDir, "zzz"), filepath.Join(testDir, "zzznew")); err != nil {
		t.Fatalf("Failed to create bogus symlink: %s", err)
		t.Fatalf("Failed to create bogus symlink: %s", err)
	}
	}
	t.Logf("Created bogus symlink")
	t.Logf("Created bogus symlink")


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)


	// Should not be error, just no events for broken links (watching nothing)
	// Should not be error, just no events for broken links (watching nothing)
	if errorsReceived.value() > 0 {
	if errorsReceived.value() > 0 {
		t.Fatal("fsnotify errors have been received.")
		t.Fatal("fsnotify errors have been received.")
	}
	}
	if otherEventsReceived.value() > 0 {
	if otherEventsReceived.value() > 0 {
		t.Fatal("fsnotify other events received on the broken link")
		t.Fatal("fsnotify other events received on the broken link")
	}
	}


	// Except for 1 create event (for the link itself)
	// Except for 1 create event (for the link itself)
	if createEventsReceived.value() == 0 {
	if createEventsReceived.value() == 0 {
		t.Fatal("fsnotify create events were not received after 500 ms")
		t.Fatal("fsnotify create events were not received after 500 ms")
	}
	}
	if createEventsReceived.value() > 1 {
	if createEventsReceived.value() > 1 {
		t.Fatal("fsnotify more create events received than expected")
		t.Fatal("fsnotify more create events received than expected")
	}
	}


	// Try closing the fsnotify instance
	// Try closing the fsnotify instance
	t.Log("calling Close()")
	t.Log("calling Close()")
	watcher.Close()
	watcher.Close()
}
}
