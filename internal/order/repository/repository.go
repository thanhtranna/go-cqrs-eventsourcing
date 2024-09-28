package repository

import (
	"context"

	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/dto"
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/order/models"
	"github.com/thanhtranna/go-cqrs-eventsourcing/pkg/utils"
)

type OrderMongoRepository interface {
	Insert(ctx context.Context, order *models.OrderProjection) (string, error)
	GetByID(ctx context.Context, orderID string) (*models.OrderProjection, error)
	UpdateOrder(ctx context.Context, order *models.OrderProjection) error

	UpdateCancel(ctx context.Context, order *models.OrderProjection) error
	UpdatePayment(ctx context.Context, order *models.OrderProjection) error
	Complete(ctx context.Context, order *models.OrderProjection) error
	UpdateDeliveryAddress(ctx context.Context, order *models.OrderProjection) error
	UpdateSubmit(ctx context.Context, order *models.OrderProjection) error
}

type ElasticOrderRepository interface {
	IndexOrder(ctx context.Context, order *models.OrderProjection) error
	GetByID(ctx context.Context, orderID string) (*models.OrderProjection, error)
	UpdateOrder(ctx context.Context, order *models.OrderProjection) error
	Search(ctx context.Context, text string, pq *utils.Pagination) (*dto.OrderSearchResponseDto, error)
}
