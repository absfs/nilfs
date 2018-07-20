package nilfs

import (
	"os"
	"time"

	"github.com/absfs/absfs"
)

type FileSystem struct{}

func NewFs() (*FileSystem, error) {
	return &FileSystem{}, nil
}

func (*FileSystem) Separator() uint8 {
	return 0
}

func (*FileSystem) ListSeparator() uint8 {
	return 0
}

func (*FileSystem) Chdir(name string) error {
	return nil
}

func (*FileSystem) Getwd() (dir string, err error) {
	return "", nil
}

func (*FileSystem) TempDir() string {
	return ""
}

func (*FileSystem) Open(name string) (absfs.File, error) {
	return File{}, nil
}

func (*FileSystem) Create(name string) (absfs.File, error) {
	return File{}, nil
}

func (*FileSystem) Truncate(name string, size int64) error {
	return nil
}

func (*FileSystem) Mkdir(name string, perm os.FileMode) error {
	return nil
}

func (*FileSystem) MkdirAll(name string, perm os.FileMode) error {
	return nil
}

func (*FileSystem) OpenFile(name string, flag int, perm os.FileMode) (absfs.File, error) {
	return File{}, nil
}

func (*FileSystem) Remove(name string) error {
	return nil
}

func (*FileSystem) RemoveAll(name string) error {
	return nil
}

func (*FileSystem) Stat(name string) (os.FileInfo, error) {
	return nil, nil
}

//Chmod changes the mode of the named file to mode.
func (*FileSystem) Chmod(name string, mode os.FileMode) error {
	return nil
}

//Chtimes changes the access and modification times of the named file
func (*FileSystem) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return nil
}

//Chown changes the owner and group ids of the named file
func (*FileSystem) Chown(name string, uid, gid int) error {
	return nil
}

func (*FileSystem) Lstat(name string) (os.FileInfo, error) {
	return nil, nil
}

func (*FileSystem) Lchown(name string, uid, gid int) error {
	return nil
}

func (*FileSystem) Readlink(name string) (string, error) {
	return "", nil
}

func (*FileSystem) Symlink(oldname, newname string) error {
	return nil
}
