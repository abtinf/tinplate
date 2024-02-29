/*
Package nolsfs provides a no-listing file system.

This is useful to prevent directory listing in an http file server.

When used with http.FileServer, it will 404 for directories.
*/
package nolsfs

import (
	"io/fs"
	"path/filepath"
)

/*
Create a new no-listing file system. The tryFile is the file to check for
if a directory is requested. If the tryFile exists, the directory will be
served. If the tryFile does not exist, the directory will 404. This mimics the behavior of http.FileServer.
*/
func New(f fs.FS, tryFile string) fs.FS {
	return noLSFS{fs: f, tryFile: tryFile}
}

type noLSFS struct {
	fs      fs.FS
	tryFile string
}

func (n noLSFS) Open(name string) (fs.File, error) {
	f, err := n.fs.Open(name)
	if err != nil {
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		if err := f.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}
	if !stat.IsDir() {
		return f, nil
	}

	tryDir, err := n.fs.Open(filepath.Join(name, n.tryFile))
	if err != nil {
		if err := f.Close(); err != nil {
			return nil, err
		}
		return nil, fs.ErrNotExist
	}
	defer func() { _ = tryDir.Close() }()
	tryStat, err := tryDir.Stat()
	if err != nil {
		if err := f.Close(); err != nil {
			return nil, err
		}
		return nil, err
	}
	if tryStat.IsDir() {
		if err := f.Close(); err != nil {
			return nil, err
		}
		return nil, fs.ErrNotExist
	}
	return f, nil
}
