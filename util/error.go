package util

import "log"

// TODO: custom logger, custom error
func Die(args ...interface{}) {
	i := len(args) - 1
	err := args[i].(error)
	if err != nil {
		log.Fatal(err)
	}
}
