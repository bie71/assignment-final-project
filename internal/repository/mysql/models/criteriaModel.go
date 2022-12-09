package models

type CriteriaModel struct {
	Id   int    `dbq:"id"`
	Name string `dbq:"name"`
}

func TableNameCriteria() string {
	return "criteria"
}
func FiledNameCriteria() []string {
	return []string{
		"id",
		"name",
	}
}
