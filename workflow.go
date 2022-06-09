package app

import (
	"go.temporal.io/sdk/workflow"
	"time"
)

func NotifyWhenAvailable(ctx workflow.Context, productId string, phoneNum string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	for {
		var quantity int
		err := workflow.ExecuteActivity(ctx, GetQuantity, productId).Get(ctx, &quantity)
		if err != nil {
			return "", err
		}

		if quantity > 0 {
			break
		}
		workflow.Sleep(ctx, 10*time.Minute)
	}

	text := "Good news, the product you wanted is now available!"

	var messageId string
	err := workflow.ExecuteActivity(ctx, NotifyCustomer, text, phoneNum).Get(ctx, &messageId)
	if err != nil {
		return "", err
	}

	result := "Successfully texted customer in message with ID=" + messageId

	return result, nil
}
