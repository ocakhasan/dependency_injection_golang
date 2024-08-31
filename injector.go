package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/samber/do"
)

func Injector() *do.Injector {
	injector := do.New()

	do.Provide(injector, func(i *do.Injector) (zerolog.Logger, error) {
		return zerolog.New(os.Stdout).Level(zerolog.DebugLevel), nil
	})

	// Bind implementations to the interfaces
	do.Provide(injector, func(i *do.Injector) (PaymentProcessor, error) {
		logger := do.MustInvoke[zerolog.Logger](i)

		return NewStripeProcessor(logger), nil
	})

	// Bind implementations to the interfaces
	do.Provide(injector, func(i *do.Injector) (NotificationService, error) {
		logger := do.MustInvoke[zerolog.Logger](i)

		return NewEmailNotification(logger), nil
	})

	do.Provide(injector, func(i *do.Injector) (*OrderService, error) {
		paymentProcessor := do.MustInvoke[PaymentProcessor](i)
		notificationService := do.MustInvoke[NotificationService](i)

		return NewOrderService(paymentProcessor, notificationService), nil
	})

	return injector
}
