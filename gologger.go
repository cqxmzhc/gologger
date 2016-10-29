package gologger

import (
    "log"
    "io"
)


var (
    trace   *log.Logger
    info    *log.Logger
    warning *log.Logger
    errorLogger   *log.Logger
)


func CustomLogger(
    traceHandle,
    infoHandle,
    warningHandle,
    errorHandle io.Writer) {

    trace = log.New(traceHandle,
                    "|T|-|R|-|A|-|C|-|E|: ",
                    log.Ldate|log.Ltime|log.Lshortfile)

    info = log.New(infoHandle,
                    "|I|-|N|-|F|-|O|: ",
                    log.Ldate|log.Ltime|log.Lshortfile)

    warning = log.New(warningHandle,
                    "|W|-|A|-|R|-|N|-|I|-|N|-|G|: ",
                    log.Ldate|log.Ltime|log.Lshortfile)

    errorLogger = log.New(errorHandle,
                    "|E|-|R|-|R|-|O|-|R|: ",
                    log.Ldate|log.Ltime|log.Lshortfile)
}

func Info(infoMsg string) {
    info.Println(infoMsg)
}

func Trace(traceMsg string) {
    trace.Println(traceMsg)
}

func Error(errorMsg string) {
    errorLogger.Println(errorMsg)
}

func Warning(warningMsg string) {
    warning.Println(warningMsg)
}
