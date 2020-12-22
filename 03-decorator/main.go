package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(args []string, stdout io.Writer) error {

	var t, c Beverager
	t = NewTea()
	displayPrice(t)
	t = WithMilk(NewTea())
	displayPrice(t)
	c = NewCoffee()
	displayPrice(c)
	c = WithCaramel(WithChocolate(WithMilk(NewCoffee())))
	displayPrice(c)
	c = WithCaramel(WithMilk(WithMilk(NewCoffee())))
	displayPrice(c)

	return nil
}

func displayPrice(bev Beverager) {
	fmt.Printf("Price %d\n", bev.cost())
}

type Beverager interface {
	cost() int
}

type Beverage struct {
	price int
	descr string
}

func (b Beverage) cost() int {
	return b.price
}

func NewCoffee() Coffee {
	return Coffee{}
}

type Coffee struct {
	Beverage
}

func (c Coffee) cost() int {
	return 10
}

func NewTea() Tea {
	return Tea{}
}

type Tea struct {
	Beverage
}

func (c Tea) cost() int {
	return 8
}

func WithMilk(w Beverager) Milk {
	return Milk{wrapped: w}
}

type Milk struct {
	wrapped Beverager
}

func (m Milk) cost() int {
	return m.wrapped.cost() + 2
}
func WithCaramel(w Beverager) Caramel {
	return Caramel{wrapped: w}
}

type Caramel struct {
	wrapped Beverager
}

func (m Caramel) cost() int {
	return m.wrapped.cost() + 3
}
func WithChocolate(w Beverager) Chocolate {
	return Chocolate{wrapped: w}
}

type Chocolate struct {
	wrapped Beverager
}

func (m Chocolate) cost() int {
	return m.wrapped.cost() + 4
}
