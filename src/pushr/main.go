package main

import (
	"flag"
	"github.com/tarm/goserial"
	"log"
	"strings"
)

func pushMsg(p string, msg string) {
    // Open the serial port.
    c := &serial.Config{Name: p, Baud: 9600}
    s, err := serial.OpenPort(c)
    if err != nil {
        log.Print(err)
        return
    }
    // Send the message.
    n, err := s.Write([]byte(msg))
    if err != nil {
        log.Print(err)
    }
    log.Printf("%v", n)
}

func main() {

	var p = flag.String("p", "", "the USB port to use")
	var l = flag.Bool("l", false, "list all avliable serial ports")
	flag.Parse()

	if *l == true {
		log.Printf("%v", listSerialPorts())
		return
	}

    var list []string

    if len(*p) > 0 {
        list = []string{*p}
    } else {
        list = listSerialPorts()
    }

    if len(list) == 0 {
        log.Print("No serial ports found.\n")
    }

    for _, port := range list {
        pushMsg(port, strings.Join(flag.Args(), " "))
    }
}
