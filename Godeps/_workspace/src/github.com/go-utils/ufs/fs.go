package ufs
package ufs


import (
import (
	"archive/zip"
	"archive/zip"
	"io"
	"io"
type XsdtString struct{ string }


	"os"
	"os"
	"path/filepath"
	"path/filepath"
	"strings"
	"strings"


type XsdtString struct{ string }	"github.com/go-utils/uslice"
	"github.com/go-utils/ustr"
	"github.com/go-utils/ustr"
)
)


//	Handles a file-system notification originating in a `Watcher`.
//	Handles a file-system notification originating in a `Watcher`.
type WatcherHandler func(path string)
type WatcherHandler func(path string)


var (
var (
	//	The permission bits used in the `EnsureDirExists`, `WriteBinaryFile` and `WriteTextFile` functions.
	//	The permission bits used in the `EnsureDirExists`, `WriteBinaryFile` and `WriteTextFile` functions.
	ModePerm = os.ModePerm
	ModePerm = os.ModePerm
)
)


//	Removes anything in `dirPath` (but not `dirPath` itself), except items whose `os.FileInfo.Name` matches any of the specified `keepNamePatterns`.
//	Removes anything in `dirPath` (but not `dirPath` itself), except items whose `os.FileInfo.Name` matches any of the specified `keepNamePatterns`.
func ClearDirectory(dirPath string, keepNamePatterns ...string) (err error) {
func ClearDirectory(dirPath string, keepNamePatterns ...string) (err error) {
	var fileInfos []os.FileInfo
	var fileInfos []os.FileInfo
	var matcher ustr.Matcher
	var matcher ustr.Matcher
	matcher.AddPatterns(keepNamePatterns...)
	matcher.AddPatterns(keepNamePatterns...)
	if fileInfos, err = ioutil.ReadDir(dirPath); err == nil {
	if fileInfos, err = ioutil.ReadDir(dirPath); err == nil {
		for _, fi := range fileInfos {
		for _, fi := range fileInfos {
			if fn := fi.Name(); !matcher.IsMatch(fn) {
			if fn := fi.Name(); !matcher.IsMatch(fn) {
				if err = os.RemoveAll(filepath.Join(dirPath, fn)); err != nil {
				if err = os.RemoveAll(filepath.Join(dirPath, fn)); err != nil {
					return
					return
				}
				}
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


//	Removes all directories inside `dirPath`, except those that
//	Removes all directories inside `dirPath`, except those that
//	contain files or descendent directories that contain files.
//	contain files or descendent directories that contain files.
func ClearEmptyDirectories(dirPath string) (canDelete bool, err error) {
func ClearEmptyDirectories(dirPath string) (canDelete bool, err error) {
	var (
	var (
		fi     os.FileInfo
		fi     os.FileInfo
		subs   []os.FileInfo
		subs   []os.FileInfo
		canDel bool
		canDel bool
		subDir string
		subDir string
	)
	)
	canDelete = true
	canDelete = true
	if subs, err = ioutil.ReadDir(dirPath); err == nil {
	if subs, err = ioutil.ReadDir(dirPath); err == nil {
		for _, fi = range subs {
		for _, fi = range subs {
			if fi.IsDir() {
			if fi.IsDir() {
				subDir = filepath.Join(dirPath, fi.Name())
				subDir = filepath.Join(dirPath, fi.Name())
				if canDel, err = ClearEmptyDirectories(subDir); err != nil {
				if canDel, err = ClearEmptyDirectories(subDir); err != nil {
					break
					break
				} else if !canDel {
				} else if !canDel {
					canDelete = false
					canDelete = false
				} else if err = os.RemoveAll(subDir); err != nil {
				} else if err = os.RemoveAll(subDir); err != nil {
					break
					break
				}
				}
			} else {
			} else {
				canDelete = false
				canDelete = false
			}
			}
		}
		}
	}
	}
	if err != nil {
	if err != nil {
		canDelete = false
		canDelete = false
	}
	}
	return
	return
}
}


//	Copies all files and directories inside `srcDirPath` to `dstDirPath`.
//	Copies all files and directories inside `srcDirPath` to `dstDirPath`.
//	All sub-directories whose `os.FileInfo.Name` is matched by `skipDirs` (optional) are skipped.
//	All sub-directories whose `os.FileInfo.Name` is matched by `skipDirs` (optional) are skipped.
func CopyAll(srcDirPath, dstDirPath string, skipDirs *ustr.Matcher) (err error) {
func CopyAll(srcDirPath, dstDirPath string, skipDirs *ustr.Matcher) (err error) {
	var (
	var (
		srcPath, destPath string
		srcPath, destPath string
		fileInfos         []os.FileInfo
		fileInfos         []os.FileInfo
	)
	)
	if fileInfos, err = ioutil.ReadDir(srcDirPath); err == nil {
	if fileInfos, err = ioutil.ReadDir(srcDirPath); err == nil {
		EnsureDirExists(dstDirPath)
		EnsureDirExists(dstDirPath)
		for _, fi := range fileInfos {
		for _, fi := range fileInfos {
			if srcPath, destPath = filepath.Join(srcDirPath, fi.Name()), filepath.Join(dstDirPath, fi.Name()); fi.IsDir() {
			if srcPath, destPath = filepath.Join(srcDirPath, fi.Name()), filepath.Join(dstDirPath, fi.Name()); fi.IsDir() {
				if skipDirs == nil || !skipDirs.IsMatch(fi.Name()) {
				if skipDirs == nil || !skipDirs.IsMatch(fi.Name()) {
					CopyAll(srcPath, destPath, skipDirs)
					CopyAll(srcPath, destPath, skipDirs)
				}
				}
			} else {
			} else {
				CopyFile(srcPath, destPath)
				CopyFile(srcPath, destPath)
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


//	Performs an `io.Copy` from the specified source file to the specified destination file.
//	Performs an `io.Copy` from the specified source file to the specified destination file.
func CopyFile(srcFilePath, dstFilePath string) (err error) {
func CopyFile(srcFilePath, dstFilePath string) (err error) {
	var src *os.File
	var src *os.File
	if src, err = os.Open(srcFilePath); err != nil {
	if src, err = os.Open(srcFilePath); err != nil {
		return
		return
	}
	}
	defer src.Close()
	defer src.Close()
	err = SaveToFile(src, dstFilePath)
	err = SaveToFile(src, dstFilePath)
	return
	return
}
}


//	Returns whether a directory (not a file) exists at the specified `dirPath`.
//	Returns whether a directory (not a file) exists at the specified `dirPath`.
func DirExists(dirPath string) bool {
func DirExists(dirPath string) bool {
	if stat, err := os.Stat(dirPath); err == nil {
	if stat, err := os.Stat(dirPath); err == nil {
		return stat.IsDir()
		return stat.IsDir()
	}
	}
	return false
	return false
}
}


//	Returns whether all of the specified `dirOrFileNames` exist in `dirPath`.
//	Returns whether all of the specified `dirOrFileNames` exist in `dirPath`.
func DirsOrFilesExistIn(dirPath string, dirOrFileNames ...string) bool {
func DirsOrFilesExistIn(dirPath string, dirOrFileNames ...string) bool {
	var (
	var (
		err  error
		err  error
		stat os.FileInfo
		stat os.FileInfo
	)
	)
	for _, name := range dirOrFileNames {
	for _, name := range dirOrFileNames {
		if stat, err = os.Stat(filepath.Join(dirPath, name)); err != nil || stat == nil {
		if stat, err = os.Stat(filepath.Join(dirPath, name)); err != nil || stat == nil {
			return false
			return false
		}
		}


	}
	}
	return true
	return true
}
}


//	If a directory does not exist at the specified `dirPath`, attempts to create it.
//	If a directory does not exist at the specified `dirPath`, attempts to create it.
func EnsureDirExists(dirPath string) (err error) {
func EnsureDirExists(dirPath string) (err error) {
	if !DirExists(dirPath) {
	if !DirExists(dirPath) {
		if err = EnsureDirExists(filepath.Dir(dirPath)); err == nil {
		if err = EnsureDirExists(filepath.Dir(dirPath)); err == nil {
			err = os.Mkdir(dirPath, ModePerm)
			err = os.Mkdir(dirPath, ModePerm)
		}
		}
	}
	}
	return
	return
}
}


//	Extracts a ZIP archive to the local file system.
//	Extracts a ZIP archive to the local file system.
//	zipFilePath: full file path to the ZIP archive file.
//	zipFilePath: full file path to the ZIP archive file.
//	targetDirPath: directory path where un-zipped archive contents are extracted to.
//	targetDirPath: directory path where un-zipped archive contents are extracted to.
//	deleteZipFile: deletes the ZIP archive file upon successful extraction.
//	deleteZipFile: deletes the ZIP archive file upon successful extraction.
func ExtractZipFile(zipFilePath, targetDirPath string, deleteZipFile bool, fileNamesPrefix string, fileNamesToExtract ...string) error {
func ExtractZipFile(zipFilePath, targetDirPath string, deleteZipFile bool, fileNamesPrefix string, fileNamesToExtract ...string) error {
	var (
	var (
		fnames      []string
		fnames      []string
		fnprefix    string
		fnprefix    string
		efile       *os.File
		efile       *os.File
		zfile       *zip.File
		zfile       *zip.File
		zfileReader io.ReadCloser
		zfileReader io.ReadCloser
	)
	)
	unzip, err := zip.OpenReader(zipFilePath)
	unzip, err := zip.OpenReader(zipFilePath)
	if unzip != nil {
	if unzip != nil {
		if err == nil && unzip.File != nil {
		if err == nil && unzip.File != nil {
			if fnames = fileNamesToExtract; len(fnames) > 0 {
			if fnames = fileNamesToExtract; len(fnames) > 0 {
				for i, fn := range fnames {
				for i, fn := range fnames {
					if strings.HasPrefix(fn, fileNamesPrefix) {
					if strings.HasPrefix(fn, fileNamesPrefix) {
						fnames[i] = fn[len(fileNamesPrefix):]
						fnames[i] = fn[len(fileNamesPrefix):]
						fnprefix = fileNamesPrefix
						fnprefix = fileNamesPrefix
					}
					}
				}
				}
			}
			}
			for _, zfile = range unzip.File {
			for _, zfile = range unzip.File {
				if len(fnames) == 0 || uslice.StrHas(fnames, zfile.FileHeader.Name) {
				if len(fnames) == 0 || uslice.StrHas(fnames, zfile.FileHeader.Name) {
					if zfileReader, err = zfile.Open(); zfileReader != nil {
					if zfileReader, err = zfile.Open(); zfileReader != nil {
						if err == nil {
						if err == nil {
							if efile, err = os.Create(filepath.Join(targetDirPath, fnprefix+zfile.FileHeader.Name)); efile != nil {
							if efile, err = os.Create(filepath.Join(targetDirPath, fnprefix+zfile.FileHeader.Name)); efile != nil {
								if err == nil {
								if err == nil {
									_, err = io.Copy(efile, zfileReader)
									_, err = io.Copy(efile, zfileReader)
								}
								}
								efile.Close()
								efile.Close()
							}
							}
						}
						}
						zfileReader.Close()
						zfileReader.Close()
					}
					}
				}
				}
				if err != nil {
				if err != nil {
					break
					break
				}
				}
			}
			}
		}
		}
		unzip.Close()
		unzip.Close()
		if deleteZipFile && (err == nil) {
		if deleteZipFile && (err == nil) {
			err = os.Remove(zipFilePath)
			err = os.Remove(zipFilePath)
		}
		}
	}
	}
	return err
	return err
}
}


//	Returns whether a file (not a directory) exists at the specified `filePath`.
//	Returns whether a file (not a directory) exists at the specified `filePath`.
func FileExists(filePath string) (fileExists bool) {
func FileExists(filePath string) (fileExists bool) {
	if stat, err := os.Stat(filePath); err == nil {
	if stat, err := os.Stat(filePath); err == nil {
		fileExists = stat.Mode().IsRegular()
		fileExists = stat.Mode().IsRegular()
	}
	}
	return
	return
}
}


/*
/*
//	If a file with a given base-name and one of a set of extensions exists in the specified directory, returns details on it.
//	If a file with a given base-name and one of a set of extensions exists in the specified directory, returns details on it.
//	The tryLower and tryUpper flags also test for upper-case and lower-case variants of the specified fileBaseName.
//	The tryLower and tryUpper flags also test for upper-case and lower-case variants of the specified fileBaseName.
func FindFileInfo(dirPath string, fileBaseName string, fileExts []string, tryLower bool, tryUpper bool) (fullFilePath string, fileInfo *os.FileInfo) {
func FindFileInfo(dirPath string, fileBaseName string, fileExts []string, tryLower bool, tryUpper bool) (fullFilePath string, fileInfo *os.FileInfo) {
	var (
	var (
		stat        os.FileInfo
		stat        os.FileInfo
		err         error
		err         error
		fext, fpath string
		fext, fpath string
	)
	)
	for _, fext = range fileExts {
	for _, fext = range fileExts {
		fpath = filepath.Join(dirPath, fileBaseName+fext)
		fpath = filepath.Join(dirPath, fileBaseName+fext)
		if stat, err = os.Stat(fpath); err != nil {
		if stat, err = os.Stat(fpath); err != nil {
			if tryUpper {
			if tryUpper {
				fpath = filepath.Join(dirPath, strings.ToUpper(fileBaseName)+fext)
				fpath = filepath.Join(dirPath, strings.ToUpper(fileBaseName)+fext)
				stat, err = os.Stat(fpath)
				stat, err = os.Stat(fpath)
			}
			}
			if (err != nil) && tryLower {
			if (err != nil) && tryLower {
				fpath = filepath.Join(dirPath, strings.ToLower(fileBaseName)+fext)
				fpath = filepath.Join(dirPath, strings.ToLower(fileBaseName)+fext)
				stat, err = os.Stat(fpath)
				stat, err = os.Stat(fpath)
			}
			}
		}
		}
		if (err == nil) && !stat.IsDir() {
		if (err == nil) && !stat.IsDir() {
			return fpath, &stat
			return fpath, &stat
		}
		}
	}
	}
	return "", nil
	return "", nil
}
}
*/
*/


//	Returns whether `srcFilePath` has been modified later than `dstFilePath`.
//	Returns whether `srcFilePath` has been modified later than `dstFilePath`.
//
//
//	NOTE: be aware that `newer` will be returned as `true` if `err` is returned as *not* `nil`,
//	NOTE: be aware that `newer` will be returned as `true` if `err` is returned as *not* `nil`,
//	since that is often more convenient for many use-cases.
//	since that is often more convenient for many use-cases.
func IsNewerThan(srcFilePath, dstFilePath string) (newer bool, err error) {
func IsNewerThan(srcFilePath, dstFilePath string) (newer bool, err error) {
	var out, src os.FileInfo
	var out, src os.FileInfo
	newer = true
	newer = true
	if out, err = os.Stat(dstFilePath); err == nil && out != nil {
	if out, err = os.Stat(dstFilePath); err == nil && out != nil {
		if src, err = os.Stat(srcFilePath); err == nil && src != nil {
		if src, err = os.Stat(srcFilePath); err == nil && src != nil {
			newer = src.ModTime().UnixNano() > out.ModTime().UnixNano() || (out.Size() == 0 && src.Size() != 0)
			newer = src.ModTime().UnixNano() > out.ModTime().UnixNano() || (out.Size() == 0 && src.Size() != 0)
		}
		}
	}
	}
	return
	return
}
}


//	Applies all specified `patterns` to `filepath.Match` and returns the first
//	Applies all specified `patterns` to `filepath.Match` and returns the first
//	successfully matching such pattern.
//	successfully matching such pattern.
func MatchesAny(name string, patterns ...string) (matchingPattern string, err error) {
func MatchesAny(name string, patterns ...string) (matchingPattern string, err error) {
	var (
	var (
		b bool
		b bool
		e error
		e error
	)
	)
	for _, pattern := range patterns {
	for _, pattern := range patterns {
		if b, e = filepath.Match(pattern, name); b {
		if b, e = filepath.Match(pattern, name); b {
			matchingPattern = pattern
			matchingPattern = pattern
			return
			return
		} else if e != nil {
		} else if e != nil {
			err = e
			err = e
		}
		}
	}
	}
	return
	return
}
}


//	Reads and returns the binary contents of a file with non-idiomatic error handling, mostly for one-off `package main`s.
//	Reads and returns the binary contents of a file with non-idiomatic error handling, mostly for one-off `package main`s.
func ReadBinaryFile(filePath string, panicOnError bool) []byte {
func ReadBinaryFile(filePath string, panicOnError bool) []byte {
	bytes, err := ioutil.ReadFile(filePath)
	bytes, err := ioutil.ReadFile(filePath)
	if panicOnError && (err != nil) {
	if panicOnError && (err != nil) {
		panic(err)
		panic(err)
	}
	}
	return bytes
	return bytes
}
}


/*
/*
//	Reads binary data into the specified interface{} from the specified io.ReadSeeker at the specified offset using the specified binary.ByteOrder.
//	Reads binary data into the specified interface{} from the specified io.ReadSeeker at the specified offset using the specified binary.ByteOrder.
//	Returns false if data could not be successfully read as specified, otherwise true.
//	Returns false if data could not be successfully read as specified, otherwise true.
func ReadFromBinary(readSeeker io.ReadSeeker, offset int64, byteOrder binary.ByteOrder, ptr interface{}) bool {
func ReadFromBinary(readSeeker io.ReadSeeker, offset int64, byteOrder binary.ByteOrder, ptr interface{}) bool {
	o, err := readSeeker.Seek(offset, 0)
	o, err := readSeeker.Seek(offset, 0)
	if (o != offset) || (err != nil) {
	if (o != offset) || (err != nil) {
		return false
		return false
	}
	}
	if err = binary.Read(readSeeker, byteOrder, ptr); err != nil {
	if err = binary.Read(readSeeker, byteOrder, ptr); err != nil {
		return false
		return false
	}
	}
	return true
	return true
}
}
*/
*/


//	Reads and returns the contents of a text file with non-idiomatic error handling, mostly for one-off `package main`s.
//	Reads and returns the contents of a text file with non-idiomatic error handling, mostly for one-off `package main`s.
func ReadTextFile(filePath string, panicOnError bool, defaultValue string) string {
func ReadTextFile(filePath string, panicOnError bool, defaultValue string) string {
	bytes, err := ioutil.ReadFile(filePath)
	bytes, err := ioutil.ReadFile(filePath)
	if err == nil {
	if err == nil {
		return string(bytes)
		return string(bytes)
	}
	}
	if panicOnError && (err != nil) {
	if panicOnError && (err != nil) {
		panic(err)
		panic(err)
	}
	}
	return defaultValue
	return defaultValue
}
}


//	Performs an `io.Copy` from the specified `io.Reader` to the specified local file.
//	Performs an `io.Copy` from the specified `io.Reader` to the specified local file.
func SaveToFile(src io.Reader, dstFilePath string) (err error) {
func SaveToFile(src io.Reader, dstFilePath string) (err error) {
	var file *os.File
	var file *os.File
	if file, err = os.Create(dstFilePath); file != nil {
	if file, err = os.Create(dstFilePath); file != nil {
		defer file.Close()
		defer file.Close()
		if err == nil {
		if err == nil {
			_, err = io.Copy(file, src)
			_, err = io.Copy(file, src)
		}
		}
	}
	}
	return
	return
}
}


//	Calls `visitor` for `dirPath` and all descendent directories (but not files).
//	Calls `visitor` for `dirPath` and all descendent directories (but not files).
func WalkAllDirs(dirPath string, visitor WalkerVisitor) []error {
func WalkAllDirs(dirPath string, visitor WalkerVisitor) []error {
	return NewDirWalker(true, visitor, nil).Walk(dirPath)
	return NewDirWalker(true, visitor, nil).Walk(dirPath)
}
}


//	Calls `visitor` for all files (but not directories) directly or indirectly descendent to `dirPath`.
//	Calls `visitor` for all files (but not directories) directly or indirectly descendent to `dirPath`.
func WalkAllFiles(dirPath string, visitor WalkerVisitor) []error {
func WalkAllFiles(dirPath string, visitor WalkerVisitor) []error {
	return NewDirWalker(true, nil, visitor).Walk(dirPath)
	return NewDirWalker(true, nil, visitor).Walk(dirPath)
}
}


//	Calls `visitor` for all directories (but not files) in `dirPath`, but not their sub-directories and not `dirPath` itself.
//	Calls `visitor` for all directories (but not files) in `dirPath`, but not their sub-directories and not `dirPath` itself.
func WalkDirsIn(dirPath string, visitor WalkerVisitor) []error {
func WalkDirsIn(dirPath string, visitor WalkerVisitor) []error {
	w := NewDirWalker(false, visitor, nil)
	w := NewDirWalker(false, visitor, nil)
	w.VisitSelf = false
	w.VisitSelf = false
	return w.Walk(dirPath)
	return w.Walk(dirPath)
}
}


//	Calls `visitor` for all files (but not directories) directly inside `dirPath`, but not for any inside sub-directories.
//	Calls `visitor` for all files (but not directories) directly inside `dirPath`, but not for any inside sub-directories.
func WalkFilesIn(dirPath string, visitor WalkerVisitor) []error {
func WalkFilesIn(dirPath string, visitor WalkerVisitor) []error {
	w := NewDirWalker(false, nil, visitor)
	w := NewDirWalker(false, nil, visitor)
	w.VisitSelf = false
	w.VisitSelf = false
	return w.Walk(dirPath)
	return w.Walk(dirPath)
}
}


//	A short-hand for `ioutil.WriteFile` using `ModePerm`.
//	A short-hand for `ioutil.WriteFile` using `ModePerm`.
//	Also ensures the target file's directory exists.
//	Also ensures the target file's directory exists.
func WriteBinaryFile(filePath string, contents []byte) error {
func WriteBinaryFile(filePath string, contents []byte) error {
	EnsureDirExists(filepath.Dir(filePath))
	EnsureDirExists(filepath.Dir(filePath))
	return ioutil.WriteFile(filePath, contents, ModePerm)
	return ioutil.WriteFile(filePath, contents, ModePerm)
}
}


//	A short-hand for `ioutil.WriteFile`, using `ModePerm`.
//	A short-hand for `ioutil.WriteFile`, using `ModePerm`.
//	Also ensures the target file's directory exists.
//	Also ensures the target file's directory exists.
func WriteTextFile(filePath, contents string) error {
func WriteTextFile(filePath, contents string) error {
	return WriteBinaryFile(filePath, []byte(contents))
	return WriteBinaryFile(filePath, []byte(contents))
}
}


func watchRunHandler(dirPath string, namePattern ustr.Pattern, handler WatcherHandler) []error {
func watchRunHandler(dirPath string, namePattern ustr.Pattern, handler WatcherHandler) []error {
	vis := func(fullPath string) (keepWalking bool) {
	vis := func(fullPath string) (keepWalking bool) {
		keepWalking = true
		keepWalking = true
		if namePattern.IsMatch(filepath.Base(fullPath)) {
		if namePattern.IsMatch(filepath.Base(fullPath)) {
			handler(fullPath)
			handler(fullPath)
		}
		}
		return
		return
	}
	}
	w := NewDirWalker(false, vis, vis)
	w := NewDirWalker(false, vis, vis)
	w.VisitSelf = false
	w.VisitSelf = false
	w.VisitDirsFirst = true
	w.VisitDirsFirst = true
	return w.Walk(dirPath)
	return w.Walk(dirPath)
}
}
