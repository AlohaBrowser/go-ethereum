/*
	This solution exists to avoid problems when using Geth under the macOS Catalyst for Intel processors.
	For some reason on this target os.ReadDir does not return the contents of folders.
*/

package geth

import (
	"io/fs"
	"os"
	"time"
)

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>
const char * _Nullable ReadDir(const char *path, const char *className, const char *selectorName) {
    Class dirReaderClass = NSClassFromString([[NSString alloc] initWithUTF8String:className]);
    SEL readDirSelector = NSSelectorFromString([[NSString alloc] initWithUTF8String:selectorName]);
    NSString *pathString = [NSString stringWithUTF8String:path];
    NSString *readResult = [[[dirReaderClass alloc] init] performSelector:readDirSelector withObject:pathString];
    return [readResult cStringUsingEncoding:NSUTF8StringEncoding];
}
*/
import "C"
import (
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// DirEntry struct. Implements os.DirEntry interface
type DirEntry struct {
	name string
}

func (dirEntry DirEntry) Name() string {
	return dirEntry.name
}

func (dirEntry DirEntry) IsDir() bool {
	return false
}

func (dirEntry DirEntry) Type() fs.FileMode {
	return 0
}

func (dirEntry DirEntry) Info() (fs.FileInfo, error) {
	return FileInfo{dirEntry.name}, nil
}

// FileInfo struct. Implements os.FileInfo interface
type FileInfo struct {
	name string
}

func (fileInfo FileInfo) Name() string {
	return fileInfo.name
}

func (fileInfo FileInfo) Size() int64 {
	return 0
}

func (fileInfo FileInfo) Mode() os.FileMode {
	return 0
}

func (fileInfo FileInfo) IsDir() bool {
	return false
}

func (fileInfo FileInfo) ModTime() time.Time {
	return time.Now()
}

func (fileInfo FileInfo) Sys() interface{} {
	return nil
}

// SetReadDir sets custom ReadDir function
func SetReadDir(class, selector string) {
	keystore.ReadDir = func(name string) ([]os.DirEntry, error) {
		readDirResult := C.ReadDir(C.CString(name), C.CString(class), C.CString(selector))
		if readDirResult != nil {
			path := C.GoString(readDirResult)
			var dirEntry DirEntry
			dirEntry.name = path
			return []os.DirEntry{dirEntry}, nil
		}

		return make([]os.DirEntry, 0), nil
	}
}
