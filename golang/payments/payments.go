package payments

type PaymentInfo struct {
	Description string
	Usd         int
	Cancelled   bool
}

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	PaymentsInfo  map[int]PaymentInfo
	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		PaymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

func (p PaymentModule) Pay(description string, usd int) int {
	id := p.paymentMethod.Pay(usd)
	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}
	p.PaymentsInfo[id] = info
	return id
}

func (p PaymentModule) Cancel(id int) {
	info, ok := p.PaymentsInfo[id]
	if !ok {
		return
	}
	p.paymentMethod.Cancel(id)
	info.Cancelled = true
	p.PaymentsInfo[id] = info
}

func (p *PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.PaymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}
	return info
}

func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.PaymentsInfo))

	for k, v := range p.PaymentsInfo {
		tempMap[k] = v
	}

	return tempMap
}
