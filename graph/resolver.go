package graph

import "github.com/neihynocnir/graphql-server/graph/model"

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	categories *model.Categories
}
