package packages

import (
	"github.com/spf13/afero"
	"container/list"
	"os"
	"ntoolkit/iter"
	"ntoolkit/errors"
)

type fileIter struct {
	fs     afero.Fs
	items  *list.List
	cursor *list.Element
	err    error
}

func newFileIter(path string, fs afero.Fs) *fileIter {
	rtn := fileIter{fs: fs,	items: list.New()}
	werr := afero.Walk(fs, path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			rtn.err = err
			return err
		}
		if !info.IsDir() {
			rtn.items.PushBack(path)
		}
		return nil
	})
	if werr != nil {
		rtn.err = werr
	}
	return &rtn
}

// Next increments the iterator cursor
func (iterator *fileIter) Next() (interface{}, error) {
	if iterator.err != nil {
		return nil, iterator.err
	}

	if iterator.cursor == nil {
		iterator.cursor = iterator.items.Front()
	} else {
		iterator.cursor = iterator.cursor.Next()
	}

	if iterator.cursor == nil {
		iterator.err = errors.Fail(iter.ErrEndIteration{}, nil, "No more values")
		return nil, iterator.err
	}

	return iterator.cursor.Value, nil
}

