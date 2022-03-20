package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"
	"onepiece/graph/generated"
	"onepiece/graph/model"
)

func (r *mutationResolver) CreatePirate(ctx context.Context, input model.NewPirate) (*model.Pirate, error) {
	pirate := &model.Pirate{
		Name:   input.Name,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Bounty: input.Bounty,
		Crew:   input.Crew,
		//User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.pirates = append(r.pirates, pirate)
	return pirate, nil
}

func (r *queryResolver) Pirates(ctx context.Context) ([]*model.Pirate, error) {
	return r.pirates, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
