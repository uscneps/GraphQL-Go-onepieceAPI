package graph

import "onepiece/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen generate
type Resolver struct{
	pirates []*model.Pirate
	strawhats []*model.Strawhats
}
