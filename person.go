package main

type Person struct {
	name string
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}
