package transactions

type Transaction struct {
	Id    int64  `json:"id"`
	Payer string `json:"payer"`
	Payee string `json:"payee"`
}
