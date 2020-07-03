package xlog

import "io"

func Expose() *TLogger {
	return logger
}

func To(writer io.Writer) *TLogger {
	return logger.To(writer)
}

func Path(path string) *TLogger {
	return logger.Path(path)
}

func Cat(category string) *TLogger {
	return logger.Cat(category)
}

func File(pattern string) *TLogger {
	return logger.File(pattern)
}

func Level(level int) *TLogger {
	return logger.Level(level)
}

func Skip(skip int) *TLogger {
	return logger.Skip(skip)
}

func Stack(enabled bool, skip ...int) *TLogger {
	return logger.Stack(enabled, skip...)
}

func Stdout(enabled ...bool) *TLogger {
	return logger.Stdout(enabled...)
}

func Header(enabled ...bool) *TLogger {
	return logger.Header(enabled...)
}

func Line(long ...bool) *TLogger {
	return logger.Line(long...)
}

func Async(enabled ...bool) *TLogger {
	return logger.Async(enabled...)
}
