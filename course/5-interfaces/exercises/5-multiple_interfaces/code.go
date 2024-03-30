package main

import (
	"fmt"
)

var costPerCharNotSubscribed float64 = 0.05
var costPerCharSubscribed float64 = 0.01

func (e email) cost() float64 {
	if e.isSubscribed {
		return costPerCharSubscribed * float64(len(e.body))

	} else {
		return costPerCharNotSubscribed * float64(len(e.body))
	}
}

func (e email) print() {
	fmt.Println(e.body)
}

// don't touch below this line

type expense interface {
	cost() float64
}

type printer interface {
	print()
}

type email struct {
	isSubscribed bool
	body         string
}

func test(e expense, p printer) {
	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
	p.print()
	fmt.Println("====================================")
}

func main() {
	m := email{
		isSubscribed: true,
		body:         "hello there",
	}
	test(m, m)
	m = email{
		isSubscribed: false,
		body:         "I want my money back",
	}
	test(m, m)
	m = email{
		isSubscribed: true,
		body:         "Are you free for a chat?",
	}
	test(m, m)
	m = email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
	}
	test(m, m)
}
