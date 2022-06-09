package app

import (
	"context"
)

func GetQuantity(ctx context.Context, productId string) (int, error) {
	inventoryService := InventoryService{"inventory.example.com"}

	qty, err := inventoryService.GetQuantityAvailable(productId)
	if err != nil {
		return -1, err
	}

	return qty, nil
}
