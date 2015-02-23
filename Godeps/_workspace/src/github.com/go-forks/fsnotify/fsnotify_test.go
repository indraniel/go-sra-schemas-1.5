// Copyright 2010 The Go Authors. All rights reserved.
// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.


package fsnotify
package fsnotify
type XsdtString struct{ string }


import (
import (
	"io/ioutil"
	"io/ioutil"
	"os"
	"os"
	"os/exec"
	"os/exec"
type XsdtString struct{ string }	"path/filepath"
	"runtime"
	"runtime"
	"sync/atomic"
	"sync/atomic"
	"testing"
	"testing"
	"time"
	"time"
)
)


// An atomic counter
// An atomic counter
type counter struct {
type counter struct {
	val int32
	val int32
}
}


func (c *counter) increment() {
func (c *counter) increment() {
	atomic.AddInt32(&c.val, 1)
	atomic.AddInt32(&c.val, 1)
}
}


func (c *counter) value() int32 {
func (c *counter) value() int32 {
	return atomic.LoadInt32(&c.val)
	return atomic.LoadInt32(&c.val)
}
}


func (c *counter) reset() {
func (c *counter) reset() {
	atomic.StoreInt32(&c.val, 0)
	atomic.StoreInt32(&c.val, 0)
}
}


// tempMkdir makes a temporary directory
// tempMkdir makes a temporary directory
func tempMkdir(t *testing.T) string {
func tempMkdir(t *testing.T) string {
	dir, err := ioutil.TempDir("", "fsnotify")
	dir, err := ioutil.TempDir("", "fsnotify")
	if err != nil {
	if err != nil {
		t.Fatalf("failed to create test directory: %s", err)
		t.Fatalf("failed to create test directory: %s", err)
	}
	}
	return dir
	return dir
}
}


// newWatcher initializes an fsnotify Watcher instance.
// newWatcher initializes an fsnotify Watcher instance.
func newWatcher(t *testing.T) *Watcher {
func newWatcher(t *testing.T) *Watcher {
	watcher, err := NewWatcher()
	watcher, err := NewWatcher()
	if err != nil {
	if err != nil {
		t.Fatalf("NewWatcher() failed: %s", err)
		t.Fatalf("NewWatcher() failed: %s", err)
	}
	}
	return watcher
	return watcher
}
}


// addWatch adds a watch for a directory
// addWatch adds a watch for a directory
func addWatch(t *testing.T, watcher *Watcher, dir string) {
func addWatch(t *testing.T, watcher *Watcher, dir string) {
	if err := watcher.Watch(dir); err != nil {
	if err := watcher.Watch(dir); err != nil {
		t.Fatalf("watcher.Watch(%q) failed: %s", dir, err)
		t.Fatalf("watcher.Watch(%q) failed: %s", dir, err)
	}
	}
}
}


func TestFsnotifyMultipleOperations(t *testing.T) {
func TestFsnotifyMultipleOperations(t *testing.T) {
	watcher := newWatcher(t)
	watcher := newWatcher(t)


	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for err := range watcher.Error {
		for err := range watcher.Error {
			t.Fatalf("error received: %s", err)
			t.Fatalf("error received: %s", err)
		}
		}
	}()
	}()


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	// Create directory that's not watched
	// Create directory that's not watched
	testDirToMoveFiles := tempMkdir(t)
	testDirToMoveFiles := tempMkdir(t)
	defer os.RemoveAll(testDirToMoveFiles)
	defer os.RemoveAll(testDirToMoveFiles)


	testFile := filepath.Join(testDir, "TestFsnotifySeq.testfile")
	testFile := filepath.Join(testDir, "TestFsnotifySeq.testfile")
	testFileRenamed := filepath.Join(testDirToMoveFiles, "TestFsnotifySeqRename.testfile")
	testFileRenamed := filepath.Join(testDirToMoveFiles, "TestFsnotifySeqRename.testfile")


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)


	// Receive events on the event channel on a separate goroutine
	// Receive events on the event channel on a separate goroutine
	eventstream := watcher.Event
	eventstream := watcher.Event
	var createReceived, modifyReceived, deleteReceived, renameReceived counter
	var createReceived, modifyReceived, deleteReceived, renameReceived counter
	done := make(chan bool)
	done := make(chan bool)
	go func() {
	go func() {
		for event := range eventstream {
		for event := range eventstream {
			// Only count relevant events
			// Only count relevant events
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) {
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) {
				t.Logf("event received: %s", event)
				t.Logf("event received: %s", event)
				if event.IsDelete() {
				if event.IsDelete() {
					deleteReceived.increment()
					deleteReceived.increment()
				}
				}
				if event.IsModify() {
				if event.IsModify() {
					modifyReceived.increment()
					modifyReceived.increment()
				}
				}
				if event.IsCreate() {
				if event.IsCreate() {
					createReceived.increment()
					createReceived.increment()
				}
				}
				if event.IsRename() {
				if event.IsRename() {
					renameReceived.increment()
					renameReceived.increment()
				}
				}
			} else {
			} else {
				t.Logf("unexpected event received: %s", event)
				t.Logf("unexpected event received: %s", event)
			}
			}
		}
		}
		done <- true
		done <- true
	}()
	}()


	// Create a file
	// Create a file
	// This should add at least one event to the fsnotify event queue
	// This should add at least one event to the fsnotify event queue
	var f *os.File
	var f *os.File
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()


	time.Sleep(time.Millisecond)
	time.Sleep(time.Millisecond)
	f.WriteString("data")
	f.WriteString("data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete


	if err := testRename(testFile, testFileRenamed); err != nil {
	if err := testRename(testFile, testFileRenamed); err != nil {
		t.Fatalf("rename failed: %s", err)
		t.Fatalf("rename failed: %s", err)
	}
	}


	// Modify the file outside of the watched dir
	// Modify the file outside of the watched dir
	f, err = os.Open(testFileRenamed)
	f, err = os.Open(testFileRenamed)
	if err != nil {
	if err != nil {
		t.Fatalf("open test renamed file failed: %s", err)
		t.Fatalf("open test renamed file failed: %s", err)
	}
	}
	f.WriteString("data")
	f.WriteString("data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete


	// Recreate the file that was moved
	// Recreate the file that was moved
	f, err = os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	f, err = os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Close()
	f.Close()
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	cReceived := createReceived.value()
	cReceived := createReceived.value()
	if cReceived != 2 {
	if cReceived != 2 {
		t.Fatalf("incorrect number of create events received after 500 ms (%d vs %d)", cReceived, 2)
		t.Fatalf("incorrect number of create events received after 500 ms (%d vs %d)", cReceived, 2)
	}
	}
	mReceived := modifyReceived.value()
	mReceived := modifyReceived.value()
	if mReceived != 1 {
	if mReceived != 1 {
		t.Fatalf("incorrect number of modify events received after 500 ms (%d vs %d)", mReceived, 1)
		t.Fatalf("incorrect number of modify events received after 500 ms (%d vs %d)", mReceived, 1)
	}
	}
	dReceived := deleteReceived.value()
	dReceived := deleteReceived.value()
	rReceived := renameReceived.value()
	rReceived := renameReceived.value()
	if dReceived+rReceived != 1 {
	if dReceived+rReceived != 1 {
		t.Fatalf("incorrect number of rename+delete events received after 500 ms (%d vs %d)", rReceived+dReceived, 1)
		t.Fatalf("incorrect number of rename+delete events received after 500 ms (%d vs %d)", rReceived+dReceived, 1)
	}
	}


	// Try closing the fsnotify instance
	// Try closing the fsnotify instance
	t.Log("calling Close()")
	t.Log("calling Close()")
	watcher.Close()
	watcher.Close()
	t.Log("waiting for the event channel to become closed...")
	t.Log("waiting for the event channel to become closed...")
	select {
	select {
	case <-done:
	case <-done:
		t.Log("event channel closed")
		t.Log("event channel closed")
	case <-time.After(2 * time.Second):
	case <-time.After(2 * time.Second):
		t.Fatal("event stream was not closed after 2 seconds")
		t.Fatal("event stream was not closed after 2 seconds")
	}
	}
}
}


func TestFsnotifyMultipleCreates(t *testing.T) {
func TestFsnotifyMultipleCreates(t *testing.T) {
	watcher := newWatcher(t)
	watcher := newWatcher(t)


	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for err := range watcher.Error {
		for err := range watcher.Error {
			t.Fatalf("error received: %s", err)
			t.Fatalf("error received: %s", err)
		}
		}
	}()
	}()


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	testFile := filepath.Join(testDir, "TestFsnotifySeq.testfile")
	testFile := filepath.Join(testDir, "TestFsnotifySeq.testfile")


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)


	// Receive events on the event channel on a separate goroutine
	// Receive events on the event channel on a separate goroutine
	eventstream := watcher.Event
	eventstream := watcher.Event
	var createReceived, modifyReceived, deleteReceived counter
	var createReceived, modifyReceived, deleteReceived counter
	done := make(chan bool)
	done := make(chan bool)
	go func() {
	go func() {
		for event := range eventstream {
		for event := range eventstream {
			// Only count relevant events
			// Only count relevant events
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) {
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) {
				t.Logf("event received: %s", event)
				t.Logf("event received: %s", event)
				if event.IsDelete() {
				if event.IsDelete() {
					deleteReceived.increment()
					deleteReceived.increment()
				}
				}
				if event.IsCreate() {
				if event.IsCreate() {
					createReceived.increment()
					createReceived.increment()
				}
				}
				if event.IsModify() {
				if event.IsModify() {
					modifyReceived.increment()
					modifyReceived.increment()
				}
				}
			} else {
			} else {
				t.Logf("unexpected event received: %s", event)
				t.Logf("unexpected event received: %s", event)
			}
			}
		}
		}
		done <- true
		done <- true
	}()
	}()


	// Create a file
	// Create a file
	// This should add at least one event to the fsnotify event queue
	// This should add at least one event to the fsnotify event queue
	var f *os.File
	var f *os.File
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()


	time.Sleep(time.Millisecond)
	time.Sleep(time.Millisecond)
	f.WriteString("data")
	f.WriteString("data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete


	os.Remove(testFile)
	os.Remove(testFile)


	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete


	// Recreate the file
	// Recreate the file
	f, err = os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	f, err = os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Close()
	f.Close()
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete


	// Modify
	// Modify
	f, err = os.OpenFile(testFile, os.O_WRONLY, 0666)
	f, err = os.OpenFile(testFile, os.O_WRONLY, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()


	time.Sleep(time.Millisecond)
	time.Sleep(time.Millisecond)
	f.WriteString("data")
	f.WriteString("data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete


	// Modify
	// Modify
	f, err = os.OpenFile(testFile, os.O_WRONLY, 0666)
	f, err = os.OpenFile(testFile, os.O_WRONLY, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()


	time.Sleep(time.Millisecond)
	time.Sleep(time.Millisecond)
	f.WriteString("data")
	f.WriteString("data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	cReceived := createReceived.value()
	cReceived := createReceived.value()
	if cReceived != 2 {
	if cReceived != 2 {
		t.Fatalf("incorrect number of create events received after 500 ms (%d vs %d)", cReceived, 2)
		t.Fatalf("incorrect number of create events received after 500 ms (%d vs %d)", cReceived, 2)
	}
	}
	mReceived := modifyReceived.value()
	mReceived := modifyReceived.value()
	if mReceived < 3 {
	if mReceived < 3 {
		t.Fatalf("incorrect number of modify events received after 500 ms (%d vs atleast %d)", mReceived, 3)
		t.Fatalf("incorrect number of modify events received after 500 ms (%d vs atleast %d)", mReceived, 3)
	}
	}
	dReceived := deleteReceived.value()
	dReceived := deleteReceived.value()
	if dReceived != 1 {
	if dReceived != 1 {
		t.Fatalf("incorrect number of rename+delete events received after 500 ms (%d vs %d)", dReceived, 1)
		t.Fatalf("incorrect number of rename+delete events received after 500 ms (%d vs %d)", dReceived, 1)
	}
	}


	// Try closing the fsnotify instance
	// Try closing the fsnotify instance
	t.Log("calling Close()")
	t.Log("calling Close()")
	watcher.Close()
	watcher.Close()
	t.Log("waiting for the event channel to become closed...")
	t.Log("waiting for the event channel to become closed...")
	select {
	select {
	case <-done:
	case <-done:
		t.Log("event channel closed")
		t.Log("event channel closed")
	case <-time.After(2 * time.Second):
	case <-time.After(2 * time.Second):
		t.Fatal("event stream was not closed after 2 seconds")
		t.Fatal("event stream was not closed after 2 seconds")
	}
	}
}
}


func TestFsnotifyDirOnly(t *testing.T) {
func TestFsnotifyDirOnly(t *testing.T) {
	watcher := newWatcher(t)
	watcher := newWatcher(t)


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	// Create a file before watching directory
	// Create a file before watching directory
	// This should NOT add any events to the fsnotify event queue
	// This should NOT add any events to the fsnotify event queue
	testFileAlreadyExists := filepath.Join(testDir, "TestFsnotifyEventsExisting.testfile")
	testFileAlreadyExists := filepath.Join(testDir, "TestFsnotifyEventsExisting.testfile")
	{
	{
		var f *os.File
		var f *os.File
		f, err := os.OpenFile(testFileAlreadyExists, os.O_WRONLY|os.O_CREATE, 0666)
		f, err := os.OpenFile(testFileAlreadyExists, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
		if err != nil {
			t.Fatalf("creating test file failed: %s", err)
			t.Fatalf("creating test file failed: %s", err)
		}
		}
		f.Sync()
		f.Sync()
		f.Close()
		f.Close()
	}
	}


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)


	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for err := range watcher.Error {
		for err := range watcher.Error {
			t.Fatalf("error received: %s", err)
			t.Fatalf("error received: %s", err)
		}
		}
	}()
	}()


	testFile := filepath.Join(testDir, "TestFsnotifyDirOnly.testfile")
	testFile := filepath.Join(testDir, "TestFsnotifyDirOnly.testfile")


	// Receive events on the event channel on a separate goroutine
	// Receive events on the event channel on a separate goroutine
	eventstream := watcher.Event
	eventstream := watcher.Event
	var createReceived, modifyReceived, deleteReceived counter
	var createReceived, modifyReceived, deleteReceived counter
	done := make(chan bool)
	done := make(chan bool)
	go func() {
	go func() {
		for event := range eventstream {
		for event := range eventstream {
			// Only count relevant events
			// Only count relevant events
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) || event.Name == filepath.Clean(testFileAlreadyExists) {
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) || event.Name == filepath.Clean(testFileAlreadyExists) {
				t.Logf("event received: %s", event)
				t.Logf("event received: %s", event)
				if event.IsDelete() {
				if event.IsDelete() {
					deleteReceived.increment()
					deleteReceived.increment()
				}
				}
				if event.IsModify() {
				if event.IsModify() {
					modifyReceived.increment()
					modifyReceived.increment()
				}
				}
				if event.IsCreate() {
				if event.IsCreate() {
					createReceived.increment()
					createReceived.increment()
				}
				}
			} else {
			} else {
				t.Logf("unexpected event received: %s", event)
				t.Logf("unexpected event received: %s", event)
			}
			}
		}
		}
		done <- true
		done <- true
	}()
	}()


	// Create a file
	// Create a file
	// This should add at least one event to the fsnotify event queue
	// This should add at least one event to the fsnotify event queue
	var f *os.File
	var f *os.File
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()


	time.Sleep(time.Millisecond)
	time.Sleep(time.Millisecond)
	f.WriteString("data")
	f.WriteString("data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete
	time.Sleep(50 * time.Millisecond) // give system time to sync write change before delete


	os.Remove(testFile)
	os.Remove(testFile)
	os.Remove(testFileAlreadyExists)
	os.Remove(testFileAlreadyExists)


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	cReceived := createReceived.value()
	cReceived := createReceived.value()
	if cReceived != 1 {
	if cReceived != 1 {
		t.Fatalf("incorrect number of create events received after 500 ms (%d vs %d)", cReceived, 1)
		t.Fatalf("incorrect number of create events received after 500 ms (%d vs %d)", cReceived, 1)
	}
	}
	mReceived := modifyReceived.value()
	mReceived := modifyReceived.value()
	if mReceived != 1 {
	if mReceived != 1 {
		t.Fatalf("incorrect number of modify events received after 500 ms (%d vs %d)", mReceived, 1)
		t.Fatalf("incorrect number of modify events received after 500 ms (%d vs %d)", mReceived, 1)
	}
	}
	dReceived := deleteReceived.value()
	dReceived := deleteReceived.value()
	if dReceived != 2 {
	if dReceived != 2 {
		t.Fatalf("incorrect number of delete events received after 500 ms (%d vs %d)", dReceived, 2)
		t.Fatalf("incorrect number of delete events received after 500 ms (%d vs %d)", dReceived, 2)
	}
	}


	// Try closing the fsnotify instance
	// Try closing the fsnotify instance
	t.Log("calling Close()")
	t.Log("calling Close()")
	watcher.Close()
	watcher.Close()
	t.Log("waiting for the event channel to become closed...")
	t.Log("waiting for the event channel to become closed...")
	select {
	select {
	case <-done:
	case <-done:
		t.Log("event channel closed")
		t.Log("event channel closed")
	case <-time.After(2 * time.Second):
	case <-time.After(2 * time.Second):
		t.Fatal("event stream was not closed after 2 seconds")
		t.Fatal("event stream was not closed after 2 seconds")
	}
	}
}
}


func TestFsnotifyDeleteWatchedDir(t *testing.T) {
func TestFsnotifyDeleteWatchedDir(t *testing.T) {
	watcher := newWatcher(t)
	watcher := newWatcher(t)
	defer watcher.Close()
	defer watcher.Close()


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	// Create a file before watching directory
	// Create a file before watching directory
	testFileAlreadyExists := filepath.Join(testDir, "TestFsnotifyEventsExisting.testfile")
	testFileAlreadyExists := filepath.Join(testDir, "TestFsnotifyEventsExisting.testfile")
	{
	{
		var f *os.File
		var f *os.File
		f, err := os.OpenFile(testFileAlreadyExists, os.O_WRONLY|os.O_CREATE, 0666)
		f, err := os.OpenFile(testFileAlreadyExists, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
		if err != nil {
			t.Fatalf("creating test file failed: %s", err)
			t.Fatalf("creating test file failed: %s", err)
		}
		}
		f.Sync()
		f.Sync()
		f.Close()
		f.Close()
	}
	}


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)


	// Add a watch for testFile
	// Add a watch for testFile
	addWatch(t, watcher, testFileAlreadyExists)
	addWatch(t, watcher, testFileAlreadyExists)


	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for err := range watcher.Error {
		for err := range watcher.Error {
			t.Fatalf("error received: %s", err)
			t.Fatalf("error received: %s", err)
		}
		}
	}()
	}()


	// Receive events on the event channel on a separate goroutine
	// Receive events on the event channel on a separate goroutine
	eventstream := watcher.Event
	eventstream := watcher.Event
	var deleteReceived counter
	var deleteReceived counter
	go func() {
	go func() {
		for event := range eventstream {
		for event := range eventstream {
			// Only count relevant events
			// Only count relevant events
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFileAlreadyExists) {
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFileAlreadyExists) {
				t.Logf("event received: %s", event)
				t.Logf("event received: %s", event)
				if event.IsDelete() {
				if event.IsDelete() {
					deleteReceived.increment()
					deleteReceived.increment()
				}
				}
			} else {
			} else {
				t.Logf("unexpected event received: %s", event)
				t.Logf("unexpected event received: %s", event)
			}
			}
		}
		}
	}()
	}()


	os.RemoveAll(testDir)
	os.RemoveAll(testDir)


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	dReceived := deleteReceived.value()
	dReceived := deleteReceived.value()
	if dReceived < 2 {
	if dReceived < 2 {
		t.Fatalf("did not receive at least %d delete events, received %d after 500 ms", 2, dReceived)
		t.Fatalf("did not receive at least %d delete events, received %d after 500 ms", 2, dReceived)
	}
	}
}
}


func TestFsnotifySubDir(t *testing.T) {
func TestFsnotifySubDir(t *testing.T) {
	watcher := newWatcher(t)
	watcher := newWatcher(t)


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	testFile1 := filepath.Join(testDir, "TestFsnotifyFile1.testfile")
	testFile1 := filepath.Join(testDir, "TestFsnotifyFile1.testfile")
	testSubDir := filepath.Join(testDir, "sub")
	testSubDir := filepath.Join(testDir, "sub")
	testSubDirFile := filepath.Join(testDir, "sub/TestFsnotifyFile1.testfile")
	testSubDirFile := filepath.Join(testDir, "sub/TestFsnotifyFile1.testfile")


	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for err := range watcher.Error {
		for err := range watcher.Error {
			t.Fatalf("error received: %s", err)
			t.Fatalf("error received: %s", err)
		}
		}
	}()
	}()


	// Receive events on the event channel on a separate goroutine
	// Receive events on the event channel on a separate goroutine
	eventstream := watcher.Event
	eventstream := watcher.Event
	var createReceived, deleteReceived counter
	var createReceived, deleteReceived counter
	done := make(chan bool)
	done := make(chan bool)
	go func() {
	go func() {
		for event := range eventstream {
		for event := range eventstream {
			// Only count relevant events
			// Only count relevant events
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testSubDir) || event.Name == filepath.Clean(testFile1) {
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testSubDir) || event.Name == filepath.Clean(testFile1) {
				t.Logf("event received: %s", event)
				t.Logf("event received: %s", event)
				if event.IsCreate() {
				if event.IsCreate() {
					createReceived.increment()
					createReceived.increment()
				}
				}
				if event.IsDelete() {
				if event.IsDelete() {
					deleteReceived.increment()
					deleteReceived.increment()
				}
				}
			} else {
			} else {
				t.Logf("unexpected event received: %s", event)
				t.Logf("unexpected event received: %s", event)
			}
			}
		}
		}
		done <- true
		done <- true
	}()
	}()


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)


	// Create sub-directory
	// Create sub-directory
	if err := os.Mkdir(testSubDir, 0777); err != nil {
	if err := os.Mkdir(testSubDir, 0777); err != nil {
		t.Fatalf("failed to create test sub-directory: %s", err)
		t.Fatalf("failed to create test sub-directory: %s", err)
	}
	}


	// Create a file
	// Create a file
	var f *os.File
	var f *os.File
	f, err := os.OpenFile(testFile1, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.OpenFile(testFile1, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	// Create a file (Should not see this! we are not watching subdir)
	// Create a file (Should not see this! we are not watching subdir)
	var fs *os.File
	var fs *os.File
	fs, err = os.OpenFile(testSubDirFile, os.O_WRONLY|os.O_CREATE, 0666)
	fs, err = os.OpenFile(testSubDirFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	fs.Sync()
	fs.Sync()
	fs.Close()
	fs.Close()


	time.Sleep(200 * time.Millisecond)
	time.Sleep(200 * time.Millisecond)


	// Make sure receive deletes for both file and sub-directory
	// Make sure receive deletes for both file and sub-directory
	os.RemoveAll(testSubDir)
	os.RemoveAll(testSubDir)
	os.Remove(testFile1)
	os.Remove(testFile1)


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	cReceived := createReceived.value()
	cReceived := createReceived.value()
	if cReceived != 2 {
	if cReceived != 2 {
		t.Fatalf("incorrect number of create events received after 500 ms (%d vs %d)", cReceived, 2)
		t.Fatalf("incorrect number of create events received after 500 ms (%d vs %d)", cReceived, 2)
	}
	}
	dReceived := deleteReceived.value()
	dReceived := deleteReceived.value()
	if dReceived != 2 {
	if dReceived != 2 {
		t.Fatalf("incorrect number of delete events received after 500 ms (%d vs %d)", dReceived, 2)
		t.Fatalf("incorrect number of delete events received after 500 ms (%d vs %d)", dReceived, 2)
	}
	}


	// Try closing the fsnotify instance
	// Try closing the fsnotify instance
	t.Log("calling Close()")
	t.Log("calling Close()")
	watcher.Close()
	watcher.Close()
	t.Log("waiting for the event channel to become closed...")
	t.Log("waiting for the event channel to become closed...")
	select {
	select {
	case <-done:
	case <-done:
		t.Log("event channel closed")
		t.Log("event channel closed")
	case <-time.After(2 * time.Second):
	case <-time.After(2 * time.Second):
		t.Fatal("event stream was not closed after 2 seconds")
		t.Fatal("event stream was not closed after 2 seconds")
	}
	}
}
}


func TestFsnotifyRename(t *testing.T) {
func TestFsnotifyRename(t *testing.T) {
	watcher := newWatcher(t)
	watcher := newWatcher(t)


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)


	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for err := range watcher.Error {
		for err := range watcher.Error {
			t.Fatalf("error received: %s", err)
			t.Fatalf("error received: %s", err)
		}
		}
	}()
	}()


	testFile := filepath.Join(testDir, "TestFsnotifyEvents.testfile")
	testFile := filepath.Join(testDir, "TestFsnotifyEvents.testfile")
	testFileRenamed := filepath.Join(testDir, "TestFsnotifyEvents.testfileRenamed")
	testFileRenamed := filepath.Join(testDir, "TestFsnotifyEvents.testfileRenamed")


	// Receive events on the event channel on a separate goroutine
	// Receive events on the event channel on a separate goroutine
	eventstream := watcher.Event
	eventstream := watcher.Event
	var renameReceived counter
	var renameReceived counter
	done := make(chan bool)
	done := make(chan bool)
	go func() {
	go func() {
		for event := range eventstream {
		for event := range eventstream {
			// Only count relevant events
			// Only count relevant events
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) || event.Name == filepath.Clean(testFileRenamed) {
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) || event.Name == filepath.Clean(testFileRenamed) {
				if event.IsRename() {
				if event.IsRename() {
					renameReceived.increment()
					renameReceived.increment()
				}
				}
				t.Logf("event received: %s", event)
				t.Logf("event received: %s", event)
			} else {
			} else {
				t.Logf("unexpected event received: %s", event)
				t.Logf("unexpected event received: %s", event)
			}
			}
		}
		}
		done <- true
		done <- true
	}()
	}()


	// Create a file
	// Create a file
	// This should add at least one event to the fsnotify event queue
	// This should add at least one event to the fsnotify event queue
	var f *os.File
	var f *os.File
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()


	f.WriteString("data")
	f.WriteString("data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	// Add a watch for testFile
	// Add a watch for testFile
	addWatch(t, watcher, testFile)
	addWatch(t, watcher, testFile)


	if err := testRename(testFile, testFileRenamed); err != nil {
	if err := testRename(testFile, testFileRenamed); err != nil {
		t.Fatalf("rename failed: %s", err)
		t.Fatalf("rename failed: %s", err)
	}
	}


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	if renameReceived.value() == 0 {
	if renameReceived.value() == 0 {
		t.Fatal("fsnotify rename events have not been received after 500 ms")
		t.Fatal("fsnotify rename events have not been received after 500 ms")
	}
	}


	// Try closing the fsnotify instance
	// Try closing the fsnotify instance
	t.Log("calling Close()")
	t.Log("calling Close()")
	watcher.Close()
	watcher.Close()
	t.Log("waiting for the event channel to become closed...")
	t.Log("waiting for the event channel to become closed...")
	select {
	select {
	case <-done:
	case <-done:
		t.Log("event channel closed")
		t.Log("event channel closed")
	case <-time.After(2 * time.Second):
	case <-time.After(2 * time.Second):
		t.Fatal("event stream was not closed after 2 seconds")
		t.Fatal("event stream was not closed after 2 seconds")
	}
	}


	os.Remove(testFileRenamed)
	os.Remove(testFileRenamed)
}
}


func TestFsnotifyRenameToCreate(t *testing.T) {
func TestFsnotifyRenameToCreate(t *testing.T) {
	watcher := newWatcher(t)
	watcher := newWatcher(t)


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	// Create directory to get file
	// Create directory to get file
	testDirFrom := tempMkdir(t)
	testDirFrom := tempMkdir(t)
	defer os.RemoveAll(testDirFrom)
	defer os.RemoveAll(testDirFrom)


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)


	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for err := range watcher.Error {
		for err := range watcher.Error {
			t.Fatalf("error received: %s", err)
			t.Fatalf("error received: %s", err)
		}
		}
	}()
	}()


	testFile := filepath.Join(testDirFrom, "TestFsnotifyEvents.testfile")
	testFile := filepath.Join(testDirFrom, "TestFsnotifyEvents.testfile")
	testFileRenamed := filepath.Join(testDir, "TestFsnotifyEvents.testfileRenamed")
	testFileRenamed := filepath.Join(testDir, "TestFsnotifyEvents.testfileRenamed")


	// Receive events on the event channel on a separate goroutine
	// Receive events on the event channel on a separate goroutine
	eventstream := watcher.Event
	eventstream := watcher.Event
	var createReceived counter
	var createReceived counter
	done := make(chan bool)
	done := make(chan bool)
	go func() {
	go func() {
		for event := range eventstream {
		for event := range eventstream {
			// Only count relevant events
			// Only count relevant events
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) || event.Name == filepath.Clean(testFileRenamed) {
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) || event.Name == filepath.Clean(testFileRenamed) {
				if event.IsCreate() {
				if event.IsCreate() {
					createReceived.increment()
					createReceived.increment()
				}
				}
				t.Logf("event received: %s", event)
				t.Logf("event received: %s", event)
			} else {
			} else {
				t.Logf("unexpected event received: %s", event)
				t.Logf("unexpected event received: %s", event)
			}
			}
		}
		}
		done <- true
		done <- true
	}()
	}()


	// Create a file
	// Create a file
	// This should add at least one event to the fsnotify event queue
	// This should add at least one event to the fsnotify event queue
	var f *os.File
	var f *os.File
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	if err := testRename(testFile, testFileRenamed); err != nil {
	if err := testRename(testFile, testFileRenamed); err != nil {
		t.Fatalf("rename failed: %s", err)
		t.Fatalf("rename failed: %s", err)
	}
	}


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	if createReceived.value() == 0 {
	if createReceived.value() == 0 {
		t.Fatal("fsnotify create events have not been received after 500 ms")
		t.Fatal("fsnotify create events have not been received after 500 ms")
	}
	}


	// Try closing the fsnotify instance
	// Try closing the fsnotify instance
	t.Log("calling Close()")
	t.Log("calling Close()")
	watcher.Close()
	watcher.Close()
	t.Log("waiting for the event channel to become closed...")
	t.Log("waiting for the event channel to become closed...")
	select {
	select {
	case <-done:
	case <-done:
		t.Log("event channel closed")
		t.Log("event channel closed")
	case <-time.After(2 * time.Second):
	case <-time.After(2 * time.Second):
		t.Fatal("event stream was not closed after 2 seconds")
		t.Fatal("event stream was not closed after 2 seconds")
	}
	}


	os.Remove(testFileRenamed)
	os.Remove(testFileRenamed)
}
}


func TestFsnotifyRenameToOverwrite(t *testing.T) {
func TestFsnotifyRenameToOverwrite(t *testing.T) {
	switch runtime.GOOS {
	switch runtime.GOOS {
	case "plan9", "windows":
	case "plan9", "windows":
		t.Skipf("skipping test on %q (os.Rename over existing file does not create event).", runtime.GOOS)
		t.Skipf("skipping test on %q (os.Rename over existing file does not create event).", runtime.GOOS)
	}
	}


	watcher := newWatcher(t)
	watcher := newWatcher(t)


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	// Create directory to get file
	// Create directory to get file
	testDirFrom := tempMkdir(t)
	testDirFrom := tempMkdir(t)
	defer os.RemoveAll(testDirFrom)
	defer os.RemoveAll(testDirFrom)


	testFile := filepath.Join(testDirFrom, "TestFsnotifyEvents.testfile")
	testFile := filepath.Join(testDirFrom, "TestFsnotifyEvents.testfile")
	testFileRenamed := filepath.Join(testDir, "TestFsnotifyEvents.testfileRenamed")
	testFileRenamed := filepath.Join(testDir, "TestFsnotifyEvents.testfileRenamed")


	// Create a file
	// Create a file
	var fr *os.File
	var fr *os.File
	fr, err := os.OpenFile(testFileRenamed, os.O_WRONLY|os.O_CREATE, 0666)
	fr, err := os.OpenFile(testFileRenamed, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	fr.Sync()
	fr.Sync()
	fr.Close()
	fr.Close()


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)


	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for err := range watcher.Error {
		for err := range watcher.Error {
			t.Fatalf("error received: %s", err)
			t.Fatalf("error received: %s", err)
		}
		}
	}()
	}()


	// Receive events on the event channel on a separate goroutine
	// Receive events on the event channel on a separate goroutine
	eventstream := watcher.Event
	eventstream := watcher.Event
	var eventReceived counter
	var eventReceived counter
	done := make(chan bool)
	done := make(chan bool)
	go func() {
	go func() {
		for event := range eventstream {
		for event := range eventstream {
			// Only count relevant events
			// Only count relevant events
			if event.Name == filepath.Clean(testFileRenamed) {
			if event.Name == filepath.Clean(testFileRenamed) {
				eventReceived.increment()
				eventReceived.increment()
				t.Logf("event received: %s", event)
				t.Logf("event received: %s", event)
			} else {
			} else {
				t.Logf("unexpected event received: %s", event)
				t.Logf("unexpected event received: %s", event)
			}
			}
		}
		}
		done <- true
		done <- true
	}()
	}()


	// Create a file
	// Create a file
	// This should add at least one event to the fsnotify event queue
	// This should add at least one event to the fsnotify event queue
	var f *os.File
	var f *os.File
	f, err = os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	f, err = os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	if err := testRename(testFile, testFileRenamed); err != nil {
	if err := testRename(testFile, testFileRenamed); err != nil {
		t.Fatalf("rename failed: %s", err)
		t.Fatalf("rename failed: %s", err)
	}
	}


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	if eventReceived.value() == 0 {
	if eventReceived.value() == 0 {
		t.Fatal("fsnotify events have not been received after 500 ms")
		t.Fatal("fsnotify events have not been received after 500 ms")
	}
	}


	// Try closing the fsnotify instance
	// Try closing the fsnotify instance
	t.Log("calling Close()")
	t.Log("calling Close()")
	watcher.Close()
	watcher.Close()
	t.Log("waiting for the event channel to become closed...")
	t.Log("waiting for the event channel to become closed...")
	select {
	select {
	case <-done:
	case <-done:
		t.Log("event channel closed")
		t.Log("event channel closed")
	case <-time.After(2 * time.Second):
	case <-time.After(2 * time.Second):
		t.Fatal("event stream was not closed after 2 seconds")
		t.Fatal("event stream was not closed after 2 seconds")
	}
	}


	os.Remove(testFileRenamed)
	os.Remove(testFileRenamed)
}
}


func TestRemovalOfWatch(t *testing.T) {
func TestRemovalOfWatch(t *testing.T) {
	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	// Create a file before watching directory
	// Create a file before watching directory
	testFileAlreadyExists := filepath.Join(testDir, "TestFsnotifyEventsExisting.testfile")
	testFileAlreadyExists := filepath.Join(testDir, "TestFsnotifyEventsExisting.testfile")
	{
	{
		var f *os.File
		var f *os.File
		f, err := os.OpenFile(testFileAlreadyExists, os.O_WRONLY|os.O_CREATE, 0666)
		f, err := os.OpenFile(testFileAlreadyExists, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
		if err != nil {
			t.Fatalf("creating test file failed: %s", err)
			t.Fatalf("creating test file failed: %s", err)
		}
		}
		f.Sync()
		f.Sync()
		f.Close()
		f.Close()
	}
	}


	watcher := newWatcher(t)
	watcher := newWatcher(t)
	defer watcher.Close()
	defer watcher.Close()


	addWatch(t, watcher, testDir)
	addWatch(t, watcher, testDir)
	if err := watcher.RemoveWatch(testDir); err != nil {
	if err := watcher.RemoveWatch(testDir); err != nil {
		t.Fatalf("Could not remove the watch: %v\n", err)
		t.Fatalf("Could not remove the watch: %v\n", err)
	}
	}


	go func() {
	go func() {
		select {
		select {
		case ev := <-watcher.Event:
		case ev := <-watcher.Event:
			t.Fatalf("We received event: %v\n", ev)
			t.Fatalf("We received event: %v\n", ev)
		case <-time.After(500 * time.Millisecond):
		case <-time.After(500 * time.Millisecond):
			t.Log("No event received, as expected.")
			t.Log("No event received, as expected.")
		}
		}
	}()
	}()


	time.Sleep(200 * time.Millisecond)
	time.Sleep(200 * time.Millisecond)
	// Modify the file outside of the watched dir
	// Modify the file outside of the watched dir
	f, err := os.Open(testFileAlreadyExists)
	f, err := os.Open(testFileAlreadyExists)
	if err != nil {
	if err != nil {
		t.Fatalf("Open test file failed: %s", err)
		t.Fatalf("Open test file failed: %s", err)
	}
	}
	f.WriteString("data")
	f.WriteString("data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()
	if err := os.Chmod(testFileAlreadyExists, 0700); err != nil {
	if err := os.Chmod(testFileAlreadyExists, 0700); err != nil {
		t.Fatalf("chmod failed: %s", err)
		t.Fatalf("chmod failed: %s", err)
	}
	}
	time.Sleep(400 * time.Millisecond)
	time.Sleep(400 * time.Millisecond)
}
}


func TestFsnotifyAttrib(t *testing.T) {
func TestFsnotifyAttrib(t *testing.T) {
	if runtime.GOOS == "windows" {
	if runtime.GOOS == "windows" {
		t.Skip("attributes don't work on Windows.")
		t.Skip("attributes don't work on Windows.")
	}
	}


	watcher := newWatcher(t)
	watcher := newWatcher(t)


	// Create directory to watch
	// Create directory to watch
	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	// Receive errors on the error channel on a separate goroutine
	// Receive errors on the error channel on a separate goroutine
	go func() {
	go func() {
		for err := range watcher.Error {
		for err := range watcher.Error {
			t.Fatalf("error received: %s", err)
			t.Fatalf("error received: %s", err)
		}
		}
	}()
	}()


	testFile := filepath.Join(testDir, "TestFsnotifyAttrib.testfile")
	testFile := filepath.Join(testDir, "TestFsnotifyAttrib.testfile")


	// Receive events on the event channel on a separate goroutine
	// Receive events on the event channel on a separate goroutine
	eventstream := watcher.Event
	eventstream := watcher.Event
	// The modifyReceived counter counts IsModify events that are not IsAttrib,
	// The modifyReceived counter counts IsModify events that are not IsAttrib,
	// and the attribReceived counts IsAttrib events (which are also IsModify as
	// and the attribReceived counts IsAttrib events (which are also IsModify as
	// a consequence).
	// a consequence).
	var modifyReceived counter
	var modifyReceived counter
	var attribReceived counter
	var attribReceived counter
	done := make(chan bool)
	done := make(chan bool)
	go func() {
	go func() {
		for event := range eventstream {
		for event := range eventstream {
			// Only count relevant events
			// Only count relevant events
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) {
			if event.Name == filepath.Clean(testDir) || event.Name == filepath.Clean(testFile) {
				if event.IsModify() {
				if event.IsModify() {
					modifyReceived.increment()
					modifyReceived.increment()
				}
				}
				if event.IsAttrib() {
				if event.IsAttrib() {
					attribReceived.increment()
					attribReceived.increment()
				}
				}
				t.Logf("event received: %s", event)
				t.Logf("event received: %s", event)
			} else {
			} else {
				t.Logf("unexpected event received: %s", event)
				t.Logf("unexpected event received: %s", event)
			}
			}
		}
		}
		done <- true
		done <- true
	}()
	}()


	// Create a file
	// Create a file
	// This should add at least one event to the fsnotify event queue
	// This should add at least one event to the fsnotify event queue
	var f *os.File
	var f *os.File
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	f, err := os.OpenFile(testFile, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
	if err != nil {
		t.Fatalf("creating test file failed: %s", err)
		t.Fatalf("creating test file failed: %s", err)
	}
	}
	f.Sync()
	f.Sync()


	f.WriteString("data")
	f.WriteString("data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	// Add a watch for testFile
	// Add a watch for testFile
	addWatch(t, watcher, testFile)
	addWatch(t, watcher, testFile)


	if err := os.Chmod(testFile, 0700); err != nil {
	if err := os.Chmod(testFile, 0700); err != nil {
		t.Fatalf("chmod failed: %s", err)
		t.Fatalf("chmod failed: %s", err)
	}
	}


	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// We expect this event to be received almost immediately, but let's wait 500 ms to be sure
	// Creating/writing a file changes also the mtime, so IsAttrib should be set to true here
	// Creating/writing a file changes also the mtime, so IsAttrib should be set to true here
	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)
	if modifyReceived.value() == 0 {
	if modifyReceived.value() == 0 {
		t.Fatal("fsnotify modify events have not received after 500 ms")
		t.Fatal("fsnotify modify events have not received after 500 ms")
	}
	}
	if attribReceived.value() == 0 {
	if attribReceived.value() == 0 {
		t.Fatal("fsnotify attribute events have not received after 500 ms")
		t.Fatal("fsnotify attribute events have not received after 500 ms")
	}
	}


	// Modifying the contents of the file does not set the attrib flag (although eg. the mtime
	// Modifying the contents of the file does not set the attrib flag (although eg. the mtime
	// might have been modified).
	// might have been modified).
	modifyReceived.reset()
	modifyReceived.reset()
	attribReceived.reset()
	attribReceived.reset()


	f, err = os.OpenFile(testFile, os.O_WRONLY, 0)
	f, err = os.OpenFile(testFile, os.O_WRONLY, 0)
	if err != nil {
	if err != nil {
		t.Fatalf("reopening test file failed: %s", err)
		t.Fatalf("reopening test file failed: %s", err)
	}
	}


	f.WriteString("more data")
	f.WriteString("more data")
	f.Sync()
	f.Sync()
	f.Close()
	f.Close()


	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)


	if modifyReceived.value() != 1 {
	if modifyReceived.value() != 1 {
		t.Fatal("didn't receive a modify event after changing test file contents")
		t.Fatal("didn't receive a modify event after changing test file contents")
	}
	}


	if attribReceived.value() != 0 {
	if attribReceived.value() != 0 {
		t.Fatal("did receive an unexpected attrib event after changing test file contents")
		t.Fatal("did receive an unexpected attrib event after changing test file contents")
	}
	}


	modifyReceived.reset()
	modifyReceived.reset()
	attribReceived.reset()
	attribReceived.reset()


	// Doing a chmod on the file should trigger an event with the "attrib" flag set (the contents
	// Doing a chmod on the file should trigger an event with the "attrib" flag set (the contents
	// of the file are not changed though)
	// of the file are not changed though)
	if err := os.Chmod(testFile, 0600); err != nil {
	if err := os.Chmod(testFile, 0600); err != nil {
		t.Fatalf("chmod failed: %s", err)
		t.Fatalf("chmod failed: %s", err)
	}
	}


	time.Sleep(500 * time.Millisecond)
	time.Sleep(500 * time.Millisecond)


	if attribReceived.value() != 1 {
	if attribReceived.value() != 1 {
		t.Fatal("didn't receive an attribute change after 500ms")
		t.Fatal("didn't receive an attribute change after 500ms")
	}
	}


	// Try closing the fsnotify instance
	// Try closing the fsnotify instance
	t.Log("calling Close()")
	t.Log("calling Close()")
	watcher.Close()
	watcher.Close()
	t.Log("waiting for the event channel to become closed...")
	t.Log("waiting for the event channel to become closed...")
	select {
	select {
	case <-done:
	case <-done:
		t.Log("event channel closed")
		t.Log("event channel closed")
	case <-time.After(1e9):
	case <-time.After(1e9):
		t.Fatal("event stream was not closed after 1 second")
		t.Fatal("event stream was not closed after 1 second")
	}
	}


	os.Remove(testFile)
	os.Remove(testFile)
}
}


func TestFsnotifyClose(t *testing.T) {
func TestFsnotifyClose(t *testing.T) {
	watcher := newWatcher(t)
	watcher := newWatcher(t)
	watcher.Close()
	watcher.Close()


	var done int32
	var done int32
	go func() {
	go func() {
		watcher.Close()
		watcher.Close()
		atomic.StoreInt32(&done, 1)
		atomic.StoreInt32(&done, 1)
	}()
	}()


	time.Sleep(50e6) // 50 ms
	time.Sleep(50e6) // 50 ms
	if atomic.LoadInt32(&done) == 0 {
	if atomic.LoadInt32(&done) == 0 {
		t.Fatal("double Close() test failed: second Close() call didn't return")
		t.Fatal("double Close() test failed: second Close() call didn't return")
	}
	}


	testDir := tempMkdir(t)
	testDir := tempMkdir(t)
	defer os.RemoveAll(testDir)
	defer os.RemoveAll(testDir)


	if err := watcher.Watch(testDir); err == nil {
	if err := watcher.Watch(testDir); err == nil {
		t.Fatal("expected error on Watch() after Close(), got nil")
		t.Fatal("expected error on Watch() after Close(), got nil")
	}
	}
}
}


func testRename(file1, file2 string) error {
func testRename(file1, file2 string) error {
	switch runtime.GOOS {
	switch runtime.GOOS {
	case "windows", "plan9":
	case "windows", "plan9":
		return os.Rename(file1, file2)
		return os.Rename(file1, file2)
	default:
	default:
		cmd := exec.Command("mv", file1, file2)
		cmd := exec.Command("mv", file1, file2)
		return cmd.Run()
		return cmd.Run()
	}
	}
}
}
