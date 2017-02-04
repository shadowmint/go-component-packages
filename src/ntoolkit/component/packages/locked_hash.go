package packages

import (
	"sync"
	"ntoolkit/component"
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

func (l *lockedHash) Add(path string, value *component.ObjectTemplate) {
	l.lock.Lock()
	l.data[path] = value
	l.lock.Unlock()
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