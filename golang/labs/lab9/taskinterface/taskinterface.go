package taskinterface

type TaskInterface interface {
	GetStatus() bool
	GetDescription() string
	SetStatus(bool)
	SetDescription(string)
}
