package payment

import (
	"context"
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment/pb/paymentpb"
	"log"
	"time"
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
	r, err := p.repository.ExecuteTransaction(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (p paymentService) GetAccountDetails(ctx context.Context, request *paymentpb.GetAccountDetailsRequest) (*paymentpb.GetAccountDetailsResponse, error) {
	log.Printf("request account number from payment service %v\n", request.AccountNumber)
	r, err := p.repository.GetAccountDetails(ctx, request.AccountNumber)
	if err != nil {
		log.Fatalf("Error at service level %v", err)
		return nil, err
	}
	log.Printf("payment service r balance %v\n", r.Balance)
	return r, nil
}

func (p paymentService) GetTransactionHistory(ctx context.Context, request *paymentpb.GetTransactionHistoryRequest) (*paymentpb.GetTransactionHistoryResponse, error) {
	r, err := p.repository.GetTransactionHistory(ctx)
	if err != nil {
		return nil, err
	}
	start, _ := time.Parse(time.RFC822, request.StartDate)
	end, _ := time.Parse(time.RFC822, request.EndDate)

	if request.StartDate != "" && request.EndDate != "" {
		tmp := make([]*paymentpb.TransactionHistory, 0)
		for _, e := range r.TransactionHistory {
			t, _ := time.Parse(time.RFC822, e.Timestamp)
			if inTimeSpan(start, end, t) {
				tmp = append(tmp, e)
			}
		}
		return &paymentpb.GetTransactionHistoryResponse{TransactionHistory: tmp}, nil
	} else if request.Filters != nil {
		// TODOS filter later
		return r, nil
	} else {
		return r, nil
	}
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)
}

func NewService(r Repository) Service {
	return &paymentService{r}
}
