package app

import (
	"context"
)

func NotifyCustomer(ctx context.Context, message string, phone string) (string, error) {
	service := NotificationService{"sms.example.com"}

	confirmation, err := service.SendText(message, phone)
	if err != nil {
		return "", err
	}
	return confirmation, nil
}
