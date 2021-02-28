package payment

import (
	"context"
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment/pb/paymentpb"
	"log"
	"time"
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

	// -----> set context for transaction
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		// everything went through commit. or else rollback
		err = tx.Commit()
	}()
	// Now lets try to implement the transaction
	// constraint A debit account should have balance greater than amount
	cbalance, e := r.getBalance(ctx, request.CreditAccountNumber)
	dbalance, ex := r.getBalance(ctx, request.DebitAccountNumber)
	if float32(request.Amount) < dbalance && e == nil && ex == nil {
		// should do an entry in transaction table
		_, err = tx.ExecContext(ctx, "INSERT INTO  transactionDetails(timestamp,debitaccountnumber,creditaccountnumber,amount,remarks) values($1,$2,$3,$4,$5)", time.Now(), request.DebitAccountNumber, request.CreditAccountNumber, request.Amount, request.Remarks)
		if err != nil {
			return nil, err
		}

		// if successful update values in accountdetails table
		_, err = tx.ExecContext(ctx, "Update accountdetails SET balance=$1 where accountnumber=$2", cbalance+float32(request.Amount), request.CreditAccountNumber)
		if err != nil {
			return nil, err
		}
		// Sleep Requirement
		time.Sleep(time.Second * 5)
		_, err = tx.ExecContext(ctx, "Update accountdetails SET balance=$1 where accountnumber=$2", dbalance-float32(request.Amount), request.DebitAccountNumber)
		if err != nil {
			return nil, err
		}

	} else {
		return nil, errors.New("Not Enough Balance in Account")
	}
	return &paymentpb.ExecuteTResponse{
		Status:  true,
		Amount:  request.Amount,
		Message: "Successfully executed trasaction",
	}, nil
}

func (r *postgresRepository) GetAccountDetails(ctx context.Context, acc int64) (*paymentpb.GetAccountDetailsResponse, error) {
	log.Printf("from repository account number %v\n", acc)
	rows, err := r.db.QueryContext(ctx, `SELECT accountid, accountnumber, name, balance,
       lastUpdated from accountdetails where accountnumber=$1`, acc)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	o := &paymentpb.GetAccountDetailsResponse{}
	for rows.Next() {
		err := rows.Scan(&o.Accountid, &o.Accountnumber, &o.AccountHolderName, &o.Balance, &o.LastUpdated)
		if err != nil {
			log.Fatal(err)
			log.Fatal("Error while fetching the account details")
			return nil, err
		}
	}
	log.Printf(" repo payment data %v\n", o.Accountnumber)
	return o, err
}

func (r *postgresRepository) getBalance(ctx context.Context, acc int64) (float32, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT balance FROM accountdetails where accountnumber=$1", acc)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var o float32
	for rows.Next() {
		err := rows.Scan(&o)
		if err != nil {
			log.Fatal("Error while fetching the amount")
			return 0, err
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
		log.Println("Error while init new Server instance")
		log.Fatal(err)
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
