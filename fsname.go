package fssup

import (
	"io/fs"

	"github.com/reiver/go-path"
)

// fsName returns the path without the directory prefix.
//
// For example, if `dir` == "once/twice" and `name` == "once/twice/thrice/fource.txt"
// then it returns "thrice/fource.txt".
//
// If `name` does not havethe directory prefix then this returns the fs.ErrNotExist error.
func fsName(dir string, name string) (string, error) {

	name = path.Canonical(name)
	dir = path.RemoveTrailingSeparators(dir)

	if dir == name {
		return "", fs.ErrNotExist
	}

	if !path.Contains(dir, name) {
		return "", fs.ErrNotExist
	}

	return name[len(dir)+len("/"):], nil
}
