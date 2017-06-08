package packages_test

import (
	"ntoolkit/assert"
	"testing"
	"ntoolkit/component/packages"
)

func TestPackageTemplates(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()
		pack := packages.NewPackageLoader(fs, 10)
		err := pack.Load("/blah/workspace1/")
		T.Assert(err == nil)
	})
}

func TestPackageKeys(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()
		pack := packages.NewPackageLoader(fs, 10)
		err := pack.Load("/blah/workspace1/")
		T.Assert(err == nil)
	})
}

func TestPackageFind(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()
		pack := packages.NewPackageLoader(fs, 10)
		err := pack.Load("/blah/workspace1/")
		T.Assert(err == nil)

		template, ok := pack.Data()["Room.1"]
		T.Assert(ok)
		T.Assert(template.Name == "Room.1")
	})
}

func TestLoadFail(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixtureBad()
		pack := packages.NewPackageLoader(fs, 10)
		err := pack.Load("/blah/workspace1/")
		T.Assert(err != nil)
	})
}
