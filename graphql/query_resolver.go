package main

import (
	"context"
	"errors"
	"github.com/shivanshsinghraghuvanshi/toll-collector/graphql/graph/model"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
	"log"
	"strconv"
	"time"
)

type queryResolver struct {
	server *Server
}

func (q queryResolver) Netc(ctx context.Context) ([]*model.Netc, error) {
	panic("implement me")
}

func (q queryResolver) TransactionHistory(ctx context.Context, startDate *string, endDate *string) ([]*model.TransactionHistory, error) {
	panic("implement me")
}

func (q queryResolver) AccountDetails(ctx context.Context, accountNumber int) (*model.AccountDetails, error) {
	if accountNumber != 0 {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		r, err := q.server.paymentClient.GetAccountDetails(ctx, int64(accountNumber))

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		accNum := int(r.Accountnumber)
		accId := int(r.Accountid)
		name := r.AccountHolderName
		b := int(r.Balance)
		lu := r.LastUpdated
		return &model.AccountDetails{
			Accountnumber:     &accNum,
			Accountid:         &accId,
			AccountHolderName: &name,
			Balance:           &b,
			LastUpdated:       &lu,
		}, err

	} else {
		return nil, errors.New("arguments not proper")
	}
}

func (q queryResolver) GenerateMatrix(ctx context.Context, num int) (*model.MatrixResponse, error) {
	if num == 0 {
		return nil, errors.New("Input Cannot be 0")
	} else {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()

		a, s := q.server.tolltaxClient.GenerateMatrix(ctx, num)

		return &model.MatrixResponse{
			Special: s,
			Matrix:  a,
		}, nil
	}

}

func (q queryResolver) Ownerinfo(ctx context.Context, rfid *string, action *int) (*model.OwnerInfoDetails, error) {
	if rfid != nil && action != nil {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		r, err := q.server.tolltaxClient.GetOwnerDetails(ctx, *rfid, tolltaxpb.ACTION(*action))

		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		accnum := strconv.Itoa(int(r.Accountnumber))
		act := r.Action.String()
		return &model.OwnerInfoDetails{
			Name:          &r.Name,
			AccountNumber: &accnum,
			Action:        &act,
		}, nil
	} else {
		return nil, errors.New("arguments not proper")
	}
}

func (q queryResolver) Tollboothinfo(ctx context.Context, id *int, action *int) (*model.TollBoothInfoDetails, error) {
	if id != nil {
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		r, err := q.server.tolltaxClient.GetTollBoothDetails(ctx, int64(*id), tolltaxpb.ACTION(*action))
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		accnum := strconv.Itoa(int(r.Accountnumber))
		act := r.Action.String()
		return &model.TollBoothInfoDetails{
			Name:          &r.Name,
			AccountNumber: &accnum,
			Action:        &act,
		}, nil
	} else {
		return nil, errors.New("arguments not proper")
	}
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
