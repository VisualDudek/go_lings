package gotchas

import "log"

// Fatal*() and Panic*() will terminate app

func FatalandPanic() {
	log.Fatalln("Fatal Level: log entry") //app exits here
	log.Println("Normal Level: log entry")
}
