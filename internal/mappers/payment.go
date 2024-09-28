package mappers

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/dto"
	"github.com/thanhtranna/go-cqrs-eventsourcing/internal/order/models"
	orderService "github.com/thanhtranna/go-cqrs-eventsourcing/proto/order"
)

func PaymentFromProto(protoPayment *orderService.Payment) dto.Payment {
	return dto.Payment{
		PaymentID: protoPayment.GetID(),
		Timestamp: protoPayment.GetTimestamp().AsTime(),
	}
}

func PaymentResponseFromModel(payment models.Payment) dto.Payment {
	return dto.Payment{
		PaymentID: payment.PaymentID,
		Timestamp: payment.Timestamp,
	}
}

func PaymentToProto(payment dto.Payment) *orderService.Payment {
	return &orderService.Payment{
		ID:        payment.PaymentID,
		Timestamp: timestamppb.New(payment.Timestamp),
	}
}
