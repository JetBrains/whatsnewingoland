package log

type Do interface {
	Info(message string, args ...string) error
}
