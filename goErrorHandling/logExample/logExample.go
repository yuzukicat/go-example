package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("I am a log")
	log.SetPrefix("main():")
	// log.Fatal("I am a fatal log")
	log.Panic("I am a panic log")
	fmt.Println("cna you see me?")
}
