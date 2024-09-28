package mappers

import (
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/dto"
	v1 "github.com/thanhtranna/go-cqrs-eventsourcing/internal/order/events/v1"
)

func UpdateOrderReqDtoToEventData(reqDto dto.UpdateShoppingItemsReqDto) v1.ShoppingCartUpdatedEvent {
	return v1.ShoppingCartUpdatedEvent{
		ShopItems: reqDto.ShopItems,
	}
}
