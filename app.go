package ecs

// App is the main struct that holds all the entities, components, and systems.
type App struct {
	Entities     []Entity
	Components   map[int][]Component
	Systems      []System
	NextEntityID int
}

func NewApp() App {
	return App{
		Entities:     make([]Entity, 0),
		Components:   make(map[int][]Component),
		Systems:      make([]System, 0),
		NextEntityID: 1,
	}
}

// CreateEntity creates a new entity and returns its ID.
func (w *App) CreateEntity() int {
	entity := Entity{ID: w.NextEntityID}
	w.Entities = append(w.Entities, entity)
	w.NextEntityID++
	return entity.ID
}

// AddComponent attaches a component to an entity.
func (w *App) AddComponent(entityID int, component Component) {
	w.Components[entityID] = append(w.Components[entityID], component)
	w.Entities[entityID-1].components = append(w.Entities[entityID-1].components, component)
}

// RemoveComponent removes a component from an entity.
func (w *App) RemoveComponent(entityID int, componentName string) {
	components := w.Components[entityID]
	for i, component := range components {
		if component.Name() == componentName {
			components = append(components[:i], components[i+1:]...)
			break
		}
	}

	w.Components[entityID] = components
}

// AddSystem adds a system to the world.
func (w *App) AddSystem(system System) {
	w.Systems = append(w.Systems, system)
}

// Update runs all systems on entities with the required components.
func (w *App) Update() {
	for _, system := range w.Systems {
		var entities []Entity
		for _, entity := range w.Entities {
			if w.EntityHasComponents(entity.ID, system.RequiredComponents()) {
				entities = append(entities, entity)
			}
		}

		system.Update(entities)
	}
}

// EntityHasComponents checks if an entity has all the required components.
func (w *App) EntityHasComponents(entityID int, requiredComponents []string) bool {
	components, ok := w.Components[entityID]
	if !ok {
		return false
	}

	componentSet := make(map[string]bool)
	for _, c := range components {
		componentSet[c.Name()] = true
	}

	for _, required := range requiredComponents {
		if !componentSet[required] {
			return false
		}
	}

	return true
}
