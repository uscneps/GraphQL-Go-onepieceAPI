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
		Image:  input.Image,
	}
	r.pirates = append(r.pirates, pirate)
	return pirate, nil
}

func (r *queryResolver) Pirates(ctx context.Context) ([]*model.Pirate, error) {
	return r.pirates, nil
}

func (r *queryResolver) Strawhats(ctx context.Context) ([]*model.Strawhats, error) {
	var strawhatsCrew = []*model.Strawhats{
		{Name: "Monkey D. Luffy", ID: "01", Bounty: "1,500,000,000", Crew: "Strawhats", Image: "https://rb.gy/xz0x9h"},
		{Name: "Roronoa Zoro", ID: "02", Bounty: "320,000,000", Crew: "Strawhats", Image: "https://rb.gy/imkdhx"},
		{Name: "Nami", ID: "03", Bounty: "66,000,000", Crew: "Strawhats", Image: "https://rb.gy/dncyia"},
		{Name: "God Usop", ID: "04", Bounty: "200,000,000", Crew: "Strawhats", Image: "https://rb.gy/xa68o1"},
		{Name: "Vinsmoke Sanji ", ID: "05", Bounty: "330,000,000", Crew: "Strawhats", Image: "https://rb.gy/sofr7o"},
		{Name: "Tony Tony Chopper", ID: "06", Bounty: "100", Crew: "Strawhats", Image: "https://rb.gy/oh2nmv"},
		{Name: "Nico Robin", ID: "07", Bounty: "130,000,000", Crew: "Strawhats", Image: "https://rb.gy/5puvys"},
		{Name: "Brook", ID: "08", Bounty: "83,000,000", Crew: "Strawhats", Image: "https://rb.gy/5kelik"},
		{Name: "Jimbei", ID: "09", Bounty: "83,000,000", Crew: "Strawhats", Image: "https://rb.gy/mafbmm"},
		{Name: "Neferutari Bibi", ID: "11", Bounty: "0", Crew: "Strawhats", Image: "https://rb.gy/zddixu"},
	}
	return strawhatsCrew, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
