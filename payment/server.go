package payment

import (
	"context"
	"fmt"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type grpcServer struct {
	service       Service
	tollTaxClient *tolltax.Client
}

func (g grpcServer) GenerateRFID(ctx context.Context, request *tolltaxpb.GenerateRFIDRequest) (*tolltaxpb.GenerateRFIDResponse, error) {
	panic("implement me")
}

func (g grpcServer) ValidateRFID(ctx context.Context, request *tolltaxpb.ValidateRFIDRequest) (*tolltaxpb.ValidateRFIDResponse, error) {
	panic("implement me")
}

func (g grpcServer) CalculateDeductibleAmount(ctx context.Context, request *tolltaxpb.CalculateAmountRequest) (*tolltaxpb.CalculateAmountResponse, error) {
	panic("implement me")
}

func (g grpcServer) GetAllOwners(ctx context.Context, request *tolltaxpb.GetAllOwnersRequest) (*tolltaxpb.GetAllOwnersResponse, error) {
	panic("implement me")
}

func (g grpcServer) GetVehicleOwnerDetails(ctx context.Context, request *tolltaxpb.VehicleOwnerDetailsRequest) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {
	panic("implement me")
}

func (g grpcServer) GetTollBoothDetails(ctx context.Context, request *tolltaxpb.TollBoothDetailsRequest) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {
	panic("implement me")
}

func (g grpcServer) CreateNewOwner(ctx context.Context, request *tolltaxpb.CreateNewOwnerRequest) (*tolltaxpb.CreateNewOwnerResponse, error) {
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
		return err
	}
	serv := grpc.NewServer()
	tolltaxpb.RegisterTollTaxServiceServer(serv, &grpcServer{s, tolltaxClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}
