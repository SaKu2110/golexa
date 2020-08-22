package golexa

import(
	"fmt"
)

func handledebugPrint(status int, request, intent string) {
	debugPrint("%-6d %-25s --> %s\n", status, request, intent)
}

func debugPrint(msg string, value ...interface{}) {
	fmt.Printf(msg, value...)
}
