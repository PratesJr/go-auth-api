package adapters

type Instrumentation interface {
	Start() error
	Stop()
}
