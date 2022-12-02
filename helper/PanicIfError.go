package helper

import "log"

func RecoverPanic() {
	err := recover()
	if err != nil {
		log.Fatalln("Error : ", err)
	}
}

func PanicIfError(err error) {
	defer RecoverPanic()
	if err != nil {
		panic(err)
	}
}
