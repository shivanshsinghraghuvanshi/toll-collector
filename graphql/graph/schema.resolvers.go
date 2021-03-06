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

func (r *mutationResolver) ValidateRfid(ctx context.Context, input model.ValidateRfid) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTollTax(ctx context.Context, input *model.NewTollTax) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PayTollTax(ctx context.Context, input *model.PayTollTax) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Owners(ctx context.Context) ([]*model.Owner, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Netc(ctx context.Context) ([]*model.Netc, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tollbooths(ctx context.Context) ([]*model.Tollbooth, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Deductions(ctx context.Context, cartype *string) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Carowners(ctx context.Context, ownerid *int) (*model.Relation, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Ownerinfo(ctx context.Context, rfid *string, action *int) (*model.OwnerInfoDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Tollboothinfo(ctx context.Context, id *int, action *int) (*model.TollBoothInfoDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TransactionHistory(ctx context.Context, startDate *string, endDate *string) ([]*model.TransactionHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AccountDetails(ctx context.Context, accountNumber int) (*model.AccountDetails, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GenerateMatrix(ctx context.Context, num int) (*model.MatrixResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
