package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/shivanshsinghraghuvanshi/toll-collector/graphql/graph/generated"
	"github.com/shivanshsinghraghuvanshi/toll-collector/graphql/graph/model"
)

func (r *mutationResolver) CreateOwner(ctx context.Context, input model.NewOwner) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCar(ctx context.Context, input *model.NewCar) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTollBooth(ctx context.Context, input *model.NewTollBooth) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateRfid(ctx context.Context, input *model.NewRfid) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
