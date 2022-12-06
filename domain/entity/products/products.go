package entity

type Products struct {
	productId         string
	name              string
	price             uint
	categoryId        string
	stock             uint
	productsCondition string
}

type DTOProducts struct {
	ProductId         string
	Name              string
	Price             uint
	CategoryId        string
	Stock             uint
	ProductsCondition string
}

func NewProducts(products *DTOProducts) *Products {
	return &Products{
		productId:         products.ProductId,
		name:              products.Name,
		price:             products.Price,
		categoryId:        products.CategoryId,
		stock:             products.Stock,
		productsCondition: products.ProductsCondition,
	}
}

func (p *Products) ProductId() string {
	return p.productId
}

func (p *Products) NameProduct() string {
	return p.name
}

func (p *Products) Price() uint {
	return p.price
}

func (p *Products) Category() string {
	return p.categoryId
}

func (p *Products) Stock() uint {
	return p.stock
}

func (p *Products) Condition() string {
	return p.productsCondition
}
