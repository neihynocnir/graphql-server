package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/neihynocnir/graphql-server/graph/generated"
	"github.com/neihynocnir/graphql-server/graph/model"
)

func (r *queryResolver) Category(ctx context.Context) ([]*model.Categories, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
