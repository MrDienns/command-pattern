package main

type Command interface {
	Execute()
}

type CommandSetName struct {
	person *Person
	name   string
}

func (c CommandSetName) Execute() {
	c.person.SetName(c.name)
}
