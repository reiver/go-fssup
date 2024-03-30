package fssup

import (
	"fmt"
	"io/fs"

	"github.com/reiver/go-path"
)

func Sup(fsys fs.FS, dir string) (fs.FS, error) {

	dir = path.Canonical(dir)

	if !fs.ValidPath(dir) {
		return nil, &fs.PathError{Op: "sup", Path: dir, Err: fmt.Errorf("%q is an invalid path", dir)}
	}

	if "." == dir {
		return fsys, nil
	}

	return supFS{fsys, dir}, nil
}
