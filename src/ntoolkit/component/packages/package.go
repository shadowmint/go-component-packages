package packages

import (
	"github.com/spf13/afero"
	"ntoolkit/threadpool"
	"ntoolkit/errors"
	"ntoolkit/iter"
	"regexp"
	"strings"
	"ntoolkit/component"
)

type Package struct {
	fs   afero.Fs
	pool *threadpool.ThreadPool
	data *lockedHash
}

// New returns a new package
func New(fs afero.Fs, workerCount int) *Package {
	rtn := &Package{
		fs: fs,
		pool: threadpool.New(),
		data: newLockedHash()}
	rtn.pool.MaxThreads = workerCount
	return rtn
}

// Load a workspace and parse all objects in it, bind them to the runtime.
func (p *Package) Load(workspacePath string) error {
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
		return errors.Fail(ErrLoadFailed{Errors: errs.All()}, nil, "Failed to load some objects")
	} else {
		p.data.Sync(data)
	}
	return nil
}

// Background load a template
func (p *Package) DeferLoadTemplate(workspacePath string, path string, data *lockedHash, errors *errorList) {
	p.pool.Run(func() {
		raw, err := afero.ReadFile(p.fs, path)
		if err != nil {
			errors.Add(err)
			return
		}

		template, err := convertYamlToTemplate(string(raw))
		if (err != nil) {
			errors.Add(err)
			return
		}

		fixed := strings.Replace(path, "\\", "/", -1)
		assetPath := strings.TrimPrefix(fixed, workspacePath)
		data.Add(assetPath, template)
	})
}

// Return all keys
func (p *Package) Keys() []string {
	p.data.lock.Lock()
	keys := make([]string, 0)
	for key := range p.data.data {
		keys = append(keys, key)
	}
	p.data.lock.Unlock()
	return keys
}

// Return a yield iterator that matches a given regex
func (p *Package) Find(pattern string) ([]*component.ObjectTemplate, error) {
	expr, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	p.data.lock.Lock()
	values := make([]*component.ObjectTemplate, 0)
	for key := range p.data.data {
		if expr.MatchString(key) {
			values = append(values, p.data.data[key])
		}
	}

	p.data.lock.Unlock()
	return values, nil
}
