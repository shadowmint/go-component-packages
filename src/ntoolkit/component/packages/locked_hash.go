package packages

import (
	"sync"
	"ntoolkit/component"
	"ntoolkit/errors"
	"fmt"
)

type lockedHash struct {
	lock *sync.Mutex
	data map[string]*component.ObjectTemplate
}

func newLockedHash() *lockedHash {
	return &lockedHash{
		lock: &sync.Mutex{},
		data: make(map[string]*component.ObjectTemplate)}
}

func (l *lockedHash) Add(path string, value *component.ObjectTemplate) error {
	var err error
	l.lock.Lock()
	if _, ok := l.data[value.Name]; ok {
		err = errors.Fail(ErrDuplicateName{}, nil, fmt.Sprintf("Duplicate name %s in object %s", value.Name, path))
	} else {
		l.data[path] = value
	}
	l.lock.Unlock()
	return err
}

func (l *lockedHash) Sync(remote *lockedHash) {
	remote.lock.Lock()
	l.lock.Lock()
	for key := range remote.data {
		l.data[key] = remote.data[key]
	}
	l.lock.Unlock()
	remote.lock.Unlock()
}