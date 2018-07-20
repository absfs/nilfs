package nilfs

import (
	"os"
)

type File struct{}

func (File) Name() string {
	return ""
}

func (File) Read(p []byte) (int, error) {
	return 0, nil
}

func (File) ReadAt(b []byte, off int64) (n int, err error) {
	return 0, nil
}

func (File) Write(p []byte) (int, error) {
	return 0, nil
}

func (File) WriteAt(b []byte, off int64) (n int, err error) {
	return 0, nil
}

func (File) Close() error {
	return nil
}

func (File) Seek(offset int64, whence int) (ret int64, err error) {
	return 0, nil
}

func (File) Stat() (os.FileInfo, error) {
	return nil, nil
}

func (File) Sync() error {
	return nil
}

func (File) Readdir(n int) ([]os.FileInfo, error) {
	return nil, nil
}

func (File) Readdirnames(n int) ([]string, error) {
	return nil, nil
}

func (File) Truncate(size int64) error {
	return nil
}

func (File) WriteString(s string) (n int, err error) {
	return 0, nil
}
