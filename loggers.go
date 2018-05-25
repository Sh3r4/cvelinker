package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

type levelOutput struct {
	// loghandles
	OverShare *log.Logger
	Print     *log.Logger
	Warn      *log.Logger
	Error     *log.Logger
	Fatality  *log.Logger

	// colours
	green   *color.Color
	yellow  *color.Color
	red     *color.Color
	magenta *color.Color
	blue    *color.Color
}

// Init sets up logging based on the log level.
// Fatality will ALWAYS output logs
// datestamps = true will add dates to all log output
// timestamps = true will add times to all log output
// loglevel >=4 outputs all logs; =3 outputs Print,Warn,Error; =2 outputs Warn,Error; =1 outputs Error, =0 suppresses all
func (l *levelOutput) Init(logLevel int, dateStamps bool, timeStamps bool) {

	// setup the colours
	l.green = color.New(color.FgGreen, color.Bold)
	l.yellow = color.New(color.FgYellow, color.Bold)
	l.red = color.New(color.FgRed, color.Bold)
	l.magenta = color.New(color.FgMagenta, color.Bold)
	l.blue = color.New(color.FgBlue, color.Bold)

	// setup log handles
	// Fatality always gets logged with all options, so set it up first
	l.Fatality = log.New(os.Stderr, "[!!!] FATAL: ", log.Ldate|log.Ltime|log.Lshortfile)

	// highest logging level
	if logLevel >= 4 {
		l.setupLoggers(
			os.Stdout,
			os.Stdout,
			os.Stdout,
			os.Stderr,
			dateStamps,
			timeStamps,
		)
	} else if logLevel == 3 {
		l.setupLoggers(
			ioutil.Discard,
			os.Stdout,
			os.Stdout,
			os.Stderr,
			dateStamps,
			timeStamps,
		)
	} else if logLevel == 2 {
		l.setupLoggers(
			ioutil.Discard,
			ioutil.Discard,
			os.Stdout,
			os.Stderr,
			dateStamps,
			timeStamps,
		)
	} else if logLevel == 1 {
		l.setupLoggers(
			ioutil.Discard,
			ioutil.Discard,
			ioutil.Discard,
			os.Stderr,
			dateStamps,
			timeStamps,
		)
	} else {
		l.setupLoggers(
			ioutil.Discard,
			ioutil.Discard,
			ioutil.Discard,
			ioutil.Discard,
			dateStamps,
			timeStamps,
		)
	}

}

func (l *levelOutput) setupLoggers(overShareWriter io.Writer, printWriter io.Writer, warnWriter io.Writer, errorWriter io.Writer, dateStamps bool, timeStamps bool) {

	if dateStamps && timeStamps {
		l.OverShare = l.setupLoggerWithDatesAndTimes(overShareWriter, l.green, "[-] ")
		l.Print = l.setupLoggerWithDatesAndTimes(printWriter, l.green, "[-] ")
		l.Warn = l.setupLoggerWithDatesAndTimes(warnWriter, l.magenta, "[!] ")
		l.Error = l.setupLoggerWithDatesAndTimes(errorWriter, l.red, "[ERROR] ")
	} else if dateStamps && !timeStamps {
		l.OverShare = l.setupLoggerWithDates(overShareWriter, l.green, "[-] ")
		l.Print = l.setupLoggerWithDates(printWriter, l.green, "[-] ")
		l.Warn = l.setupLoggerWithDates(warnWriter, l.magenta, "[!] ")
		l.Error = l.setupLoggerWithDates(errorWriter, l.red, "[ERROR] ")
	} else if timeStamps && !dateStamps {
		l.OverShare = l.setupLoggerWithTimes(overShareWriter, l.green, "[-] ")
		l.Print = l.setupLoggerWithTimes(printWriter, l.green, "[-] ")
		l.Warn = l.setupLoggerWithTimes(warnWriter, l.magenta, "[!] ")
		l.Error = l.setupLoggerWithTimes(errorWriter, l.red, "[ERROR] ")
	} else {
		l.OverShare = l.setupLogger(overShareWriter, l.green, "")
		l.Print = l.setupLogger(printWriter, l.green, "")
		l.Warn = l.setupLogger(warnWriter, l.magenta, "[!] ")
		l.Error = l.setupLogger(errorWriter, l.red, "[ERROR] ")
	}
}

func (l *levelOutput) setupLoggerWithDatesAndTimes(outputWriter io.Writer, colouriser *color.Color, prefixText string) (logger *log.Logger) {
	return log.New(outputWriter, colouriser.Sprintf(prefixText), log.Ldate|log.Ltime)
}

func (l *levelOutput) setupLoggerWithDates(outputWriter io.Writer, colouriser *color.Color, prefixText string) (logger *log.Logger) {
	return log.New(outputWriter, colouriser.Sprintf(prefixText), log.Ldate)
}

func (l *levelOutput) setupLoggerWithTimes(outputWriter io.Writer, colouriser *color.Color, prefixText string) (logger *log.Logger) {
	return log.New(outputWriter, colouriser.Sprintf(prefixText), log.Ltime)
}

func (l *levelOutput) setupLogger(outputWriter io.Writer, colouriser *color.Color, prefixText string) (logger *log.Logger) {
	return log.New(outputWriter, colouriser.Sprintf(prefixText), 0)
}
