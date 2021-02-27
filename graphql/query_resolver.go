package main

import (
	"context"
	"github.com/shivanshsinghraghuvanshi/toll-collector/graphql/graph/model"
	"log"
	"strconv"
	"time"
)

type queryResolver struct {
	server *Server
}

func (q queryResolver) Cars(ctx context.Context) ([]*model.Car, error) {
	panic("implement me")
}

func (q queryResolver) Owners(ctx context.Context) ([]*model.Owner, error) {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Get single
	//if id != nil {
	//	r, err := r.server.tolltaxClient.GetOwners(ctx, *id)
	//	if err != nil {
	//		log.Println(err)
	//		return nil, err
	//	}
	//	return []*Account{{
	//		ID:   r.ID,
	//		Name: r.Name,
	//	}}, nil
	//}

	//skip, take := uint64(0), uint64(0)
	//if pagination != nil {
	//	skip, take = pagination.bounds()
	//}

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

	return o, nil
}

func (q queryResolver) Tollbooths(ctx context.Context) ([]*model.Tollbooth, error) {
	panic("implement me")
}
