package tolltax

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
	"log"
	"strings"
)

type Repository interface {
	Close()
	GenerateRFID(ctx context.Context, rfid string, ownerid, carid int64) (string, error)
	ValidateRFID(ctx context.Context, rfid string, carid int64) (bool, error)
	CalculateDeductibleAmount(ctx context.Context, cartype string) (int32, error)
	GetAllOwners(ctx context.Context) ([]*tolltaxpb.Owner, error)
	CreateNewOwner(ctx context.Context, o *tolltaxpb.CreateNewOwnerRequest) (*tolltaxpb.CreateNewOwnerResponse, error)
	GetTollBoothDetails(ctx context.Context, tollboothid int64, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error)
	GetVehicleOwnerDetails(ctx context.Context, rfid string, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error)
}

type postgresRepository struct {
	db *sql.DB
}

func (r *postgresRepository) GetTollBoothDetails(ctx context.Context, tollboothid int64, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT name, accountnumber from tollbooth where tollboothid=$1", tollboothid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	v := &tolltaxpb.VehicleOwnerDetailsResponse{}
	for rows.Next() {
		err := rows.Scan(&v.Name, &v.Accountnumber)
		if err != nil {
			log.Fatal("Error while fetching the amount")
			return nil, err
		}
	}
	v.Action = action
	return v, nil
}

func (r *postgresRepository) GetVehicleOwnerDetails(ctx context.Context, rfid string, action tolltaxpb.ACTION) (*tolltaxpb.VehicleOwnerDetailsResponse, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT name, accountnumber from owner where ownerid = (SELECT ownerid from netc where rfid=$1)", rfid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	v := &tolltaxpb.VehicleOwnerDetailsResponse{}
	for rows.Next() {
		err := rows.Scan(&v.Name, &v.Accountnumber)
		if err != nil {
			log.Fatal("Error while fetching the amount")
			return nil, err
		}
	}
	v.Action = action
	return v, nil
}

func (r *postgresRepository) CreateNewOwner(ctx context.Context, o *tolltaxpb.CreateNewOwnerRequest) (*tolltaxpb.CreateNewOwnerResponse, error) {
	panic("implement me")
}

func (r *postgresRepository) GenerateRFID(ctx context.Context, rfid string, ownerid, carid int64) (string, error) {
	_, err := r.db.ExecContext(ctx, "INSERT INTO netc(ownerid,carid,rfid) VALUES($1,$2,$3)", ownerid, carid, rfid)
	if err != nil {
		return "", err
	}
	return rfid, nil
}

func (r *postgresRepository) ValidateRFID(ctx context.Context, rfid string, carid int64) (bool, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT COUNT(*) from netc where carid = $1 and rfid =$2", carid, rfid)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	count := checkCount(rows)

	log.Printf("count is %v\n", count)
	if count == 0 {
		return false, nil
	} else {
		return true, err
	}
}

func (r *postgresRepository) CalculateDeductibleAmount(ctx context.Context, carType string) (int32, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT amount from deductible where cartype=$1", strings.ToUpper(carType))
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	amount := 0
	for rows.Next() {
		err := rows.Scan(&amount)
		if err != nil {
			log.Fatal("Error while fetching the amount")
			return 0, err
		}
	}
	return int32(amount), nil
}

func (r *postgresRepository) GetAllOwners(ctx context.Context) ([]*tolltaxpb.Owner, error) {

	rows, err := r.db.QueryContext(ctx, "SELECT * FROM owner")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	owners := []*tolltaxpb.Owner{}
	for rows.Next() {
		o := &tolltaxpb.Owner{}
		if err = rows.Scan(&o.Ownerid, &o.Accountnumber, &o.Name); err == nil {
			owners = append(owners, o)
		}
	}
	return owners, err
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
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

func checkCount(rows *sql.Rows) (count int) {
	for rows.Next() {
		err := rows.Scan(&count)
		if err != nil {
			log.Fatal("Some issue while scanning the count")
			return 0
		}
	}
	return count
}
