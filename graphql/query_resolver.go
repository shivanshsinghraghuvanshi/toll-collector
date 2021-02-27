package main

import (
	"context"
	"github.com/shivanshsinghraghuvanshi/toll-collector/graphql/graph/model"
)

type queryResolver struct {
	server *Server
}

func (q queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	panic("implement me")
}

func (q queryResolver) Owners(ctx context.Context) ([]*model.Owner, error) {
	panic("implement me")
}

func (q queryResolver) Tollbooths(ctx context.Context) ([]*model.Tollbooth, error) {
	panic("implement me")
}
