# TimeWriter

TimeWriter implements io.Writer to roll daily and comporess log file time

## Overview

* implements io.Writer. You can easily use in golang log, GORM, grpclog etc.
* daily roll log, you can specific the log file name's prefix, default is process name
* compress to gz for old file

## Getting Started

**Example**

To use TimeWriter, you can git clone [https://github.com/longbozhan/timewriter](https://github.com/longbozhan/timewriter), and import like this:

```go
package main

import (
	"log"
)

func main() {
	timeWriter := &TimeWriter{
		Dir:        "./log",
		Compress:   true,
		ReserveDay: 30,
	}
	logDebug := log.New(timeWriter, " [Debug] ", log.LstdFlags)
	logDebug.Println("this is debug")
}

```

## Reference

* [lumberject](https://github.com/natefinch/lumberjack)
