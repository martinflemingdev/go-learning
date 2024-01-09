package logger

import (
    "log"
    "os"
)

var Logger *log.Logger

func init() {
    // Initialize the global logger
    // Flags used:
    // log.Ldate - include the date in the log message
    // log.Ltime - include the time in the log message
    // log.Lshortfile - include the file name and line number of the log call
    // The flags are combined using the bitwise OR operator '|'
    Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}
