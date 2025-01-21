package gateway

import "github.com/alladiobb/ticket/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
