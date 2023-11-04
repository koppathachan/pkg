package util

import "log"

// TODO: custom logger, custom error
func Die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
