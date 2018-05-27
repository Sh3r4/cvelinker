package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/fatih/color"
)

type levelOutput struct {
	// logHandles
	OverShare   *log.Logger
	OverShareNP *log.Logger
	Print       *log.Logger
	PrintNP     *log.Logger
	Warn        *log.Logger
	Error       *log.Logger
	Fatality    *log.Logger

	// colours
	green   *color.Color
	yellow  *color.Color
	red     *color.Color
	magenta *color.Color
	blue    *color.Color
}

func (l *levelOutput) InitColours() {
	// setup the colours
	l.green = color.New(color.FgGreen, color.Bold)
	l.yellow = color.New(color.FgYellow, color.Bold)
	l.red = color.New(color.FgRed, color.Bold)
	l.magenta = color.New(color.FgMagenta, color.Bold)
	l.blue = color.New(color.FgBlue, color.Bold)
}

// Init sets up logging based on the log level.
// Fatality will ALWAYS output logs
// dateStamps = true will add dates to all log output
// timeStamps = true will add times to all log output
// logLevel >=4 outputs all logs; =3 outputs Print,Warn,Error; =2 outputs Warn,Error; =1 outputs Error, =0 suppresses all
func (l *levelOutput) Init(logLevel int, dateStamps bool, timeStamps bool) {

	// init the colours in case it hasn't been done already
	l.InitColours()

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

	OverSharePrefix := "[-] "
	PrintPrefix := "[-] "
	WarnPrefix := "[!] "
	ErrorPrefix := "[ERROR] "

	if dateStamps && timeStamps {
		l.OverShare = l.setupLoggerWithDatesAndTimes(overShareWriter, l.green, OverSharePrefix)
		l.Print = l.setupLoggerWithDatesAndTimes(printWriter, l.green, PrintPrefix)
		l.OverShareNP = l.setupLoggerWithDatesAndTimes(overShareWriter, l.green, "")
		l.PrintNP = l.setupLoggerWithDatesAndTimes(printWriter, l.green, "")
		l.Warn = l.setupLoggerWithDatesAndTimes(warnWriter, l.magenta, WarnPrefix)
		l.Error = l.setupLoggerWithDatesAndTimes(errorWriter, l.red, ErrorPrefix)
	} else if dateStamps && !timeStamps {
		l.OverShare = l.setupLoggerWithDates(overShareWriter, l.green, OverSharePrefix)
		l.Print = l.setupLoggerWithDates(printWriter, l.green, PrintPrefix)
		l.OverShareNP = l.setupLoggerWithDates(overShareWriter, l.green, "")
		l.PrintNP = l.setupLoggerWithDates(printWriter, l.green, "")
		l.Warn = l.setupLoggerWithDates(warnWriter, l.magenta, WarnPrefix)
		l.Error = l.setupLoggerWithDates(errorWriter, l.red, ErrorPrefix)
	} else if timeStamps && !dateStamps {
		l.OverShare = l.setupLoggerWithTimes(overShareWriter, l.green, OverSharePrefix)
		l.Print = l.setupLoggerWithTimes(printWriter, l.green, PrintPrefix)
		l.OverShareNP = l.setupLoggerWithTimes(overShareWriter, l.green, "")
		l.PrintNP = l.setupLoggerWithTimes(printWriter, l.green, "")
		l.Warn = l.setupLoggerWithTimes(warnWriter, l.magenta, WarnPrefix)
		l.Error = l.setupLoggerWithTimes(errorWriter, l.red, ErrorPrefix)
	} else {
		l.OverShare = l.setupLogger(overShareWriter, l.green, OverSharePrefix)
		l.Print = l.setupLogger(printWriter, l.green, PrintPrefix)
		l.OverShareNP = l.setupLogger(overShareWriter, l.green, "")
		l.PrintNP = l.setupLogger(printWriter, l.green, "")
		l.Warn = l.setupLogger(warnWriter, l.magenta, WarnPrefix)
		l.Error = l.setupLogger(errorWriter, l.red, ErrorPrefix)
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
