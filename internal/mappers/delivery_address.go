package mappers

import (
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/dto"
	v1 "github.com/thanhtranna/go-cqrs-eventsourcing/internal/order/events/v1"
)

func ChangeDeliveryAddressReqDtoToEventData(reqDto dto.ChangeDeliveryAddressReqDto) v1.OrderDeliveryAddressChangedEvent {
	return v1.OrderDeliveryAddressChangedEvent{
		DeliveryAddress: reqDto.DeliveryAddress,
	}
}
