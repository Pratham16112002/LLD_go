package absfact

import (
	"fmt"
)

type PaymentGateway interface {
	processPayment(amount float64)
}

type Invoice interface {
	generateInvoice()
}

type RazorPayGateway struct{}

func NewRazorPayGateway() *RazorPayGateway {
	return &RazorPayGateway{}
}

func (r *RazorPayGateway) processPayment(amount float64) {
	fmt.Println("Prcessing INR payment via razorpay payment gateway")
}

type StripePayGateway struct{}

func NewStripePayGateway() *StripePayGateway {
	return &StripePayGateway{}
}

func (s *StripePayGateway) processPayment(amount float64) {
	fmt.Println("Processing payment via stripe payment gateway")
}

type PaypalPayGateway struct{}

func NewPaypalPayGateway() *PaypalPayGateway {
	return &PaypalPayGateway{}
}

func (p *PaypalPayGateway) processPayment(amount float64) {
	fmt.Println("Processing paymane via paypal payment gateway")
}

type PayUGateway struct{}

func NewPayUGateway() *PayUGateway {
	return &PayUGateway{}
}

func (p *PayUGateway) processPayment(amount float64) {
	fmt.Println("Processing INR payment via Pay U")
}

type PaymentGatewayFactory interface {
	CreatePaymentGateway() PaymentGateway
}

type PayUGatewayFactory struct{}

func (f *PayUGatewayFactory) CreatePaymentGateway() *PayUGateway {
	return NewPayUGateway()
}

type RazorPayGatewayFactor struct{}

func (f *RazorPayGatewayFactor) CreatePaymentGateway() *RazorPayGateway {
	return NewRazorPayGateway()
}

type GSTInvoice struct{}

func NewGSTInvoice() *GSTInvoice {
	return &GSTInvoice{}
}

func (r *GSTInvoice) generateInvoice() {
	fmt.Println("Generating GST Invoice for India")
}

type USAInvoice struct{}

func NewUSInvoice() *USAInvoice {
	return &USAInvoice{}
}

func (u *USAInvoice) generateInvoice() {
	fmt.Println("Generating USA Invoice for USA")
}

// Abstract factory pattern interface
// It defines the methods for creating families of related objects
// Since Payment Gateway and Creating Invoice are related to each other
type ECommerceFactory interface {
	CreatePaymentGateway(string) PaymentGateway
	CreateInvoice() Invoice
}

type IndiaFactory struct{}

func (f *IndiaFactory) CreatePaymentGateway(gatewayType string) PaymentGateway {
	switch gatewayType {
	case "razorpay":
		return NewRazorPayGateway()
	case "payu":
		return NewPayUGateway()
	default:
		return nil
	}
}

func (f *IndiaFactory) CreateInvoice() Invoice {
	return NewGSTInvoice()
}

type USAFactory struct{}

func (u *USAFactory) CreatePaymentGateway(gatewayType string) PaymentGateway {
	switch gatewayType {
	case "stripe":
		return NewStripePayGateway()
	case "paypal":
		return NewStripePayGateway()
	default:
		fmt.Println("unsupported paymentgateway")
		return nil
	}
}

func (u *USAFactory) CreateInvoice() Invoice {
	return NewUSInvoice()
}

type CheckoutService struct {
	factory ECommerceFactory
}

func NewCheckoutService(factory ECommerceFactory) *CheckoutService {
	return &CheckoutService{
		factory: factory,
	}
}

func (cs *CheckoutService) CheckOut(amount float64, gatewayType string) {
	paymentGateway := cs.factory.CreatePaymentGateway(gatewayType)
	paymentGateway.processPayment(amount)
	invoice := cs.factory.CreateInvoice()
	invoice.generateInvoice()
}
