package chains

import "fmt"

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Request) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
	c.next.Execute(p)
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
