package ecs

import "testing"

func TestApp_AddComponent(t *testing.T) {
	app := NewApp()

	entity := app.CreateEntity()

	component := NewTestComponent("TestComponent")
	app.AddComponent(entity, component)

	if len(app.Components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(app.Components))
	}
}

func TestApp_RemoveComponent(t *testing.T) {
	app := NewApp()

	entity := app.CreateEntity()

	component := NewTestComponent("TestComponent")
	app.AddComponent(entity, component)
	app.RemoveComponent(entity, "TestComponent")

	if len(app.Components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(app.Components))
	}
}

func TestApp_AddSystem(t *testing.T) {
	app := NewApp()

	system := NewTestSystem("TestSystem", []string{"TestComponent"})
	app.AddSystem(system)

	if len(app.Systems) != 1 {
		t.Errorf("Expected 1 system, got %d", len(app.Systems))
	}
}

func TestApp_Update(t *testing.T) {
	app := NewApp()

	entity := app.CreateEntity()

	component := NewTestComponent("TestComponent")
	app.AddComponent(entity, component)

	system := NewTestSystem("TestSystem", []string{"TestComponent"})
	app.AddSystem(system)

	app.Update()

	if len(app.Systems) != 1 {
		t.Errorf("Expected 1 system, got %d", len(app.Systems))
	}
}

func TestApp_EntityHasComponents(t *testing.T) {
	app := NewApp()

	entity := app.CreateEntity()

	component := NewTestComponent("TestComponent")
	app.AddComponent(entity, component)

	if !app.EntityHasComponents(entity, []string{"TestComponent"}) {
		t.Error("Expected entity to have TestComponent")
	}
}

func TestApp_EntityHasComponents_False(t *testing.T) {
	app := NewApp()

	entity := app.CreateEntity()

	component := NewTestComponent("TestComponent")
	app.AddComponent(entity, component)

	if app.EntityHasComponents(entity, []string{"TestComponent2"}) {
		t.Error("Expected entity to not have TestComponent2")
	}
}

func TestApp_EntityHasComponents_Multiple(t *testing.T) {
	app := NewApp()

	entity := app.CreateEntity()

	component1 := NewTestComponent("TestComponent")
	app.AddComponent(entity, component1)

	component2 := NewTestComponent("TestComponent2")
	app.AddComponent(entity, component2)

	if !app.EntityHasComponents(entity, []string{"TestComponent", "TestComponent2"}) {
		t.Error("Expected entity to have TestComponent and TestComponent2")
	}
}

func TestApp_EntityHasComponents_Multiple_False(t *testing.T) {
	app := NewApp()

	entity := app.CreateEntity()

	component1 := NewTestComponent("TestComponent")
	app.AddComponent(entity, component1)

	component2 := NewTestComponent("TestComponent2")
	app.AddComponent(entity, component2)

	if app.EntityHasComponents(entity, []string{"TestComponent", "TestComponent3"}) {
		t.Error("Expected entity to not have TestComponent and TestComponent3")
	}
}

func TestApp_EntityHasComponents_Multiple_Extra(t *testing.T) {
	app := NewApp()

	entity := app.CreateEntity()

	component1 := NewTestComponent("TestComponent")
	app.AddComponent(entity, component1)

	component2 := NewTestComponent("TestComponent2")
	app.AddComponent(entity, component2)

	if !app.EntityHasComponents(entity, []string{"TestComponent"}) {
		t.Error("Expected entity to have TestComponent")
	}
}
