package chains

import (
	"github.com/sirupsen/logrus"
)

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Request) {
	if p.doctorCheckUpDone {
		logrus.Info("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	logrus.Info("Doctor checking patient")
	p.doctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}
