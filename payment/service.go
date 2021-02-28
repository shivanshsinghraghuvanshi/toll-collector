package payment

type Service interface {
}

type paymentService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &paymentService{}
}
