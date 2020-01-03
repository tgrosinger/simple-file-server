package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// directoryHidingFileSystem is a modified http.FileSystem to prevent listing
// directories.
type directoryHidingFileSystem struct {
	http.FileSystem
	root string
}

// Open implements the http.FileSystem Open method with a modification to
// prevent listing directories.
func (fs directoryHidingFileSystem) Open(name string) (http.File, error) {
	if filepath.Separator != '/' && strings.ContainsRune(name, filepath.Separator) {
		return nil, errors.New("http: invalid character in file path")
	}
	if strings.Contains(name, "..") {
		return nil, errors.New("http: invalid path")
	}
	fullName := filepath.Join(fs.root, filepath.FromSlash(path.Clean("/"+name)))
	f, err := os.Open(fullName)
	if err != nil {
		return nil, err
	}

	d, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if d.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

func main() {
	root := flag.String("root", "", "Root of the static files to serve")
	flag.Parse()

	if root == nil || *root == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	fmt.Println("Starting file server in", *root)

	fs := directoryHidingFileSystem{
		FileSystem: http.Dir(*root),
		root:       *root,
	}
	fmt.Println(http.ListenAndServe(":80", http.FileServer(fs)))
}
