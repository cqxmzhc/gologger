# gologger
custom golang logger based on log package just for fun 

# usage
```
go get github.com/cqxmzhc/gologger

import "github.com/cqxmzhc/gologger"

// arguments specify the log destination that support the io.Writer interface.
gologger.CustomLogger(traceHandle, infoHandle, warningHandle, errorHandle)

gologger.Info("info "test)
```
