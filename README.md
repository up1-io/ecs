# Game ECS Library for Go

![Version](https://img.shields.io/badge/Version-Prototype-red)
[![GoDoc](https://godoc.org/github.com/up1-io/ecs?status.svg)](https://godoc.org/github.com/up1-io/ecs)

> Note: This library is currently in prototype stage. It is not recommended to use this library in production.

Welcome to the Game ECS (Entity-Component-System) library for Go! This library provides a straightforward and efficient
ECS implementation to help you simplify game development. With our ECS framework, you can manage entities, components,
and systems with ease, allowing you to focus on creating engaging gameplay.

## Features

- **Basic ECS**: A simple and generic ECS framework in Go.
- **Versatile**: Easily adaptable to various game development needs.
- **Open Source**: Licensed under [LICENSE](LICENSE) for your free use and modification.

## Getting Started

1. **Installation**: Get the package using `go get`:

```bash
   go get github.com/up1-io/ecs
```

2. **Usage**: Import the package in your code and start using it:

```go
   import "github.com/up1-io/ecs"
```

3. **Examples**:

```go
package example

import (
	"fmt"
	"github.com/up1-io/ecs"
)

type PositionComponent struct {
	X, Y float64
}

func (p PositionComponent) Name() string {
	return "PositionComponent"
}

type MovementSystem struct{}

func (s MovementSystem) Name() string {
	return "MovementSystem"
}

func (s MovementSystem) RequiredComponents() []string {
	return []string{"PositionComponent"}
}

func (s MovementSystem) Update(entities []ecs.Entity) {
	for _, entity := range entities {
		component := entity.GetComponent("PositionComponent")

		if position, ok := component.(PositionComponent); ok {
			fmt.Printf("Moving entity %d to X: %f, Y: %f\n", entity.ID, position.X, position.Y)
		}
	}
}

func main() {
	app := ecs.NewApp()

	entity := app.CreateEntity()
	app.AddComponent(entity, PositionComponent{X: 1.0, Y: 2.0})

	app.AddSystem(MovementSystem{})

	app.Update()
}
```

