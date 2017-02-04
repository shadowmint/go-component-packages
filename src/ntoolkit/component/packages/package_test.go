package packages_test

import (
	"ntoolkit/assert"
	"testing"
	"github.com/spf13/afero"
	"ntoolkit/component/packages"
	"ntoolkit/component"
)

func TestPackageTemplates(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()
		pack := packages.New(fs, 10)
		err := pack.Load("/blah/workspace1/")
		T.Assert(err == nil)
	})
}

func TestPackageKeys(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()
		pack := packages.New(fs, 10)
		err := pack.Load("/blah/workspace1/")
		T.Assert(err == nil)
		T.Assert(len(pack.Keys()) > 0)
	})
}

func TestPackageFind(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()
		pack := packages.New(fs, 10)
		err := pack.Load("/blah/workspace1/")
		T.Assert(err == nil)

		templates, err := pack.Find("rooms/.*\\.yaml")
		T.Assert(err == nil)
		T.Assert(len(templates) > 0)
	})
}

func TestPackageThaw(T *testing.T) {
	assert.Test(T, func(T *assert.T) {
		fs := fixture()
		pack := packages.New(fs, 10)
		pack.Load("/blah/workspace1/")
		templates, _ := pack.Find("rooms/.*\\.yaml")

		factory := component.NewObjectFactory()
		factory.Register(&SimpleComponent{})
		factory.Register(&SerializedComponent{})

		obj, err := factory.Deserialize(templates[0])
		T.Assert(obj != nil)
		T.Assert(err == nil)
	})
}

func fixture() afero.Fs {
	fs := afero.NewMemMapFs()

	// Workspace
	workspacePath := "/blah/workspace1/"

	// Rooms
	roomPath := workspacePath + "rooms/"
	fs.MkdirAll(roomPath, 0755)
	afero.WriteFile(fs, roomPath + "room1.yaml", []byte(roomFixture1), 0755)

	return fs
}