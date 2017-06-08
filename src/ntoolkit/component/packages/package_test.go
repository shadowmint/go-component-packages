package packages_test

import (
	"ntoolkit/assert"
	"testing"
	"ntoolkit/component/packages"
	"ntoolkit/component"
)

func TestNew(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()
		pack := packages.New(packages.Config{Fs: fs})
		err := pack.Load("/blah/workspace1/")
		T.Assert(err == nil)
	})
}

func TestCombineWorkspaces(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()
		pack := packages.New(packages.Config{Fs: fs})

		err := pack.Load("/blah/workspace1/")
		T.Assert(err == nil)

		err = pack.Load("/blah/workspace2/")
		T.Assert(err == nil)

		T.Assert(pack.Templates["Room.1"] != nil)
		T.Assert(pack.Templates["Room.2"] != nil)
	})
}

func TestPackageThaw(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()

		factory := component.NewObjectFactory()
		factory.Register(&SimpleComponent{})
		factory.Register(&SerializedComponent{})

		pack := packages.New(packages.Config{Fs: fs, Factory: factory})

		pack.Load("/blah/workspace1/")

		obj, err := pack.Spawn("Room.1")
		T.Assert(obj != nil)
		T.Assert(err == nil)
	})
}