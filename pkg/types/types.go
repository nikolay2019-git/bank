package types

//Money представляет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д.)
type Money int64

// Currency представляет код валюты
type Currency string

const (
	TJS Currency = "TJS"
	USD Currency = "USD"
	RUB Currency = "RUB"
	EUR Currency = "EUR"
)

//PAN представляет номер карты
type PAN string

//Card представляет информацию о платёжной карте
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type Category string
type Status string

const (
	StatusOk         Status = "OK"
	StatusFail       Status = "Fail"
	StatusInProgress Status = "INPROGRESS"
)

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

type PaymentSource struct {
	Type    string //'card'
	Number  string // номер вида "5058 xxxx xxxx 8888"
	Balance Money  // Баланс в дирамах

}
