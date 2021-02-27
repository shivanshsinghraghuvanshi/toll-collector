package tolltax

import (
	"context"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
)

type Service interface {
	GenerateRFID(ctx context.Context, rfid string, ownerid, carid int64) (string, error)
	ValidateRFID(ctx context.Context, rfid string, carid int64) bool
	DeductTransaction(ctx context.Context, amount int32, owner *owner) bool
	CreditTransaction(ctx context.Context, amount int32, tollbooth *tollbooth) bool
	CalculateDeductibleAmount(ctx context.Context, amount int32, carnumber string) int32
	GetAllOwners(ctx context.Context) ([]*tolltaxpb.Owner, error)
}

type netc struct {
	NETCID    int64  `json:"netcid"`
	RFID      string `json:"rfid"`
	FKOWNERID int64  `json:"fkownerid"`
	FKCARID   int64  `json:"fkcarid"`
}

type owner struct {
	OWNERID       int64  `json:"ownerid"`
	NAME          string `json:"name"`
	ACCOUNTNUMBER int64  `json:"accountnumber"`
}

type tollbooth struct {
	TOLLBOOTHID   int64  `json:"tollboothid"`
	NAME          string `json:"name"`
	ACCOUNTNUMBER int64  `json:"accountnumber"`
}
type tolltaxService struct {
	repository Repository
}

func (t *tolltaxService) GetAllOwners(ctx context.Context) ([]*tolltaxpb.Owner, error) {
	r, err := t.repository.GetAllOwners(ctx)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (t *tolltaxService) GenerateRFID(ctx context.Context, rfid string, ownerid, carid int64) (string, error) {
	r, err := t.repository.GenerateRFID(ctx, rfid, ownerid, carid)
	if err != nil {
		return "", err
	}
	return r, nil
}

func (t *tolltaxService) ValidateRFID(ctx context.Context, rfid string, carid int64) bool {
	panic("implement me")
}

func (t *tolltaxService) DeductTransaction(ctx context.Context, amount int32, owner *owner) bool {
	panic("implement me")
}

func (t *tolltaxService) CreditTransaction(ctx context.Context, amount int32, tollbooth *tollbooth) bool {
	panic("implement me")
}

func (t *tolltaxService) CalculateDeductibleAmount(ctx context.Context, amount int32, carnumber string) int32 {
	panic("implement me")
}

func NewService(r Repository) Service {
	return &tolltaxService{r}
}
