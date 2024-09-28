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

type CompleteOrderCommandHandler interface {
	Handle(ctx context.Context, command *CompleteOrderCommand) error
}

type completeOrderCommandHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewCompleteOrderCommandHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *completeOrderCommandHandler {
	return &completeOrderCommandHandler{log: log, cfg: cfg, es: es}
}

func (c *completeOrderCommandHandler) Handle(ctx context.Context, command *CompleteOrderCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "completeOrderCommandHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.CompleteOrder(ctx, command.DeliveryTimestamp); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
