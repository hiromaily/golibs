package log

import (
	"log"
	"os"
)

const (
	debugStatus uint = iota + 1
	infoStatus
	warningStatus
	errorStatus
	fatalStatus
)

/*
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger

*/

var (
	logLevel    uint   = 1
	logFilePath string = ""
)

type LogObject struct {
	logger *log.Logger
	f      *os.File
}

var (
	logDebug LogObject = LogObject{}
	logInfo  LogObject = LogObject{}
	logWarn  LogObject = LogObject{}
	logError LogObject = LogObject{}
	logFatal LogObject = LogObject{}
)

//for output log file
func (self *LogObject) openFile(fileName string) {
	if fileName == "" {
		return
	}

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Error opening file :", err.Error())
	}
	self.f = f

	self.logger.SetOutput(self.f)
}

//Create New Original Object
func New(logTitle string, logFmt int, fileName string) (*log.Logger, error) {
	//e.g. handlelog.New("[JamesTest] ", Ltime|Lshortfile, "jamesdebug.log")
	logObj := log.New(os.Stderr, logTitle, logFmt)

	if fileName != "" {
		f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal("Error opening logfile :", err.Error())
			return nil, err
		}
		logObj.SetOutput(f)
	}

	return logObj, nil
}

//Initialize base log object using default setting
func InitLog(level uint, logFmt int, filePath string) {
	//e.g. handlelog.InitializeDefaultLog(1, 0)
	logLevel = level //default(info)

	//Log Format
	if logFmt == 0 {
		logFmt = log.LstdFlags | log.Lshortfile
		//logFmt = log.Ldate | log.Ltime | log.Lshortfile
	}

	//Log Object
	logDebug.logger = log.New(os.Stderr, "[DEBUG] ", logFmt)
	logInfo.logger = log.New(os.Stderr, "[INFO]  ", logFmt)
	logWarn.logger = log.New(os.Stderr, "[WARN]  ", logFmt)
	logError.logger = log.New(os.Stderr, "[ERROR] ", logFmt)
	logFatal.logger = log.New(os.Stderr, "[FATAL] ", logFmt)

	if filePath != "" {
		logDebug.openFile(filePath + "debug.log")
		logInfo.openFile(filePath + "info.log")
		logWarn.openFile(filePath + "warn.log")
		logError.openFile(filePath + "error.log")
		logFatal.openFile(filePath + "fatal.log")
	}
}

//Debug
func Debug(v ...interface{}) {
	if logLevel == debugStatus {
		//logDebug.logger.Printf(": %s", msg)
		logDebug.logger.Print(v...)
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel == debugStatus {
		//logDebug.logger.Printf(": %s", msg)
		logDebug.logger.Printf(format, v...)
	}
}

//Info
func Info(v ...interface{}) {
	if logLevel <= infoStatus {
		logInfo.logger.Print(v...)
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel <= infoStatus {
		logInfo.logger.Printf(format, v...)
	}
}

//Warn
func Warn(v ...interface{}) {
	if logLevel <= warningStatus {
		logWarn.logger.Print(v...)
	}
}

func Warnf(format string, v ...interface{}) {
	if logLevel <= warningStatus {
		logWarn.logger.Printf(format, v...)
	}
}

//Error
func Error(v ...interface{}) {
	if logLevel <= errorStatus {
		logError.logger.Print(v...)
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel <= errorStatus {
		logError.logger.Printf(format, v...)
	}
}

//Fatal
func Fatal(v ...interface{}) {
	if logLevel <= fatalStatus {
		logFatal.logger.Fatal(v...)
	}
}

func Fatalf(format string, v ...interface{}) {
	if logLevel <= fatalStatus {
		logFatal.logger.Fatalf(format, v...)
	}
}
