package gologger

import (
    "log"
    "io"
    "runtime"
    "fmt"
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
                    log.Ldate|log.Ltime)

    info = log.New(infoHandle,
                    "|I|-|N|-|F|-|O|: ",
                    log.Ldate|log.Ltime)

    warning = log.New(warningHandle,
                    "|W|-|A|-|R|-|N|-|I|-|N|-|G|: ",
                    log.Ldate|log.Ltime)

    errorLogger = log.New(errorHandle,
                    "|E|-|R|-|R|-|O|-|R|: ",
                    log.Ldate|log.Ltime)
}

func getDetail(logMsg interface{}) (detail string) {
    pc, fn, line, _ := runtime.Caller(2)
    detail = fmt.Sprintf("in [%s:%d] <%s>: %#v",
                        fn, line, runtime.FuncForPC(pc).Name(), logMsg)

    return
}

func Info(v ...interface{}) {
    info.Println(getDetail(v))
}

func Trace(v ...interface{}) {
    trace.Println(getDetail(v))
}

func Error(v ...interface{}) {
    errorLogger.Println(getDetail(v))
}

func Warning(v ...interface{}) {
    warning.Println(getDetail(v))
}
