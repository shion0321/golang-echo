package utils

import "log"

// Check function
func Check(err error){
	if err != nil {
		log.Fatal(err)
	}
}