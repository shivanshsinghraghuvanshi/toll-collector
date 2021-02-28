package payment

import (
	"context"
	"database/sql"
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment/pb/paymentpb"
	"log"
)
import _ "github.com/lib/pq"

type Repository interface {
	Close()
	ExecuteTransaction(ctx context.Context, request *paymentpb.ExecuteTRequest) (*paymentpb.ExecuteTResponse, error)
	GetAccountDetails(ctx context.Context, acc int64) (*paymentpb.GetAccountDetailsResponse, error)
	GetTransactionHistory(ctx context.Context) (*paymentpb.GetTransactionHistoryResponse, error)
}

type postgresRepository struct {
	db *sql.DB
}

func (r *postgresRepository) ExecuteTransaction(ctx context.Context, request *paymentpb.ExecuteTRequest) (*paymentpb.ExecuteTResponse, error) {
	panic("implement me")
}

func (r *postgresRepository) GetAccountDetails(ctx context.Context, acc int64) (*paymentpb.GetAccountDetailsResponse, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT accountid, accountnumber, name, balance, lastUpdated FROM  where accountnumber=$1", acc)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	o := &paymentpb.GetAccountDetailsResponse{}
	for rows.Next() {
		err := rows.Scan(&o.Accountid, &o.Accountnumber, &o.AccountHolderName, &o.Balance, &o.LastUpdated)
		if err != nil {
			log.Fatal("Error while fetching the amount")
			return nil, err
		}
	}
	return o, err
}

func (r *postgresRepository) GetTransactionHistory(ctx context.Context) (*paymentpb.GetTransactionHistoryResponse, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id,timestamp,debitaccountnumber,creditaccountnumber,amount,remarks from transactionDetails")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	transactions := []*paymentpb.TransactionHistory{}

	for rows.Next() {
		t := &paymentpb.TransactionHistory{}
		if err = rows.Scan(); err == nil {
			transactions = append(transactions, t)
		}
	}
	return &paymentpb.GetTransactionHistoryResponse{TransactionHistory: transactions}, err
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}
