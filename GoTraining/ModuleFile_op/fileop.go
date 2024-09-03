package main

import (
	"log"
	"os"

	"github.com/magiconair/properties"
)

func main() {
	f, err := os.OpenFile("logfile.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	log.Println("This is a test log entry")

	p := properties.MustLoadFile("config.properties", properties.UTF8)
	if port, ok := p.Get("port"); ok {
		log.Println(port)
	}
}
