package mappers

import (
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/dto"
	v1 "github.com/thanhtranna/go-cqrs-eventsourcing/internal/order/events/v1"
)

func CreateOrderDtoToEventData(createDto dto.CreateOrderReqDto) v1.OrderCreatedEvent {
	return v1.OrderCreatedEvent{
		ShopItems:       createDto.ShopItems,
		AccountEmail:    createDto.AccountEmail,
		DeliveryAddress: createDto.DeliveryAddress,
	}
}
