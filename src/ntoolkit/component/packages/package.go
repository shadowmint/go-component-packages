package packages

import (
	"github.com/spf13/afero"
	"ntoolkit/component"
	"ntoolkit/errors"
	"fmt"
)

const DEFAULT_WORKER_COUNT = 10

type Config struct {
	Fs      afero.Fs
	Workers int
	Factory *component.ObjectFactory
}

type Package struct {
	Config    Config
	Templates map[string]*component.ObjectTemplate
}

// New returns a new package instance
func New(config ...Config) *Package {
	rtn := &Package{
		Templates: make(map[string]*component.ObjectTemplate),
	}
	if len(config) > 0 {
		rtn.Config = config[0]
	}
	packageConfigDefaults(&rtn.Config)
	return rtn
}

func packageConfigDefaults(config *Config) {
	if config.Workers <= 0 {
		config.Workers = DEFAULT_WORKER_COUNT
	}
	if config.Fs == nil {
		config.Fs = afero.NewOsFs()
	}
	if config.Factory == nil {
		config.Factory = component.NewObjectFactory()
	}
}

// Load a workspace path.
func (pack *Package) Load(path string) error {
	loader := NewPackageLoader(pack.Config.Fs, pack.Config.Workers)
	if err := loader.Load(path); err != nil {
		return err
	}
	values := loader.Data()
	for key := range values {
		if _, ok := pack.Templates[key]; ok {
			return errors.Fail(ErrDuplicateName{}, nil, fmt.Sprintf("Duplicate name %s in workspace %s", key, path))
		} else {
			pack.Templates[key] = values[key]
		}
	}
	return nil
}

// Spawn a new instance of the given template by name, if it exists.
func (pack *Package) Spawn(templateName string) (*component.Object, error) {
	if template, ok := pack.Templates[templateName]; ok {
		obj, err := pack.Config.Factory.Deserialize(template)
		if err != nil {
			return nil, errors.Fail(ErrBadTemplate{}, err, fmt.Sprintf("Failed to thraw template: %s", templateName))
		}
		return obj, nil
	}
	return nil, errors.Fail(ErrBadTemplate{}, nil, fmt.Sprintf("No such template: %s", templateName))
}