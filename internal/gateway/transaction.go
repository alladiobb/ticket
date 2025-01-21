package gateway

import "github.com/alladiobb/ticket/internal/entity"

type Trasactiongateway interface {
	Create(transaction *entity.Transaction) error
}
