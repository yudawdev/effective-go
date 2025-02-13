package error

import (
	"os"
)

// panic: no value for $USER
//
// goroutine 1 [running]:
// main.init.0()
//
//	/tmp/sandbox2523853571/prog.go:15 +0x31
var user = os.Getenv("USER")

func init() {
	if user == "" {
		panic("no value for $USER")
	}
}
