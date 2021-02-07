package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func debugf(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
}

func debugln(a ...interface{}) {
	fmt.Fprintln(os.Stderr, a...)
}

func toJSON(a interface{}) string {
	bytes, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
