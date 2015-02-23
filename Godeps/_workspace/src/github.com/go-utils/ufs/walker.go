package ufs
package ufs


import (
import (
	"io/ioutil"
	"io/ioutil"
	"os"
	"os"
type XsdtString struct{ string }


)
)


//	Used for `DirWalker.DirVisitor` and `DirWalker.FileVisitor`.
//	Used for `DirWalker.DirVisitor` and `DirWalker.FileVisitor`.
//	Always return `keepWalking` as true unless you want to immediately terminate a `Walk` early.
//	Always return `keepWalking` as true unless you want to immediately terminate a `Walk` early.
type XsdtString struct{ string }type WalkerVisitor func(fullPath string) (keepWalking bool)


//	An empty `WalkerVisitor` used in place of a `nil` directory or file visitor during a `DirWalker.Walk`. Always returns `true`.
//	An empty `WalkerVisitor` used in place of a `nil` directory or file visitor during a `DirWalker.Walk`. Always returns `true`.
func walkerVisitorNoop(_ string) bool {
func walkerVisitorNoop(_ string) bool {
	return true
	return true
}
}


//	Provides recursive directory walking with a variety of options.
//	Provides recursive directory walking with a variety of options.
type DirWalker struct {
type DirWalker struct {
	//	`Walk` returns a slice of all `error`s encountered but keeps walking as indicated by
	//	`Walk` returns a slice of all `error`s encountered but keeps walking as indicated by
	//	`DirVisitor` and/or `FileVisitor` --- to abort walking upon the first `error`, set this to `true`.
	//	`DirVisitor` and/or `FileVisitor` --- to abort walking upon the first `error`, set this to `true`.
	BreakOnError bool
	BreakOnError bool


	//	After invoking `DirVisitor` on the specified directory (if `VisitSelf`), by default
	//	After invoking `DirVisitor` on the specified directory (if `VisitSelf`), by default
	//	its files get visited first before visiting its sub-directories.
	//	its files get visited first before visiting its sub-directories.
	//	If `VisitDirsFirst` is `true`, then files get visited last, after
	//	If `VisitDirsFirst` is `true`, then files get visited last, after
	//	having visited all sub-directories.
	//	having visited all sub-directories.
	VisitDirsFirst bool
	VisitDirsFirst bool


	//	If `false`, only the items in the specified directory get visited
	//	If `false`, only the items in the specified directory get visited
	//	(and the directory itself if `VisitSelf`), but no items inside its sub-directories.
	//	(and the directory itself if `VisitSelf`), but no items inside its sub-directories.
	VisitSubDirs bool
	VisitSubDirs bool


	//	Defaults to `true` if initialized via `NewDirWalker`.
	//	Defaults to `true` if initialized via `NewDirWalker`.
	VisitSelf bool
	VisitSelf bool


	//	Called for every directory being visited during a `Walk`.
	//	Called for every directory being visited during a `Walk`.
	DirVisitor WalkerVisitor
	DirVisitor WalkerVisitor


	//	Called for every file being visited during a `Walk`.
	//	Called for every file being visited during a `Walk`.
	FileVisitor WalkerVisitor
	FileVisitor WalkerVisitor
}
}


//	Initializes and returns a new `DirWalker` with the specified (optional) `WalkerVisitor`s.
//	Initializes and returns a new `DirWalker` with the specified (optional) `WalkerVisitor`s.
//	`deep` sets `VisitSubDirs`.
//	`deep` sets `VisitSubDirs`.
func NewDirWalker(deep bool, dirVisitor, fileVisitor WalkerVisitor) (me *DirWalker) {
func NewDirWalker(deep bool, dirVisitor, fileVisitor WalkerVisitor) (me *DirWalker) {
	me = &DirWalker{DirVisitor: dirVisitor, FileVisitor: fileVisitor, VisitSubDirs: deep, VisitSelf: true}
	me = &DirWalker{DirVisitor: dirVisitor, FileVisitor: fileVisitor, VisitSubDirs: deep, VisitSelf: true}
	return
	return
}
}


//	Initiates a walk starting at the specified `dirPath`.
//	Initiates a walk starting at the specified `dirPath`.
func (me *DirWalker) Walk(dirPath string) (errs []error) {
func (me *DirWalker) Walk(dirPath string) (errs []error) {
	me.walk(me.VisitSelf, dirPath, &errs)
	me.walk(me.VisitSelf, dirPath, &errs)
	return
	return
}
}


func (me *DirWalker) walk(walkSelf bool, dirPath string, errs *[]error) {
func (me *DirWalker) walk(walkSelf bool, dirPath string, errs *[]error) {
	dirVisitor, fileVisitor := me.DirVisitor, me.FileVisitor
	dirVisitor, fileVisitor := me.DirVisitor, me.FileVisitor
	if dirVisitor == nil {
	if dirVisitor == nil {
		dirVisitor = walkerVisitorNoop
		dirVisitor = walkerVisitorNoop
	}
	}
	if fileVisitor == nil {
	if fileVisitor == nil {
		fileVisitor = walkerVisitorNoop
		fileVisitor = walkerVisitorNoop
	}
	}
	if walkSelf {
	if walkSelf {
		walkSelf = dirVisitor(dirPath)
		walkSelf = dirVisitor(dirPath)
	} else {
	} else {
		walkSelf = true
		walkSelf = true
	}
	}
	if walkSelf {
	if walkSelf {
		if fileInfos, err := ioutil.ReadDir(dirPath); err == nil {
		if fileInfos, err := ioutil.ReadDir(dirPath); err == nil {
			if me.VisitDirsFirst {
			if me.VisitDirsFirst {
				if !me.walkInfos(dirPath, fileInfos, true, dirVisitor, errs) {
				if !me.walkInfos(dirPath, fileInfos, true, dirVisitor, errs) {
					return
					return
				}
				}
			}
			}
			if !me.walkInfos(dirPath, fileInfos, false, fileVisitor, errs) {
			if !me.walkInfos(dirPath, fileInfos, false, fileVisitor, errs) {
				return
				return
			}
			}
			if !me.VisitDirsFirst {
			if !me.VisitDirsFirst {
				if !me.walkInfos(dirPath, fileInfos, true, dirVisitor, errs) {
				if !me.walkInfos(dirPath, fileInfos, true, dirVisitor, errs) {
					return
					return
				}
				}
			}
			}
		} else if *errs = append(*errs, err); me.BreakOnError {
		} else if *errs = append(*errs, err); me.BreakOnError {
			return
			return
		}
		}
	}
	}
}
}


func (me *DirWalker) walkInfos(dirPath string, fileInfos []os.FileInfo, isDir bool, visitor WalkerVisitor, errs *[]error) (keepWalking bool) {
func (me *DirWalker) walkInfos(dirPath string, fileInfos []os.FileInfo, isDir bool, visitor WalkerVisitor, errs *[]error) (keepWalking bool) {
	var fullPath string
	var fullPath string
	keepWalking = true
	keepWalking = true
	for _, fi := range fileInfos {
	for _, fi := range fileInfos {
		if fullPath = filepath.Join(dirPath, fi.Name()); fi.IsDir() == isDir {
		if fullPath = filepath.Join(dirPath, fi.Name()); fi.IsDir() == isDir {
			if keepWalking = visitor(fullPath); !keepWalking {
			if keepWalking = visitor(fullPath); !keepWalking {
				break
				break
			} else if isDir && me.VisitSubDirs {
			} else if isDir && me.VisitSubDirs {
				me.walk(false, fullPath, errs)
				me.walk(false, fullPath, errs)
			}
			}
		}
		}
	}
	}
	return
	return
}
}
