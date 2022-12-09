package entity

type Criteria struct {
	id   int
	name string
}

type DTOCriteria struct {
	Id   int
	Name string
}

func NewCriteria(DTO *DTOCriteria) *Criteria {
	return &Criteria{id: DTO.Id, name: DTO.Name}
}

func (c *Criteria) Id() int {
	return c.id
}

func (c *Criteria) CriteriaName() string {
	return c.name
}
