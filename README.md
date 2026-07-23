# Golang Payment Module
 
Простой модуль оплаты на Go с поддержкой нескольких платёжных методов (банк, крипта, PayPal) через единый интерфейс.
 
## Структура проекта
 
```
paypal/
├── main.go                  # точка входа, пример использования
├── payments/
│   ├── payment_module.go    # основная логика модуля
│   ├── payment_info.go      # структура информации об оплате
│   └── methods/
│       ├── bank.go          # оплата через банк
│       ├── crypto.go        # оплата криптовалютой
│       └── paypal.go        # оплата через PayPal
```
 
## Как использовать
 
### 1. Выбери платёжный метод
 
```go
method := methods.NewPayPal()
// или
method := methods.NewBank()
// или
method := methods.NewCrypto()
```
 
### 2. Создай модуль оплаты
 
```go
paymentModule := payments.NewPaymentModule(method)
```
 
### 3. Проведи оплату
 
```go
id := paymentModule.Pay("описание покупки", сумма)
```
Метод вернёт `id` операции — сохрани его, если понадобится отменить платёж или получить информацию о нём.
 
### 4. Отмени оплату (если нужно)
 
```go
paymentModule.Cancel(id)
```
 
### 5. Получи информацию об оплате
 
По конкретному id:
```go
info := paymentModule.Info(id)
```
 
По всем оплатам:
```go
allinfo := paymentModule.AllInfo()
```
 
## Пример
 
```go
func main() {
	method := methods.NewPayPal()
	paymentModule := payments.NewPaymentModule(method)
 
	paymentModule.Pay("бургер", 5)
	idPhone := paymentModule.Pay("телефон", 500)
	idGame := paymentModule.Pay("игра", 20)
 
	paymentModule.Cancel(idPhone)
 
	gameInfo := paymentModule.Info(idGame)
	allInfo := paymentModule.AllInfo()
 
	fmt.Println("все наши оплаты", allInfo)
	fmt.Println("game info", gameInfo)
}
```
 
## Как добавить свой платёжный метод
 
Реализуй интерфейс `PaymentMethod`:
 
```go
type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}
```
 
Создай новый файл в `payments/methods/` со своим типом и этими двумя методами — модуль оплаты сможет использовать его без изменений в основной логике.
 
