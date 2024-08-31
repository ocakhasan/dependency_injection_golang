package main

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type NotificationService interface {
	SendNotification(orderID string) error
}
