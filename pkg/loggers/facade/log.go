package facade

type Log interface {
	Info(string,map[string]string)
	Error(string,map[string]string)
}
