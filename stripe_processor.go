package main

import "github.com/rs/zerolog"

type StripeProcessor struct {
	logger zerolog.Logger
}

func (p *StripeProcessor) ProcessPayment(amount float64) error {
	p.logger.Printf("Processing payment of $%.2f through Stripe", amount)

	return nil
}

func NewStripeProcessor(logger zerolog.Logger) *StripeProcessor {
	return &StripeProcessor{logger: logger}
}
