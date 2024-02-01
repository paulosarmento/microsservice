package gateway

import "github.com/paulosarmento/microsservice/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
