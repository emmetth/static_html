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

// help function that appends path to static html directory.
func join_html(path string) string {
	return filepath.Join(root, path)
}

// Copies file into the static html directory.
func Copy(src string, dest string) {
	cmd := exec.Command("cp", src, join_html(dest))
	check(cmd.Run())
}

// Create file in the static html directory and returns a file handle.
func Create(path string) *os.File {
	f, err := os.Create(join_html(path))
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
		check(os.RemoveAll(join_html(name)))
	}
}

// Makes directory in the static html directory
func Mkdir(path string) {
	check(os.Mkdir(join_html(path), 0755))
}
