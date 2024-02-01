package gateway

import "github.com/paulosarmento/microsservice/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
