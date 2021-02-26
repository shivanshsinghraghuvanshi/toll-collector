package tolltax

import (
	"context"
	"fmt"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type grpcServer struct {
	service Service
}

func (g grpcServer) GenerateRFID(ctx context.Context, request *tolltaxpb.GenerateRFIDRequest) (*tolltaxpb.GenerateRFIDResponse, error) {
	panic("implement me")
}

func (g grpcServer) ValidateRFID(ctx context.Context, request *tolltaxpb.ValidateRFIDRequest) (*tolltaxpb.ValidateRFIDResponse, error) {
	panic("implement me")
}

func (g grpcServer) DeductTransaction(ctx context.Context, request *tolltaxpb.DeductRequest) (*tolltaxpb.DeductResponse, error) {
	panic("implement me")
}

func (g grpcServer) CreditTransaction(ctx context.Context, request *tolltaxpb.CreditRequest) (*tolltaxpb.CreditResponse, error) {
	panic("implement me")
}

func (g grpcServer) CalculateDeductibleAmount(ctx context.Context, request *tolltaxpb.CalculateAmountRequest) (*tolltaxpb.CalculateAmountResponse, error) {
	panic("implement me")
}

func ListenGRPC(s Service, port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	serv := grpc.NewServer()
	tolltaxpb.RegisterTollTaxServiceServer(serv, &grpcServer{s})
	reflection.Register(serv)
	return serv.Serve(lis)
}
