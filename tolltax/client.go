package tolltax

import (
	"context"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
	"google.golang.org/grpc"
	"log"
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

func (c *Client) ValidateRFID(ctx context.Context, rfid string, carnumber string) (*tolltaxpb.ValidateRFIDResponse, error) {
	log.Printf("%v is rfid and %v is carnumber\n", rfid, carnumber)
	in := &tolltaxpb.ValidateRFIDRequest{
		Rfid:      rfid,
		Carnumber: carnumber,
	}
	r, err := c.service.ValidateRFID(ctx, in)
	if err != nil {
		return nil, err
	}
	log.Printf("return is %v\n", r.Ok)
	return r, nil
}
func (c *Client) GenerateRFID(ctx context.Context, ownerid int64, carnumber string) (*tolltaxpb.GenerateRFIDResponse, error) {

	p := &tolltaxpb.GenerateRFIDRequest{Netc: &tolltaxpb.Netc{
		Fkownerid: ownerid,
		Carnumber: carnumber,
	}}
	r, err := c.service.GenerateRFID(ctx, p)
	if err != nil {
		return nil, err
	}
	return &tolltaxpb.GenerateRFIDResponse{Rfid: r.Rfid, Status: r.Status}, nil

}

func (c *Client) GetAllOwners(ctx context.Context) (*tolltaxpb.GetAllOwnersResponse, error) {
	r, err := c.service.GetAllOwners(ctx, &tolltaxpb.GetAllOwnersRequest{Ref: 1})
	if err != nil {
		return nil, err
	}
	log.Printf("all the owners are %v\n", r.Owner)
	return &tolltaxpb.GetAllOwnersResponse{Owner: r.Owner}, err
}
func (c *Client) CalculateDeductibleAmount(ctx context.Context, cartype string) (*tolltaxpb.CalculateAmountResponse, error) {
	r, err := c.service.CalculateDeductibleAmount(ctx, &tolltaxpb.CalculateAmountRequest{Cartype: cartype})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) GetOwnerDetails(ctx context.Context, rfid string, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {
	r, err := c.service.GetVehicleOwnerDetails(ctx, &tolltaxpb.VehicleOwnerDetailsRequest{
		Rfid:   rfid,
		Action: action,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) GetTollBoothDetails(ctx context.Context, id int64, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {
	r, err := c.service.GetTollBoothDetails(ctx, &tolltaxpb.TollBoothDetailsRequest{
		Tollboothid: id,
		Action:      action,
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Client) PayTollTax(ctx context.Context, rfid string, tollid int64, amount int32, remarks string) (bool, error) {
	in := &tolltaxpb.PayTollTaxRequest{
		Rfid:    rfid,
		Tollid:  tollid,
		Amount:  amount,
		Remarks: remarks,
	}
	r, err := c.service.PayTollTax(ctx, in)
	if err != nil {
		return false, err
	}
	return r.Ok, nil
}

func (c *Client) GenerateMatrix(ctx context.Context, n int) ([]int, int) {
	// TODO optimization if function can be reuse ?
	// TODO optimization if can be run all movement in concurrent ?
	var s int
	if (n*n)%2 == 1 {
		s = ((n * n) / 2) + 1
		return odd(n), s
	} else {
		return even(n), 0
	}
}

func even(n int) []int {
	left, top, right, bottom := 0, 0, n-1, n-1
	s := make([]int, n*n)
	i := 1
	for left < right {
		// move top to right
		for c := left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++
		//move top to bottom
		for r := top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--
		if top == bottom {
			break
		}
		//move bottom to left
		for c := right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--
		// move left to top
		for r := bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++
	}
	s[top*n+left] = i

	return s
}

func odd(n int) []int {
	left, top, right, bottom := 0, 0, n-1, n-1
	s := make([]int, n*n)
	i := 1
	for left < right {
		//start from bottom to left
		for c := right; c >= left; c-- {
			s[bottom*n+c] = i
			i++
		}
		bottom--
		//move left to top
		for r := bottom; r >= top; r-- {
			s[r*n+left] = i
			i++
		}
		left++

		if top == bottom {
			break
		}
		//move top to right
		for c := left; c <= right; c++ {
			s[top*n+c] = i
			i++
		}
		top++
		// move right to bottom
		for r := top; r <= bottom; r++ {
			s[r*n+right] = i
			i++
		}
		right--
	}
	s[top*n+left] = i

	return s
}
