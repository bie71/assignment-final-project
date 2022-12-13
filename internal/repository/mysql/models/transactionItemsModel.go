package models

type TransactionItemsModel struct {
	Id            int    `dbq:"id"`
	TransactionId string `dbq:"transaction_id"`
	ProductId     string `dbq:"product_id"`
	Quantity      int    `dbq:"quantity"`
}

func TableNameTransactionItems() string {
	return "transaction_items"
}
func FieldNameTransactionItems() []string {
	return []string{
		"id",
		"transaction_id",
		"product_id",
		"quantity",
	}
}
