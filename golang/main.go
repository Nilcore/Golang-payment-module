package main

import (
	"fmt"
	"paypal/payments"
	"paypal/payments/methods"
)

func main() {
	method := methods.NewPayPal()

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("бургер", 5)
	IdPhone := paymentModule.Pay("телефон", 500)
	IdGame := paymentModule.Pay("стим", 20)

	paymentModule.Cancel(IdPhone)
	gameinfo := paymentModule.Info(IdGame)
	allinfo := paymentModule.AllInfo()

	fmt.Println("все наши оплаты", allinfo)
	fmt.Println("гейм инфо", gameinfo)
}
