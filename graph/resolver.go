//go:generate go run github.com/99designs/gqlgen
package graph

import (
	"graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver handles the arrays of objects
type Resolver struct {
	todos    []*model.Todo
}
