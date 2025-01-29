package pinproxy

type Logger interface {
	Printf(format string, v ...interface{})
}
