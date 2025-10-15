package chains

import (
	"github.com/sirupsen/logrus"
)

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Request) {
	if p.paymentDone {
		logrus.Info("Payment Done")
	}
	logrus.Info("Cashier getting money from patient patient")
	c.next.Execute(p)
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}
