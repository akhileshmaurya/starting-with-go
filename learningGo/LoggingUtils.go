package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	//Init(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	logfile, err := os.Create("/tmp/goag.log")
	if err != nil {
		log.Fatalln("Fail to open log file")
	}
	initLog(ioutil.Discard, logfile, logfile, logfile)

	Trace.Println("Trace I have something standard to say")
	Info.Println("Info Special Information")
	Warning.Println("Warning There is something you need to know about")
	Error.Println("Error Something has failed")
}
