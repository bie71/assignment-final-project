package models

type CategoriesModel struct {
	CategoryId string `dbq:"category_id"`
	Name       string `dbq:"name"`
}

func TableNameCategories() string {
	return "categories"
}

func FieldNameCategories() []string {
	return []string{
		"category_id",
		"name",
	}
}
