package ecs

type TestComponent struct {
	name string
}

func NewTestComponent(name string) *TestComponent {
	return &TestComponent{
		name: name,
	}
}

func (c *TestComponent) Name() string {
	return c.name
}
