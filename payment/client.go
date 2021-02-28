package payment

import (
	"context"
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment/pb/paymentpb"
	"google.golang.org/grpc"
	"log"
)

type Client struct {
	conn    *grpc.ClientConn
	service paymentpb.PaymentServiceClient
}

func NewClient(url string) (*Client, error) {
	log.Printf("payment client service url %v\n", url)
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := paymentpb.NewPaymentServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) ExecuteTransaction(ctx context.Context, dAccNum, cAccNum int64, amount int32, remarks string) (*paymentpb.ExecuteTResponse, error) {
	log.Printf("Parameters for execute Transaction are %v %v %v %v\n", dAccNum, cAccNum, amount, remarks)
	in := &paymentpb.ExecuteTRequest{
		DebitAccountNumber:  dAccNum,
		CreditAccountNumber: cAccNum,
		Amount:              amount,
		Remarks:             remarks,
	}

	r, err := c.service.ExecuteTransaction(ctx, in)
	if err != nil {
		return nil, err
	}
	log.Printf("return is %v\n", r.Status)
	return r, nil
}

func (c *Client) GetAccountDetails(ctx context.Context, accNum int64) (*paymentpb.GetAccountDetailsResponse, error) {
	in := &paymentpb.GetAccountDetailsRequest{AccountNumber: accNum}
	r, err := c.service.GetAccountDetails(ctx, in)
	if err != nil {
		return nil, err
	}
	log.Printf("return is %v\n", r.Accountid)
	return r, nil
}
func (c *Client) GetTransactionHistory(ctx context.Context, sDate, endDate *string, accNumber *int64, remarks *string, tID *int64, act *paymentpb.PAYMENTACTION) (*paymentpb.GetTransactionHistoryResponse, error) {
	in := &paymentpb.GetTransactionHistoryRequest{
		StartDate: *sDate,
		EndDate:   *endDate,
		Filters: &paymentpb.TransactionFilter{
			AccountNumber: *accNumber,
			Remarks:       *remarks,
			Transactionid: nil,
			Action:        *act,
		},
	}
	r, err := c.service.GetTransactionHistory(ctx, in)
	if err != nil {
		return nil, err
	}
	log.Printf("return is %v\n", len(r.TransactionHistory))
	return r, nil
}
