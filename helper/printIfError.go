package helper

import "log"

func PrintIfError(err error) {
	if err != nil {
		log.Println(err)
	}
}
