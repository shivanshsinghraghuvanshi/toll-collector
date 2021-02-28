package main

import (
	"context"
	"errors"
	"github.com/shivanshsinghraghuvanshi/toll-collector/graphql/graph/model"
	"log"
	"strconv"
	"time"
)

var (
	ErrInvalidParameter = errors.New("invalid parameter")
)

type mutationResolver struct {
	server *Server
}

func (m mutationResolver) CreateTollTax(ctx context.Context, input *model.NewTollTax) (bool, error) {
	panic("implement me")
}

func (m mutationResolver) ValidateRfid(ctx context.Context, input model.ValidateRfid) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	r, err := m.server.tolltaxClient.ValidateRFID(ctx, input.Rfid, int64(input.Carid))
	if err != nil {
		log.Println(err)
		return false, err
	}
	return r.Ok, nil
}

func (m mutationResolver) CreateOwner(ctx context.Context, input model.NewOwner) (string, error) {
	panic("implement me")
}

func (m mutationResolver) CreateCar(ctx context.Context, input *model.NewCar) (string, error) {
	panic("implement me")
}

func (m mutationResolver) CreateTollBooth(ctx context.Context, input *model.NewTollBooth) (string, error) {
	panic("implement me")
}

func (m mutationResolver) CreateRfid(ctx context.Context, input *model.NewRfid) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	oId, _ := strconv.ParseInt(input.Ownerid, 10, 64)
	cID, _ := strconv.ParseInt(input.Carid, 10, 64)
	r, err := m.server.tolltaxClient.GenerateRFID(ctx, oId, cID)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return r.Rfid, err
}
