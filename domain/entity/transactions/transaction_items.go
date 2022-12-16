package entity

type TransactionItems struct {
	id            int
	transactionId string
	productId     string
	quantity      uint
}

type DTOTransactionItems struct {
	Id            int
	TransactionId string
	ProductId     string
	Quantity      uint
}

type TransactionItemsProduct struct {
	Id            int
	TransactionId string
	ProductId     string
	Name          string
	Price         uint
	Quantity      uint
}

func NewTransactionItems(DTO *DTOTransactionItems) *TransactionItems {
	return &TransactionItems{
		id:            DTO.Id,
		transactionId: DTO.TransactionId,
		productId:     DTO.ProductId,
		quantity:      DTO.Quantity,
	}
}

func (t *TransactionItems) Id() int {
	return t.id
}

func (t *TransactionItems) TransactionId() string {
	return t.transactionId
}

func (t *TransactionItems) ProductId() string {
	return t.productId
}

func (t *TransactionItems) Quantity() uint {
	return t.quantity
}
