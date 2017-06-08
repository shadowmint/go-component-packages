package packages

import (
	"github.com/spf13/afero"
	"ntoolkit/threadpool"
	"ntoolkit/errors"
	"ntoolkit/iter"
	"ntoolkit/component"
	"fmt"
)

type packageLoader struct {
	fs   afero.Fs
	pool *threadpool.ThreadPool
	data *lockedHash
}

// NewPackageLoader returns a new package
func NewPackageLoader(fs afero.Fs, workerCount int) *packageLoader {
	rtn := &packageLoader{
		fs:   fs,
		pool: threadpool.New(),
		data: newLockedHash()}
	rtn.pool.MaxThreads = workerCount
	return rtn
}

// Load a workspace and parse all objects in it, bind them to the runtime.
func (p *packageLoader) Load(workspacePath string) error {
	items := newFileIter(workspacePath, p.fs)
	var err error
	var v interface{}
	data := newLockedHash()
	errs := newErrorList()
	for v, err = items.Next(); err == nil; v, err = items.Next() {
		p.DeferLoadTemplate(workspacePath, v.(string), data, errs)
	}
	p.pool.Wait()
	if !errors.Is(err, iter.ErrEndIteration{}) {
		return err
	}
	if errs.HasErrors() {
		return errors.Fail(ErrLoadFailed{}, errors.Data(errs.All()), "Failed to load some objects")
	} else {
		p.data.Sync(data)
	}
	return nil
}

// Background load a template
func (p *packageLoader) DeferLoadTemplate(workspacePath string, path string, data *lockedHash, errs *errorList) {
	p.pool.Run(func() {
		raw, err := afero.ReadFile(p.fs, path)
		if err != nil {
			errs.Add(err)
			return
		}

		template, err := convertYamlToTemplate(string(raw))
		if err != nil {
			errs.Add(errors.Fail(ErrBadFile{}, err, fmt.Sprintf("Bad file: %s", path)))
			return
		}

		if err := data.Add(template.Name, template); err != nil {
			errs.Add(err)
		}
	})
}

// Return all keys
func (p *packageLoader) Data() map[string]*component.ObjectTemplate {
	p.data.lock.Lock()
	rtn := make(map[string]*component.ObjectTemplate)
	for key := range p.data.data {
		rtn[key] = p.data.data[key]
	}
	p.data.lock.Unlock()
	return rtn
}