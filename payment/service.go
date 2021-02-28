package payment

import (
	"context"
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment/pb/paymentpb"
)

type Service interface {
	ExecuteTransaction(ctx context.Context, request *paymentpb.ExecuteTRequest) (*paymentpb.ExecuteTResponse, error)
	GetAccountDetails(ctx context.Context, request *paymentpb.GetAccountDetailsRequest) (*paymentpb.GetAccountDetailsResponse, error)
	GetTransactionHistory(ctx context.Context, request *paymentpb.GetTransactionHistoryRequest) (*paymentpb.GetTransactionHistoryResponse, error)
}

type paymentService struct {
	repository Repository
}

func (p paymentService) ExecuteTransaction(ctx context.Context, request *paymentpb.ExecuteTRequest) (*paymentpb.ExecuteTResponse, error) {
	panic("implement me")
}

func (p paymentService) GetAccountDetails(ctx context.Context, request *paymentpb.GetAccountDetailsRequest) (*paymentpb.GetAccountDetailsResponse, error) {
	panic("implement me")
}

func (p paymentService) GetTransactionHistory(ctx context.Context, request *paymentpb.GetTransactionHistoryRequest) (*paymentpb.GetTransactionHistoryResponse, error) {
	panic("implement me")
}

func NewService(r Repository) Service {
	return &paymentService{}
}
