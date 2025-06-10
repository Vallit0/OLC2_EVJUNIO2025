package repl

import (
	"fmt"
)

type Console struct {
	output string // Va a ser la salida de la consola
}

// tiene el "\n"
func (c *Console) Println(s string) {
	c.output += s + "\n"
}

// no tiene el "\n"
func (c *Console) Print(s string) {
	c.output += s
}

func (c *Console) Show() {
	fmt.Println(c.output)
}

func (c *Console) Clear() {
	c.output = ""
}

func NewConsole() *Console {
	return &Console{}
}

func (c *Console) GetOutput() string {
	return c.output
}
