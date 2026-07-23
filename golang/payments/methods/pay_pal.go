package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (c PayPal) Pay(usd int) int {
	fmt.Println("оплата пейпал")
	fmt.Println("сумма:", usd, "монеты")

	return rand.Int()
}

func (c PayPal) Cancel(id int) {
	fmt.Println("отмена пейпал опирация id=", id)
}
