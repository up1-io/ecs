package ecs

import "testing"

type TestSystem struct {
	systemName         string
	requiredComponents []string
}

func NewTestSystem(systemName string, requiredComponents []string) *TestSystem {
	return &TestSystem{
		systemName:         systemName,
		requiredComponents: requiredComponents,
	}
}

func (s *TestSystem) Name() string {
	return s.systemName
}

func (s *TestSystem) RequiredComponents() []string {
	return s.requiredComponents
}

func (s *TestSystem) Update(entities []Entity) {
	// Do nothing
}

func BenchmarkApp_AddComponent(b *testing.B) {
	app := NewApp()

	entity := app.CreateEntity()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.AddComponent(entity, &TestComponent{})
	}
}

func BenchmarkApp_RemoveComponent(b *testing.B) {
	app := NewApp()

	entity := app.CreateEntity()

	component := NewTestComponent("TestComponent")
	app.AddComponent(entity, component)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.RemoveComponent(entity, "TestComponent")
	}
}

func BenchmarkApp_AddSystem(b *testing.B) {
	app := NewApp()

	system := NewTestSystem("TestSystem", []string{"TestComponent"})

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.AddSystem(system)
	}
}

func BenchmarkApp_Update(b *testing.B) {
	app := NewApp()

	entity := app.CreateEntity()

	component := NewTestComponent("TestComponent")
	app.AddComponent(entity, component)

	system := NewTestSystem("TestSystem", []string{"TestComponent"})
	app.AddSystem(system)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.Update()
	}
}

func BenchmarkApp_EntityHasComponents(b *testing.B) {
	app := NewApp()

	entity := app.CreateEntity()

	component := NewTestComponent("TestComponent")
	app.AddComponent(entity, component)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app.EntityHasComponents(entity, []string{"TestComponent"})
	}
}
