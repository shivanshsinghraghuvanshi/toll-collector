package payment

import (
	"context"
	"fmt"
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment/pb/paymentpb"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type grpcServer struct {
	service       Service
	tollTaxClient *tolltax.Client
}

func (g grpcServer) ExecuteTransaction(ctx context.Context, request *paymentpb.ExecuteTRequest) (*paymentpb.ExecuteTResponse, error) {
	panic("implement me")
}

func (g grpcServer) GetAccountDetails(ctx context.Context, request *paymentpb.GetAccountDetailsRequest) (*paymentpb.GetAccountDetailsResponse, error) {
	panic("implement me")
}

func (g grpcServer) GetTransactionHistory(ctx context.Context, request *paymentpb.GetTransactionHistoryRequest) (*paymentpb.GetTransactionHistoryResponse, error) {
	panic("implement me")
}

func ListenGRPC(s Service, tolltaxserviceurl string, port int) error {

	tolltaxClient, err := tolltax.NewClient(tolltaxserviceurl)
	if err != nil {
		log.Fatal("Cannot Connect to tolltax service")
		return err
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		tolltaxClient.Close()
		return err
	}
	serv := grpc.NewServer()
	paymentpb.RegisterPaymentServiceServer(serv, &grpcServer{s, tolltaxClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}
