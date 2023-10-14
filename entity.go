package ecs

// Entity represents an entity in the ECS. It is a collection of components.
type Entity struct {
	ID         int
	components []Component
}

// GetComponent returns a component from the entity. If the component does not exist, nil is returned.
func (e *Entity) GetComponent(target Component) Component {
	for _, component := range e.components {
		if component.Name() == target.Name() {
			return component
		}
	}

	return nil
}
