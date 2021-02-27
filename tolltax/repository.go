package tolltax

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	GenerateRFID(ctx context.Context, rfid string, ownerid, carid int64) (string, error)
	ValidateRFID(ctx context.Context, rfid string, carid int64) bool
	DeductTransaction(ctx context.Context, amount int32, owner *owner) bool
	CreditTransaction(ctx context.Context, amount int32, tollbooth *tollbooth) bool
	CalculateDeductibleAmount(ctx context.Context, amount int32, carnumber string) int32
}

type postgresRepository struct {
	db *sql.DB
}

func (r *postgresRepository) GenerateRFID(ctx context.Context, rfid string, ownerid, carid int64) (string, error) {
	_, err := r.db.ExecContext(ctx, "INSERT INTO NETC(ownerid,carid,rfid) VALUES($1,$2,$3)", ownerid, carid, rfid)
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
