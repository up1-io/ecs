package ecs

// Component is an interface that defines the methods that a component must implement.
type Component interface {
	Name() string
}
