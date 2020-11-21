package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandSetName_Execute(t *testing.T) {
	p := &Person{name: "Dennis"}
	command := CommandSetName{
		person: p,
		name:   "Martin",
	}
	command.Execute()
	assert.Equal(t, "Martin", p.Name())
}
