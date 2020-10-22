package main

import (
	"github.com/albinofreitas/linkapi-golang/internal/orders"
	"github.com/jasonlvhit/gocron"
)

func main() {
	s := gocron.NewScheduler()
	s.Every(5).Seconds().Do(orders.HandlePendingOrders)
	<-s.Start()
}
