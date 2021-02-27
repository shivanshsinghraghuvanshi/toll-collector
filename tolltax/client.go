package tolltax

import (
	"context"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service tolltaxpb.TollTaxServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := tolltaxpb.NewTollTaxServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) GenerateRFID(ctx context.Context, ownerid, carid int64) (*tolltaxpb.GenerateRFIDResponse, error) {

	p := &tolltaxpb.GenerateRFIDRequest{Netc: &tolltaxpb.Netc{
		Fkownerid: ownerid,
		Fkcarid:   carid,
	}}
	r, err := c.service.GenerateRFID(ctx, p)
	if err != nil {
		return nil, err
	}
	return &tolltaxpb.GenerateRFIDResponse{Rfid: r.Rfid, Status: r.Status}, nil

}

func (c *Client) GetAllOwners(ctx context.Context) (*tolltaxpb.GetAllOwnersResponse, error) {
	r, err := c.service.GetAllOwners(ctx, nil)
	if err != nil {
		return nil, err
	}
	return &tolltaxpb.GetAllOwnersResponse{Owner: r.Owner}, err
}

func (c *Client) GenerateMatrix(ctx context.Context, n int) ([][]int, int) {
	a := make([][]int, n)
	for i, _ := range a {
		a[i] = make([]int, n)
	}
	var s int

	//TODO Core Logic to create Matrix
	if (n*n)%2 == 1 {
		s = ((n * n) / 2) + 1
	} else {
		s = 0
	}
	return a, s
}
