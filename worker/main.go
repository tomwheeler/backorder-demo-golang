package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"wfe-go-example/app"
)

func main() {
	c, err := client.NewClient(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, "PRODUCT_AVAILABILITY_QUEUE", worker.Options{})
	w.RegisterWorkflow(app.NotifyWhenAvailable)
	w.RegisterActivity(app.GetQuantity)
	w.RegisterActivity(app.NotifyCustomer)

	err = w.Run(worker.InterruptCh())
	if err != nil {
		log.Fatalln("unable to start Worker", err)
	}
}
