package xlog

import (
	"github.com/go-xe2/x/os/xfile"
	"io"
)

func (l *TLogger) To(writer io.Writer) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetWriter(writer)
	return logger
}

func (l *TLogger) Path(path string) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if path != "" {
		logger.SetPath(path)
	}
	return logger
}

func (l *TLogger) Cat(category string) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if logger.path != "" {
		logger.SetPath(logger.path + xfile.Separator + category)
	}
	return logger
}

func (l *TLogger) File(file string) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetFile(file)
	return logger
}

func (l *TLogger) Level(level int) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetLevel(level)
	return logger
}

func (l *TLogger) Skip(skip int) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetStackSkip(skip)
	return logger
}

func (l *TLogger) Stack(enabled bool, skip ...int) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	logger.SetStack(enabled)
	if len(skip) > 0 {
		logger.SetStackSkip(skip[0])
	}
	return logger
}

func (l *TLogger) Stdout(enabled ...bool) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// stdout printing is enabled if <enabled> is not passed.
	if len(enabled) > 0 && !enabled[0] {
		logger.stdoutPrint = false
	} else {
		logger.stdoutPrint = true
	}
	return logger
}

func (l *TLogger) StdPrint(enabled ...bool) *TLogger {
	return l.Stdout(enabled...)
}

func (l *TLogger) Header(enabled ...bool) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// header is enabled if <enabled> is not passed.
	if len(enabled) > 0 && !enabled[0] {
		logger.SetHeaderPrint(false)
	} else {
		logger.SetHeaderPrint(true)
	}
	return logger
}

func (l *TLogger) Line(long ...bool) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	if len(long) > 0 && long[0] {
		logger.flags |= F_FILE_LONG
	} else {
		logger.flags |= F_FILE_SHORT
	}
	return logger
}

func (l *TLogger) Async(enabled ...bool) *TLogger {
	logger := (*TLogger)(nil)
	if l.parent == nil {
		logger = l.Clone()
	} else {
		logger = l
	}
	// async feature is enabled if <enabled> is not passed.
	if len(enabled) > 0 && !enabled[0] {
		logger.SetAsync(false)
	} else {
		logger.SetAsync(true)
	}
	return logger
}
