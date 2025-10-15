package chains

import (
	"github.com/sirupsen/logrus"
)

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Request) {
	if p.registrationDone {
		logrus.Info("Patient registration already done")
		r.next.Execute(p)
		return
	}
	logrus.Info("Reception registering patient")
	p.registrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}
