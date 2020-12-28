package model

type Product struct {
	Id   string
	Name string
	Age  int
}

// sets

func (p *Product) SetId(value string) {
	p.Id = value
}

func (p *Product) SetName(value string) {
	p.Name = value
}

func (p *Product) SetAge(value int) {
	p.Age = value
}
