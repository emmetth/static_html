// Generic helpers for generating static html sites.
// These helpers default to a static html directory.
package static_html

import (
	"os"
	"os/exec"
	"path/filepath"
)

const root = "html"

// panic if err is not nil
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// Copies file into the static html directory.
func Copy(src string, dest string) {
	dest = filepath.Join(root, dest)
	check(os.MkdirAll(filepath.Dir(dest), 0755))
	cmd := exec.Command("cp", src, dest)
	check(cmd.Run())
}

// Create file in the static html directory and returns a file handle.
func Create(path string) *os.File {
	path = filepath.Join(root, path)
	check(os.MkdirAll(filepath.Dir(path), 0755))
	f, err := os.Create(path)
	check(err)
	return f
}

// Deletes all files and sub directories from the static html directory.
func Empty() {
	d, err := os.Open(root)
	check(err)
	defer d.Close()

	names, err := d.Readdirnames(-1)
	check(err)
	for _, name := range names {
		check(os.RemoveAll(filepath.Join(root, name)))
	}
}
