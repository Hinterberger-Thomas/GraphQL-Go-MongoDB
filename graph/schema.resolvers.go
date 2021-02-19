package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Salomon-Novachrono/graphQL-test/graph/generated"
	"github.com/Salomon-Novachrono/graphQL-test/graph/model"
)

func (r *mutationResolver) CreateHuman(ctx context.Context, input model.NewHuman) (*model.Human, error) {
	var user model.Human
	user.ID = "hey"
	user.Name = input.Name
	return &user, nil
}

func (r *queryResolver) Humans(ctx context.Context) ([]*model.Human, error) {
	var humans []*model.Human
	human := model.Human{
		ID:   "je",
		Name: "john",
	}
	humans = append(humans, &human)
	return humans, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
