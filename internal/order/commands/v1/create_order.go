package v1

import (
	"context"
	"errors"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/thanhtranna/go-cqrs-eventsourcing/config"
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/order/aggregate"
	"github.com/thanhtranna/go-cqrs-eventsourcing/pkg/es"
	"github.com/thanhtranna/go-cqrs-eventsourcing/pkg/logger"
)

type CreateOrderCommandHandler interface {
	Handle(ctx context.Context, command *CreateOrderCommand) error
}

type createOrderHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewCreateOrderHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *createOrderHandler {
	return &createOrderHandler{log: log, cfg: cfg, es: es}
}

func (c *createOrderHandler) Handle(ctx context.Context, command *CreateOrderCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createOrderHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order := aggregate.NewOrderAggregateWithID(command.AggregateID)
	err := c.es.Exists(ctx, order.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := order.CreateOrder(ctx, command.ShopItems, command.AccountEmail, command.DeliveryAddress); err != nil {
		return err
	}

	span.LogFields(log.String("order", order.String()))
	return c.es.Save(ctx, order)
}
