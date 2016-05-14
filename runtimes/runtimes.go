package runtimes

import (
	"log"
	"runtime"
	"runtime/debug"
)

//Caller
func Caller() {
	programCounter, sourceFileName, sourceFileLineNum, ok := runtime.Caller(1)
	log.Printf("ok: %t\n", ok)
	log.Printf("programCounter: %v\n", programCounter)
	log.Printf("sourceFileName: %s\n", sourceFileName)
	log.Printf("sourceFileLineNum: %d\n", sourceFileLineNum)

	//PrintStack prints to standard error the stack trace returned by runtime.Stack.
	debug.PrintStack()
}


