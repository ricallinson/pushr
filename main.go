package main

import (
	"github.com/tarm/goserial"
	"log"
	"flag"
	"strings"
)

func main() {
	var port = flag.String("p", "/dev/cu.usbserial-A901LLNP", "the USB port to use")
	flag.Parse()
	// Open the serial port.
	c := &serial.Config{Name: *port, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	// Send the message.
	n, err := s.Write([]byte(strings.Join(flag.Args(), " ")))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", n)
}
