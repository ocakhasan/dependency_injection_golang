package main

import (
	"github.com/samber/do"
)

func main() {
	injector := Injector()

	// Resolve the OrderService and use it
	orderService := do.MustInvoke[*OrderService](injector)
	err := orderService.ProcessOrder("12345", 100.00)
	if err != nil {
		panic(err)
	}
}
