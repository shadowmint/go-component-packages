package packages

import "sync"

type errorList struct {
	lock   *sync.Mutex
	errors []error
}

func newErrorList() *errorList {
	return &errorList{
		lock: &sync.Mutex{},
		errors: make([]error, 0)}
}

func (e *errorList) HasErrors() bool {
	return len(e.errors) > 0
}

func (e *errorList) All() []error {
	return e.errors
}

func (e *errorList) Add(err error) {
	e.lock.Lock()
	e.errors = append(e.errors, err)
	e.lock.Unlock()
}