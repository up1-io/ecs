package ecs

// System is an interface that defines the methods that a system must implement.
type System interface {
	Name() string
	RequiredComponents() []string
	Update(entities []Entity)
}
