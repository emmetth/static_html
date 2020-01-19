// Generic helpers for generating static html sites.
// These helpers default to a static html directory.
package static_html

import (
	"os"
	"os/exec"
	"path/filepath"
)

// help function that appends path to static html directory.
func join_html(path string) string {
	return filepath.Join("html", path)
}

// Copies file into the static html directory.
func Copy(src string, dest string) {
	cmd := exec.Command("cp", src, join_html(dest))
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

// Create file in the static html directory and returns a file handle.
func Create(path string) *os.File {
	f, err := os.Create(join_html(path))
	if err != nil {
		panic(err)
	}
	return f
}

// Deletes all files and sub directories from the static html directory.
func Empty() {
	d, err := os.Open("html")
	if err != nil {
		panic(err)
	}
	defer d.Close()

	names, err := d.Readdirnames(-1)
	if err != nil {
		panic(err)
	}

	for _, name := range names {
		err = os.RemoveAll(join_html(name))
		if err != nil {
			panic(err)
		}
	}
}

// Makes directory in the static html directory
func Mkdir(path string) {
	err := os.Mkdir(join_html(path), 0755)
	if err != nil {
		panic(err)
	}
}
