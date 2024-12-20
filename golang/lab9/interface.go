package lab9

type TaskInterface interface {
	GetDescription() string
	GetStatus() string
	SetDescription(string)
	SetStatus(string)
}
