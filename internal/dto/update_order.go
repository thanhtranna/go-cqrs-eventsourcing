package dto

import "github.com/thanhtranna/go-cqrs-eventsourcing/internal/order/models"

type UpdateShoppingItemsReqDto struct {
	ShopItems []*models.ShopItem `json:"shopItems" bson:"shopItems,omitempty" validate:"required"`
}
