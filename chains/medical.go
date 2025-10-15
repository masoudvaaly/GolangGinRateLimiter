package chains

import (
	"github.com/sirupsen/logrus"
)

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Request) {
	if p.medicineDone {
		logrus.Info("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	logrus.Info("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.Execute(p)
}

func (d *Medical) SetNext(next Department) {
	d.next = next
}
