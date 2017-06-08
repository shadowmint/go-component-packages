package packages_test

import (
	"reflect"
	"ntoolkit/component"
	"github.com/spf13/afero"
)

type SimpleComponent struct {
}

func (c *SimpleComponent) Type() reflect.Type {
	return reflect.TypeOf(c)
}

func (c *SimpleComponent) New() component.Component {
	return &SimpleComponent{}
}

type SerializedComponent struct {
}

func (c *SerializedComponent) Type() reflect.Type {
	return reflect.TypeOf(c)
}

func (c *SerializedComponent) New() component.Component {
	return &SerializedComponent{}
}

var roomFixture1 string = `
name: Room.1

components:
- type: '*ntoolkit/component/packages_test.SimpleComponent'
- type: '*ntoolkit/component/packages_test.SerializedComponent'
	data:
		nodes:
			key: Hello
			value: World

objects:
	- name: Fixture.Statue
		components:
		- type: '*ntoolkit/component/packages_test.SimpleComponent'
`

var roomFixture2 string = `
name: Room.2

components:
- type: '*ntoolkit/component/packages_test.SimpleComponent'
- type: '*ntoolkit/component/packages_test.SerializedComponent'
	data:
		nodes:
			key: Hello
			value: World
`

var otherFixture1 string = `
name: Other.1

components:
- type: '*ntoolkit/component/packages_test.SimpleComponent'
`

var otherFixture2 string = `
name: Other.2

components:
- type: '*ntoolkit/component/packages_test.SimpleComponent'
`

var otherFixture3 string = `
name: Other.3

components:
- type: '*ntoolkit/component/packages_test.SimpleComponent'
`


func fixture() afero.Fs {
	fs := afero.NewMemMapFs()

	// Workspace
	workspacePath := "/blah/workspace1/"

	// Rooms
	roomPath := workspacePath + "rooms/"
	fs.MkdirAll(roomPath, 0755)
	afero.WriteFile(fs, roomPath + "room1.yaml", []byte(roomFixture1), 0755)

	// Rooms
	othersPath := workspacePath + "others/"
	fs.MkdirAll(othersPath, 0755)
	afero.WriteFile(fs, othersPath + "other1.yaml", []byte(otherFixture1), 0755)
	afero.WriteFile(fs, othersPath + "other2.yaml", []byte(otherFixture2), 0755)

	// Workspace
	workspacePath = "/blah/workspace2/"

	// Rooms
	roomPath = workspacePath + "rooms/"
	fs.MkdirAll(roomPath, 0755)
	afero.WriteFile(fs, roomPath + "room1.yaml", []byte(roomFixture2), 0755)

	// Rooms
	othersPath = workspacePath + "others/"
	fs.MkdirAll(othersPath, 0755)
	afero.WriteFile(fs, othersPath + "other2.yaml", []byte(otherFixture3), 0755)

	return fs
}

func fixtureBad() afero.Fs {
	fs := afero.NewMemMapFs()

	// Workspace
	workspacePath := "/blah/workspace3/"

	// Rooms
	roomPath := workspacePath + "rooms/"
	fs.MkdirAll(roomPath, 0755)
	afero.WriteFile(fs, roomPath + "room1.yaml", []byte(roomFixture1), 0755)
	afero.WriteFile(fs, roomPath + "room2.yaml", []byte(roomFixture1), 0755)

	return fs
}