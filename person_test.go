package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerson_Name(t *testing.T) {
	p := &Person{name: "Dennis"}
	assert.Equal(t, "Dennis", p.Name())
}

func TestPerson_SetName(t *testing.T) {
	p := &Person{name: "Dennis"}
	p.SetName("Martin")
	assert.Equal(t, "Martin", p.Name())
}
