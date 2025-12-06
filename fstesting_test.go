package nilfs_test

import (
	"io"
	"testing"
	"time"

	"github.com/absfs/nilfs"
)

// TestNilFS_Operations verifies that nilfs accepts all operations without error.
// nilfs is a no-op filesystem - it accepts all operations but stores nothing.
func TestNilFS_Operations(t *testing.T) {
	fs, err := nilfs.NewFs()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Filesystem", func(t *testing.T) {
		// Test basic filesystem properties
		if sep := fs.Separator(); sep != 0 {
			t.Errorf("Separator() = %d, want 0", sep)
		}
		if sep := fs.ListSeparator(); sep != 0 {
			t.Errorf("ListSeparator() = %d, want 0", sep)
		}
		if dir := fs.TempDir(); dir != "" {
			t.Errorf("TempDir() = %q, want empty string", dir)
		}

		// Getwd and Chdir should succeed
		if err := fs.Chdir("/some/path"); err != nil {
			t.Errorf("Chdir() error = %v, want nil", err)
		}
		if dir, err := fs.Getwd(); err != nil || dir != "" {
			t.Errorf("Getwd() = (%q, %v), want (\"\", nil)", dir, err)
		}
	})

	t.Run("FileCreation", func(t *testing.T) {
		// Create should succeed
		f, err := fs.Create("test.txt")
		if err != nil {
			t.Fatalf("Create() error = %v, want nil", err)
		}
		if err := f.Close(); err != nil {
			t.Errorf("Close() error = %v, want nil", err)
		}

		// Open should succeed
		f, err = fs.Open("test.txt")
		if err != nil {
			t.Fatalf("Open() error = %v, want nil", err)
		}
		if err := f.Close(); err != nil {
			t.Errorf("Close() error = %v, want nil", err)
		}

		// OpenFile should succeed
		f, err = fs.OpenFile("test.txt", 0, 0644)
		if err != nil {
			t.Fatalf("OpenFile() error = %v, want nil", err)
		}
		if err := f.Close(); err != nil {
			t.Errorf("Close() error = %v, want nil", err)
		}
	})

	t.Run("FileOperations", func(t *testing.T) {
		f, _ := fs.Create("test.txt")

		// Write accepts data but returns 0 (no-op)
		data := []byte("hello world")
		n, err := f.Write(data)
		if err != nil {
			t.Errorf("Write() error = %v, want nil", err)
		}
		if n != 0 {
			t.Errorf("Write() = %d bytes, want 0 (no-op)", n)
		}

		// WriteAt accepts data but returns 0 (no-op)
		n, err = f.WriteAt(data, 0)
		if err != nil {
			t.Errorf("WriteAt() error = %v, want nil", err)
		}
		if n != 0 {
			t.Errorf("WriteAt() = %d bytes, want 0 (no-op)", n)
		}

		// WriteString accepts data but returns 0 (no-op)
		n, err = f.WriteString("test")
		if err != nil {
			t.Errorf("WriteString() error = %v, want nil", err)
		}
		if n != 0 {
			t.Errorf("WriteString() = %d bytes, want 0 (no-op)", n)
		}

		// Read returns EOF immediately (no data stored)
		buf := make([]byte, 10)
		n, err = f.Read(buf)
		if err != io.EOF {
			t.Errorf("Read() error = %v, want io.EOF", err)
		}
		if n != 0 {
			t.Errorf("Read() = %d bytes, want 0", n)
		}

		// ReadAt returns EOF immediately (no data stored)
		n, err = f.ReadAt(buf, 0)
		if err != io.EOF {
			t.Errorf("ReadAt() error = %v, want io.EOF", err)
		}
		if n != 0 {
			t.Errorf("ReadAt() = %d bytes, want 0", n)
		}

		// Seek should succeed
		pos, err := f.Seek(0, 0)
		if err != nil {
			t.Errorf("Seek() error = %v, want nil", err)
		}
		if pos != 0 {
			t.Errorf("Seek() = %d, want 0", pos)
		}

		// Truncate should succeed
		if err := f.Truncate(10); err != nil {
			t.Errorf("Truncate() error = %v, want nil", err)
		}

		// Sync should succeed
		if err := f.Sync(); err != nil {
			t.Errorf("Sync() error = %v, want nil", err)
		}

		// Stat should return nil (no info)
		info, err := f.Stat()
		if err != nil || info != nil {
			t.Errorf("Stat() = (%v, %v), want (nil, nil)", info, err)
		}

		f.Close()
	})

	t.Run("DirectoryOperations", func(t *testing.T) {
		// Mkdir should succeed
		if err := fs.Mkdir("/dir", 0755); err != nil {
			t.Errorf("Mkdir() error = %v, want nil", err)
		}

		// MkdirAll should succeed
		if err := fs.MkdirAll("/path/to/dir", 0755); err != nil {
			t.Errorf("MkdirAll() error = %v, want nil", err)
		}

		// Readdir should return nil (no entries)
		f, _ := fs.Open("/dir")
		entries, err := f.Readdir(-1)
		if err != nil || entries != nil {
			t.Errorf("Readdir() = (%v, %v), want (nil, nil)", entries, err)
		}

		// Readdirnames should return nil (no names)
		names, err := f.Readdirnames(-1)
		if err != nil || names != nil {
			t.Errorf("Readdirnames() = (%v, %v), want (nil, nil)", names, err)
		}
		f.Close()
	})

	t.Run("FileManipulation", func(t *testing.T) {
		// Remove should succeed
		if err := fs.Remove("test.txt"); err != nil {
			t.Errorf("Remove() error = %v, want nil", err)
		}

		// RemoveAll should succeed
		if err := fs.RemoveAll("/path"); err != nil {
			t.Errorf("RemoveAll() error = %v, want nil", err)
		}

		// Rename should succeed
		if err := fs.Rename("old.txt", "new.txt"); err != nil {
			t.Errorf("Rename() error = %v, want nil", err)
		}

		// Truncate should succeed
		if err := fs.Truncate("test.txt", 100); err != nil {
			t.Errorf("Truncate() error = %v, want nil", err)
		}

		// Stat should return nil (no info)
		info, err := fs.Stat("test.txt")
		if err != nil || info != nil {
			t.Errorf("Stat() = (%v, %v), want (nil, nil)", info, err)
		}

		// Lstat should return nil (no info)
		info, err = fs.Lstat("test.txt")
		if err != nil || info != nil {
			t.Errorf("Lstat() = (%v, %v), want (nil, nil)", info, err)
		}
	})

	t.Run("Attributes", func(t *testing.T) {
		// Chmod should succeed
		if err := fs.Chmod("test.txt", 0600); err != nil {
			t.Errorf("Chmod() error = %v, want nil", err)
		}

		// Chown should succeed
		if err := fs.Chown("test.txt", 1000, 1000); err != nil {
			t.Errorf("Chown() error = %v, want nil", err)
		}

		// Lchown should succeed
		if err := fs.Lchown("test.txt", 1000, 1000); err != nil {
			t.Errorf("Lchown() error = %v, want nil", err)
		}

		// Chtimes should succeed
		now := time.Now()
		if err := fs.Chtimes("test.txt", now, now); err != nil {
			t.Errorf("Chtimes() error = %v, want nil", err)
		}
	})

	t.Run("Symlinks", func(t *testing.T) {
		// Symlink should succeed
		if err := fs.Symlink("target", "link"); err != nil {
			t.Errorf("Symlink() error = %v, want nil", err)
		}

		// Readlink should return empty string
		target, err := fs.Readlink("link")
		if err != nil || target != "" {
			t.Errorf("Readlink() = (%q, %v), want (\"\", nil)", target, err)
		}
	})
}
