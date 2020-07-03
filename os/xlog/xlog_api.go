package xlog

func Print(v ...interface{}) {
	logger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Println(v ...interface{}) {
	logger.Println(v...)
}

func Printfln(format string, v ...interface{}) {
	logger.Printfln(format, v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func Fatalfln(format string, v ...interface{}) {
	logger.Fatalfln(format, v...)
}

func Panic(v ...interface{}) {
	logger.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}

func Panicfln(format string, v ...interface{}) {
	logger.Panicfln(format, v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Infof(format string, v ...interface{}) {
	logger.Infof(format, v...)
}

func Infofln(format string, v ...interface{}) {
	logger.Infofln(format, v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	logger.Debugf(format, v...)
}

func Debugfln(format string, v ...interface{}) {
	logger.Debugfln(format, v...)
}

func Notice(v ...interface{}) {
	logger.Notice(v...)
}

func Noticef(format string, v ...interface{}) {
	logger.Noticef(format, v...)
}

func Noticefln(format string, v ...interface{}) {
	logger.Noticefln(format, v...)
}

func Warning(v ...interface{}) {
	logger.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	logger.Warningf(format, v...)
}

func Warningfln(format string, v ...interface{}) {
	logger.Warningfln(format, v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v...)
}

func Errorfln(format string, v ...interface{}) {
	logger.Errorfln(format, v...)
}

func Critical(v ...interface{}) {
	logger.Critical(v...)
}

func Criticalf(format string, v ...interface{}) {
	logger.Criticalf(format, v...)
}

func Criticalfln(format string, v ...interface{}) {
	logger.Criticalfln(format, v...)
}
