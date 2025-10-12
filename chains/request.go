package chains

type Request struct {
	Name              string
	Status            Status
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

type Status int

const (
	Pending Status = iota
	Approved
	Rejected
)

const (
	SEMAT_CHECK = true
	NAV_ENABLED = 1
)
