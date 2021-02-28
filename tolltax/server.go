package tolltax

import (
	"context"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"github.com/shivanshsinghraghuvanshi/toll-collector/payment"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type grpcServer struct {
	service       Service
	paymentClient *payment.Client
}

func (g *grpcServer) PayTollTax(ctx context.Context, request *tolltaxpb.PayTollTaxRequest) (*tolltaxpb.PayTollTaxResponse, error) {
	vOd, e1 := g.GetVehicleOwnerDetails(ctx, &tolltaxpb.VehicleOwnerDetailsRequest{
		Rfid:   request.Rfid,
		Action: 0,
	})
	if e1 != nil {
		log.Fatal("Cannot get ID")
		return nil, e1
	}
	tOd, e2 := g.GetVehicleOwnerDetails(ctx, &tolltaxpb.VehicleOwnerDetailsRequest{
		Rfid:   request.Rfid,
		Action: 0,
	})
	if e2 != nil {
		log.Fatal("Cannot get tollbooth Details")
		return nil, e2
	}
	t, err := g.paymentClient.ExecuteTransaction(ctx, vOd.Accountnumber, tOd.Accountnumber, request.Amount, request.Remarks)
	if err != nil {
		log.Fatal("Cannot get tollbooth Details")
		return nil, err
	}
	n := &tolltaxpb.PayTollTaxResponse{Ok: t.Status}
	return n, nil
}

func (g *grpcServer) GetVehicleOwnerDetails(ctx context.Context, request *tolltaxpb.VehicleOwnerDetailsRequest) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {

	r, err := g.service.GetVehicleOwnerDetails(ctx, request.Rfid, request.Action)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (g *grpcServer) GetTollBoothDetails(ctx context.Context, request *tolltaxpb.TollBoothDetailsRequest) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {
	r, err := g.service.GetTollBoothDetails(ctx, request.Tollboothid, request.Action)
	if err != nil {
		return nil, err
	}
	return r, nil
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

func ListenGRPC(s Service, paymentserviceURL string, port int) error {

	paymentClient, err := payment.NewClient(paymentserviceURL)
	if err != nil {
		log.Fatal("Cannot Connect to tolltax service")
		return err
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		paymentClient.Close()
		return err
	}
	serv := grpc.NewServer()
	tolltaxpb.RegisterTollTaxServiceServer(serv, &grpcServer{s, paymentClient})
	reflection.Register(serv)
	return serv.Serve(lis)
}
