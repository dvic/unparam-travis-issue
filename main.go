package unparamtest

type Processor interface {
	Process(taskSource string, limit int) (bool, error)
}
