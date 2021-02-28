package main

import (
	"context"
	"errors"
	"github.com/shivanshsinghraghuvanshi/toll-collector/graphql/graph/model"
	"log"
	"strconv"
	"time"
)

type queryResolver struct {
	server *Server
}

func (q queryResolver) Deductions(ctx context.Context, cartype *string) (int, error) {
	if cartype != nil {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		r, err := q.server.tolltaxClient.CalculateDeductibleAmount(ctx, *cartype)

		if err != nil {
			log.Fatal(err)
			return 0, err
		}

		return int(r.Deducible), nil
	} else {
		return 0, errors.New("no id provided as an argument.")
	}
}

func (q queryResolver) Carowners(ctx context.Context, ownerid *int) (*model.Relation, error) {
	panic("implement me")
}

func (q queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	panic("implement me")
}

func (q queryResolver) Owners(ctx context.Context) ([]*model.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	r, err := q.server.tolltaxClient.GetAllOwners(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var o []*model.Owner
	for _, a := range r.Owner {
		oID := strconv.Itoa(int(a.Ownerid))
		acc := strconv.Itoa(int(a.Accountnumber))
		owner := &model.Owner{
			Ownerid:       oID,
			Accountnumber: acc,
			Name:          a.Name,
		}
		o = append(o, owner)
	}
	log.Printf("number of item in owner is : %d\n", len(o))
	return o, nil
}

func (q queryResolver) Tollbooths(ctx context.Context) ([]*model.Tollbooth, error) {
	panic("implement me")
}
