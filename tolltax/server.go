package tolltax

import (
	"context"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type grpcServer struct {
	service Service
}

func (g *grpcServer) CalculateDeductibleAmount(ctx context.Context, request *tolltaxpb.CalculateAmountRequest) (*tolltaxpb.CalculateAmountResponse, error) {
	r, err := g.service.CalculateDeductibleAmount(ctx, request.Cartype)
	if err != nil {
		return nil, err
	}
	return &tolltaxpb.CalculateAmountResponse{Deducible: r}, nil
}

func (g *grpcServer) CreateNewOwner(ctx context.Context, request *tolltaxpb.CreateNewOwnerRequest) (*tolltaxpb.CreateNewOwnerResponse, error) {
	panic("implement me")
}

func (g *grpcServer) GetAllOwners(ctx context.Context, request *tolltaxpb.GetAllOwnersRequest) (*tolltaxpb.GetAllOwnersResponse, error) {
	r, err := g.service.GetAllOwners(ctx)
	if err != nil {
		return nil, err
	}
	return &tolltaxpb.GetAllOwnersResponse{Owner: r}, err
}

func (g *grpcServer) GenerateRFID(ctx context.Context, request *tolltaxpb.GenerateRFIDRequest) (*tolltaxpb.GenerateRFIDResponse, error) {
	u, _ := uuid.NewV4()
	_, err := g.service.GenerateRFID(ctx, u.String(), request.Netc.Fkownerid, request.Netc.Fkcarid)
	if err != nil {
		return nil, err
	}
	return &tolltaxpb.GenerateRFIDResponse{
		Status: "OK",
		Rfid:   u.String(),
	}, nil
}

func (g *grpcServer) ValidateRFID(ctx context.Context, request *tolltaxpb.ValidateRFIDRequest) (*tolltaxpb.ValidateRFIDResponse, error) {
	log.Printf("at server inputs are %v %v\n", request.Rfid, request.Carid)
	r, err := g.service.ValidateRFID(ctx, request.Rfid, request.Carid)
	if err != nil {
		return nil, err
	}
	return &tolltaxpb.ValidateRFIDResponse{Ok: r}, nil
}

func (g *grpcServer) DeductTransaction(ctx context.Context, request *tolltaxpb.DeductRequest) (*tolltaxpb.DeductResponse, error) {
	panic("implement me")
}

func (g *grpcServer) CreditTransaction(ctx context.Context, request *tolltaxpb.CreditRequest) (*tolltaxpb.CreditResponse, error) {
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
