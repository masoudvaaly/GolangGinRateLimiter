package chains

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Request) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.Execute(p)
}

func (d *Medical) SetNext(next Department) {
	d.next = next
}
