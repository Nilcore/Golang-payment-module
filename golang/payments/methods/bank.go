package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("оплата банком")
	fmt.Println("сумма:", usd, "доларов")

	return rand.Int()
}

func (c Bank) Cancel(id int) {
	fmt.Println("отмена банк опирация id=", id)
}
