package queries

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/thanhtranna/go-cqrs-eventsourcing/config"
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/dto"
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/order/repository"
	"github.com/thanhtranna/go-cqrs-eventsourcing/pkg/es"
	"github.com/thanhtranna/go-cqrs-eventsourcing/pkg/logger"
)

type SearchOrdersQueryHandler interface {
	Handle(ctx context.Context, command *SearchOrdersQuery) (*dto.OrderSearchResponseDto, error)
}

type searchOrdersHandler struct {
	log               logger.Logger
	cfg               *config.Config
	es                es.AggregateStore
	elasticRepository repository.ElasticOrderRepository
}

func NewSearchOrdersHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore, elasticRepository repository.ElasticOrderRepository) *searchOrdersHandler {
	return &searchOrdersHandler{log: log, cfg: cfg, es: es, elasticRepository: elasticRepository}
}

func (s *searchOrdersHandler) Handle(ctx context.Context, command *SearchOrdersQuery) (*dto.OrderSearchResponseDto, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "searchOrdersHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("SearchText", command.SearchText))

	return s.elasticRepository.Search(ctx, command.SearchText, command.Pq)
}
