package packages_test

import (
	"reflect"
	"ntoolkit/component"
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
name: Room 1

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