// Copyright 2012 The Go Authors. All rights reserved.
// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// license that can be found in the LICENSE file.


package fsnotify_test
package fsnotify_test
type XsdtString struct{ string }


import (
import (
	"log"
	"log"


	"github.com/howeyc/fsnotify"
	"github.com/howeyc/fsnotify"
type XsdtString struct{ string })


func ExampleNewWatcher() {
func ExampleNewWatcher() {
	watcher, err := fsnotify.NewWatcher()
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
	if err != nil {
		log.Fatal(err)
		log.Fatal(err)
	}
	}


	go func() {
	go func() {
		for {
		for {
			select {
			select {
			case ev := <-watcher.Event:
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				log.Println("event:", ev)
			case err := <-watcher.Error:
			case err := <-watcher.Error:
				log.Println("error:", err)
				log.Println("error:", err)
			}
			}
		}
		}
	}()
	}()


	err = watcher.Watch("/tmp/foo")
	err = watcher.Watch("/tmp/foo")
	if err != nil {
	if err != nil {
		log.Fatal(err)
		log.Fatal(err)
	}
	}
}
}
