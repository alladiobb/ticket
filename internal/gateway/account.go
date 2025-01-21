package gateway

import "github.com/alladiobb/ticket/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindByID(id string) (*entity.Account, error)
	// FindByEmail(email string) (*entity.Account, error)
	// FindByName(name string) (*entity.Account, error)
}
