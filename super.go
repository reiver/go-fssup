package fssup

import (
	"fmt"
	"io/fs"

	"github.com/reiver/go-path"
)

func Super(fsys fs.FS, dir string) (fs.FS, error) {

	dir = path.Canonical(dir)

	if !fs.ValidPath(dir) {
		return nil, &fs.PathError{Op: "super", Path: dir, Err: fmt.Errorf("%q is an invalid path", dir)}
	}

	if "." == dir {
		return fsys, nil
	}

	return superFS{fsys, dir}, nil
}
