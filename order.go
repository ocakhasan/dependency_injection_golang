package main

type OrderService struct {
	paymentProcessor    PaymentProcessor
	notificationService NotificationService
}

func NewOrderService(paymentProcessor PaymentProcessor, notificationService NotificationService) *OrderService {
	return &OrderService{
		paymentProcessor:    paymentProcessor,
		notificationService: notificationService,
	}
}

func (s *OrderService) ProcessOrder(orderID string, amount float64) error {
	err := s.paymentProcessor.ProcessPayment(amount)
	if err != nil {
		return err
	}

	return s.notificationService.SendNotification(orderID)
}
