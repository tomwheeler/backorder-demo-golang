package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	"wfe-go-example/app"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "product-availability-workflow",
		TaskQueue: "PRODUCT_AVAILABILITY_QUEUE",
	}

	we, err := c.ExecuteWorkflow(context.Background(), options,
		app.NotifyWhenAvailable, "OU812", "213-867-5309")
	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}

	var result string
	err = we.Get(context.Background(), &result)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	log.Println(result)
}
