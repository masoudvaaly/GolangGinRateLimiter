package chains

type Department interface {
	Execute(*Request)
	SetNext(Department)
}
