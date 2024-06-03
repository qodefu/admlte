package dbstore

import (
	"context"
	"goth/internal/store/models"
)

type ClientStoreImpl struct {
	q *models.Queries
}

func NewCliStore(q *models.Queries) *ClientStoreImpl {
	return &ClientStoreImpl{
		q,
	}

}
func (thing ClientStoreImpl) ListClients() []models.Client {
	ret, _ := thing.q.ListClients(context.Background())
	return ret
}
