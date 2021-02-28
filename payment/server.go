package payment

import (
	"context"
	"fmt"
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment/pb/paymentpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type grpcServer struct {
	service Service
}

func (g grpcServer) ExecuteTransaction(ctx context.Context, request *paymentpb.ExecuteTRequest) (*paymentpb.ExecuteTResponse, error) {
	r, err := g.service.ExecuteTransaction(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g grpcServer) GetAccountDetails(ctx context.Context, request *paymentpb.GetAccountDetailsRequest) (*paymentpb.GetAccountDetailsResponse, error) {
	r, err := g.service.GetAccountDetails(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g grpcServer) GetTransactionHistory(ctx context.Context, request *paymentpb.GetTransactionHistoryRequest) (*paymentpb.GetTransactionHistoryResponse, error) {
	r, err := g.service.GetTransactionHistory(ctx, request)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	paymentpb.RegisterPaymentServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}
