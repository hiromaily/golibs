package log

import (
	"log"
	"os"
)

const (
	DEBUG_STATUS uint = iota + 1
	INFO_STATUS
	WARNING_STATUS
	ERROR_STATUS
	FATAL_STATUS
)

const (
	DEBUG_PREFIX   string = "[DEBUG]"
	INFO_PREFIX    string = "[INFO]"
	WARNING_PREFIX string = "[WARNING]"
	ERROR_PREFIX   string = "[ERROR]"
	FATAL_PREFIX   string = "[FATAL]"
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
	logLevel     int    = 1
	logFileLevel int    = 4
	filePathName string = "/var/log/go/xxxx.log"
)

type LogObject struct {
	logger *log.Logger
}

var (
	logStdOut  LogObject = LogObject{}
	logFileOut LogObject = LogObject{}
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
	self.logger.SetOutput(f)
}

//Create New Original Object
func New(prefix string, logFmt int, fileName string) (*log.Logger, error) {
	//e.g. handlelog.New("[ProjectName] ", Ltime|Lshortfile, "/var/log/jamesdebug.log")
	logObj := log.New(os.Stderr, prefix, logFmt)

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
//filePath have to include file name.
func InitializeLog(level, fileLevel , logFmt int, prefix, fileName string) {
	logLevel = level
	logFileLevel = fileLevel

	//Log File Path
	if fileName == "" {
		fileName = filePathName
	}

	//Log Format
	if logFmt == 0 {
		logFmt = log.LstdFlags | log.Lshortfile
		//logFmt = log.Ldate | log.Ltime | log.Lshortfile
	}

	//Log Object
	logStdOut.logger = log.New(os.Stderr, prefix, logFmt)

	logFileOut.logger = log.New(os.Stderr, prefix, logFmt)
	logFileOut.openFile(fileName)
}

//Debug
func Debug(v ...interface{}) {
	//TODO: Sliceの頭にprefixをセットし、渡す必要がある。
	//nv := []interface{}{DEBUG_PREFIX}
	//nv = append(nv, v...)

	if logLevel == DEBUG_STATUS {
		if logFileLevel == DEBUG_STATUS{
			//file
			//logFileOut.logger.Print(nv...)
			logFileOut.logger.Print(DEBUG_PREFIX, v...)
		}else{
			logStdOut.logger.Print(DEBUG_PREFIX, v...)
		}
	}
}

func Debugf(format string, v ...interface{}) {
	if logLevel == DEBUG_STATUS {
		if logFileLevel == DEBUG_STATUS{
			//file
			logFileOut.logger.Printf(DEBUG_PREFIX + format, v...)
		}else{
			logStdOut.logger.Printf(DEBUG_PREFIX + format, v...)
		}
	}
}

//Info
func Info(v ...interface{}) {
	if logLevel <= INFO_STATUS {
		if logFileLevel == INFO_STATUS{
			//file
			logFileOut.logger.Print(INFO_PREFIX, v...)
		}else{
			logStdOut.logger.Print(INFO_PREFIX, v...)
		}
	}
}

func Infof(format string, v ...interface{}) {
	if logLevel == INFO_STATUS {
		if logFileLevel <= INFO_STATUS{
			//file
			logFileOut.logger.Printf(INFO_PREFIX + format, v...)
		}else{
			logStdOut.logger.Printf(INFO_PREFIX + format, v...)
		}
	}
}

//Warn
func Warn(v ...interface{}) {
	if logLevel <= WARNING_STATUS {
		if logFileLevel == WARNING_STATUS{
			//file
			logFileOut.logger.Print(WARNING_PREFIX, v...)
		}else{
			logStdOut.logger.Print(WARNING_PREFIX, v...)
		}
	}
}

func Warnf(format string, v ...interface{}) {
	if logLevel == WARNING_STATUS {
		if logFileLevel <= WARNING_STATUS{
			//file
			logFileOut.logger.Printf(WARNING_PREFIX + format, v...)
		}else{
			logStdOut.logger.Printf(WARNING_PREFIX + format, v...)
		}
	}
}

//Error
func Error(v ...interface{}) {
	if logLevel <= ERROR_STATUS {
		if logFileLevel == ERROR_STATUS{
			//file
			logFileOut.logger.Print(ERROR_PREFIX, v...)
		}else{
			logStdOut.logger.Print(ERROR_PREFIX, v...)
		}
	}
}

func Errorf(format string, v ...interface{}) {
	if logLevel == ERROR_STATUS {
		if logFileLevel <= ERROR_STATUS{
			//file
			logFileOut.logger.Printf(ERROR_PREFIX + format, v...)
		}else{
			logStdOut.logger.Printf(ERROR_PREFIX + format, v...)
		}
	}
}

//Fatal
func Fatal(v ...interface{}) {
	if logLevel <= FATAL_STATUS {
		if logFileLevel == FATAL_STATUS{
			//file
			logFileOut.logger.Print(FATAL_PREFIX, v...)
		}else{
			logStdOut.logger.Print(FATAL_PREFIX, v...)
		}
	}
}

func Fatalf(format string, v ...interface{}) {
	if logLevel == FATAL_STATUS {
		if logFileLevel <= FATAL_STATUS{
			//file
			logFileOut.logger.Printf(FATAL_PREFIX + format, v...)
		}else{
			logStdOut.logger.Printf(FATAL_PREFIX + format, v...)
		}
	}
}
