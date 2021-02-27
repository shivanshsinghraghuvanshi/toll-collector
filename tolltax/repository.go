package tolltax

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/shivanshsinghraghuvanshi/toll-collector/tolltax/pb/tolltaxpb"
)

type Repository interface {
	Close()
	GenerateRFID(ctx context.Context, rfid string, ownerid, carid int64) (string, error)
	ValidateRFID(ctx context.Context, rfid string, carid int64) bool
	DeductTransaction(ctx context.Context, amount int32, owner *owner) bool
	CreditTransaction(ctx context.Context, amount int32, tollbooth *tollbooth) bool
	CalculateDeductibleAmount(ctx context.Context, amount int32, carnumber string) int32
	GetAllOwners(ctx context.Context) ([]*tolltaxpb.Owner, error)
}

type postgresRepository struct {
	db *sql.DB
}

func (r *postgresRepository) GenerateRFID(ctx context.Context, rfid string, ownerid, carid int64) (string, error) {
	_, err := r.db.ExecContext(ctx, "INSERT INTO netc(ownerid,carid,rfid) VALUES($1,$2,$3)", ownerid, carid, rfid)
	if err != nil {
		return "", err
	}
	return rfid, nil
}

func (r *postgresRepository) ValidateRFID(ctx context.Context, rfid string, carid int64) bool {
	panic("implement me")
}

func (r *postgresRepository) DeductTransaction(ctx context.Context, amount int32, owner *owner) bool {
	panic("implement me")
}

func (r *postgresRepository) CreditTransaction(ctx context.Context, amount int32, tollbooth *tollbooth) bool {
	panic("implement me")
}

func (r *postgresRepository) CalculateDeductibleAmount(ctx context.Context, amount int32, carnumber string) int32 {
	panic("implement me")
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
