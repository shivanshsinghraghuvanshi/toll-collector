package tolltax

import (
	"context"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
	"log"
)

type Service interface {
	GenerateRFID(ctx context.Context, rfid string, ownerid int64, carnumber string) (string, error)
	ValidateRFID(ctx context.Context, rfid string, carnumber string) (bool, error)
	DeductTransaction(ctx context.Context, amount int32, owner *owner) bool
	CreditTransaction(ctx context.Context, amount int32, tollbooth *tollbooth) bool
	CalculateDeductibleAmount(ctx context.Context, cartype string) (int32, error)
	GetAllOwners(ctx context.Context) ([]*tolltaxpb.Owner, error)
	GetTollBoothDetails(ctx context.Context, tollboothid int64, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error)
	GetVehicleOwnerDetails(ctx context.Context, rfid string, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error)
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

func (t *tolltaxService) GetTollBoothDetails(ctx context.Context, tollboothid int64, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {
	r, err := t.repository.GetTollBoothDetails(ctx, tollboothid, action)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (t *tolltaxService) GetVehicleOwnerDetails(ctx context.Context, rfid string, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {
	log.Printf("Action Value is %v\n", action)
	r, err := t.repository.GetVehicleOwnerDetails(ctx, rfid, action)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (t *tolltaxService) GetAllOwners(ctx context.Context) ([]*tolltaxpb.Owner, error) {
	r, err := t.repository.GetAllOwners(ctx)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (t *tolltaxService) GenerateRFID(ctx context.Context, rfid string, ownerid int64, carnumber string) (string, error) {
	r, err := t.repository.GenerateRFID(ctx, rfid, ownerid, carnumber)
	if err != nil {
		return "", err
	}
	return r, nil
}

func (t *tolltaxService) ValidateRFID(ctx context.Context, rfid string, carnumber string) (bool, error) {
	r, err := t.repository.ValidateRFID(ctx, rfid, carnumber)
	if err != nil {
		return false, err
	}
	return r, nil
}

func (t *tolltaxService) DeductTransaction(ctx context.Context, amount int32, owner *owner) bool {
	panic("implement me")
}

func (t *tolltaxService) CreditTransaction(ctx context.Context, amount int32, tollbooth *tollbooth) bool {
	panic("implement me")
}

func (t *tolltaxService) CalculateDeductibleAmount(ctx context.Context, cartype string) (int32, error) {
	r, err := t.repository.CalculateDeductibleAmount(ctx, cartype)
	if err != nil {
		return 0, err
	}
	return r, nil
}

func NewService(r Repository) Service {
	return &tolltaxService{r}
}
