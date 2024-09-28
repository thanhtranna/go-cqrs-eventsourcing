package v1

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/thanhtranna/go-cqrs-eventsourcing/config"
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/order/aggregate"
	"github.com/thanhtranna/go-cqrs-eventsourcing/pkg/es"
	"github.com/thanhtranna/go-cqrs-eventsourcing/pkg/logger"
)

type ChangeDeliveryAddressCommandHandler interface {
	Handle(ctx context.Context, command *ChangeDeliveryAddressCommand) error
}

type changeDeliveryAddressCmdHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewChangeDeliveryAddressCmdHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *changeDeliveryAddressCmdHandler {
	return &changeDeliveryAddressCmdHandler{log: log, cfg: cfg, es: es}
}

func (c *changeDeliveryAddressCmdHandler) Handle(ctx context.Context, command *ChangeDeliveryAddressCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "changeDeliveryAddressCmdHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.ChangeDeliveryAddress(ctx, command.DeliveryAddress); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
