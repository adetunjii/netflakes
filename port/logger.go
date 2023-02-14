package port

type Logger interface {
	Info(msg string, args ...interface{})
	Error(msg string, err error)
	Fatal(msg string, err error)
}
