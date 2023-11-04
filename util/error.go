package util

import "log"

// TODO: custom logger, custom error
func Die(args ...interface{}) {
	i := len(args) - 1
	arg := args[i]
	if arg != nil {
		err := arg.(error)
		log.Fatal(err)
	}
}
