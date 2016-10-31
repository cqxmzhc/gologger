# gologger
custom golang logger based on log package just for fun 

# usage
```
go get github.com/cqxmzhc/gologger

import "github.com/cqxmzhc/gologger"

// arguments specify the log destination that support the io.Writer interface.
gologger.CustomLogger(traceHandle, infoHandle, warningHandle, errorHandle)

gologger.Info("info test")
// output
|I|-|N|-|F|-|O|: 2016/10/31 09:17:19 in [/root/etek/purifiersServer/src/wdserver/wdserver.go:63] <main.main>: []interface {}{"info test"}
```

