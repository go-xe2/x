package xlog

import (
	"fmt"
	"os"
)

func (l *TLogger) Print(v ...interface{}) {
	l.printStd("", v...)
}

func (l *TLogger) Printf(format string, v ...interface{}) {
	l.printStd(l.format(format, v...))
}

func (l *TLogger) Println(v ...interface{}) {
	l.Print(v...)
}

func (l *TLogger) Printfln(format string, v ...interface{}) {
	l.printStd(l.format(format, v...))
}

func (l *TLogger) Fatal(v ...interface{}) {
	l.printErr("[FATA]", v...)
	os.Exit(1)
}

func (l *TLogger) Fatalf(format string, v ...interface{}) {
	l.printErr("[FATA]", l.format(format, v...))
	os.Exit(1)
}

func (l *TLogger) Fatalfln(format string, v ...interface{}) {
	l.Fatalf(format, v...)
	os.Exit(1)
}

func (l *TLogger) Panic(v ...interface{}) {
	l.printErr("[PANI]", v...)
	panic(fmt.Sprint(v...))
}

func (l *TLogger) Panicf(format string, v ...interface{}) {
	l.printErr("[PANI]", l.format(format, v...))
	panic(l.format(format, v...))
}

func (l *TLogger) Panicfln(format string, v ...interface{}) {
	l.Panicf(format, v...)
}

func (l *TLogger) Info(v ...interface{}) {
	if l.checkLevel(LEVEL_INFO) {
		l.printStd("[INFO]", v...)
	}
}

func (l *TLogger) Infof(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_INFO) {
		l.printStd("[INFO]", l.format(format, v...))
	}
}

func (l *TLogger) Infofln(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_INFO) {
		l.Infof(format, v...)
	}
}

func (l *TLogger) Debug(v ...interface{}) {
	if l.checkLevel(LEVEL_DEBU) {
		l.printStd("[DEBU]", v...)
	}
}

func (l *TLogger) Debugf(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_DEBU) {
		l.printStd("[DEBU]", l.format(format, v...))
	}
}

func (l *TLogger) Debugfln(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_DEBU) {
		l.Debugf(format, v...)
	}
}

func (l *TLogger) Notice(v ...interface{}) {
	if l.checkLevel(LEVEL_NOTI) {
		l.printErr("[NOTI]", v...)
	}
}

func (l *TLogger) Noticef(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_NOTI) {
		l.printErr("[NOTI]", l.format(format, v...))
	}
}

func (l *TLogger) Noticefln(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_NOTI) {
		l.Noticef(format, v...)
	}
}

func (l *TLogger) Warning(v ...interface{}) {
	if l.checkLevel(LEVEL_WARN) {
		l.printErr("[WARN]", v...)
	}
}

func (l *TLogger) Warningf(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_WARN) {
		l.printErr("[WARN]", l.format(format, v...))
	}
}

func (l *TLogger) Warningfln(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_WARN) {
		l.Warningf(format, v...)
	}
}

func (l *TLogger) Error(v ...interface{}) {
	if l.checkLevel(LEVEL_ERRO) {
		l.printErr("[ERRO]", v...)
	}
}

func (l *TLogger) Errorf(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_ERRO) {
		l.printErr("[ERRO]", l.format(format, v...))
	}
}

func (l *TLogger) Errorfln(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_ERRO) {
		l.Errorf(format, v...)
	}
}

func (l *TLogger) Critical(v ...interface{}) {
	if l.checkLevel(LEVEL_CRIT) {
		l.printErr("[CRIT]", v...)
	}
}

func (l *TLogger) Criticalf(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_CRIT) {
		l.printErr("[CRIT]", l.format(format, v...))
	}
}

func (l *TLogger) Criticalfln(format string, v ...interface{}) {
	if l.checkLevel(LEVEL_CRIT) {
		l.Criticalf(format, v...)
	}
}

func (l *TLogger) checkLevel(level int) bool {
	return l.level&level > 0
}
