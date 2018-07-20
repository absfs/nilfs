package nilfs_test

import (
	"io"
	"testing"

	"github.com/absfs/absfs"
	"github.com/absfs/nilfs"
)

func TestNilFs(t *testing.T) {

	var testfs absfs.SymlinkFileSystem

	t.Run("NewFs", func(t *testing.T) {
		fs, err := nilfs.NewFs()
		if err != nil {
			t.Fatal(err)
		}

		testfs = fs
	})

	if testfs.Separator() != 0 {
		t.Fatal("nilfs should not define a Separator")
	}

	if testfs.ListSeparator() != 0 {
		t.Fatal("nilfs should not define a ListSeparator")
	}

	var err error
	name := "/foo.bar"
	err = testfs.Chdir(name)
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}

	dir, err := testfs.Getwd()
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}
	if dir != "" {
		t.Error("nilfs should not return any values")
	}

	dir = testfs.TempDir()
	if dir != "" {
		t.Error("nilfs should not return any values")
	}

	f, err := testfs.Open(name)
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}

	dir = f.Name()
	if dir != "" {
		t.Error("nilfs file handle should have an empty string for name")
	}

	p := []byte("The quick brown fox, jumped over the lazy dog.")

	n, err := f.Read(p)
	if err != nil && err != io.EOF {
		t.Error("nilfs should only return <nil>, or io.EOF errors")
	}
	if n != 0 {
		t.Error("nilfs should not return any values")
	}

	data := make([]byte, 512)
	n, err = f.ReadAt(data, 42)
	if err != nil && err != io.EOF {
		t.Error("nilfs should only return <nil>, or io.EOF errors")
	}
	if n != 0 {
		t.Error("nilfs should not return any values")
	}

	n, err = f.Write(p)
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}
	if n != 0 {
		t.Error("nilfs should not return any values")
	}

	n, err = f.WriteAt(p, 42)
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}
	if n != 0 {
		t.Error("nilfs should not return any values")
	}

	err = f.Close()
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}

	ret, err := f.Seek(0, io.SeekStart)
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}
	if ret != 0 {
		t.Error("nilfs should not return any values")
	}

	info, err := f.Stat()
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}
	if info != nil {
		t.Error("nilfs should not return any values")
	}

	err = f.Sync()
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}

	infos, err := f.Readdir(0)
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}
	if infos != nil {
		t.Error("nilfs should not return any values")
	}

	list, err := f.Readdirnames(0)
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}
	if list != nil {
		t.Error("nilfs should not return any values")
	}

	err = f.Truncate(0)
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}

	n, err = f.WriteString("foo")
	if err != nil {
		t.Error("nilfs should only return <nil> errors")
	}
	if n != 0 {
		t.Error("nilfs should not return any values")
	}

}
