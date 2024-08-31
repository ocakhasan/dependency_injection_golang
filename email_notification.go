package main

import "github.com/rs/zerolog"

type EmailNotification struct {
	logger zerolog.Logger
}

func (n *EmailNotification) SendNotification(orderID string) error {
	n.logger.Printf("Sending email notification for order %s", orderID)

	return nil
}

func NewEmailNotification(logger zerolog.Logger) *EmailNotification {
	return &EmailNotification{logger: logger}
}
