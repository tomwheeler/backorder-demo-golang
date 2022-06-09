# wfe-go-example
Simple Two-Activity Workflow, for demonstrating event history

The example is currently one that checks an inventory service
every ten minutes for product availability, and then notifies
a customer by SMS when it becomes available.

To run, make sure that you have a Temporal cluster running 
locally, and then execute the following commands in separate 
terminals:

```
go run microservices/inventory-service.go
go run microservices/sms-service.go
go run worker/main.go
go run start/main.go
```

Experiment by killing the inventory service for a while and/or 
changing the quantity of items that it returns (alternating 
between zero and some value greater than zero). For faster
results, change the unit of time in the workflow's sleep
statement from Minute to Second.
