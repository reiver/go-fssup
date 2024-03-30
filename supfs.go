package fssup

import (
	"io/fs"
)

type supFS struct {
	fsys fs.FS
	dir string
}

func (receiver supFS) Open(name string) (fs.File, error) {
	var fsys fs.FS = receiver.fsys
	if nil == fsys {
		return nil, errNilFieSystem
	}

	var dir string = receiver.dir

	var fsname string
	{
		var err error

		fsname, err  = fsName(dir, name)
		if nil != err {
			return nil, err
		}
	}

	return fsys.Open(fsname)
}
